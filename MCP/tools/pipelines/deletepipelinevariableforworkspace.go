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

func DeletepipelinevariableforworkspaceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		variable_uuidVal, ok := args["variable_uuid"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: variable_uuid"), nil
		}
		variable_uuid, ok := variable_uuidVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: variable_uuid"), nil
		}
		url := fmt.Sprintf("%s/workspaces/%s/pipelines-config/variables/%s", cfg.BaseURL, workspace, variable_uuid)
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

func CreateDeletepipelinevariableforworkspaceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_workspaces_workspace_pipelines-config_variables_variable_uuid",
		mcp.WithDescription("Delete a variable for a workspace"),
		mcp.WithString("workspace", mcp.Required(), mcp.Description("This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.")),
		mcp.WithString("variable_uuid", mcp.Required(), mcp.Description("The UUID of the variable to delete.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletepipelinevariableforworkspaceHandler(cfg),
	}
}
