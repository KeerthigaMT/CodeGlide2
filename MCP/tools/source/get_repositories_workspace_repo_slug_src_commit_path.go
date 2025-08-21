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

func Get_repositories_workspace_repo_slug_src_commit_pathHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["format"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("format=%v", val))
		}
		if val, ok := args["q"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("q=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		if val, ok := args["max_depth"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("max_depth=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/repositories/%s/%s/src/%s/%s%s", cfg.BaseURL, queryString)
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
		var result models.Paginatedtreeentries
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

func CreateGet_repositories_workspace_repo_slug_src_commit_pathTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_repositories_workspace_repo_slug_src_commit_path",
		mcp.WithDescription("Get file or directory contents"),
		mcp.WithString("format", mcp.Description("If 'meta' is provided, returns the (json) meta data for the contents of the file.  If 'rendered' is provided, returns the contents of a non-binary file in HTML-formatted rendered markup. Since Git does not generally track what text encoding scheme is used, this endpoint attempts to detect the most appropriate character encoding. While usually correct, determining the character encoding can be ambiguous which in exceptional cases can lead to misinterpretation of the characters. As such, the raw element in the response object should not be treated as equivalent to the file's actual contents.")),
		mcp.WithString("q", mcp.Description("Optional filter expression as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering).")),
		mcp.WithString("sort", mcp.Description("Optional sorting parameter as per [filtering and sorting](/cloud/bitbucket/rest/intro/#sorting-query-results).")),
		mcp.WithNumber("max_depth", mcp.Description("If provided, returns the contents of the repository and its subdirectories recursively until the specified max_depth of nested directories. When omitted, this defaults to 1.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_repositories_workspace_repo_slug_src_commit_pathHandler(cfg),
	}
}
