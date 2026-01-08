package githubactivity

import (
	"fmt"
	"strings"
)

func ToUserAction(e Event) UserAction {
	switch e.Type {

	case "PushEvent":
		if e.Payload.Push == nil {
			break
		}

		ref := strings.TrimPrefix(e.Payload.Push.Ref, "refs/heads/")
		return UserAction{
			Verb:    "Pushed",
			Subject: "commits",
			Repo:    e.Repo.Name,
			Ref:     ref,
		}

	case "PullRequestEvent":
		if e.Payload.PullRequest == nil {
			break
		}

		action := e.Payload.PullRequest.Action
		verb := strings.Title(action)

		return UserAction{
			Verb:    verb,
			Subject: "Pull request",
			Repo:    e.Repo.Name,
		}

	case "CreateEvent":
		if e.Payload.Create == nil {
			break
		}

		return UserAction{
			Verb:    "Created",
			Subject: e.Payload.Create.RefType,
			Repo:    e.Repo.Name,
		}

	case "WatchEvent":
		if e.Payload.Watch == nil {
			break
		}

		return UserAction{
			Verb: "Starred",
			Repo: e.Repo.Name,
		}

	case "MemberEvent":
		if e.Payload.Member == nil {
			break
		}

		return UserAction{
			Verb:    "Added",
			Subject: "a collaborator to",
			Repo:    e.Repo.Name,
		}
	}

	// Safe fallback (never panic)
	return UserAction{
		Verb: "Performed an action in",
		Repo: e.Repo.Name,
	}
}


func (a UserAction) String() string {
	switch {
	case a.Count > 0 && a.Ref != "":
		return fmt.Sprintf(
			"%s %d %s to %s (%s)",
			a.Verb, a.Count, a.Subject, a.Repo, a.Ref,
		)

	case a.Count > 0:
		return fmt.Sprintf(
			"%s %d %s to %s",
			a.Verb, a.Count, a.Subject, a.Repo,
		)

	case a.Extra != "":
		return fmt.Sprintf(
			"%s %s in %s",
			a.Verb, a.Extra, a.Repo,
		)

	default:
		return fmt.Sprintf(
			"%s %s %s",
			a.Verb, a.Subject, a.Repo,
		)
	}
}
