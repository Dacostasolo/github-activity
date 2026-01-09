package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/Dacostasolo/github-activity/cmd/githubactivity"
)

func Run() {
	var helpFlag = flag.Bool("help", false, "Show help message")
	flag.Usage = showUsageMessage
	flag.Parse()

	if *helpFlag {
		showUsageMessage()
		return
	}

	username := flag.Arg(0)
	if username == "" {
		showUsageMessage()
		return
	}

	events, err := githubactivity.GetUserActivity(username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("GitHub activity for user %s:\n", username)
	for _, event := range events {
		fmt.Println(event.String())
	}
}
