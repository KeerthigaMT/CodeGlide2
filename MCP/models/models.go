package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Ref represents the Ref schema from the OpenAPI specification
type Ref struct {
	Links map[string]interface{} `json:"links,omitempty"`
	Name string `json:"name,omitempty"` // The name of the ref.
	Target Commit `json:"target,omitempty"`
	TypeField string `json:"type"`
}

// Pullrequestcomment represents the Pullrequestcomment schema from the OpenAPI specification
type Pullrequestcomment struct {
	TypeField string `json:"type"`
	Deleted bool `json:"deleted,omitempty"`
	Id int `json:"id,omitempty"`
	Inline map[string]interface{} `json:"inline,omitempty"`
	Content map[string]interface{} `json:"content,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Updated_on string `json:"updated_on,omitempty"`
	User Account `json:"user,omitempty"`
	Parent Comment `json:"parent,omitempty"`
	Created_on string `json:"created_on,omitempty"`
	Pullrequest Pullrequest `json:"pullrequest,omitempty"`
}

// Pipelinestepstateinprogress represents the Pipelinestepstateinprogress schema from the OpenAPI specification
type Pipelinestepstateinprogress struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of pipeline step state (IN_PROGRESS).
}

// Pipelinevariable represents the Pipelinevariable schema from the OpenAPI specification
type Pipelinevariable struct {
	TypeField string `json:"type"`
	Key string `json:"key,omitempty"` // The unique name of the variable.
	Secured bool `json:"secured,omitempty"` // If true, this variable will be treated as secured. The value will never be exposed in the logs or the REST API.
	Uuid string `json:"uuid,omitempty"` // The UUID identifying the variable.
	Value string `json:"value,omitempty"` // The value of the variable. If the variable is secured, this will be empty.
}

// Paginateddeployments represents the Paginateddeployments schema from the OpenAPI specification
type Paginateddeployments struct {
	Values []Deployment `json:"values,omitempty"` // The values of the current page.
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
}

// Issuechange represents the Issuechange schema from the OpenAPI specification
type Issuechange struct {
	Changes map[string]interface{} `json:"changes,omitempty"`
	Created_on string `json:"created_on,omitempty"`
	Issue Issue `json:"issue,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Message map[string]interface{} `json:"message,omitempty"`
	Name string `json:"name,omitempty"`
	TypeField string `json:"type"`
	User Account `json:"user,omitempty"`
}

// Projectdeploykey represents the Projectdeploykey schema from the OpenAPI specification
type Projectdeploykey struct {
	TypeField string `json:"type"`
	Comment string `json:"comment,omitempty"` // The comment parsed from the deploy key (if present)
	Created_by Account `json:"created_by,omitempty"`
	Key string `json:"key,omitempty"` // The deploy key value.
	Label string `json:"label,omitempty"` // The user-defined label for the deploy key
	Last_used string `json:"last_used,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Project Project `json:"project,omitempty"`
	Added_on string `json:"added_on,omitempty"`
}

// Paginatedpullrequests represents the Paginatedpullrequests schema from the OpenAPI specification
type Paginatedpullrequests struct {
	Values []Pullrequest `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
}

// Paginatedrepositorygrouppermissions represents the Paginatedrepositorygrouppermissions schema from the OpenAPI specification
type Paginatedrepositorygrouppermissions struct {
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Repositorygrouppermission `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
}

// Paginatedprojectdeploykeys represents the Paginatedprojectdeploykeys schema from the OpenAPI specification
type Paginatedprojectdeploykeys struct {
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Projectdeploykey `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
}

// Pipelinestatecompletedsuccessful represents the Pipelinestatecompletedsuccessful schema from the OpenAPI specification
type Pipelinestatecompletedsuccessful struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the successful result (SUCCESSFUL).
}

// Pipelinestatecompletedexpired represents the Pipelinestatecompletedexpired schema from the OpenAPI specification
type Pipelinestatecompletedexpired struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the stopped result (EXPIRED).
}

// Commit represents the Commit schema from the OpenAPI specification
type Commit struct {
	TypeField string `json:"type"`
	Hash string `json:"hash,omitempty"`
	Message string `json:"message,omitempty"`
	Parents []Basecommit `json:"parents,omitempty"`
	Summary map[string]interface{} `json:"summary,omitempty"`
	Author Author `json:"author,omitempty"`
	Date string `json:"date,omitempty"`
	Participants []Participant `json:"participants,omitempty"`
	Repository Repository `json:"repository,omitempty"`
}

// Pipelinesconfig represents the Pipelinesconfig schema from the OpenAPI specification
type Pipelinesconfig struct {
	TypeField string `json:"type"`
	Enabled bool `json:"enabled,omitempty"` // Whether Pipelines is enabled for the repository.
	Repository Repository `json:"repository,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Links map[string]interface{} `json:"links,omitempty"`
	Name string `json:"name,omitempty"` // The name of the ref.
	Target Commit `json:"target,omitempty"`
	TypeField string `json:"type"`
	Tagger Author `json:"tagger,omitempty"`
	Date string `json:"date,omitempty"` // The date that the tag was created, if available
	Message string `json:"message,omitempty"` // The message associated with the tag, if available.
}

// Pipelinescheduleexecutionexecuted represents the Pipelinescheduleexecutionexecuted schema from the OpenAPI specification
type Pipelinescheduleexecutionexecuted struct {
	TypeField string `json:"type"`
	Pipeline Pipeline `json:"pipeline,omitempty"`
}

// Pipelinesteperror represents the Pipelinesteperror schema from the OpenAPI specification
type Pipelinesteperror struct {
	TypeField string `json:"type"`
	Key string `json:"key,omitempty"` // The error key.
	Message string `json:"message,omitempty"` // The error message.
}

// Pipelinestepstateready represents the Pipelinestepstateready schema from the OpenAPI specification
type Pipelinestepstateready struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of pipeline step state (READY).
}

// Pipelineerror represents the Pipelineerror schema from the OpenAPI specification
type Pipelineerror struct {
	TypeField string `json:"type"`
	Key string `json:"key,omitempty"` // The error key.
	Message string `json:"message,omitempty"` // The error message.
}

// Deploymentstate represents the Deploymentstate schema from the OpenAPI specification
type Deploymentstate struct {
	TypeField string `json:"type"`
}

// Deploymentsddevdeploymentenvironment represents the Deploymentsddevdeploymentenvironment schema from the OpenAPI specification
type Deploymentsddevdeploymentenvironment struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the environment.
	Uuid string `json:"uuid,omitempty"` // The UUID identifying the environment.
}

// Deploymentstatecompletedstatussuccessful represents the Deploymentstatecompletedstatussuccessful schema from the OpenAPI specification
type Deploymentstatecompletedstatussuccessful struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the completed deployment status (SUCCESSFUL).
}

// Issueattachment represents the Issueattachment schema from the OpenAPI specification
type Issueattachment struct {
	TypeField string `json:"type"`
	Links map[string]interface{} `json:"links,omitempty"`
	Name string `json:"name,omitempty"`
}

// Exportoptions represents the Exportoptions schema from the OpenAPI specification
type Exportoptions struct {
	Send_email bool `json:"send_email,omitempty"`
	TypeField string `json:"type"`
	Include_attachments bool `json:"include_attachments,omitempty"`
	Project_key string `json:"project_key,omitempty"`
	Project_name string `json:"project_name,omitempty"`
}

// Paginatedlogentries represents the Paginatedlogentries schema from the OpenAPI specification
type Paginatedlogentries struct {
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Issuechange `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
}

// Snippetcommit represents the Snippetcommit schema from the OpenAPI specification
type Snippetcommit struct {
	TypeField string `json:"type"`
	Summary map[string]interface{} `json:"summary,omitempty"`
	Author Author `json:"author,omitempty"`
	Date string `json:"date,omitempty"`
	Hash string `json:"hash,omitempty"`
	Message string `json:"message,omitempty"`
	Parents []Basecommit `json:"parents,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Snippet Snippet `json:"snippet,omitempty"`
}

// Pipelinetriggerpush represents the Pipelinetriggerpush schema from the OpenAPI specification
type Pipelinetriggerpush struct {
	TypeField string `json:"type"`
}

// Pipelinecommand represents the Pipelinecommand schema from the OpenAPI specification
type Pipelinecommand struct {
	Name string `json:"name,omitempty"` // The name of the command.
	Command string `json:"command,omitempty"` // The executable command.
}

// Author represents the Author schema from the OpenAPI specification
type Author struct {
	TypeField string `json:"type"`
	Raw string `json:"raw,omitempty"` // The raw author value from the repository. This may be the only value available if the author does not match a user in Bitbucket.
	User Account `json:"user,omitempty"`
}

// Team represents the Team schema from the OpenAPI specification
type Team struct {
	TypeField string `json:"type"`
	Username string `json:"username,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Created_on string `json:"created_on,omitempty"`
	Display_name string `json:"display_name,omitempty"`
	Links Accountlinks `json:"links,omitempty"` // Links related to an Account.
	Links Teamlinks `json:"links,omitempty"`
}

