package drone

import (
	"encoding/json"
	"testing"
)

func TestParse(t *testing.T) {
	// Sample from http://readme.drone.io/api/build-endpoint/
	sampleBuildResponse := `{
		"id": 1,
		"number": 1,
		"event": "push",
		"status": "success",
		"created_at": 1443677151,
		"enqueued_at": 1443677151,
		"started_at": 1443677151,
		"finished_at": 1443677255,
		"commit": "2deb7e0d0cbac357eeb110c8a2f2f32ce037e0d5",
		"branch": "master",
		"ref": "refs/heads/master",
		"remote": "https://github.com/octocat/hello-world.git",
		"message": "New line at end of file. --Signed off by Spaceghost",
		"timestamp": 1443677255,
		"author": "Spaceghost",
		"author_avatar": "https://avatars0.githubusercontent.com/u/251370?v=3",
		"author_email": "octocat@github.com",
		"link_url": "https://github.com/octocat/hello-world/commit/762941318ee16e59dabbacb1b4049eec22f0d303",
		"jobs": [
		  {
			"id": 1,
			"number": 1,
			"status": "success",
			"enqueued_at": 1443677151,
			"started_at": 1443677151,
			"finished_at": 1443677255,
			"exit_code": 0,
			"environment": { "GO_VERSION": "1.4" }
		  },
		  {
			"id": 2,
			"number": 2,
			"status": "success",
			"enqueued_at": 1443677151,
			"started_at": 1443677151,
			"finished_at": 1443677255,
			"exit_code": 0,
			"environment": { "GO_VERSION": "1.5" }
		  }
		]
	  }
	  `

	var build Build
	err := json.Unmarshal([]byte(sampleBuildResponse), &build)
	if err != nil {
		t.Errorf("unable to unmarshal: %s", err)
	}

	if len(build.Jobs) != 2 {
		t.Errorf("error unmarshaling jobs, expected len %d, got %d", 2, len(build.Jobs))
	}
}
