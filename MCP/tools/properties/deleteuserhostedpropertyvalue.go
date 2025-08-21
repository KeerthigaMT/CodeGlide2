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

func DeleteuserhostedpropertyvalueHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/users/%s/properties/%s/%s", cfg.BaseURL, selected_user, app_key, property_name)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BasicAuth != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Basic %s", cfg.BasicAuth))
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

func CreateDeleteuserhostedpropertyvalueTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_users_selected_user_properties_app_key_property_name",
		mcp.WithDescription("Delete a user application property"),
		mcp.WithString("selected_user", mcp.Required(), mcp.Description("Either the UUID of the account surrounded by curly-braces, for example `{account UUID}`, OR an Atlassian Account ID.")),
		mcp.WithString("app_key", mcp.Required(), mcp.Description("The key of the Connect app.")),
		mcp.WithString("property_name", mcp.Required(), mcp.Description("The name of the property.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeleteuserhostedpropertyvalueHandler(cfg),
	}
}
