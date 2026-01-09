package githubactivity

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const githubAPIVersion = "2022-11-28"

func GetUserActivity(username string) ([]Event, error) {
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET",
		"https://api.github.com/users/"+username+"/events",
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "github-activity-cli")
	req.Header.Set("X-GitHub-Api-Version", githubAPIVersion)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: status code %d", resp.StatusCode)
	}

	var events []Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return events, nil
}

func (p Payload) String() string {
	switch {
	case p.Push != nil:
		return fmt.Sprintf("pushed to %s", p.Push.Ref)
	case p.Create != nil:
		return fmt.Sprintf("created %s", p.Create.RefType)
	case p.PullRequest != nil:
		return fmt.Sprintf("%s pull request #%d", p.PullRequest.Action, p.PullRequest.Number)
	case p.Member != nil:
		return fmt.Sprintf("%s member %s", p.Member.Action, p.Member.Member.Login)
	case p.Watch != nil:
		return fmt.Sprintf("%s repository", p.Watch.Action)
	default:
		return "did something"
	}
}

func (e Event) String() string {
	action := ToUserAction(e)
	return fmt.Sprintf("- %s", action.String())
}
