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

func DeleteannotationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		reportIdVal, ok := args["reportId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: reportId"), nil
		}
		reportId, ok := reportIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: reportId"), nil
		}
		annotationIdVal, ok := args["annotationId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: annotationId"), nil
		}
		annotationId, ok := annotationIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: annotationId"), nil
		}
		url := fmt.Sprintf("%s/repositories/%s/%s/commit/%s/reports/%s/annotations/%s", cfg.BaseURL, workspace, repo_slug, commit, reportId, annotationId)
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

func CreateDeleteannotationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_repositories_workspace_repo_slug_commit_commit_reports_reportId_annotations_annotationId",
		mcp.WithDescription("Delete an annotation"),
		mcp.WithString("workspace", mcp.Required(), mcp.Description("This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.")),
		mcp.WithString("repo_slug", mcp.Required(), mcp.Description("The repository.")),
		mcp.WithString("commit", mcp.Required(), mcp.Description("The commit the annotation belongs to.")),
		mcp.WithString("reportId", mcp.Required(), mcp.Description("Either the uuid or external-id of the annotation.")),
		mcp.WithString("annotationId", mcp.Required(), mcp.Description("Either the uuid or external-id of the annotation.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeleteannotationHandler(cfg),
	}
}
