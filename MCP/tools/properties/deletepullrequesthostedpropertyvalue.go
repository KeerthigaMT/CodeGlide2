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

func DeletepullrequesthostedpropertyvalueHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		pullrequest_idVal, ok := args["pullrequest_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: pullrequest_id"), nil
		}
		pullrequest_id, ok := pullrequest_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: pullrequest_id"), nil
		}
		app_keyVal, ok := args["app_key"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: app_key"), nil
		}
		app_key, ok := app_keyVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: app_key"), nil
		}
		property_nameVal, ok := args["property_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: property_name"), nil
		}
		property_name, ok := property_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: property_name"), nil
		}
		url := fmt.Sprintf("%s/repositories/%s/%s/pullrequests/%s/properties/%s/%s", cfg.BaseURL, workspace, repo_slug, pullrequest_id, app_key, property_name)
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

func CreateDeletepullrequesthostedpropertyvalueTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_repositories_workspace_repo_slug_pullrequests_pullrequest_id_properties_app_key_property_name",
		mcp.WithDescription("Delete a pull request application property"),
		mcp.WithString("workspace", mcp.Required(), mcp.Description("The repository container; either the workspace slug or the UUID in curly braces.")),
		mcp.WithString("repo_slug", mcp.Required(), mcp.Description("The repository.")),
		mcp.WithString("pullrequest_id", mcp.Required(), mcp.Description("The pull request ID.")),
		mcp.WithString("app_key", mcp.Required(), mcp.Description("The key of the Connect app.")),
		mcp.WithString("property_name", mcp.Required(), mcp.Description("The name of the property.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletepullrequesthostedpropertyvalueHandler(cfg),
	}
}
