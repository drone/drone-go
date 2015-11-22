package template

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/aymerick/raymond"
)

func init() {
	raymond.RegisterHelpers(funcs)
}

// Render parses and executes a template, returning the results
// in string format.
func Render(template string, args interface{}) (string, error) {

	return raymond.Render(template, args)
}

// Render parses and executes a template, returning the results
// in string format. The result is trimmed to remove left and right
// padding and newlines that may be added unintentially in the
// template markup.
func RenderTrim(template string, args interface{}) (string, error) {
	out, err := Render(template, normalize(args))
	return strings.Trim(out, " \n"), err
}

// Write parses and executes a template, writing the results to
// writer w.
func Write(w io.Writer, template string, args interface{}) error {
	out, err := Render(template, args)
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, out)
	return err
}

var funcs = map[string]interface{}{
	"uppercase": strings.ToUpper,
	"lowercase": strings.ToLower,
	"duration":  toDuration,
	"datetime":  toDatetime,
	"success":   isSuccess,
	"failure":   isFailure,
}

// toDuration is a helper function that calculates a duration for a start and
// and end time, and returns the duration in string format.
func toDuration(started, finished float64) string {
	dur := time.Duration(int64(finished - started))
	return fmt.Sprintln(dur)
}

// toDatetime is a helper function that converts a unix timestamp to a string.
func toDatetime(timestamp float64, layout string) string {
	return time.Unix(int64(timestamp), 0).Format(layout)
}

// isSuccess is a helper function that executes a block iff the status
// is success, else it executes the else block.
func isSuccess(conditional bool, options *raymond.Options) string {
	if !conditional {
		return options.Inverse()
	}

	switch options.ParamStr(0) {
	case "success":
		return options.Fn()
	default:
		return options.Inverse()
	}
}

// isFailure is a helper function that executes a block iff the status
// is a form of failure, else it executes the else block.
func isFailure(conditional bool, options *raymond.Options) string {
	if !conditional {
		return options.Inverse()
	}

	switch options.ParamStr(0) {
	case "failure", "error", "killed":
		return options.Fn()
	default:
		return options.Inverse()
	}
}

// normalize takes a Go representation of the variable, marshals
// to json and then unmarshals to a map[string]interfacce{}. This
// is important because it let's us use the JSON variable names
// in our template
func normalize(in interface{}) map[string]interface{} {
	data, _ := json.Marshal(in) // we own the types, so this should never fail

	out := map[string]interface{}{}
	json.Unmarshal(data, &out)
	return out
}
