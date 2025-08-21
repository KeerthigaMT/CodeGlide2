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

func GetpipelinevariableforuserHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		selected_userVal, ok := args["selected_user"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: selected_user"), nil
		}
		selected_user, ok := selected_userVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: selected_user"), nil
		}
		variable_uuidVal, ok := args["variable_uuid"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: variable_uuid"), nil
		}
		variable_uuid, ok := variable_uuidVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: variable_uuid"), nil
		}
		url := fmt.Sprintf("%s/users/%s/pipelines_config/variables/%s", cfg.BaseURL, selected_user, variable_uuid)
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
		var result models.Pipelinevariable
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

func CreateGetpipelinevariableforuserTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_users_selected_user_pipelines_config_variables_variable_uuid",
		mcp.WithDescription("Get a variable for a user"),
		mcp.WithString("selected_user", mcp.Required(), mcp.Description("Either the UUID of the account surrounded by curly-braces, for example `{account UUID}`, OR an Atlassian Account ID.")),
		mcp.WithString("variable_uuid", mcp.Required(), mcp.Description("The UUID of the variable to retrieve.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetpipelinevariableforuserHandler(cfg),
	}
}
