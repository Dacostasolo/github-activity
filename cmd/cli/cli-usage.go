package cli

import "fmt"


func showUsageMessage()  {
	desc := `github-activity CLI
Usage:
  github-activity [username]
Description:
  Fetch and display GitHub activity for the specified username.
Arguments:
  username     GitHub username to fetch activity for.
Example:
  github-activity dacostaaboagye
`

fmt.Println(desc)
}
