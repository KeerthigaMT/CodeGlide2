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

func UpdatecommithostedpropertyvalueHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		commitVal, ok := args["commit"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: commit"), nil
		}
		commit, ok := commitVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: commit"), nil
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
		// Create properly typed request body using the generated schema
		var requestBody models.Applicationproperty
		
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
		url := fmt.Sprintf("%s/repositories/%s/%s/commit/%s/properties/%s/%s", cfg.BaseURL, workspace, repo_slug, commit, app_key, property_name)
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

func CreateUpdatecommithostedpropertyvalueTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_repositories_workspace_repo_slug_commit_commit_properties_app_key_property_name",
		mcp.WithDescription("Update a commit application property"),
		mcp.WithString("workspace", mcp.Required(), mcp.Description("The repository container; either the workspace slug or the UUID in curly braces.")),
		mcp.WithString("repo_slug", mcp.Required(), mcp.Description("The repository.")),
		mcp.WithString("commit", mcp.Required(), mcp.Description("The commit.")),
		mcp.WithString("app_key", mcp.Required(), mcp.Description("The key of the Connect app.")),
		mcp.WithString("property_name", mcp.Required(), mcp.Description("The name of the property.")),
		mcp.WithArray("_attributes", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdatecommithostedpropertyvalueHandler(cfg),
	}
}