// Branchingmodelsettings represents the Branchingmodelsettings schema from the OpenAPI specification
type Branchingmodelsettings struct {
	TypeField string `json:"type"`
	Branch_types []map[string]interface{} `json:"branch_types,omitempty"`
	Development map[string]interface{} `json:"development,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Production map[string]interface{} `json:"production,omitempty"`
}

// Pipelinestateinprogressstage represents the Pipelinestateinprogressstage schema from the OpenAPI specification
type Pipelinestateinprogressstage struct {
	TypeField string `json:"type"`
}

// Paginateddiffstats represents the Paginateddiffstats schema from the OpenAPI specification
type Paginateddiffstats struct {
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 500 with 5000 being the maximum allowed value.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Diffstat `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
}

// Repositoryuserpermission represents the Repositoryuserpermission schema from the OpenAPI specification
type Repositoryuserpermission struct {
	User User `json:"user,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Permission string `json:"permission,omitempty"`
	Repository Repository `json:"repository,omitempty"`
	TypeField string `json:"type"`
}

// Pipelinestepstatepending represents the Pipelinestepstatepending schema from the OpenAPI specification
type Pipelinestepstatepending struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of pipeline step state (PENDING).
}

// Paginatedpipelineknownhosts represents the Paginatedpipelineknownhosts schema from the OpenAPI specification
type Paginatedpipelineknownhosts struct {
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Pipelineknownhost `json:"values,omitempty"` // The values of the current page.
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
}

// Workspacemembership represents the Workspacemembership schema from the OpenAPI specification
type Workspacemembership struct {
	TypeField string `json:"type"`
	Workspace Workspace `json:"workspace,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	User Account `json:"user,omitempty"`
}

// Repositorygrouppermission represents the Repositorygrouppermission schema from the OpenAPI specification
type Repositorygrouppermission struct {
	Group Group `json:"group,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Permission string `json:"permission,omitempty"`
	Repository Repository `json:"repository,omitempty"`
	TypeField string `json:"type"`
}

// Paginatedchangeset represents the Paginatedchangeset schema from the OpenAPI specification
type Paginatedchangeset struct {
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Basecommit `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
}

// Paginateddeploymentvariable represents the Paginateddeploymentvariable schema from the OpenAPI specification
type Paginateddeploymentvariable struct {
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Deploymentvariable `json:"values,omitempty"` // The values of the current page.
}

// Snippet represents the Snippet schema from the OpenAPI specification
type Snippet struct {
	TypeField string `json:"type"`
	Created_on string `json:"created_on,omitempty"`
	Creator Account `json:"creator,omitempty"`
	Id int `json:"id,omitempty"`
	Is_private bool `json:"is_private,omitempty"`
	Owner Account `json:"owner,omitempty"`
	Scm string `json:"scm,omitempty"` // The DVCS used to store the snippet.
	Title string `json:"title,omitempty"`
	Updated_on string `json:"updated_on,omitempty"`
}

// Page represents the Page schema from the OpenAPI specification
type Page struct {
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
}

// Participant represents the Participant schema from the OpenAPI specification
type Participant struct {
	TypeField string `json:"type"`
	State string `json:"state,omitempty"`
	User Account `json:"user,omitempty"`
	Approved bool `json:"approved,omitempty"`
	Participated_on string `json:"participated_on,omitempty"` // The ISO8601 timestamp of the participant's action. For approvers, this is the time of their approval. For commenters and pull request reviewers who are not approvers, this is the time they last commented, or null if they have not commented.
	Role string `json:"role,omitempty"`
}

// Paginatedworkspaces represents the Paginatedworkspaces schema from the OpenAPI specification
type Paginatedworkspaces struct {
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Workspace `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
}

// Pipelinereftarget represents the Pipelinereftarget schema from the OpenAPI specification
type Pipelinereftarget struct {
	TypeField string `json:"type"`
	Selector Pipelineselector `json:"selector,omitempty"`
	Commit Commit `json:"commit,omitempty"`
	Ref_name string `json:"ref_name,omitempty"` // The name of the reference.
	Ref_type string `json:"ref_type,omitempty"` // The type of reference (branch/tag).
}

// Searchsegment represents the Searchsegment schema from the OpenAPI specification
type Searchsegment struct {
	Text string `json:"text,omitempty"`
	Match bool `json:"match,omitempty"`
}

// Deploymentvariable represents the Deploymentvariable schema from the OpenAPI specification
type Deploymentvariable struct {
	TypeField string `json:"type"`
	Value string `json:"value,omitempty"` // The value of the variable. If the variable is secured, this will be empty.
	Key string `json:"key,omitempty"` // The unique name of the variable.
	Secured bool `json:"secured,omitempty"` // If true, this variable will be treated as secured. The value will never be exposed in the logs or the REST API.
	Uuid string `json:"uuid,omitempty"` // The UUID identifying the variable.
}

// Repositoryinheritancestate represents the Repositoryinheritancestate schema from the OpenAPI specification
type Repositoryinheritancestate struct {
	Override_settings map[string]interface{} `json:"override_settings,omitempty"`
	TypeField string `json:"type"`
}

// Paginatedpipelines represents the Paginatedpipelines schema from the OpenAPI specification
type Paginatedpipelines struct {
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Pipeline `json:"values,omitempty"` // The values of the current page.
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
}

// Pipelineschedule represents the Pipelineschedule schema from the OpenAPI specification
type Pipelineschedule struct {
	TypeField string `json:"type"`
	Created_on string `json:"created_on,omitempty"` // The timestamp when the schedule was created.
	Cron_pattern string `json:"cron_pattern,omitempty"` // The cron expression with second precision (7 fields) that the schedule applies. For example, for expression: 0 0 12 * * ? *, will execute at 12pm UTC every day.
	Enabled bool `json:"enabled,omitempty"` // Whether the schedule is enabled.
	Target Pipelinereftarget `json:"target,omitempty"`
	Updated_on string `json:"updated_on,omitempty"` // The timestamp when the schedule was updated.
	Uuid string `json:"uuid,omitempty"` // The UUID identifying the schedule.
}

// Paginatedworkspacememberships represents the Paginatedworkspacememberships schema from the OpenAPI specification
type Paginatedworkspacememberships struct {
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Workspacemembership `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
}

// Group represents the Group schema from the OpenAPI specification
type Group struct {
	TypeField string `json:"type"`
	Slug string `json:"slug,omitempty"` // The "sluggified" version of the group's name. This contains only ASCII characters and can therefore be slightly different than the name
	Workspace Workspace `json:"workspace,omitempty"`
	Full_slug string `json:"full_slug,omitempty"` // The concatenation of the workspace's slug and the group's slug, separated with a colon (e.g. `acme:developers`)
	Links map[string]interface{} `json:"links,omitempty"`
	Name string `json:"name,omitempty"`
	Owner Account `json:"owner,omitempty"`
}

// Pipelinestepstatecompletednotrun represents the Pipelinestepstatecompletednotrun schema from the OpenAPI specification
type Pipelinestepstatecompletednotrun struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the result (NOT_RUN)
}

// User represents the User schema from the OpenAPI specification
type User struct {
	TypeField string `json:"type"`
	Display_name string `json:"display_name,omitempty"`
	Links Accountlinks `json:"links,omitempty"` // Links related to an Account.
	Username string `json:"username,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Created_on string `json:"created_on,omitempty"`
	Account_status string `json:"account_status,omitempty"` // The status of the account. Currently the only possible value is "active", but more values may be added in the future.
	Has_2fa_enabled bool `json:"has_2fa_enabled,omitempty"`
	Is_staff bool `json:"is_staff,omitempty"`
	Links Userlinks `json:"links,omitempty"`
	Nickname string `json:"nickname,omitempty"` // Account name defined by the owner. Should be used instead of the "username" field. Note that "nickname" cannot be used in place of "username" in URLs and queries, as "nickname" is not guaranteed to be unique.
	Website string `json:"website,omitempty"`
	Account_id string `json:"account_id,omitempty"` // The user's Atlassian account ID.
}

// Paginatedwebhooksubscriptions represents the Paginatedwebhooksubscriptions schema from the OpenAPI specification
type Paginatedwebhooksubscriptions struct {
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Webhooksubscription `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
}

