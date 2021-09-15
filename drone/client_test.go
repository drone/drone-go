// Copyright 2018 Drone.IO Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package drone

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//
// user tests.
//

func TestSelf(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.Self()
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/user.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := new(User)
	err = json.Unmarshal(in, want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

func TestUser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.User("octocat")
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/user.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := new(User)
	err = json.Unmarshal(in, want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

func TestUserList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.UserList()
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/users.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := []*User{}
	err = json.Unmarshal(in, &want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

func TestUserDelete(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	err := client.UserDelete("octocat")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestUserCreate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.UserCreate(&User{})
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/user.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := new(User)
	err = json.Unmarshal(in, want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

func TestUserUpdate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.UserUpdate("octocat", &UserPatch{})
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/user.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := new(User)
	err = json.Unmarshal(in, want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

//
// repos
//

func TestRepo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.Repo("octocat", "hello-world")
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/repo.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := new(Repo)
	err = json.Unmarshal(in, want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

func TestRepoList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.RepoList()
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/repos.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := []*Repo{}
	err = json.Unmarshal(in, &want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

func TestRepoListSync(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.RepoListSync()
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/repos.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := []*Repo{}
	err = json.Unmarshal(in, &want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

func TestRepoEnable(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.RepoEnable("octocat", "hello-world")
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/repo.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := new(Repo)
	err = json.Unmarshal(in, want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

func TestRepoDisable(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	err := client.RepoDisable("octocat", "hello-world")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestRepoRepair(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	err := client.RepoRepair("octocat", "hello-world")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestRepoChown(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.RepoChown("octocat", "hello-world")
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/repo.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := new(Repo)
	err = json.Unmarshal(in, want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

func TestRepoUpdate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.RepoUpdate("octocat", "hello-world", &RepoPatch{})
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/repo.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := new(Repo)
	err = json.Unmarshal(in, want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

//
// cron jobs
//

func TestCron(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.Cron("octocat", "hello-world", "nightly")
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/cron.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := new(Cron)
	err = json.Unmarshal(in, want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

func TestCronList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.CronList("octocat", "hello-world")
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/crons.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := []*Cron{}
	err = json.Unmarshal(in, &want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

// func TestCronDisable(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
// 	defer ts.Close()

// 	client := New(ts.URL)
// 	err := client.CronDisable("octocat", "hello-world", "nightly")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// }

// func TestCronEnable(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
// 	defer ts.Close()

// 	client := New(ts.URL)
// 	err := client.CronEnable("octocat", "hello-world", "nightly")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// }

//
// builds
//

func TestBuild(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.Build("octocat", "hello-world", 1)
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/build.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := new(Build)
	err = json.Unmarshal(in, want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

func TestBuildLast(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.BuildLast("octocat", "hello-world", "master")
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/build.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := new(Build)
	err = json.Unmarshal(in, want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

func TestBuildList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.BuildList("octocat", "hello-world", ListOptions{})
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/builds.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := []*Build{}
	err = json.Unmarshal(in, &want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

// func TestBuildQueue(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
// 	defer ts.Close()

// 	client := New(ts.URL)
// 	got, err := client.BuildQueue()
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	in, err := ioutil.ReadFile("testdata/builds.json.golden")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	want := []*Build{}
// 	err = json.Unmarshal(in, &want)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	if diff := cmp.Diff(got, want); diff != "" {
// 		t.Errorf("Unexpected response")
// 		t.Log(diff)
// 	}
// }

func TestBuildRestart(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.BuildRestart("octocat", "hello-world", 99, nil)
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/build.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := new(Build)
	err = json.Unmarshal(in, want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

func TestBuildCancel(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	err := client.BuildCancel("octocat", "hello-world", 1)
	if err != nil {
		t.Error(err)
	}
}

func TestApprove(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	err := client.Approve("octocat", "hello-world", 1, 2)
	if err != nil {
		t.Error(err)
	}
}

func TestDecline(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	err := client.Decline("octocat", "hello-world", 1, 3)
	if err != nil {
		t.Error(err)
	}
}

//
// logs
//

func TestLogs(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	got, err := client.Logs("octocat", "hello-world", 1, 2, 3)
	if err != nil {
		t.Error(err)
		return
	}

	in, err := ioutil.ReadFile("testdata/logs.json.golden")
	if err != nil {
		t.Error(err)
		return
	}
	want := []*Line{}
	err = json.Unmarshal(in, &want)
	if err != nil {
		t.Error(err)
		return
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected response")
		t.Log(diff)
	}
}

func TestLogsPurge(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer ts.Close()

	client := New(ts.URL)
	err := client.LogsPurge("octocat", "hello-world", 1, 2, 3)
	if err != nil {
		t.Error(err)
		return
	}
}

//
// mock server and testdata.
//
func mockHandler(w http.ResponseWriter, r *http.Request) {
	routes := []struct {
		verb string
		path string
		body string
		code int
	}{
		//
		// users
		//
		{
			verb: "GET",
			path: "/api/user",
			body: "testdata/user.json",
			code: 200,
		},
		{
			verb: "GET",
			path: "/api/users/octocat",
			body: "testdata/user.json",
			code: 200,
		},
		{
			verb: "DELETE",
			path: "/api/users/octocat",
			code: 204,
		},
		{
			verb: "POST",
			path: "/api/users",
			body: "testdata/user.json",
			code: 200,
		},
		{
			verb: "PATCH",
			path: "/api/users/octocat",
			body: "testdata/user.json",
			code: 200,
		},
		{
			verb: "GET",
			path: "/api/users",
			body: "testdata/users.json",
			code: 200,
		},
		//
		// repos
		//
		{
			verb: "GET",
			path: "/api/repos/octocat/hello-world",
			body: "testdata/repo.json",
			code: 200,
		},
		{
			verb: "GET",
			path: "/api/user/repos",
			body: "testdata/repos.json",
			code: 200,
		},
		{
			verb: "POST",
			path: "/api/user/repos",
			body: "testdata/repos.json",
			code: 200,
		},
		{
			verb: "POST",
			path: "/api/repos/octocat/hello-world/repair",
			code: 204,
		},
		{
			verb: "POST",
			path: "/api/repos/octocat/hello-world/chown",
			body: "testdata/repo.json",
			code: 200,
		},
		{
			verb: "PATCH",
			path: "/api/repos/octocat/hello-world",
			body: "testdata/repo.json",
			code: 200,
		},
		{
			verb: "POST",
			path: "/api/repos/octocat/hello-world",
			body: "testdata/repo.json",
			code: 200,
		},
		{
			verb: "DELETE",
			path: "/api/repos/octocat/hello-world",
			code: 204,
		},
		//
		// crons
		//
		{
			verb: "GET",
			path: "/api/repos/octocat/hello-world/cron/nightly",
			body: "testdata/cron.json",
			code: 200,
		},
		{
			verb: "GET",
			path: "/api/repos/octocat/hello-world/cron",
			body: "testdata/crons.json",
			code: 200,
		},
		{
			verb: "POST",
			path: "/api/repos/octocat/hello-world/cron/nightly",
			code: 204,
		},
		{
			verb: "DELETE",
			path: "/api/repos/octocat/hello-world/cron/nightly",
			code: 204,
		},
		//
		// builds
		//
		{
			verb: "GET",
			path: "/api/system/builds",
			body: "testdata/builds.json",
			code: 200,
		},
		{
			verb: "GET",
			path: "/api/repos/octocat/hello-world/builds",
			body: "testdata/builds.json",
			code: 200,
		},
		{
			verb: "GET",
			path: "/api/repos/octocat/hello-world/builds/1",
			body: "testdata/build.json",
			code: 200,
		},
		{
			verb: "GET",
			path: "/api/repos/octocat/hello-world/builds/latest",
			body: "testdata/build.json",
			code: 200,
		},
		{
			verb: "POST",
			path: "/api/repos/octocat/hello-world/builds/99",
			body: "testdata/build.json",
			code: 200,
		},
		{
			verb: "DELETE",
			path: "/api/repos/octocat/hello-world/builds/1",
			code: 204,
		},
		{
			verb: "POST",
			path: "/api/repos/octocat/hello-world/builds/1/approve/2",
			code: 204,
		},
		{
			verb: "POST",
			path: "/api/repos/octocat/hello-world/builds/1/decline/3",
			code: 204,
		},
		//
		// logs
		//
		{
			verb: "GET",
			path: "/api/repos/octocat/hello-world/builds/1/logs/2/3",
			body: "testdata/logs.json",
			code: 200,
		},
		{
			verb: "DELETE",
			path: "/api/repos/octocat/hello-world/builds/1/logs/2/3",
			code: 204,
		},
	}

	path := r.URL.Path
	verb := r.Method
	for _, route := range routes {
		if route.verb != verb {
			continue
		}
		if route.path != path {
			continue
		}
		if route.code == 204 {
			w.WriteHeader(204)
			return
		}
		body, err := ioutil.ReadFile(route.body)
		if err != nil {
			break
		}
		w.WriteHeader(route.code)
		_, _ = w.Write(body)
		return
	}
	w.WriteHeader(404)
}
