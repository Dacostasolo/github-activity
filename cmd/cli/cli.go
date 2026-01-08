package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/Dacostasolo/github-activity/cmd/githubactivity"
)

func Run() {
	if len(os.Args) < 1 {
		showUsageMessage()
		return
	}

	flag.Usage = showUsageMessage

	flag.Parse()
	username := os.Args[1]

	events, err := githubactivity.GetUserActivity(username)
	if err != nil {
		panic(err)
	}

	// - Pushed 3 commits to kamranahmedse/developer-roadmap
	// - Opened a new issue in kamranahmedse/developer-roadmap
	// - Starred kamranahmedse/developer-roadmap

	fmt.Println("Output:")
	for _, event := range events {
		fmt.Println(event)
	}
}