// Treeentry represents the Treeentry schema from the OpenAPI specification
type Treeentry struct {
	Commit Commit `json:"commit,omitempty"`
	Path string `json:"path,omitempty"` // The path in the repository
	TypeField string `json:"type"`
}

// Pipelinestepstate represents the Pipelinestepstate schema from the OpenAPI specification
type Pipelinestepstate struct {
	TypeField string `json:"type"`
}

// Deploymentstatecompletedstatusfailed represents the Deploymentstatecompletedstatusfailed schema from the OpenAPI specification
type Deploymentstatecompletedstatusfailed struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the completed deployment status (FAILED).
}

// Pullrequest represents the Pullrequest schema from the OpenAPI specification
type Pullrequest struct {
	TypeField string `json:"type"`
	Destination Pullrequestendpoint `json:"destination,omitempty"`
	Reviewers []Account `json:"reviewers,omitempty"` // The list of users that were added as reviewers on this pull request when it was created. For performance reasons, the API only includes this list on a pull request's `self` URL.
	Task_count int `json:"task_count,omitempty"` // The number of open tasks for a specific pull request.
	Links map[string]interface{} `json:"links,omitempty"`
	Summary map[string]interface{} `json:"summary,omitempty"`
	Merge_commit map[string]interface{} `json:"merge_commit,omitempty"`
	Rendered map[string]interface{} `json:"rendered,omitempty"` // User provided pull request text, interpreted in a markup language and rendered in HTML
	Title string `json:"title,omitempty"` // Title of the pull request.
	Comment_count int `json:"comment_count,omitempty"` // The number of comments for a specific pull request.
	Updated_on string `json:"updated_on,omitempty"` // The ISO8601 timestamp the request was last updated.
	Participants []Participant `json:"participants,omitempty"` // The list of users that are collaborating on this pull request. Collaborators are user that: * are added to the pull request as a reviewer (part of the reviewers list) * are not explicit reviewers, but have commented on the pull request * are not explicit reviewers, but have approved the pull request Each user is wrapped in an object that indicates the user's role and whether they have approved the pull request. For performance reasons, the API only returns this list when an API requests a pull request by id.
	Reason string `json:"reason,omitempty"` // Explains why a pull request was declined. This field is only applicable to pull requests in rejected state.
	Id int `json:"id,omitempty"` // The pull request's unique ID. Note that pull request IDs are only unique within their associated repository.
	State string `json:"state,omitempty"` // The pull request's current status.
	Closed_by Account `json:"closed_by,omitempty"`
	Close_source_branch bool `json:"close_source_branch,omitempty"` // A boolean flag indicating if merging the pull request closes the source branch.
	Author Account `json:"author,omitempty"`
	Created_on string `json:"created_on,omitempty"` // The ISO8601 timestamp the request was created.
	Source Pullrequestendpoint `json:"source,omitempty"`
}

// Pipelinestatepending represents the Pipelinestatepending schema from the OpenAPI specification
type Pipelinestatepending struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of pipeline state (PENDING).
}

// Sshkey represents the Sshkey schema from the OpenAPI specification
type Sshkey struct {
	TypeField string `json:"type"`
	Links map[string]interface{} `json:"links,omitempty"`
	Uuid string `json:"uuid,omitempty"` // The SSH key's immutable ID.
	Comment string `json:"comment,omitempty"` // The comment parsed from the SSH key (if present)
	Created_on string `json:"created_on,omitempty"`
	Key string `json:"key,omitempty"` // The SSH public key value in OpenSSH format.
	Label string `json:"label,omitempty"` // The user-defined label for the SSH key
	Last_used string `json:"last_used,omitempty"`
}

// Repository represents the Repository schema from the OpenAPI specification
type Repository struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"`
	Owner Account `json:"owner,omitempty"`
	Project Project `json:"project,omitempty"`
	Full_name string `json:"full_name,omitempty"` // The concatenation of the repository owner's username and the slugified name, e.g. "evzijst/interruptingcow". This is the same string used in Bitbucket URLs.
	Has_issues bool `json:"has_issues,omitempty"`
	Is_private bool `json:"is_private,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Has_wiki bool `json:"has_wiki,omitempty"`
	Description string `json:"description,omitempty"`
	Scm string `json:"scm,omitempty"`
	Mainbranch Branch `json:"mainbranch,omitempty"`
	Updated_on string `json:"updated_on,omitempty"`
	Size int `json:"size,omitempty"`
	Language string `json:"language,omitempty"`
	Created_on string `json:"created_on,omitempty"`
	Fork_policy string `json:"fork_policy,omitempty"` // Controls the rules for forking this repository. * **allow_forks**: unrestricted forking * **no_public_forks**: restrict forking to private forks (forks cannot be made public later) * **no_forks**: deny all forking
	Parent Repository `json:"parent,omitempty"`
	Uuid string `json:"uuid,omitempty"` // The repository's immutable id. This can be used as a substitute for the slug segment in URLs. Doing this guarantees your URLs will survive renaming of the repository by its owner, or even transfer of the repository to a different user.
}

// Paginatedissuecomments represents the Paginatedissuecomments schema from the OpenAPI specification
type Paginatedissuecomments struct {
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Issuecomment `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
}

// Pipelinecache represents the Pipelinecache schema from the OpenAPI specification
type Pipelinecache struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the cache.
	Path string `json:"path,omitempty"` // The path where the cache contents were retrieved from.
	Pipeline_uuid string `json:"pipeline_uuid,omitempty"` // The UUID of the pipeline that created the cache.
	Step_uuid string `json:"step_uuid,omitempty"` // The uuid of the step that created the cache.
	Uuid string `json:"uuid,omitempty"` // The UUID identifying the pipeline cache.
	Created_on string `json:"created_on,omitempty"` // The timestamp when the cache was created.
	File_size_bytes int `json:"file_size_bytes,omitempty"` // The size of the file containing the archive of the cache.
	Key_hash string `json:"key_hash,omitempty"` // The key hash of the cache version.
}

// Commitcomment represents the Commitcomment schema from the OpenAPI specification
type Commitcomment struct {
	TypeField string `json:"type"`
	Deleted bool `json:"deleted,omitempty"`
	Id int `json:"id,omitempty"`
	Inline map[string]interface{} `json:"inline,omitempty"`
	Content map[string]interface{} `json:"content,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Updated_on string `json:"updated_on,omitempty"`
	User Account `json:"user,omitempty"`
	Parent Comment `json:"parent,omitempty"`
	Created_on string `json:"created_on,omitempty"`
	Commit Commit `json:"commit,omitempty"`
}

// Deployment represents the Deployment schema from the OpenAPI specification
type Deployment struct {
	TypeField string `json:"type"`
	Uuid string `json:"uuid,omitempty"` // The UUID identifying the deployment.
	Environment Deploymentenvironment `json:"environment,omitempty"`
	Release Deploymentrelease `json:"release,omitempty"`
	State Deploymentstate `json:"state,omitempty"`
}

// Snippetcomment represents the Snippetcomment schema from the OpenAPI specification
type Snippetcomment struct {
	TypeField string `json:"type"`
	Snippet Snippet `json:"snippet,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
}

// Jirasite represents the Jirasite schema from the OpenAPI specification
type Jirasite struct {
	TypeField string `json:"type"`
}

// Paginatedtreeentries represents the Paginatedtreeentries schema from the OpenAPI specification
type Paginatedtreeentries struct {
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Treeentry `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
}

// Accountlinks represents the Accountlinks schema from the OpenAPI specification
type Accountlinks struct {
	Avatar Link `json:"avatar,omitempty"` // A link to a resource related to this object.
}

// Pipelinestepstatecompletedfailed represents the Pipelinestepstatecompletedfailed schema from the OpenAPI specification
type Pipelinestepstatecompletedfailed struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the result (FAILED)
}

// Paginatedsnippetcommit represents the Paginatedsnippetcommit schema from the OpenAPI specification
type Paginatedsnippetcommit struct {
	Values []Snippetcommit `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
}

