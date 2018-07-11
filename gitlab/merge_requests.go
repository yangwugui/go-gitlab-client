package gitlab

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

const (
	mergeRequestsUrl                  = "/merge_requests"
	projectMergeRequestsUrl           = "/projects/:id/merge_requests"
	groupMergeRequestsUrl             = "/groups/:id/merge_requests"
	projectMergeRequestUrl            = "/projects/:id/merge_requests/:merge_request_id"                                  // Get information about a single merge request
	projectMergeRequestCommitsUrl     = "/projects/:id/merge_requests/:merge_request_id/commits"                          // Get a list of merge request commits
	projectMergeRequestChangesUrl     = "/projects/:id/merge_requests/:merge_request_id/changes"                          // Shows information about the merge request including its files and changes
	projectMergeRequestMergeUrl       = "/projects/:id/merge_requests/:merge_request_id/merge"                            // Merge changes submitted with MR
	projectMergeRequestCancelMergeUrl = "/projects/:id/merge_requests/:merge_request_id/cancel_merge_when_build_succeeds" // Cancel Merge When Build Succeeds
	projectMergeRequestCommentsUrl    = "/projects/:id/merge_requests/:merge_request_id/comments"                         // Lists all comments associated with a merge request
)

type MergeRequest struct {
	Id                        int               `json:"id,omitempty" yaml:"id,omitempty"`
	Iid                       int               `json:"iid,omitempty" yaml:"iid,omitempty"`
	ProjectId                 int               `json:"project_id,omitempty" yaml:"project_id,omitempty"`
	WebUrl                    string            `json:"web_url,omitempty" yaml:"web_url,omitempty"`
	Title                     string            `json:"title,omitempty" yaml:"title,omitempty"`
	Description               string            `json:"description,omitempty" yaml:"description,omitempty"`
	State                     string            `json:"state,omitempty" yaml:"state,omitempty"`
	CreatedAt                 string            `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt                 string            `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	SourceBranch              string            `json:"source_branch,omitempty" yaml:"source_branch,omitempty"`
	TargetBranch              string            `json:"target_branch,omitempty" yaml:"target_branch,omitempty"`
	Upvotes                   int               `json:"upvotes,omitempty" yaml:"upvotes,omitempty"`
	Downvotes                 int               `json:"downvotes,omitempty" yaml:"downvotes,omitempty"`
	SourceProjectID           int               `json:"source_project_id,omitempty" yaml:"source_project_id,omitempty"`
	TargetProjectID           int               `json:"target_project_id,omitempty" yaml:"target_project_id,omitempty"`
	Sha                       string            `json:"sha,omitempty" yaml:"sha,omitempty"`
	MergeCommitSha            string            `json:"merge_commit_sha,omitempty" yaml:"merge_commit_sha,omitempty"`
	WorkInProgress            bool              `json:"work_in_progress,omitempty" yaml:"work_in_progress,omitempty"`
	MergeStatus               string            `json:"merge_status,omitempty" yaml:"merge_status,omitempty"`
	Squash                    bool              `json:"squash,omitempty" yaml:"squash,omitempty"`
	MergeWhenPipelineSucceeds bool              `json:"merge_when_pipeline_succeeds,omitempty" yaml:"merge_when_pipeline_succeeds,omitempty"`
	ShouldRemoveSourceBranch  bool              `json:"should_remove_source_branch,omitempty" yaml:"should_remove_source_branch,omitempty"`
	ForceRemoveSourceBranch   bool              `json:"force_remove_source_branch,omitempty" yaml:"force_remove_source_branch,omitempty"`
	DiscussionLocked          bool              `json:"discussion_locked,omitempty" yaml:"discussion_locked,omitempty"`
	UserNotesCount            int               `json:"user_notes_count,omitempty" yaml:"user_notes_count,omitempty"`
	Pipeline                  *Pipeline         `json:"pipeline,omitempty" yaml:"pipeline,omitempty"`
	Author                    *MergeRequestUser `json:"author,omitempty" yaml:"author,omitempty"`
	Assignee                  *MergeRequestUser `json:"assignee,omitempty" yaml:"assignee,omitempty"`
	Labels                    []string          `json:"labels,omitempty" yaml:"labels,omitempty"`
	TimeStats                 *TimeStats        `json:"time_stats,omitempty" yaml:"time_stats,omitempty"`
	Milestone                 *Milestone        `json:"milestone,omitempty" yaml:"milestone,omitempty"`
}

