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

func GetoidcconfigurationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/workspaces/%s/pipelines-config/identity/oidc/.well-known/openid-configuration", cfg.BaseURL, workspace)
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

func CreateGetoidcconfigurationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_workspaces_workspace_pipelines-config_identity_oidc_.well-known_openid-configuration",
		mcp.WithDescription("Get OpenID configuration for OIDC in Pipelines"),
		mcp.WithString("workspace", mcp.Required(), mcp.Description("This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetoidcconfigurationHandler(cfg),
	}
}
