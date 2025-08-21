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

func Post_repositories_workspace_repo_slug_srcHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["message"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("message=%v", val))
		}
		if val, ok := args["author"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("author=%v", val))
		}
		if val, ok := args["parents"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("parents=%v", val))
		}
		if val, ok := args["files"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("files=%v", val))
		}
		if val, ok := args["branch"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("branch=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/repositories/%s/%s/src%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, nil)
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

func CreatePost_repositories_workspace_repo_slug_srcTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_repositories_workspace_repo_slug_src",
		mcp.WithDescription("Create a commit by uploading a file"),
		mcp.WithString("message", mcp.Description("The commit message. When omitted, Bitbucket uses a canned string.")),
		mcp.WithString("author", mcp.Description("\nThe raw string to be used as the new commit's author.\nThis string follows the format\n`Erik van Zijst <evzijst@atlassian.com>`.\n\nWhen omitted, Bitbucket uses the authenticated user's\nfull/display name and primary email address. Commits cannot\nbe created anonymously.")),
		mcp.WithString("parents", mcp.Description("\nA comma-separated list of SHA1s of the commits that should\nbe the parents of the newly created commit.\n\nWhen omitted, the new commit will inherit from and become\na child of the main branch's tip/HEAD commit.\n\nWhen more than one SHA1 is provided, the first SHA1\nidentifies the commit from which the content will be\ninherited.\".")),
		mcp.WithString("files", mcp.Description("\nOptional field that declares the files that the request is\nmanipulating. When adding a new file to a repo, or when\noverwriting an existing file, the client can just upload\nthe full contents of the file in a normal form field and\nthe use of this `files` meta data field is redundant.\nHowever, when the `files` field contains a file path that\ndoes not have a corresponding, identically-named form\nfield, then Bitbucket interprets that as the client wanting\nto replace the named file with the null set and the file is\ndeleted instead.\n\nPaths in the repo that are referenced in neither files nor\nan individual file field, remain unchanged and carry over\nfrom the parent to the new commit.\n\nThis API does not support renaming as an explicit feature.\nTo rename a file, simply delete it and recreate it under\nthe new name in the same commit.\n")),
		mcp.WithString("branch", mcp.Description("\nThe name of the branch that the new commit should be\ncreated on. When omitted, the commit will be created on top\nof the main branch and will become the main branch's new\nhead.\n\nWhen a branch name is provided that already exists in the\nrepo, then the commit will be created on top of that\nbranch. In this case, *if* a parent SHA1 was also provided,\nthen it is asserted that the parent is the branch's\ntip/HEAD at the time the request is made. When this is not\nthe case, a 409 is returned.\n\nWhen a new branch name is specified (that does not already\nexist in the repo), and no parent SHA1s are provided, then\nthe new commit will inherit from the current main branch's\ntip/HEAD commit, but not advance the main branch. The new\ncommit will be the new branch. When the request *also*\nspecifies a parent SHA1, then the new commit and branch\nare created directly on top of the parent commit,\nregardless of the state of the main branch.\n\nWhen a branch name is not specified, but a parent SHA1 is\nprovided, then Bitbucket asserts that it represents the\nmain branch's current HEAD/tip, or a 409 is returned.\n\nWhen a branch name is not specified and the repo is empty,\nthe new commit will become the repo's root commit and will\nbe on the main branch.\n\nWhen a branch name is specified and the repo is empty, the\nnew commit will become the repo's root commit and also\ndefine the repo's main branch going forward.\n\nThis API cannot be used to create additional root commits\nin non-empty repos.\n\nThe branch field cannot be repeated.\n\nAs a side effect, this API can be used to create a new\nbranch without modifying any files, by specifying a new\nbranch name in this field, together with `parents`, but\nomitting the `files` fields, while not sending any files.\nThis will create a new commit and branch with the same\ncontents as the first parent. The diff of this commit\nagainst its first parent will be empty.\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_repositories_workspace_repo_slug_srcHandler(cfg),
	}
}
