package githubactivity

import (
	"encoding/json"
	"fmt"
	"net/http"
)



func GetUserActivity(username string) ([]Event, error) {
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET",
		"https://api.github.com/users/"+username+"/events",
		nil,
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "github-activity-cli")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var events []Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, err
	}

	
	return events, nil
}

func (p Payload) String() (string) {
	switch {
	case p.Push != nil:
		return fmt.Sprintf("Push: %s", p.Push.Ref)
	case p.Create != nil:
		return fmt.Sprintf("Create: %s", p.Create.RefType)
	case p.PullRequest != nil:
		return fmt.Sprintf("Pull Request: %d", p.PullRequest.Number)
	case p.Member != nil:
		return fmt.Sprintf("Member: %s", p.Member.Action)
	case p.Watch != nil:
		return fmt.Sprintf("Watch: %s", p.Watch.Action)
	default:
		return "Unknown"
	}
}

func (e Event) String() string {
	action := ToUserAction(e)
	return fmt.Sprintf("- %s", action.String())
}

