package githubactivity

import (
	"fmt"
	"strings"
)

func ToUserAction(e Event) UserAction {
	switch e.Type {
	case "PushEvent":
		var ref string
		if e.Payload.Push != nil {
			ref = strings.TrimPrefix(e.Payload.Push.Ref, "refs/heads/")
		}
		return UserAction{
			Verb:    "pushed",
			Subject: "commits",
			Repo:    e.Repo.Name,
			Ref:     ref,
		}
	case "PullRequestEvent":
		action := "made"
		if e.Payload.PullRequest != nil {
			action = e.Payload.PullRequest.Action
		}
		return UserAction{
			Verb:    action,
			Subject: "pull request",
			Repo:    e.Repo.Name,
		}
	case "CreateEvent":
		if e.Payload.Create == nil {
			break
		}
		return UserAction{
			Verb:    "created",
			Subject: e.Payload.Create.RefType,
			Repo:    e.Repo.Name,
			Ref:     e.Payload.Create.Ref,
		}
	case "WatchEvent":
		return UserAction{
			Verb:    "starred",
			Subject: "repository",
			Repo:    e.Repo.Name,
		}
	case "MemberEvent":
		if e.Payload.Member == nil {
			break
		}
		return UserAction{
			Verb:    "added",
			Subject: "a collaborator to",
			Repo:    e.Repo.Name,
		}
	}
	return UserAction{
		Verb: "performed an action in",
		Repo: e.Repo.Name,
	}
}

func (a UserAction) String() string {
	var parts []string
	if a.Verb != "" {
		parts = append(parts, a.Verb)
	}
	if a.Count > 0 {
		parts = append(parts, fmt.Sprintf("%d", a.Count))
	}
	if a.Subject != "" {
		parts = append(parts, a.Subject)
	}
	if a.Repo != "" {
		parts = append(parts, "in "+a.Repo)
	}
	if a.Ref != "" {
		parts = append(parts, "("+a.Ref+")")
	}
	if a.Extra != "" {
		parts = append(parts, a.Extra)
	}
	return strings.Join(parts, " ")
}
