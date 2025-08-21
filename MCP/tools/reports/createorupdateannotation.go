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

func CreateorupdateannotationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody models.Reportannotation
		
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
		url := fmt.Sprintf("%s/repositories/%s/%s/commit/%s/reports/%s/annotations/%s", cfg.BaseURL, workspace, repo_slug, commit, reportId, annotationId)
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
		var result models.Reportannotation
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

func CreateCreateorupdateannotationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_repositories_workspace_repo_slug_commit_commit_reports_reportId_annotations_annotationId",
		mcp.WithDescription("Create or update an annotation"),
		mcp.WithString("workspace", mcp.Required(), mcp.Description("This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.")),
		mcp.WithString("repo_slug", mcp.Required(), mcp.Description("The repository.")),
		mcp.WithString("commit", mcp.Required(), mcp.Description("The commit the report belongs to.")),
		mcp.WithString("reportId", mcp.Required(), mcp.Description("Either the uuid or external-id of the report.")),
		mcp.WithString("annotationId", mcp.Required(), mcp.Description("Either the uuid or external-id of the annotation.")),
		mcp.WithString("type", mcp.Description("")),
		mcp.WithString("uuid", mcp.Description("Input parameter: The UUID that can be used to identify the annotation.")),
		mcp.WithString("created_on", mcp.Description("Input parameter: The timestamp when the report was created.")),
		mcp.WithString("details", mcp.Description("Input parameter: The details to show to users when clicking on the annotation.")),
		mcp.WithString("link", mcp.Description("Input parameter: A URL linking to the annotation in an external tool.")),
		mcp.WithString("path", mcp.Description("Input parameter: The path of the file on which this annotation should be placed. This is the path of the file relative to the git repository. If no path is provided, then it will appear in the overview modal on all pull requests where the tip of the branch is the given commit, regardless of which files were modified.")),
		mcp.WithString("result", mcp.Description("Input parameter: The state of the report. May be set to PENDING and later updated.")),
		mcp.WithString("summary", mcp.Description("Input parameter: The message to display to users.")),
		mcp.WithString("severity", mcp.Description("Input parameter: The severity of the annotation.")),
		mcp.WithString("updated_on", mcp.Description("Input parameter: The timestamp when the report was updated.")),
		mcp.WithString("annotation_type", mcp.Description("Input parameter: The type of the report.")),
		mcp.WithString("external_id", mcp.Description("Input parameter: ID of the annotation provided by the annotation creator. It can be used to identify the annotation as an alternative to it's generated uuid. It is not used by Bitbucket, but only by the annotation creator for updating or deleting this specific annotation. Needs to be unique.")),
		mcp.WithNumber("line", mcp.Description("Input parameter: The line number that the annotation should belong to. If no line number is provided, then it will default to 0 and in a pull request it will appear at the top of the file specified by the path field.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreateorupdateannotationHandler(cfg),
	}
}
