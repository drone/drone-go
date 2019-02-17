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

package admission

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/drone/drone-go/drone"

	"github.com/99designs/httpsignatures-go"
)

func TestHandler(t *testing.T) {
	key := "xVKAGlWQiY3sOp8JVc0nbuNId3PNCgWh"

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&Request{})

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", buf)
	req.Header.Add("Date", time.Now().UTC().Format(http.TimeFormat))

	err := httpsignatures.DefaultSha256Signer.AuthRequest("hmac-key", key, req)
	if err != nil {
		t.Error(err)
		return
	}

	want := &drone.User{
		Login: "octocat",
		Admin: false,
	}
	plugin := &mockPlugin{
		res: want,
		err: nil,
	}

	handler := Handler(plugin, key, nil)
	handler.ServeHTTP(res, req)

	if got, want := res.Code, 200; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}

	resp := &drone.User{}
	json.Unmarshal(res.Body.Bytes(), resp)
	if got, want := resp.Login, want.Login; got != want {
		t.Errorf("Want user data %q, got %q", want, got)
	}
}

func TestHandler_NoContent(t *testing.T) {
	key := "xVKAGlWQiY3sOp8JVc0nbuNId3PNCgWh"

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&Request{})

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", buf)
	req.Header.Add("Date", time.Now().UTC().Format(http.TimeFormat))

	err := httpsignatures.DefaultSha256Signer.AuthRequest("hmac-key", key, req)
	if err != nil {
		t.Error(err)
		return
	}

	plugin := &mockPlugin{
		res: nil,
		err: nil,
	}

	handler := Handler(plugin, key, nil)
	handler.ServeHTTP(res, req)

	if got, want := res.Code, 204; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestHandler_AccessDenied(t *testing.T) {
	key := "xVKAGlWQiY3sOp8JVc0nbuNId3PNCgWh"

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&Request{})

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", buf)
	req.Header.Add("Date", time.Now().UTC().Format(http.TimeFormat))

	err := httpsignatures.DefaultSha256Signer.AuthRequest("hmac-key", key, req)
	if err != nil {
		t.Error(err)
		return
	}

	plugin := &mockPlugin{
		res: nil,
		err: errors.New("access denied"),
	}

	handler := Handler(plugin, key, nil)
	handler.ServeHTTP(res, req)

	if got, want := res.Code, 403; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}

	got, want := strings.TrimSpace(res.Body.String()), plugin.err.Error()
	if got != want {
		t.Errorf("Want error %q, got %q", want, got)
	}
}

func TestHandler_MissingSignature(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	handler := Handler(nil, "xVKAGlWQiY3sOp8JVc0nbuNId3PNCgWh", nil)
	handler.ServeHTTP(res, req)

	got, want := res.Body.String(), "Invalid or Missing Signature\n"
	if got != want {
		t.Errorf("Want response body %q, got %q", want, got)
	}
}

func TestHandler_InvalidSignature(t *testing.T) {
	sig := `keyId="hmac-key",algorithm="hmac-sha256",signature="QrS16+RlRsFjXn5IVW8tWz+3ZRAypjpNgzehEuvJksk=",headers="(request-target) accept accept-encoding content-type date digest"`
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Signature", sig)

	handler := Handler(nil, "xVKAGlWQiY3sOp8JVc0nbuNId3PNCgWh", nil)
	handler.ServeHTTP(res, req)

	got, want := res.Body.String(), "Invalid Signature\n"
	if got != want {
		t.Errorf("Want response body %q, got %q", want, got)
	}
}

type mockPlugin struct {
	res *drone.User
	err error
}

func (m *mockPlugin) Admit(ctx context.Context, req *Request) (*drone.User, error) {
	return m.res, m.err
}
