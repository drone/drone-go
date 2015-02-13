# drone-go

drone-go is a Go client library for accessing the Drone [API](http://readme.drone.io/api/overview/).


```Go
go get github.com/drone/drone-go/drone
```


```Go
import "github.com/drone/drone-go/drone"
```

```Go
token := "my-user-token"
url := "https://my-drone-url.com"
client := drone.NewClient(token, url)
```