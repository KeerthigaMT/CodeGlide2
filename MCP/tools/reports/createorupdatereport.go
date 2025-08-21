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

func CreateorupdatereportHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody models.Report
		
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
		url := fmt.Sprintf("%s/repositories/%s/%s/commit/%s/reports/%s", cfg.BaseURL, workspace, repo_slug, commit, reportId)
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
		var result models.Report
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

func CreateCreateorupdatereportTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_repositories_workspace_repo_slug_commit_commit_reports_reportId",
		mcp.WithDescription("Create or update a report"),
		mcp.WithString("workspace", mcp.Required(), mcp.Description("This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example `{workspace UUID}`.")),
		mcp.WithString("repo_slug", mcp.Required(), mcp.Description("The repository.")),
		mcp.WithString("commit", mcp.Required(), mcp.Description("The commit the report belongs to.")),
		mcp.WithString("reportId", mcp.Required(), mcp.Description("Either the uuid or external-id of the report.")),
		mcp.WithString("type", mcp.Description("")),
		mcp.WithString("reporter", mcp.Description("Input parameter: A string to describe the tool or company who created the report.")),
		mcp.WithString("details", mcp.Description("Input parameter: A string to describe the purpose of the report.")),
		mcp.WithBoolean("remote_link_enabled", mcp.Description("Input parameter: If enabled, a remote link is created in Jira for the issue associated with the commit the report belongs to.")),
		mcp.WithString("report_type", mcp.Description("Input parameter: The type of the report.")),
		mcp.WithString("result", mcp.Description("Input parameter: The state of the report. May be set to PENDING and later updated.")),
		mcp.WithString("uuid", mcp.Description("Input parameter: The UUID that can be used to identify the report.")),
		mcp.WithString("external_id", mcp.Description("Input parameter: ID of the report provided by the report creator. It can be used to identify the report as an alternative to it's generated uuid. It is not used by Bitbucket, but only by the report creator for updating or deleting this specific report. Needs to be unique.")),
		mcp.WithString("title", mcp.Description("Input parameter: The title of the report.")),
		mcp.WithString("updated_on", mcp.Description("Input parameter: The timestamp when the report was updated.")),
		mcp.WithString("created_on", mcp.Description("Input parameter: The timestamp when the report was created.")),
		mcp.WithArray("data", mcp.Description("Input parameter: An array of data fields to display information on the report. Maximum 10.")),
		mcp.WithString("link", mcp.Description("Input parameter: A URL linking to the results of the report in an external tool.")),
		mcp.WithString("logo_url", mcp.Description("Input parameter: A URL to the report logo. If none is provided, the default insights logo will be used.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreateorupdatereportHandler(cfg),
	}
}
