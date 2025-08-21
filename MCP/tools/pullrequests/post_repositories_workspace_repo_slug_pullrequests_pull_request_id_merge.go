package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/bitbucket-api/mcp-server/config"
	"github.com/bitbucket-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Post_repositories_workspace_repo_slug_pullrequests_pull_request_id_mergeHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["async"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("async=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Pullrequestmergeparameters
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/repositories/%s/%s/pullrequests/%s/merge%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.Pullrequest
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

func CreatePost_repositories_workspace_repo_slug_pullrequests_pull_request_id_mergeTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_repositories_workspace_repo_slug_pullrequests_pull_request_id_merge",
		mcp.WithDescription("Merge a pull request"),
		mcp.WithBoolean("async", mcp.Description("Default value is false.\n\n\nWhen set to true, runs merge asynchronously and\nimmediately returns a 202 with polling link to\nthe task-status API in the Location header.\n\n\nWhen set to false, runs merge and waits for it to\ncomplete, returning 200 when it succeeds. If the\nduration of the merge exceeds a timeout threshold,\nthe API returns a 202 with polling link to the\ntask-status API in the Location header.")),
		mcp.WithString("message", mcp.Description("Input parameter: The commit message that will be used on the resulting commit.")),
		mcp.WithString("type", mcp.Required(), mcp.Description("")),
		mcp.WithBoolean("close_source_branch", mcp.Description("Input parameter: Whether the source branch should be deleted. If this is not provided, we fallback to the value used when the pull request was created, which defaults to False")),
		mcp.WithString("merge_strategy", mcp.Description("Input parameter: The merge strategy that will be used to merge the pull request.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_repositories_workspace_repo_slug_pullrequests_pull_request_id_mergeHandler(cfg),
	}
}
