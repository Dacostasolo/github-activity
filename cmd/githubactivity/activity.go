package githubactivity

// Event represents a GitHub event.
type Event struct {
	ID      string   `json:"id"`
	Type    string  `json:"type"`
	Repo    Repo    `json:"repo"`
	Payload Payload `json:"payload"`
}

// Repo represents a GitHub repository.
type Repo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Payload represents the payload of a GitHub event.
// Only ONE field is set depending on Event.Type
type Payload struct {
	Push        *PushPayload        `json:"push,omitempty"`
	Create      *CreatePayload      `json:"create,omitempty"`
	PullRequest *PullRequestPayload `json:"pull_request,omitempty"`
	Member      *MemberPayload      `json:"member,omitempty"`
	Watch       *WatchPayload       `json:"watch,omitempty"`
}

// PushPayload represents the payload of a PushEvent.
type PushPayload struct {
	Ref          string `json:"ref"`
	Before       string `json:"before"`
	Head         string `json:"head"`
	PushID       string  `json:"push_id"`
	RepositoryID string  `json:"repository_id"`
}

// CreatePayload represents the payload of a CreateEvent.
type CreatePayload struct {
	Ref          string `json:"ref"`
	RefType      string `json:"ref_type"` // branch | tag
	MasterBranch string `json:"master_branch"`
	PusherType   string `json:"pusher_type"`
}

// PullRequestPayload represents the payload of a PullRequestEvent.
type PullRequestPayload struct {
	Action string      `json:"action"` // opened | closed | merged
	Number int         `json:"number"`
	PR     PullRequest `json:"pull_request"`
}

// PullRequest represents a GitHub pull request.
type PullRequest struct {
	ID     string `json:"id"`
	Number int   `json:"number"`
	URL    string `json:"url"`

	Base PRBranch `json:"base"`
	Head PRBranch `json:"head"`
}

// PRBranch represents a branch in a pull request.
type PRBranch struct {
	Ref  string `json:"ref"`
	SHA  string `json:"sha"`
	Repo Repo   `json:"repo"`
}

// MemberPayload represents the payload of a MemberEvent.
type MemberPayload struct {
	Action string `json:"action"` // added
	Member User   `json:"member"`
}

// User represents a GitHub user.
type User struct {
	ID    string  `json:"id"`
	Login string `json:"login"`
	URL   string `json:"html_url"`
	Type  string `json:"type"`
}

// WatchPayload represents the payload of a WatchEvent.
type WatchPayload struct {
	Action string `json:"action"`
}

// UserAction represents a user action. It is a simplified representation of an event.
type UserAction struct {
	Verb       string // Pushed, Opened, Starred, Merged, Created, Added
	Subject    string // commits, issue, pull request, branch, repository
	Count      int    // optional (e.g. commits count)
	Repo       string // owner/name
	Ref        string // branch name (optional)
	Extra      string // free-form extra info (optional)
}
