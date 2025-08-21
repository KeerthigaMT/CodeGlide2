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

func UpdatepipelinevariableforteamHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		usernameVal, ok := args["username"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: username"), nil
		}
		username, ok := usernameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: username"), nil
		}
		variable_uuidVal, ok := args["variable_uuid"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: variable_uuid"), nil
		}
		variable_uuid, ok := variable_uuidVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: variable_uuid"), nil
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
		url := fmt.Sprintf("%s/teams/%s/pipelines_config/variables/%s", cfg.BaseURL, username, variable_uuid)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
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

func CreateUpdatepipelinevariableforteamTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_teams_username_pipelines_config_variables_variable_uuid",
		mcp.WithDescription("Update a variable for a team"),
		mcp.WithString("username", mcp.Required(), mcp.Description("The account.")),
		mcp.WithString("variable_uuid", mcp.Required(), mcp.Description("The UUID of the variable.")),
		mcp.WithString("type", mcp.Description("")),
		mcp.WithString("key", mcp.Description("Input parameter: The unique name of the variable.")),
		mcp.WithBoolean("secured", mcp.Description("Input parameter: If true, this variable will be treated as secured. The value will never be exposed in the logs or the REST API.")),
		mcp.WithString("uuid", mcp.Description("Input parameter: The UUID identifying the variable.")),
		mcp.WithString("value", mcp.Description("Input parameter: The value of the variable. If the variable is secured, this will be empty.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdatepipelinevariableforteamHandler(cfg),
	}
}
