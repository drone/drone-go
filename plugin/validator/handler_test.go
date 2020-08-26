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
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/drone/drone-go/drone"

	"github.com/99designs/httpsignatures-go"
)

func TestHandler(t *testing.T) {
	key := "xVKAGlWQiY3sOp8JVc0nbuNId3PNCgWh"

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&Request{
		Config: drone.Config{
			Data: "{kind: pipeline, type: docker}",
		},
	})

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", buf)
	req.Header.Add("Date", time.Now().UTC().Format(http.TimeFormat))

	err := httpsignatures.DefaultSha256Signer.AuthRequest("hmac-key", key, req)
	if err != nil {
		t.Error(err)
		return
	}

	plugin := &mockPlugin{
		err: nil,
	}

	handler := Handler(key, plugin, nil)
	handler.ServeHTTP(res, req)

	if got, want := res.Code, 204; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestHandler_Error(t *testing.T) {
	key := "xVKAGlWQiY3sOp8JVc0nbuNId3PNCgWh"

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&Request{
		Config: drone.Config{
			Data: "{kind: pipeline, type: docker}",
		},
	})

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", buf)
	req.Header.Add("Date", time.Now().UTC().Format(http.TimeFormat))

	err := httpsignatures.DefaultSha256Signer.AuthRequest("hmac-key", key, req)
	if err != nil {
		t.Error(err)
		return
	}

	err = errors.New("insufficient permission to mount volumes")
	plugin := &mockPlugin{err: err}

	handler := Handler(key, plugin, nil)
	handler.ServeHTTP(res, req)

	if got, want := res.Code, 400; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}

	resp := &drone.Error{}
	json.Unmarshal(res.Body.Bytes(), resp)
	if got, want := resp.Message, err.Error(); got != want {
		t.Errorf("Want validation error %s, got %s", want, got)
	}
}

func TestHandler_ErrorSkip(t *testing.T) {
	key := "xVKAGlWQiY3sOp8JVc0nbuNId3PNCgWh"

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&Request{
		Config: drone.Config{
			Data: "{kind: pipeline, type: docker}",
		},
	})

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", buf)
	req.Header.Add("Date", time.Now().UTC().Format(http.TimeFormat))

	err := httpsignatures.DefaultSha256Signer.AuthRequest("hmac-key", key, req)
	if err != nil {
		t.Error(err)
		return
	}

	err = ErrSkip
	plugin := &mockPlugin{err: err}

	handler := Handler(key, plugin, nil)
	handler.ServeHTTP(res, req)

	if got, want := res.Code, 498; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestHandler_ErrorBlock(t *testing.T) {
	key := "xVKAGlWQiY3sOp8JVc0nbuNId3PNCgWh"

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&Request{
		Config: drone.Config{
			Data: "{kind: pipeline, type: docker}",
		},
	})

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", buf)
	req.Header.Add("Date", time.Now().UTC().Format(http.TimeFormat))

	err := httpsignatures.DefaultSha256Signer.AuthRequest("hmac-key", key, req)
	if err != nil {
		t.Error(err)
		return
	}

	err = ErrBlock
	plugin := &mockPlugin{err: err}

	handler := Handler(key, plugin, nil)
	handler.ServeHTTP(res, req)

	if got, want := res.Code, 499; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestHandler_MissingSignature(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	handler := Handler("xVKAGlWQiY3sOp8JVc0nbuNId3PNCgWh", nil, nil)
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

	handler := Handler("xVKAGlWQiY3sOp8JVc0nbuNId3PNCgWh", nil, nil)
	handler.ServeHTTP(res, req)

	got, want := res.Body.String(), "Invalid Signature\n"
	if got != want {
		t.Errorf("Want response body %q, got %q", want, got)
	}
}

type mockPlugin struct {
	err error
}

func (m *mockPlugin) Validate(ctx context.Context, req *Request) error {
	return m.err
}
