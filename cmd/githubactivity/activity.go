package githubactivity

// ===== Root Event =====

type Event struct {
	ID      string   `json:"id"`
	Type    string  `json:"type"`
	Repo    Repo    `json:"repo"`
	Payload Payload `json:"payload"`
}

// ===== Repo =====

type Repo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// ===== Payload Dispatcher =====
// Only ONE field is set depending on Event.Type

type Payload struct {
	Push        *PushPayload        `json:"push,omitempty"`
	Create      *CreatePayload      `json:"create,omitempty"`
	PullRequest *PullRequestPayload `json:"pull_request,omitempty"`
	Member      *MemberPayload      `json:"member,omitempty"`
	Watch       *WatchPayload       `json:"watch,omitempty"`
}

// ===== PushEvent =====

type PushPayload struct {
	Ref          string `json:"ref"`
	Before       string `json:"before"`
	Head         string `json:"head"`
	PushID       string  `json:"push_id"`
	RepositoryID string  `json:"repository_id"`
}

// ===== CreateEvent =====

type CreatePayload struct {
	Ref          string `json:"ref"`
	RefType      string `json:"ref_type"` // branch | tag
	MasterBranch string `json:"master_branch"`
	PusherType   string `json:"pusher_type"`
}

// ===== PullRequestEvent =====

type PullRequestPayload struct {
	Action string      `json:"action"` // opened | closed | merged
	Number int         `json:"number"`
	PR     PullRequest `json:"pull_request"`
}

type PullRequest struct {
	ID     string `json:"id"`
	Number int   `json:"number"`
	URL    string `json:"url"`

	Base PRBranch `json:"base"`
	Head PRBranch `json:"head"`
}

type PRBranch struct {
	Ref  string `json:"ref"`
	SHA  string `json:"sha"`
	Repo Repo   `json:"repo"`
}

// ===== MemberEvent =====

type MemberPayload struct {
	Action string `json:"action"` // added
	Member User   `json:"member"`
}

type User struct {
	ID    string  `json:"id"`
	Login string `json:"login"`
	URL   string `json:"html_url"`
	Type  string `json:"type"`
}

// ===== WatchEvent =====

type WatchPayload struct {
	Action string `json:"action"`
}

type UserAction struct {
	Verb       string // Pushed, Opened, Starred, Merged, Created, Added
	Subject    string // commits, issue, pull request, branch, repository
	Count      int    // optional (e.g. commits count)
	Repo       string // owner/name
	Ref        string // branch name (optional)
	Extra      string // free-form extra info (optional)
}


