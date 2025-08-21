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

func Post_repositories_workspace_repo_slug_pullrequestsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Pullrequest
		
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
		url := fmt.Sprintf("%s/repositories/%s/%s/pullrequests", cfg.BaseURL)
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
		var result models.Pullrequest
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

func CreatePost_repositories_workspace_repo_slug_pullrequestsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_repositories_workspace_repo_slug_pullrequests",
		mcp.WithDescription("Create a pull request"),
		mcp.WithString("type", mcp.Description("")),
		mcp.WithObject("summary", mcp.Description("")),
		mcp.WithObject("merge_commit", mcp.Description("")),
		mcp.WithObject("rendered", mcp.Description("Input parameter: User provided pull request text, interpreted in a markup language and rendered in HTML")),
		mcp.WithString("title", mcp.Description("Input parameter: Title of the pull request.")),
		mcp.WithNumber("comment_count", mcp.Description("Input parameter: The number of comments for a specific pull request.")),
		mcp.WithString("updated_on", mcp.Description("Input parameter: The ISO8601 timestamp the request was last updated.")),
		mcp.WithArray("participants", mcp.Description("Input parameter:         The list of users that are collaborating on this pull request.\n        Collaborators are user that:\n\n        * are added to the pull request as a reviewer (part of the reviewers\n          list)\n        * are not explicit reviewers, but have commented on the pull request\n        * are not explicit reviewers, but have approved the pull request\n\n        Each user is wrapped in an object that indicates the user's role and\n        whether they have approved the pull request. For performance reasons,\n        the API only returns this list when an API requests a pull request by\n        id.\n        ")),
		mcp.WithString("reason", mcp.Description("Input parameter: Explains why a pull request was declined. This field is only applicable to pull requests in rejected state.")),
		mcp.WithNumber("id", mcp.Description("Input parameter: The pull request's unique ID. Note that pull request IDs are only unique within their associated repository.")),
		mcp.WithString("state", mcp.Description("Input parameter: The pull request's current status.")),
		mcp.WithObject("closed_by", mcp.Description("")),
		mcp.WithBoolean("close_source_branch", mcp.Description("Input parameter: A boolean flag indicating if merging the pull request closes the source branch.")),
		mcp.WithObject("author", mcp.Description("")),
		mcp.WithString("created_on", mcp.Description("Input parameter: The ISO8601 timestamp the request was created.")),
		mcp.WithObject("source", mcp.Description("")),
		mcp.WithObject("destination", mcp.Description("")),
		mcp.WithArray("reviewers", mcp.Description("Input parameter: The list of users that were added as reviewers on this pull request when it was created. For performance reasons, the API only includes this list on a pull request's `self` URL.")),
		mcp.WithNumber("task_count", mcp.Description("Input parameter: The number of open tasks for a specific pull request.")),
		mcp.WithObject("links", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_repositories_workspace_repo_slug_pullrequestsHandler(cfg),
	}
}
