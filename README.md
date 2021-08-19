# drone-go

[![Go.dev](https://pkg.go.dev/badge/github.com/drone/drone-go)](https://pkg.go.dev/github.com/drone/drone-go?tab=doc)

```Go
package main

import (
	"fmt"

	"github.com/drone/drone-go/drone"
	"golang.org/x/oauth2"
)

const (
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	host  = "http://drone.company.com"
)

func main() {
	// create an http client with oauth authentication.
	config := new(oauth2.Config)
	auther := config.Client(
		oauth2.NoContext,
		&oauth2.Token{
			AccessToken: token,
		},
	)

	// create the drone client with authenticator
	client := drone.NewClient(host, auther)

	// gets the current user
	user, err := client.Self()
	fmt.Println(user, err)

	// gets the named repository information
	repo, err := client.Repo("drone", "drone-go")
	fmt.Println(repo, err)
}
```
## Release procedure

Run the changelog generator.

```BASH
docker run -it --rm -v "$(pwd)":/usr/local/src/your-app githubchangeloggenerator/github-changelog-generator -u drone -p drone-go -t <secret github token>
```

You can generate a token by logging into your GitHub account and going to Settings -> Personal access tokens.

Next we tag the PR's with the fixes or enhancements labels. If the PR does not fufil the requirements, do not add a label.

Run the changelog generator again with the future version according to semver.

```BASH
docker run -it --rm -v "$(pwd)":/usr/local/src/your-app githubchangeloggenerator/github-changelog-generator -u drone -p drone-go -t <secret token> --future-release v1.0.0
```

Create your pull request for the release. Get it merged then tag the release.