// Deploymentenvironmentlock represents the Deploymentenvironmentlock schema from the OpenAPI specification
type Deploymentenvironmentlock struct {
	TypeField string `json:"type"`
	Environmentuuid string `json:"environmentUuid,omitempty"` // The UUID identifying the environment.
}

// Pullrequestendpoint represents the Pullrequestendpoint schema from the OpenAPI specification
type Pullrequestendpoint struct {
	Branch map[string]interface{} `json:"branch,omitempty"`
	Commit map[string]interface{} `json:"commit,omitempty"`
	Repository Repository `json:"repository,omitempty"`
}

// Pipelinestepstatecompletedsuccessful represents the Pipelinestepstatecompletedsuccessful schema from the OpenAPI specification
type Pipelinestepstatecompletedsuccessful struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the result (SUCCESSFUL)
}

// Deploymentsddevpaginatedenvironments represents the Deploymentsddevpaginatedenvironments schema from the OpenAPI specification
type Deploymentsddevpaginatedenvironments struct {
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Deploymentsddevdeploymentenvironment `json:"values,omitempty"` // The values of the current page.
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
}

// Pipelinestateinprogresspaused represents the Pipelinestateinprogresspaused schema from the OpenAPI specification
type Pipelinestateinprogresspaused struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the stage (PAUSED)
}

// Pipelinetarget represents the Pipelinetarget schema from the OpenAPI specification
type Pipelinetarget struct {
	TypeField string `json:"type"`
}

// Branchingmodel represents the Branchingmodel schema from the OpenAPI specification
type Branchingmodel struct {
	TypeField string `json:"type"`
	Production map[string]interface{} `json:"production,omitempty"`
	Branch_types []map[string]interface{} `json:"branch_types,omitempty"` // The active branch types.
	Development map[string]interface{} `json:"development,omitempty"`
}

// Paginatedissueattachments represents the Paginatedissueattachments schema from the OpenAPI specification
type Paginatedissueattachments struct {
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Issueattachment `json:"values,omitempty"`
}

// Pipelinestepstatecompletedstopped represents the Pipelinestepstatecompletedstopped schema from the OpenAPI specification
type Pipelinestepstatecompletedstopped struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the result (STOPPED)
}

// Deploykey represents the Deploykey schema from the OpenAPI specification
type Deploykey struct {
	TypeField string `json:"type"`
	Label string `json:"label,omitempty"` // The user-defined label for the deploy key
	Last_used string `json:"last_used,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Owner Account `json:"owner,omitempty"`
	Repository Repository `json:"repository,omitempty"`
	Added_on string `json:"added_on,omitempty"`
	Comment string `json:"comment,omitempty"` // The comment parsed from the deploy key (if present)
	Key string `json:"key,omitempty"` // The deploy key value.
}

// Deploymentstateinprogress represents the Deploymentstateinprogress schema from the OpenAPI specification
type Deploymentstateinprogress struct {
	TypeField string `json:"type"`
	Url string `json:"url,omitempty"` // Link to the deployment result.
	Deployer Account `json:"deployer,omitempty"`
	Name string `json:"name,omitempty"` // The name of deployment state (IN_PROGRESS).
	Start_date string `json:"start_date,omitempty"` // The timestamp when the deployment was started.
}

// Paginatedissues represents the Paginatedissues schema from the OpenAPI specification
type Paginatedissues struct {
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Issue `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
}

// Reportannotation represents the Reportannotation schema from the OpenAPI specification
type Reportannotation struct {
	TypeField string `json:"type"`
	Severity string `json:"severity,omitempty"` // The severity of the annotation.
	Updated_on string `json:"updated_on,omitempty"` // The timestamp when the report was updated.
	Annotation_type string `json:"annotation_type,omitempty"` // The type of the report.
	External_id string `json:"external_id,omitempty"` // ID of the annotation provided by the annotation creator. It can be used to identify the annotation as an alternative to it's generated uuid. It is not used by Bitbucket, but only by the annotation creator for updating or deleting this specific annotation. Needs to be unique.
	Line int `json:"line,omitempty"` // The line number that the annotation should belong to. If no line number is provided, then it will default to 0 and in a pull request it will appear at the top of the file specified by the path field.
	Uuid string `json:"uuid,omitempty"` // The UUID that can be used to identify the annotation.
	Created_on string `json:"created_on,omitempty"` // The timestamp when the report was created.
	Details string `json:"details,omitempty"` // The details to show to users when clicking on the annotation.
	Link string `json:"link,omitempty"` // A URL linking to the annotation in an external tool.
	Path string `json:"path,omitempty"` // The path of the file on which this annotation should be placed. This is the path of the file relative to the git repository. If no path is provided, then it will appear in the overview modal on all pull requests where the tip of the branch is the given commit, regardless of which files were modified.
	Result string `json:"result,omitempty"` // The state of the report. May be set to PENDING and later updated.
	Summary string `json:"summary,omitempty"` // The message to display to users.
}

// Deploymentsddevdeploymentenvironmentlock represents the Deploymentsddevdeploymentenvironmentlock schema from the OpenAPI specification
type Deploymentsddevdeploymentenvironmentlock struct {
	TypeField string `json:"type"`
	Environmentuuid string `json:"environmentUuid,omitempty"` // The UUID identifying the environment.
}

// Paginatedsnippetcomments represents the Paginatedsnippetcomments schema from the OpenAPI specification
type Paginatedsnippetcomments struct {
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Snippetcomment `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
}

// Paginatedsshuserkeys represents the Paginatedsshuserkeys schema from the OpenAPI specification
type Paginatedsshuserkeys struct {
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Sshaccountkey `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
}

// Paginatedprojects represents the Paginatedprojects schema from the OpenAPI specification
type Paginatedprojects struct {
	Values []Project `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
}

// Paginatedpipelinevariables represents the Paginatedpipelinevariables schema from the OpenAPI specification
type Paginatedpipelinevariables struct {
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Pipelinevariable `json:"values,omitempty"` // The values of the current page.
}

// Pipelinestateinprogress represents the Pipelinestateinprogress schema from the OpenAPI specification
type Pipelinestateinprogress struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of pipeline state (IN_PROGRESS).
	Stage Pipelinestateinprogressstage `json:"stage,omitempty"`
}

// Deploymentstatecompletedstatusstopped represents the Deploymentstatecompletedstatusstopped schema from the OpenAPI specification
type Deploymentstatecompletedstatusstopped struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the completed deployment status (STOPPED).
}

// Teamlinks represents the Teamlinks schema from the OpenAPI specification
type Teamlinks struct {
	Avatar Link `json:"avatar,omitempty"` // A link to a resource related to this object.
	Projects Link `json:"projects,omitempty"` // A link to a resource related to this object.
	Repositories Link `json:"repositories,omitempty"` // A link to a resource related to this object.
	Self Link `json:"self,omitempty"` // A link to a resource related to this object.
	Html Link `json:"html,omitempty"` // A link to a resource related to this object.
	Members Link `json:"members,omitempty"` // A link to a resource related to this object.
}

// Pipelineselector represents the Pipelineselector schema from the OpenAPI specification
type Pipelineselector struct {
	TypeField string `json:"type"`
	Pattern string `json:"pattern,omitempty"` // The name of the matching pipeline definition.
	TypeField string `json:"type,omitempty"` // The type of selector.
}

// Hookevent represents the Hookevent schema from the OpenAPI specification
type Hookevent struct {
	Category string `json:"category,omitempty"` // The category this event belongs to.
	Description string `json:"description,omitempty"` // More detailed description of the webhook event type.
	Event string `json:"event,omitempty"` // The event identifier.
	Label string `json:"label,omitempty"` // Summary of the webhook event type.
}

// Paginatedpipelinescheduleexecutions represents the Paginatedpipelinescheduleexecutions schema from the OpenAPI specification
type Paginatedpipelinescheduleexecutions struct {
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Pipelinescheduleexecution `json:"values,omitempty"` // The values of the current page.
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
}

// Pipelinescheduleexecutionerrored represents the Pipelinescheduleexecutionerrored schema from the OpenAPI specification
type Pipelinescheduleexecutionerrored struct {
	TypeField string `json:"type"`
	ErrorField Pipelineerror `json:"error,omitempty"`
}

// Version represents the Version schema from the OpenAPI specification
type Version struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"`
	Id int `json:"id,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
}

// Pipelinestate represents the Pipelinestate schema from the OpenAPI specification
type Pipelinestate struct {
	TypeField string `json:"type"`
}

// Paginatedcommitstatuses represents the Paginatedcommitstatuses schema from the OpenAPI specification
type Paginatedcommitstatuses struct {
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Commitstatus `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
}

