package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bitbucket-api/mcp-server/config"
	"github.com/bitbucket-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_repositories_workspace_repo_slug_diff_specHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["context"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("context=%v", val))
		}
		if val, ok := args["path"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("path=%v", val))
		}
		if val, ok := args["ignore_whitespace"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ignore_whitespace=%v", val))
		}
		if val, ok := args["binary"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("binary=%v", val))
		}
		if val, ok := args["renames"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("renames=%v", val))
		}
		if val, ok := args["merge"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("merge=%v", val))
		}
		if val, ok := args["topic"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("topic=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/repositories/%s/%s/diff/%s%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGet_repositories_workspace_repo_slug_diff_specTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_repositories_workspace_repo_slug_diff_spec",
		mcp.WithDescription("Compare two commits"),
		mcp.WithNumber("context", mcp.Description("Generate diffs with <n> lines of context instead of the usual three.")),
		mcp.WithString("path", mcp.Description("Limit the diff to a particular file (this parameter\ncan be repeated for multiple paths).")),
		mcp.WithBoolean("ignore_whitespace", mcp.Description("Generate diffs that ignore whitespace.")),
		mcp.WithBoolean("binary", mcp.Description("Generate diffs that include binary files, true if omitted.")),
		mcp.WithBoolean("renames", mcp.Description("Whether to perform rename detection, true if omitted.")),
		mcp.WithBoolean("merge", mcp.Description("This parameter is deprecated and will be removed at the end\nof 2022. The 'topic' parameter should be used instead. The\n'merge' and 'topic' parameters cannot be both used at the same\ntime.\n\nIf true, the source commit is merged into the\ndestination commit, and then a diff from the\ndestination to the merge result is returned. If false,\na simple 'two dot' diff between the source and\ndestination is returned. True if omitted.")),
		mcp.WithBoolean("topic", mcp.Description("If true, returns 2-way 'three-dot' diff.\nThis is a diff between the source commit and the merge base\nof the source commit and the destination commit.\nIf false, a simple 'two dot' diff between the source and\ndestination is returned.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_repositories_workspace_repo_slug_diff_specHandler(cfg),
	}
}
