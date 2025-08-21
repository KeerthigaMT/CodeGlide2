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

func Put_repositories_workspace_repo_slug_commit_commit_statuses_build_keyHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Commitstatus
		
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
		url := fmt.Sprintf("%s/repositories/%s/%s/commit/%s/statuses/build/%s", cfg.BaseURL)
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
		var result models.Commitstatus
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

func CreatePut_repositories_workspace_repo_slug_commit_commit_statuses_build_keyTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_repositories_workspace_repo_slug_commit_commit_statuses_build_key",
		mcp.WithDescription("Update a build status for a commit"),
		mcp.WithString("type", mcp.Description("")),
		mcp.WithString("state", mcp.Description("Input parameter: Provides some indication of the status of this commit")),
		mcp.WithString("description", mcp.Description("Input parameter: A description of the build (e.g. \"Unit tests in Bamboo\")")),
		mcp.WithString("key", mcp.Description("Input parameter: An identifier for the status that's unique to\n        its type (current \"build\" is the only supported type) and the vendor,\n        e.g. BB-DEPLOY")),
		mcp.WithString("refname", mcp.Description("Input parameter: \nThe name of the ref that pointed to this commit at the time the status\nobject was created. Note that this the ref may since have moved off of\nthe commit. This optional field can be useful for build systems whose\nbuild triggers and configuration are branch-dependent (e.g. a Pipeline\nbuild).\nIt is legitimate for this field to not be set, or even apply (e.g. a\nstatic linting job).")),
		mcp.WithString("url", mcp.Description("Input parameter: A URL linking back to the vendor or build system, for providing more information about whatever process produced this status. Accepts context variables `repository` and `commit` that Bitbucket will evaluate at runtime whenever at runtime. For example, one could use https://foo.com/builds/{repository.full_name} which Bitbucket will turn into https://foo.com/builds/foo/bar at render time.")),
		mcp.WithObject("links", mcp.Description("")),
		mcp.WithString("name", mcp.Description("Input parameter: An identifier for the build itself, e.g. BB-DEPLOY-1")),
		mcp.WithString("updated_on", mcp.Description("")),
		mcp.WithString("uuid", mcp.Description("Input parameter: The commit status' id.")),
		mcp.WithString("created_on", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Put_repositories_workspace_repo_slug_commit_commit_statuses_build_keyHandler(cfg),
	}
}