// Pipeline represents the Pipeline schema from the OpenAPI specification
type Pipeline struct {
	TypeField string `json:"type"`
	Build_seconds_used int `json:"build_seconds_used,omitempty"` // The number of build seconds used by this pipeline.
	Completed_on string `json:"completed_on,omitempty"` // The timestamp when the Pipeline was completed. This is not set if the pipeline is still in progress.
	Target Pipelinetarget `json:"target,omitempty"`
	Created_on string `json:"created_on,omitempty"` // The timestamp when the pipeline was created.
	Creator Account `json:"creator,omitempty"`
	Repository Repository `json:"repository,omitempty"`
	State Pipelinestate `json:"state,omitempty"`
	Trigger Pipelinetrigger `json:"trigger,omitempty"`
	Build_number int `json:"build_number,omitempty"` // The build number of the pipeline.
	Uuid string `json:"uuid,omitempty"` // The UUID identifying the pipeline.
}

// Deploymentsstgwestdeploymentenvironmentlock represents the Deploymentsstgwestdeploymentenvironmentlock schema from the OpenAPI specification
type Deploymentsstgwestdeploymentenvironmentlock struct {
	TypeField string `json:"type"`
	Environmentuuid string `json:"environmentUuid,omitempty"` // The UUID identifying the environment.
}

// Issuecomment represents the Issuecomment schema from the OpenAPI specification
type Issuecomment struct {
	TypeField string `json:"type"`
	Deleted bool `json:"deleted,omitempty"`
	Id int `json:"id,omitempty"`
	Inline map[string]interface{} `json:"inline,omitempty"`
	Content map[string]interface{} `json:"content,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Updated_on string `json:"updated_on,omitempty"`
	User Account `json:"user,omitempty"`
	Parent Comment `json:"parent,omitempty"`
	Created_on string `json:"created_on,omitempty"`
	Issue Issue `json:"issue,omitempty"`
}

// Milestone represents the Milestone schema from the OpenAPI specification
type Milestone struct {
	TypeField string `json:"type"`
	Id int `json:"id,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Name string `json:"name,omitempty"`
}

// Component represents the Component schema from the OpenAPI specification
type Component struct {
	TypeField string `json:"type"`
	Id int `json:"id,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Name string `json:"name,omitempty"`
}

// Deploymentrelease represents the Deploymentrelease schema from the OpenAPI specification
type Deploymentrelease struct {
	TypeField string `json:"type"`
	Commit Commit `json:"commit,omitempty"`
	Created_on string `json:"created_on,omitempty"` // The timestamp when the release was created.
	Name string `json:"name,omitempty"` // The name of the release.
	Url string `json:"url,omitempty"` // Link to the pipeline that produced the release.
	Uuid string `json:"uuid,omitempty"` // The UUID identifying the release.
}

// Pipelinestatecompletederror represents the Pipelinestatecompletederror schema from the OpenAPI specification
type Pipelinestatecompletederror struct {
	TypeField string `json:"type"`
	ErrorField Pipelineerror `json:"error,omitempty"`
	Name string `json:"name,omitempty"` // The name of the result (ERROR)
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	TypeField string `json:"type"`
	Key string `json:"key,omitempty"` // The project's key.
	Updated_on string `json:"updated_on,omitempty"`
	Created_on string `json:"created_on,omitempty"`
	Description string `json:"description,omitempty"`
	Has_publicly_visible_repos bool `json:"has_publicly_visible_repos,omitempty"` // Indicates whether the project contains publicly visible repositories. Note that private projects cannot contain public repositories.
	Links map[string]interface{} `json:"links,omitempty"`
	Uuid string `json:"uuid,omitempty"` // The project's immutable id.
	Name string `json:"name,omitempty"` // The name of the project.
	Owner Team `json:"owner,omitempty"`
	Is_private bool `json:"is_private,omitempty"` // Indicates whether the project is publicly accessible, or whether it is private to the team and consequently only visible to team members. Note that private projects cannot contain public repositories.
}

// Pipelinestep represents the Pipelinestep schema from the OpenAPI specification
type Pipelinestep struct {
	TypeField string `json:"type"`
	Setup_commands []Pipelinecommand `json:"setup_commands,omitempty"` // The list of commands that are executed as part of the setup phase of the build. These commands are executed outside the build container.
	Started_on string `json:"started_on,omitempty"` // The timestamp when the step execution was started. This is not set when the step hasn't executed yet.
	State Pipelinestepstate `json:"state,omitempty"`
	Uuid string `json:"uuid,omitempty"` // The UUID identifying the step.
	Completed_on string `json:"completed_on,omitempty"` // The timestamp when the step execution was completed. This is not set if the step is still in progress.
	Image Pipelineimage `json:"image,omitempty"` // The definition of a Docker image that can be used for a Bitbucket Pipelines step execution context.
	Script_commands []Pipelinecommand `json:"script_commands,omitempty"` // The list of build commands. These commands are executed in the build container.
}

// Pipelinescheduleputrequestbody represents the Pipelinescheduleputrequestbody schema from the OpenAPI specification
type Pipelinescheduleputrequestbody struct {
	TypeField string `json:"type"`
	Enabled bool `json:"enabled,omitempty"` // Whether the schedule is enabled.
}

// Stgwestreport represents the Stgwestreport schema from the OpenAPI specification
type Stgwestreport struct {
	TypeField string `json:"type"`
}

// Paginateddeploykeys represents the Paginateddeploykeys schema from the OpenAPI specification
type Paginateddeploykeys struct {
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Deploykey `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
}

// Searchline represents the Searchline schema from the OpenAPI specification
type Searchline struct {
	Segments []Searchsegment `json:"segments,omitempty"`
	Line int `json:"line,omitempty"`
}

// Searchcontentmatch represents the Searchcontentmatch schema from the OpenAPI specification
type Searchcontentmatch struct {
	Lines []Searchline `json:"lines,omitempty"`
}

// Workspace represents the Workspace schema from the OpenAPI specification
type Workspace struct {
	TypeField string `json:"type"`
	Created_on string `json:"created_on,omitempty"`
	Is_private bool `json:"is_private,omitempty"` // Indicates whether the workspace is publicly accessible, or whether it is private to the members and consequently only visible to members.
	Links map[string]interface{} `json:"links,omitempty"`
	Name string `json:"name,omitempty"` // The name of the workspace.
	Slug string `json:"slug,omitempty"` // The short label that identifies this workspace.
	Updated_on string `json:"updated_on,omitempty"`
	Uuid string `json:"uuid,omitempty"` // The workspace's immutable id.
}

// Diffstat represents the Diffstat schema from the OpenAPI specification
type Diffstat struct {
	Old Commitfile `json:"old,omitempty"` // A file object, representing a file at a commit in a repository
	Status string `json:"status,omitempty"`
	TypeField string `json:"type"`
	Lines_added int `json:"lines_added,omitempty"`
	Lines_removed int `json:"lines_removed,omitempty"`
	NewField Commitfile `json:"new,omitempty"` // A file object, representing a file at a commit in a repository
}

// Pipelinescheduleexecution represents the Pipelinescheduleexecution schema from the OpenAPI specification
type Pipelinescheduleexecution struct {
	TypeField string `json:"type"`
}

// Pipelinecachecontenturi represents the Pipelinecachecontenturi schema from the OpenAPI specification
type Pipelinecachecontenturi struct {
	Uri string `json:"uri,omitempty"` // The uri for pipeline cache content.
}

// Branch represents the Branch schema from the OpenAPI specification
type Branch struct {
	Links map[string]interface{} `json:"links,omitempty"`
	Name string `json:"name,omitempty"` // The name of the ref.
	Target Commit `json:"target,omitempty"`
	TypeField string `json:"type"`
	Default_merge_strategy string `json:"default_merge_strategy,omitempty"` // The default merge strategy for pull requests targeting this branch.
	Merge_strategies []string `json:"merge_strategies,omitempty"` // Available merge strategies for pull requests targeting this branch.
}

