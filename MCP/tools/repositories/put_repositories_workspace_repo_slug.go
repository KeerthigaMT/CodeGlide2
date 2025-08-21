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

func Put_repositories_workspace_repo_slugHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Repository
		
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
		url := fmt.Sprintf("%s/repositories/%s/%s", cfg.BaseURL)
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
		var result models.Repository
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

func CreatePut_repositories_workspace_repo_slugTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_repositories_workspace_repo_slug",
		mcp.WithDescription("Update a repository"),
		mcp.WithString("type", mcp.Description("")),
		mcp.WithBoolean("has_issues", mcp.Description("")),
		mcp.WithBoolean("is_private", mcp.Description("")),
		mcp.WithObject("links", mcp.Description("")),
		mcp.WithBoolean("has_wiki", mcp.Description("")),
		mcp.WithString("description", mcp.Description("")),
		mcp.WithString("scm", mcp.Description("")),
		mcp.WithObject("mainbranch", mcp.Description("")),
		mcp.WithString("updated_on", mcp.Description("")),
		mcp.WithNumber("size", mcp.Description("")),
		mcp.WithString("language", mcp.Description("")),
		mcp.WithString("created_on", mcp.Description("")),
		mcp.WithString("fork_policy", mcp.Description("Input parameter: \nControls the rules for forking this repository.\n\n* **allow_forks**: unrestricted forking\n* **no_public_forks**: restrict forking to private forks (forks cannot\n  be made public later)\n* **no_forks**: deny all forking\n")),
		mcp.WithObject("parent", mcp.Description("")),
		mcp.WithString("uuid", mcp.Description("Input parameter: The repository's immutable id. This can be used as a substitute for the slug segment in URLs. Doing this guarantees your URLs will survive renaming of the repository by its owner, or even transfer of the repository to a different user.")),
		mcp.WithString("name", mcp.Description("")),
		mcp.WithObject("owner", mcp.Description("")),
		mcp.WithObject("project", mcp.Description("")),
		mcp.WithString("full_name", mcp.Description("Input parameter: The concatenation of the repository owner's username and the slugified name, e.g. \"evzijst/interruptingcow\". This is the same string used in Bitbucket URLs.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Put_repositories_workspace_repo_slugHandler(cfg),
	}
}
