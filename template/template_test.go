package template

import (
	"testing"

	"github.com/drone/drone-go/drone"
)

var tests = []struct {
	Build  drone.Build
	Input  string
	Output string
}{
	{
		drone.Build{Number: 1},
		"build #{{number}}",
		"build #1",
	},
	{
		drone.Build{Status: drone.StatusSuccess},
		"{{uppercase status}}",
		"SUCCESS",
	},
	{
		drone.Build{Author: "Octocat"},
		"{{lowercase author}}",
		"octocat",
	},
	{
		drone.Build{Status: drone.StatusSuccess},
		"{{uppercasefirst status}}",
		"Success",
	},
	{
		drone.Build{Started: 1448127131, Finished: 1448127505},
		"{{ duration started_at finished_at }}",
		"374ns",
	},
	{
		drone.Build{Finished: 1448127505},
		`finished at {{ datetime finished_at "Mon Jan _2 15:04:05 2006" }}`,
		"finished at Sat Nov 21 11:38:25 2015",
	},
	// verify the success if / else block works
	{
		drone.Build{Status: drone.StatusSuccess},
		"{{#success status}}SUCCESS{{/success}}",
		"SUCCESS",
	},
	{
		drone.Build{Status: drone.StatusFailure},
		"{{#success status}}SUCCESS{{/success}}",
		"",
	},
	{
		drone.Build{Status: drone.StatusFailure},
		"{{#success status}}SUCCESS{{else}}NOT SUCCESS{{/success}}",
		"NOT SUCCESS",
	},
	// verify the failure if / else block works
	{
		drone.Build{Status: drone.StatusFailure},
		"{{#failure status}}FAILURE{{/failure}}",
		"FAILURE",
	},
	{
		drone.Build{Status: drone.StatusSuccess},
		"{{#failure status}}FAILURE{{/failure}}",
		"",
	},
	{
		drone.Build{Status: drone.StatusSuccess},
		"{{#failure status}}FAILURE{{else}}NOT FAILURE{{/failure}}",
		"NOT FAILURE",
	},
}

func TestTemplate(t *testing.T) {

	for _, test := range tests {
		got, err := RenderTrim(test.Input, &test.Build)
		if err != nil {
			t.Errorf("Failed rendering template %q, got error %s.", test.Input, err)
		}
		if got != test.Output {
			t.Errorf("Wanted rendered template %q, got %q", test.Output, got)
		}
	}
}