type MergeRequestUser struct {
	Id        int    `json:"id,omitempty" yaml:"id,omitempty"`
	Username  string `json:"username,omitempty" yaml:"username,omitempty"`
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	State     string `json:"state,omitempty" yaml:"state,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty" yaml:"avatar_url,omitempty"`
	WebUrl    string `json:"web_url,omitempty" yaml:"web_url,omitempty"`
}

type TimeStats struct {
	TimeEstimate        int    `json:"time_estimate,omitempty" yaml:"time_estimate,omitempty"`
	TotalTimeSpent      int    `json:"total_time_spent,omitempty" yaml:"total_time_spent,omitempty"`
	HumanTimeEstimate   string `json:"human_time_estimate,omitempty" yaml:"human_time_estimate,omitempty"`
	HumanTotalTimeSpent string `json:"human_total_time_spent,omitempty" yaml:"human_total_time_spent,omitempty"`
}

type ChangeItem struct {
	OldPath     string `json:"old_path,omitempty"`
	NewPath     string `json:"new_path,omitempty"`
	AMode       string `json:"a_mode,omitempty"`
	BMode       string `json:"b_mode,omitempty"`
	Diff        string `json:"diff,omitempty"`
	NewFile     bool   `json:"new_file,omitempty"`
	RenamedFile bool   `json:"renamed_file,omitempty"`
	DeletedFile bool   `json:"deleted_file,omitempty"`
}

type MergeRequestChanges struct {
	*MergeRequest
	CreatedAt       string       `json:"created_at,omitempty"`
	UpdatedAt       string       `json:"updated_at,omitempty"`
	SourceProjectId int          `json:"source_project_id,omitempty"`
	TargetProjectId int          `json:"target_project_id,omitempty"`
	Labels          []string     `json:"labels,omitempty"`
	Milestone       Milestone    `json:"milestone,omitempty"`
	Changes         []ChangeItem `json:"changes,omitempty"`
}

type AddMergeRequestRequest struct {
	SourceBranch    string   `json:"source_branch"`
	TargetBranch    string   `json:"target_branch"`
	AssigneeId      int      `json:"assignee_id,omitempty"`
	Title           string   `json:"title"`
	Description     string   `json:"description,omitempty"`
	TargetProjectId int      `json:"target_project_id,omitempty"`
	Lables          []string `json:"lables,omitempty"`
}

type AcceptMergeRequestRequest struct {
	MergeCommitMessage       string `json:"merge_commit_message,omitempty"`
	ShouldRemoveSourceBranch bool   `json:"should_remove_source_branch,omitempty"`
	MergedWhenBuildSucceeds  bool   `json:"merged_when_build_succeeds,omitempty"`
}

type MergeRequestScope string

// For versions before 11.0, use the now deprecated
// created-by-me or assigned-to-me scopes
const (
	MergeRequestScopeCreatedByMe        MergeRequestScope = "created_by_me"
	LegacyMergeRequestScopeCreatedByMe  MergeRequestScope = "created-by-me"
	MergeRequestScopeAssignedToMe       MergeRequestScope = "assigned_to_me"
	LegacyMergeRequestScopeAssignedToMe MergeRequestScope = "assigned-to-me"
	MergeRequestScopeAll                MergeRequestScope = "all"
)