// Deploymentsstgwestdeploymentenvironment represents the Deploymentsstgwestdeploymentenvironment schema from the OpenAPI specification
type Deploymentsstgwestdeploymentenvironment struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the environment.
	Uuid string `json:"uuid,omitempty"` // The UUID identifying the environment.
}

// Repositorypermission represents the Repositorypermission schema from the OpenAPI specification
type Repositorypermission struct {
	Permission string `json:"permission,omitempty"`
	Repository Repository `json:"repository,omitempty"`
	TypeField string `json:"type"`
	User User `json:"user,omitempty"`
}

// Subjecttypes represents the Subjecttypes schema from the OpenAPI specification
type Subjecttypes struct {
	Repository map[string]interface{} `json:"repository,omitempty"`
	Workspace map[string]interface{} `json:"workspace,omitempty"`
}

// Paginatedrepositories represents the Paginatedrepositories schema from the OpenAPI specification
type Paginatedrepositories struct {
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Repository `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
}

// Pipelinestepstatecompletedexpired represents the Pipelinestepstatecompletedexpired schema from the OpenAPI specification
type Pipelinestepstatecompletedexpired struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the result (EXPIRED)
}

// Paginatedcommitcomments represents the Paginatedcommitcomments schema from the OpenAPI specification
type Paginatedcommitcomments struct {
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Commitcomment `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
}

// Paginatedhookevents represents the Paginatedhookevents schema from the OpenAPI specification
type Paginatedhookevents struct {
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Hookevent `json:"values,omitempty"`
}

// Ddevreport represents the Ddevreport schema from the OpenAPI specification
type Ddevreport struct {
	TypeField string `json:"type"`
}

// Appuser represents the Appuser schema from the OpenAPI specification
type Appuser struct {
	TypeField string `json:"type"`
	Display_name string `json:"display_name,omitempty"`
	Links Accountlinks `json:"links,omitempty"` // Links related to an Account.
	Username string `json:"username,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Created_on string `json:"created_on,omitempty"`
	Account_id string `json:"account_id,omitempty"` // The user's Atlassian account ID.
	Account_status string `json:"account_status,omitempty"` // The status of the account. Currently the only possible value is "active", but more values may be added in the future.
	Kind string `json:"kind,omitempty"` // The kind of App User.
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	ErrorField map[string]interface{} `json:"error,omitempty"`
	TypeField string `json:"type"`
}

// Comment represents the Comment schema from the OpenAPI specification
type Comment struct {
	TypeField string `json:"type"`
	Created_on string `json:"created_on,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Id int `json:"id,omitempty"`
	Inline map[string]interface{} `json:"inline,omitempty"`
	Content map[string]interface{} `json:"content,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Updated_on string `json:"updated_on,omitempty"`
	User Account `json:"user,omitempty"`
	Parent Comment `json:"parent,omitempty"`
}

// Paginatedbranches represents the Paginatedbranches schema from the OpenAPI specification
type Paginatedbranches struct {
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Branch `json:"values,omitempty"`
}

// Deploymentsstgwestpaginatedenvironments represents the Deploymentsstgwestpaginatedenvironments schema from the OpenAPI specification
type Deploymentsstgwestpaginatedenvironments struct {
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Deploymentsstgwestdeploymentenvironment `json:"values,omitempty"` // The values of the current page.
}

// Pipelinesshkeypair represents the Pipelinesshkeypair schema from the OpenAPI specification
type Pipelinesshkeypair struct {
	TypeField string `json:"type"`
	Private_key string `json:"private_key,omitempty"` // The SSH private key. This value will be empty when retrieving the SSH key pair.
	Public_key string `json:"public_key,omitempty"` // The SSH public key.
}

// Userlinks represents the Userlinks schema from the OpenAPI specification
type Userlinks struct {
	Avatar Link `json:"avatar,omitempty"` // A link to a resource related to this object.
	Html Link `json:"html,omitempty"` // A link to a resource related to this object.
	Repositories Link `json:"repositories,omitempty"` // A link to a resource related to this object.
	Self Link `json:"self,omitempty"` // A link to a resource related to this object.
}

// Projectbranchingmodel represents the Projectbranchingmodel schema from the OpenAPI specification
type Projectbranchingmodel struct {
	TypeField string `json:"type"`
	Branch_types []map[string]interface{} `json:"branch_types,omitempty"` // The active branch types.
	Development map[string]interface{} `json:"development,omitempty"`
	Production map[string]interface{} `json:"production,omitempty"`
}

// Paginatedmilestones represents the Paginatedmilestones schema from the OpenAPI specification
type Paginatedmilestones struct {
	Values []Milestone `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
}

// Sshaccountkey represents the Sshaccountkey schema from the OpenAPI specification
type Sshaccountkey struct {
	TypeField string `json:"type"`
	Last_used string `json:"last_used,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Uuid string `json:"uuid,omitempty"` // The SSH key's immutable ID.
	Comment string `json:"comment,omitempty"` // The comment parsed from the SSH key (if present)
	Created_on string `json:"created_on,omitempty"`
	Key string `json:"key,omitempty"` // The SSH public key value in OpenSSH format.
	Label string `json:"label,omitempty"` // The user-defined label for the SSH key
	Owner Account `json:"owner,omitempty"`
}

// Paginatedaccounts represents the Paginatedaccounts schema from the OpenAPI specification
type Paginatedaccounts struct {
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Account `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
}

// Paginatedannotations represents the Paginatedannotations schema from the OpenAPI specification
type Paginatedannotations struct {
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Reportannotation `json:"values,omitempty"` // The values of the current page.
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
}

// Paginatedcomponents represents the Paginatedcomponents schema from the OpenAPI specification
type Paginatedcomponents struct {
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Component `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
}

// Pipelinestepstatecompletedresult represents the Pipelinestepstatecompletedresult schema from the OpenAPI specification
type Pipelinestepstatecompletedresult struct {
	TypeField string `json:"type"`
}

// Commitfile represents the Commitfile schema from the OpenAPI specification
type Commitfile struct {
	Commit Commit `json:"commit,omitempty"`
	Escaped_path string `json:"escaped_path,omitempty"` // The escaped version of the path as it appears in a diff. If the path does not require escaping this will be the same as path.
	Path string `json:"path,omitempty"` // The path in the repository
	TypeField string `json:"type"`
	Attributes string `json:"attributes,omitempty"`
}

// Deploymentstatecompletedstatus represents the Deploymentstatecompletedstatus schema from the OpenAPI specification
type Deploymentstatecompletedstatus struct {
	TypeField string `json:"type"`
}

// Deploymentenvironment represents the Deploymentenvironment schema from the OpenAPI specification
type Deploymentenvironment struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the environment.
	Uuid string `json:"uuid,omitempty"` // The UUID identifying the environment.
}

// Paginatedreports represents the Paginatedreports schema from the OpenAPI specification
type Paginatedreports struct {
	Values []Report `json:"values,omitempty"` // The values of the current page.
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
}

// Paginatedfiles represents the Paginatedfiles schema from the OpenAPI specification
type Paginatedfiles struct {
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Commitfile `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
}

// Pipelinestepstatecompletederror represents the Pipelinestepstatecompletederror schema from the OpenAPI specification
type Pipelinestepstatecompletederror struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the result (ERROR)
	ErrorField Pipelinesteperror `json:"error,omitempty"`
}

// Pipelinestepstatecompleted represents the Pipelinestepstatecompleted schema from the OpenAPI specification
type Pipelinestepstatecompleted struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of pipeline step state (COMPLETED).
	Result Pipelinestepstatecompletedresult `json:"result,omitempty"`
}

// Paginatedtags represents the Paginatedtags schema from the OpenAPI specification
type Paginatedtags struct {
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Tag `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
}

// Jiraproject represents the Jiraproject schema from the OpenAPI specification
type Jiraproject struct {
	TypeField string `json:"type"`
}

// Issuejobstatus represents the Issuejobstatus schema from the OpenAPI specification
type Issuejobstatus struct {
	Count int `json:"count,omitempty"` // The total number of issues already imported/exported
	Pct float64 `json:"pct,omitempty"` // The percentage of issues already imported/exported
	Phase string `json:"phase,omitempty"` // The phase of the import/export job
	Status string `json:"status,omitempty"` // The status of the import/export job
	Total int `json:"total,omitempty"` // The total number of issues being imported/exported
	TypeField string `json:"type,omitempty"`
}

// Pipelinetriggermanual represents the Pipelinetriggermanual schema from the OpenAPI specification
type Pipelinetriggermanual struct {
	TypeField string `json:"type"`
}

// Pipelinetrigger represents the Pipelinetrigger schema from the OpenAPI specification
type Pipelinetrigger struct {
	TypeField string `json:"type"`
}

// Paginatedbranchrestrictions represents the Paginatedbranchrestrictions schema from the OpenAPI specification
type Paginatedbranchrestrictions struct {
	Values []Branchrestriction `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
}

