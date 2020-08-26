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

package validator

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"
)

var noContext = context.Background()

func TestErrSkip(t *testing.T) {
	client := http.Client{}
	client.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		buf := bytes.NewBuffer(nil)
		buf.WriteString("skip")
		return &http.Response{
			Body:       ioutil.NopCloser(buf),
			StatusCode: 498,
		}, nil
	})

	plugin := Client("http://localhost", "top-secret", false)
	plugin.(*pluginClient).client.Client = &client

	err := plugin.Validate(noContext, &Request{})
	if err != ErrSkip {
		t.Errorf("Expect skip error, got %v", err)
	}
}

type roundTripFunc func(r *http.Request) (*http.Response, error)

func (s roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return s(r)
}
