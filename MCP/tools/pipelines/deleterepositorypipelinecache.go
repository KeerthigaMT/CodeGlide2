package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bitbucket-api/mcp-server/config"
	"github.com/bitbucket-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DeleterepositorypipelinecacheHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		workspaceVal, ok := args["workspace"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: workspace"), nil
		}
		workspace, ok := workspaceVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: workspace"), nil
		}
		repo_slugVal, ok := args["repo_slug"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: repo_slug"), nil
		}
		repo_slug, ok := repo_slugVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: repo_slug"), nil
		}
		cache_uuidVal, ok := args["cache_uuid"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: cache_uuid"), nil
		}
		cache_uuid, ok := cache_uuidVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: cache_uuid"), nil
		}
		url := fmt.Sprintf("%s/repositories/%s/%s/pipelines-config/caches/%s", cfg.BaseURL, workspace, repo_slug, cache_uuid)
		req, err := http.NewRequest("DELETE", url, nil)
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

func CreateDeleterepositorypipelinecacheTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_repositories_workspace_repo_slug_pipelines-config_caches_cache_uuid",
		mcp.WithDescription("Delete a cache"),
		mcp.WithString("workspace", mcp.Required(), mcp.Description("The account.")),
		mcp.WithString("repo_slug", mcp.Required(), mcp.Description("The repository.")),
		mcp.WithString("cache_uuid", mcp.Required(), mcp.Description("The UUID of the cache to delete.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeleterepositorypipelinecacheHandler(cfg),
	}
}
