package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bitbucket-api/mcp-server/config"
	"github.com/bitbucket-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func SearchaccountHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["search_query"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("search_query=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		if val, ok := args["pagelen"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pagelen=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/users/%s/search/code%s", cfg.BaseURL, selected_user, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.Searchresultpage
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

func CreateSearchaccountTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_users_selected_user_search_code",
		mcp.WithDescription("Search for code in a user's repositories"),
		mcp.WithString("selected_user", mcp.Required(), mcp.Description("Either the UUID of the account surrounded by curly-braces, for example `{account UUID}`, OR an Atlassian Account ID.")),
		mcp.WithString("search_query", mcp.Required(), mcp.Description("The search query")),
		mcp.WithNumber("page", mcp.Description("Which page of the search results to retrieve")),
		mcp.WithNumber("pagelen", mcp.Description("How many search results to retrieve per page")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    SearchaccountHandler(cfg),
	}
}
