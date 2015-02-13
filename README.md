# drone-go

drone-go is a Go client library for accessing the Drone [API](http://readme.drone.io/api/overview/).


Download the package using `go get`:

```Go
go get "github.com/drone/drone-go/drone"
```

Import the package:

```Go
import "github.com/drone/drone-go/drone"
```

Create the client:

```Go
token := "my-user-token"
url := "https://my-drone-url.com"
client := drone.NewClient(token, url)
```

Get the current user:

```Go
user, err := client.Users.GetCurrent()
fmt.Println(user)
```

Get the repository list:

```Go
repos, err := client.Repos.List()
fmt.Println(repos)
```

Get the named repository:

```Go
repo, err := client.Repos.Get("github.com", "drone", "drone-go")
fmt.Println(repo)
```