// Paginateddefaultreviewerandtype represents the Paginateddefaultreviewerandtype schema from the OpenAPI specification
type Paginateddefaultreviewerandtype struct {
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Defaultreviewerandtype `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
}

// Paginatedversions represents the Paginatedversions schema from the OpenAPI specification
type Paginatedversions struct {
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Version `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
}

// Branchrestriction represents the Branchrestriction schema from the OpenAPI specification
type Branchrestriction struct {
	TypeField string `json:"type"`
	Groups []Group `json:"groups,omitempty"`
	Users []Account `json:"users,omitempty"`
}

// Pipelinesstgwestpipelinestep represents the Pipelinesstgwestpipelinestep schema from the OpenAPI specification
type Pipelinesstgwestpipelinestep struct {
	TypeField string `json:"type"`
}

// Paginatedrepositoryuserpermissions represents the Paginatedrepositoryuserpermissions schema from the OpenAPI specification
type Paginatedrepositoryuserpermissions struct {
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Repositoryuserpermission `json:"values,omitempty"`
}

// Pipelineschedulepostrequestbody represents the Pipelineschedulepostrequestbody schema from the OpenAPI specification
type Pipelineschedulepostrequestbody struct {
	TypeField string `json:"type"`
	Cron_pattern string `json:"cron_pattern"` // The cron expression with second precision (7 fields) that the schedule applies. For example, for expression: 0 0 12 * * ? *, will execute at 12pm UTC every day.
	Enabled bool `json:"enabled,omitempty"` // Whether the schedule is enabled.
	Target map[string]interface{} `json:"target"` // The target on which the schedule will be executed.
}

// Paginatedpipelinecaches represents the Paginatedpipelinecaches schema from the OpenAPI specification
type Paginatedpipelinecaches struct {
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Pipelinecache `json:"values,omitempty"` // The values of the current page.
}

// Applicationproperty represents the Applicationproperty schema from the OpenAPI specification
type Applicationproperty struct {
	Attributes []string `json:"_attributes,omitempty"`
}

// Deploymentstateundeployed represents the Deploymentstateundeployed schema from the OpenAPI specification
type Deploymentstateundeployed struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of deployment state (UNDEPLOYED).
	Trigger_url string `json:"trigger_url,omitempty"` // Link to trigger the deployment.
}

// Object represents the Object schema from the OpenAPI specification
type Object struct {
	TypeField string `json:"type"`
}

// Pipelinebuildnumber represents the Pipelinebuildnumber schema from the OpenAPI specification
type Pipelinebuildnumber struct {
	TypeField string `json:"type"`
	Next int `json:"next,omitempty"` // The next number that will be used as build number.
}

// Searchcodesearchresult represents the Searchcodesearchresult schema from the OpenAPI specification
type Searchcodesearchresult struct {
	TypeField string `json:"type,omitempty"`
	Content_match_count int64 `json:"content_match_count,omitempty"`
	Content_matches []Searchcontentmatch `json:"content_matches,omitempty"`
	File Commitfile `json:"file,omitempty"` // A file object, representing a file at a commit in a repository
	Path_matches []Searchsegment `json:"path_matches,omitempty"`
}

// Pipelinesshpublickey represents the Pipelinesshpublickey schema from the OpenAPI specification
type Pipelinesshpublickey struct {
	TypeField string `json:"type"`
	Key string `json:"key,omitempty"` // The base64 encoded public key.
	Key_type string `json:"key_type,omitempty"` // The type of the public key.
	Md5_fingerprint string `json:"md5_fingerprint,omitempty"` // The MD5 fingerprint of the public key.
	Sha256_fingerprint string `json:"sha256_fingerprint,omitempty"` // The SHA-256 fingerprint of the public key.
}

// Pipelinestatecompletedstopped represents the Pipelinestatecompletedstopped schema from the OpenAPI specification
type Pipelinestatecompletedstopped struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the stopped result (STOPPED).
}

// Deploymentstatecompleted represents the Deploymentstatecompleted schema from the OpenAPI specification
type Deploymentstatecompleted struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of deployment state (COMPLETED).
	Start_date string `json:"start_date,omitempty"` // The timestamp when the deployment was started.
	Status Deploymentstatecompletedstatus `json:"status,omitempty"`
	Url string `json:"url,omitempty"` // Link to the deployment result.
	Completion_date string `json:"completion_date,omitempty"` // The timestamp when the deployment completed.
	Deployer Account `json:"deployer,omitempty"`
}

// Paginatedenvironments represents the Paginatedenvironments schema from the OpenAPI specification
type Paginatedenvironments struct {
	Values []Deploymentenvironment `json:"values,omitempty"` // The values of the current page.
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
}

// Effectiverepobranchingmodel represents the Effectiverepobranchingmodel schema from the OpenAPI specification
type Effectiverepobranchingmodel struct {
	TypeField string `json:"type"`
	Development map[string]interface{} `json:"development,omitempty"`
	Production map[string]interface{} `json:"production,omitempty"`
	Branch_types []map[string]interface{} `json:"branch_types,omitempty"` // The active branch types.
}

// Pipelineknownhost represents the Pipelineknownhost schema from the OpenAPI specification
type Pipelineknownhost struct {
	TypeField string `json:"type"`
	Public_key Pipelinesshpublickey `json:"public_key,omitempty"`
	Uuid string `json:"uuid,omitempty"` // The UUID identifying the known host.
	Hostname string `json:"hostname,omitempty"` // The hostname of the known host.
}

// Pipelinestatecompletedfailed represents the Pipelinestatecompletedfailed schema from the OpenAPI specification
type Pipelinestatecompletedfailed struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the failed result (FAILED).
}

// Basecommit represents the Basecommit schema from the OpenAPI specification
type Basecommit struct {
	TypeField string `json:"type"`
	Date string `json:"date,omitempty"`
	Hash string `json:"hash,omitempty"`
	Message string `json:"message,omitempty"`
	Parents []Basecommit `json:"parents,omitempty"`
	Summary map[string]interface{} `json:"summary,omitempty"`
	Author Author `json:"author,omitempty"`
}

// Pipelinestatecompletedresult represents the Pipelinestatecompletedresult schema from the OpenAPI specification
type Pipelinestatecompletedresult struct {
	TypeField string `json:"type"`
}

// Link represents the Link schema from the OpenAPI specification
type Link struct {
	Href string `json:"href,omitempty"`
	Name string `json:"name,omitempty"`
}

// Account represents the Account schema from the OpenAPI specification
type Account struct {
	TypeField string `json:"type"`
	Username string `json:"username,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Created_on string `json:"created_on,omitempty"`
	Display_name string `json:"display_name,omitempty"`
	Links Accountlinks `json:"links,omitempty"` // Links related to an Account.
}

// Reportdata represents the Reportdata schema from the OpenAPI specification
type Reportdata struct {
	TypeField string `json:"type,omitempty"` // The type of data contained in the value field. If not provided, then the value will be detected as a boolean, number or string.
	Value map[string]interface{} `json:"value,omitempty"` // The value of the data element.
	Title string `json:"title,omitempty"` // A string describing what this data field represents.
}

// Pipelinestateinprogressrunning represents the Pipelinestateinprogressrunning schema from the OpenAPI specification
type Pipelinestateinprogressrunning struct {
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"` // The name of the stage (RUNNING)
}

// Pipelinesddevpipelinestep represents the Pipelinesddevpipelinestep schema from the OpenAPI specification
type Pipelinesddevpipelinestep struct {
	TypeField string `json:"type"`
}

// Paginatedpipelinesteps represents the Paginatedpipelinesteps schema from the OpenAPI specification
type Paginatedpipelinesteps struct {
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Pipelinestep `json:"values,omitempty"` // The values of the current page.
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
}

