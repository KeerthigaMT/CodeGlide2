package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/bitbucket-api/mcp-server/config"
	"github.com/bitbucket-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreatepipelinevariableforuserHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody models.Pipelinevariable
		
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
		url := fmt.Sprintf("%s/users/%s/pipelines_config/variables", cfg.BaseURL, selected_user)
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

func CreateCreatepipelinevariableforuserTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_users_selected_user_pipelines_config_variables",
		mcp.WithDescription("Create a variable for a user"),
		mcp.WithString("selected_user", mcp.Required(), mcp.Description("Either the UUID of the account surrounded by curly-braces, for example `{account UUID}`, OR an Atlassian Account ID.")),
		mcp.WithString("type", mcp.Description("")),
		mcp.WithString("uuid", mcp.Description("Input parameter: The UUID identifying the variable.")),
		mcp.WithString("value", mcp.Description("Input parameter: The value of the variable. If the variable is secured, this will be empty.")),
		mcp.WithString("key", mcp.Description("Input parameter: The unique name of the variable.")),
		mcp.WithBoolean("secured", mcp.Description("Input parameter: If true, this variable will be treated as secured. The value will never be exposed in the logs or the REST API.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreatepipelinevariableforuserHandler(cfg),
	}
}