type MergeRequestsOptions struct {
	PaginationOptions

	// Return the request having the given iid
	// only available for generic merge requests API endpoint
	iids []int

	// Return all merge requests or just those that are
	// opened, closed, locked, or merged
	State string

	// Return merge requests ordered by created_at or
	// updated_at fields. Default is created_at
	OrderBY string

	// sort	string	no	Return requests sorted in asc or desc order
	// Default is desc
	Sort string

	// Return merge requests for a specific milestone
	Milestone string

	// If simple, returns the iid, URL, title, description,
	// and basic state of merge request
	View string

	// Return merge requests matching a comma separated
	// list of labels
	Labels []string

	// Return merge requests created on or after the given time
	CreatedAfter *time.Time

	// Return merge requests created on or before the given time
	CreatedBefore *time.Time

	// Return merge requests updated on or after the given time
	UpdatedAfter *time.Time

	// Return merge requests updated on or before the given time
	UpdatedBefore *time.Time

	// Return merge requests with the given source branch
	SourceBranch string

	// Return merge requests with the given target branch
	TargetBranch string

	// Search merge requests against their title and description
	Search string

	// Returns merge requests created by the given user id
	AuthorId int

	// Returns merge requests assigned to the given user id
	AssigneeId int

	// Return merge requests reacted by the authenticated user by the given emoji
	MyReactionEmoji string

	// Return merge requests for the given scope: created_by_me, assigned_to_me or all,
	// For versions before 11.0, use the now deprecated created-by-me or assigned-to-me scopes instead.
	Scope MergeRequestScope
}

func (g *Gitlab) getMergeRequests(u *url.URL, o *MergeRequestsOptions) ([]*MergeRequest, *ResponseMeta, error) {
	if o != nil {
		q := u.Query()

		if o.Page != 1 {
			q.Set("page", strconv.Itoa(o.Page))
		}
		if o.PerPage != 0 {
			q.Set("per_page", strconv.Itoa(o.PerPage))
		}

		u.RawQuery = q.Encode()
	}

	var err error
	var mergeRequests []*MergeRequest

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &mergeRequests)
	}

	return mergeRequests, meta, err
}

func (g *Gitlab) MergeRequests(o *MergeRequestsOptions) ([]*MergeRequest, *ResponseMeta, error) {
	u := g.ResourceUrl(mergeRequestsUrl, nil)

	return g.getMergeRequests(u, o)
}

func (g *Gitlab) ProjectMergeRequests(projectId string, o *MergeRequestsOptions) ([]*MergeRequest, *ResponseMeta, error) {
	u := g.ResourceUrl(projectMergeRequestsUrl, map[string]string{":id": projectId})

	return g.getMergeRequests(u, o)
}

func (g *Gitlab) GroupMergeRequests(groupId int, o *MergeRequestsOptions) ([]*MergeRequest, *ResponseMeta, error) {
	u := g.ResourceUrl(groupMergeRequestsUrl, map[string]string{":id": strconv.Itoa(groupId)})

	return g.getMergeRequests(u, o)
}

func (g *Gitlab) ProjectMergeRequest(projectId string, mergeRequestId int) (*MergeRequest, *ResponseMeta, error) {
	u := g.ResourceUrl(projectMergeRequestUrl, map[string]string{
		":id":               projectId,
		":merge_request_id": strconv.Itoa(mergeRequestId),
	})

	var err error
	mr := new(MergeRequest)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &mr)
	}

	return mr, meta, err
}

/*
Get a list of merge request commits.

    GET /projects/:id/merge_request/:merge_request_id/commits

Parameters:

    id               The ID of a project
    merge_request_id The ID of a merge request

*/
func (g *Gitlab) ProjectMergeRequestCommits(id, merge_request_id string) ([]*Commit, *ResponseMeta, error) {
	u := g.ResourceUrl(projectMergeRequestCommitsUrl, map[string]string{
		":id":               id,
		":merge_request_id": merge_request_id,
	})

	var err error
	var commits []*Commit

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &commits)
		if err == nil {
			for _, commit := range commits {
				t, _ := time.Parse(dateLayout, commit.Created_At)
				commit.CreatedAt = t
			}
		}
	}

	return commits, meta, err
}