// Pipelineimage represents the Pipelineimage schema from the OpenAPI specification
type Pipelineimage struct {
	Name string `json:"name,omitempty"` // The name of the image. If the image is hosted on DockerHub the short name can be used, otherwise the fully qualified name is required here.
	Password string `json:"password,omitempty"` // The password needed to authenticate with the Docker registry. Only required when using a private Docker image.
	Username string `json:"username,omitempty"` // The username needed to authenticate with the Docker registry. Only required when using a private Docker image.
	Email string `json:"email,omitempty"` // The email needed to authenticate with the Docker registry. Only required when using a private Docker image.
}

// Pullrequestmergeparameters represents the Pullrequestmergeparameters schema from the OpenAPI specification
type Pullrequestmergeparameters struct {
	Close_source_branch bool `json:"close_source_branch,omitempty"` // Whether the source branch should be deleted. If this is not provided, we fallback to the value used when the pull request was created, which defaults to False
	Merge_strategy string `json:"merge_strategy,omitempty"` // The merge strategy that will be used to merge the pull request.
	Message string `json:"message,omitempty"` // The commit message that will be used on the resulting commit.
	TypeField string `json:"type"`
}

// Paginatedrefs represents the Paginatedrefs schema from the OpenAPI specification
type Paginatedrefs struct {
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Ref `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
}

// Paginatedrepositorypermissions represents the Paginatedrepositorypermissions schema from the OpenAPI specification
type Paginatedrepositorypermissions struct {
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Values []Repositorypermission `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
}

// Report represents the Report schema from the OpenAPI specification
type Report struct {
	TypeField string `json:"type"`
	Uuid string `json:"uuid,omitempty"` // The UUID that can be used to identify the report.
	External_id string `json:"external_id,omitempty"` // ID of the report provided by the report creator. It can be used to identify the report as an alternative to it's generated uuid. It is not used by Bitbucket, but only by the report creator for updating or deleting this specific report. Needs to be unique.
	Title string `json:"title,omitempty"` // The title of the report.
	Updated_on string `json:"updated_on,omitempty"` // The timestamp when the report was updated.
	Created_on string `json:"created_on,omitempty"` // The timestamp when the report was created.
	Data []Reportdata `json:"data,omitempty"` // An array of data fields to display information on the report. Maximum 10.
	Link string `json:"link,omitempty"` // A URL linking to the results of the report in an external tool.
	Logo_url string `json:"logo_url,omitempty"` // A URL to the report logo. If none is provided, the default insights logo will be used.
	Reporter string `json:"reporter,omitempty"` // A string to describe the tool or company who created the report.
	Details string `json:"details,omitempty"` // A string to describe the purpose of the report.
	Remote_link_enabled bool `json:"remote_link_enabled,omitempty"` // If enabled, a remote link is created in Jira for the issue associated with the commit the report belongs to.
	Report_type string `json:"report_type,omitempty"` // The type of the report.
	Result string `json:"result,omitempty"` // The state of the report. May be set to PENDING and later updated.
}

// Webhooksubscription represents the Webhooksubscription schema from the OpenAPI specification
type Webhooksubscription struct {
	TypeField string `json:"type"`
	Events []string `json:"events,omitempty"` // The events this webhook is subscribed to.
	Subject Object `json:"subject,omitempty"` // Base type for most resource objects. It defines the common `type` element that identifies an object's type. It also identifies the element as Swagger's `discriminator`.
	Subject_type string `json:"subject_type,omitempty"` // The type of entity. Set to either `repository` or `workspace` based on where the subscription is defined.
	Url string `json:"url,omitempty"` // The URL events get delivered to.
	Uuid string `json:"uuid,omitempty"` // The webhook's id
	Active bool `json:"active,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Description string `json:"description,omitempty"` // A user-defined description of the webhook.
}

// Pipelinecommittarget represents the Pipelinecommittarget schema from the OpenAPI specification
type Pipelinecommittarget struct {
	TypeField string `json:"type"`
	Commit Commit `json:"commit,omitempty"`
	Selector Pipelineselector `json:"selector,omitempty"`
}

// Paginatedpullrequestcomments represents the Paginatedpullrequestcomments schema from the OpenAPI specification
type Paginatedpullrequestcomments struct {
	Values []Pullrequestcomment `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
}

// Commentresolution represents the Commentresolution schema from the OpenAPI specification
type Commentresolution struct {
	TypeField string `json:"type"`
	User Account `json:"user,omitempty"`
	Created_on string `json:"created_on,omitempty"` // The ISO8601 timestamp the resolution was created.
}

// Defaultreviewerandtype represents the Defaultreviewerandtype schema from the OpenAPI specification
type Defaultreviewerandtype struct {
	Reviewer_type string `json:"reviewer_type,omitempty"`
	TypeField string `json:"type"`
	User User `json:"user,omitempty"`
}

// Paginatedsnippets represents the Paginatedsnippets schema from the OpenAPI specification
type Paginatedsnippets struct {
	Values []Snippet `json:"values,omitempty"`
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
}

// Pipelinestatecompleted represents the Pipelinestatecompleted schema from the OpenAPI specification
type Pipelinestatecompleted struct {
	TypeField string `json:"type"`
	Result Pipelinestatecompletedresult `json:"result,omitempty"`
	Name string `json:"name,omitempty"` // The name of pipeline state (COMPLETED).
}

// Commitstatus represents the Commitstatus schema from the OpenAPI specification
type Commitstatus struct {
	TypeField string `json:"type"`
	Created_on string `json:"created_on,omitempty"`
	State string `json:"state,omitempty"` // Provides some indication of the status of this commit
	Description string `json:"description,omitempty"` // A description of the build (e.g. "Unit tests in Bamboo")
	Key string `json:"key,omitempty"` // An identifier for the status that's unique to its type (current "build" is the only supported type) and the vendor, e.g. BB-DEPLOY
	Refname string `json:"refname,omitempty"` // The name of the ref that pointed to this commit at the time the status object was created. Note that this the ref may since have moved off of the commit. This optional field can be useful for build systems whose build triggers and configuration are branch-dependent (e.g. a Pipeline build). It is legitimate for this field to not be set, or even apply (e.g. a static linting job).
	Url string `json:"url,omitempty"` // A URL linking back to the vendor or build system, for providing more information about whatever process produced this status. Accepts context variables `repository` and `commit` that Bitbucket will evaluate at runtime whenever at runtime. For example, one could use https://foo.com/builds/{repository.full_name} which Bitbucket will turn into https://foo.com/builds/foo/bar at render time.
	Links map[string]interface{} `json:"links,omitempty"`
	Name string `json:"name,omitempty"` // An identifier for the build itself, e.g. BB-DEPLOY-1
	Updated_on string `json:"updated_on,omitempty"`
	Uuid string `json:"uuid,omitempty"` // The commit status' id.
}

// Searchresultpage represents the Searchresultpage schema from the OpenAPI specification
type Searchresultpage struct {
	Query_substituted bool `json:"query_substituted,omitempty"`
	Size int64 `json:"size,omitempty"`
	Values []Searchcodesearchresult `json:"values,omitempty"`
	Next string `json:"next,omitempty"`
	Page int `json:"page,omitempty"`
	Pagelen int `json:"pagelen,omitempty"`
	Previous string `json:"previous,omitempty"`
}

// Paginatedpipelineschedules represents the Paginatedpipelineschedules schema from the OpenAPI specification
type Paginatedpipelineschedules struct {
	Values []Pipelineschedule `json:"values,omitempty"` // The values of the current page.
	Next string `json:"next,omitempty"` // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Page int `json:"page,omitempty"` // Page number of the current results. This is an optional element that is not provided in all responses.
	Pagelen int `json:"pagelen,omitempty"` // Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Previous string `json:"previous,omitempty"` // Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Size int `json:"size,omitempty"` // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
}

// Issue represents the Issue schema from the OpenAPI specification
type Issue struct {
	TypeField string `json:"type"`
	Component Component `json:"component,omitempty"`
	Repository Repository `json:"repository,omitempty"`
	Votes int `json:"votes,omitempty"`
	State string `json:"state,omitempty"`
	Reporter Account `json:"reporter,omitempty"`
	Content map[string]interface{} `json:"content,omitempty"`
	Edited_on string `json:"edited_on,omitempty"`
	Title string `json:"title,omitempty"`
	Kind string `json:"kind,omitempty"`
	Assignee Account `json:"assignee,omitempty"`
	Created_on string `json:"created_on,omitempty"`
	Updated_on string `json:"updated_on,omitempty"`
	Id int `json:"id,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Version Version `json:"version,omitempty"`
	Milestone Milestone `json:"milestone,omitempty"`
	Priority string `json:"priority,omitempty"`
}
