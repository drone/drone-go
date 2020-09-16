# drone-go

[![Build Status](https://cloud.drone.io/api/badges/drone/drone-go/status.svg)](https://cloud.drone.io/drone/drone-go)

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