/*
Get information about the merge request including its files and changes.

    GET /projects/:id/merge_request/:merge_request_id/changes

Parameters:

    id               The ID of a project
    merge_request_id The ID of a merge request

*/
func (g *Gitlab) ProjectMergeRequestChanges(id, merge_request_id string) (*MergeRequestChanges, *ResponseMeta, error) {
	u := g.ResourceUrl(projectMergeRequestChangesUrl, map[string]string{
		":id":               id,
		":merge_request_id": merge_request_id,
	})

	var err error
	changes := new(MergeRequestChanges)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &changes)
	}

	return changes, meta, err
}

/*
Creates a new merge request.

    POST /projects/:id/merge_requests

Parameters:

    id               The ID of a project

*/
func (g *Gitlab) AddMergeRequest(req *AddMergeRequestRequest) (*MergeRequest, error) {
	u := g.ResourceUrl(projectMergeRequestsUrl, map[string]string{
		":id": string(req.TargetProjectId),
	})

	encodedRequest, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	data, _, err := g.buildAndExecRequest("POST", u.String(), encodedRequest)
	if err != nil {
		return nil, err
	}

	mr := new(MergeRequest)
	err = json.Unmarshal(data, mr)
	if err != nil {
		panic(err)
	}
	return mr, nil
}

/*
Updates an existing merge request.

    PUT /projects/:id/merge_request/:merge_request_id

Parameters:

    id               The ID of a project

*/
func (g *Gitlab) EditMergeRequest(mr *MergeRequest) error {
	u := g.ResourceUrl(projectMergeRequestUrl, map[string]string{
		":id":               string(mr.ProjectId),
		":merge_request_id": string(mr.Id),
	})

	encodedRequest, err := json.Marshal(mr)
	if err != nil {
		return err
	}

	data, _, err := g.buildAndExecRequest("PUT", u.String(), encodedRequest)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, mr)
	if err != nil {
		panic(err)
	}

	return nil
}

/*
Merge changes submitted with MR.

    PUT /projects/:id/merge_request/:merge_request_id/merge

Parameters:

    id               The ID of a project
    merge_request_id The ID of a merge request

*/
func (g *Gitlab) ProjectMergeRequestAccept(id, merge_request_id string, req *AcceptMergeRequestRequest) (*MergeRequest, error) {
	u := g.ResourceUrl(projectMergeRequestMergeUrl, map[string]string{
		":id":               id,
		":merge_request_id": merge_request_id,
	})

	encodedRequest, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	data, _, err := g.buildAndExecRequest("PUT", u.String(), encodedRequest)
	if err != nil {
		return nil, err
	}

	mr := new(MergeRequest)
	err = json.Unmarshal(data, mr)
	if err != nil {
		panic(err)
	}
	return mr, nil
}

/*
Cancel Merge When Build Succeeds.

    PUT /projects/:id/merge_request/:merge_request_id/cancel_merge_when_build_succeeds

Parameters:

    id               The ID of a project
    merge_request_id The ID of a merge request

*/
func (g *Gitlab) ProjectMergeRequestCancelMerge(id, merge_request_id string) (*MergeRequest, *ResponseMeta, error) {
	u := g.ResourceUrl(projectMergeRequestCancelMergeUrl, map[string]string{
		":id":               id,
		":merge_request_id": merge_request_id,
	})

	data, meta, err := g.buildAndExecRequest("PUT", u.String(), []byte{})
	if err != nil {
		return nil, meta, err
	}

	mr := new(MergeRequest)
	err = json.Unmarshal(data, mr)
	if err != nil {
		panic(err)
	}

	return mr, meta, nil
}