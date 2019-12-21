// Copyright 2019 Drone.IO Inc.
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

package environ

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/drone/drone-go/plugin/internal/aesgcm"
	"github.com/drone/drone-go/plugin/logger"

	"github.com/99designs/httpsignatures-go"
)

// Handler returns a http.Handler that accepts JSON-encoded
// HTTP requests for environment variables, invokes the underlying
// plugin, and writes the JSON-encoded secret to the HTTP response.
//
// The handler verifies the authenticity of the HTTP request
// using the http-signature, and returns a 400 Bad Request if
// the signature is missing or invalid.
//
// The handler can optionally encrypt the response body using
// aesgcm if the HTTP request includes the Accept-Encoding header
// set to aesgcm.
func Handler(secret string, plugin Plugin, logs logger.Logger) http.Handler {
	handler := &handler{
		secret: secret,
		plugin: plugin,
		logger: logs,
	}
	if handler.logger == nil {
		handler.logger = logger.Discard()
	}
	return handler
}

type handler struct {
	secret string
	plugin Plugin
	logger logger.Logger
}

func (p *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	signature, err := httpsignatures.FromRequest(r)
	if err != nil {
		p.logger.Debugf("environment: invalid or missing signature in http.Request")
		http.Error(w, "Invalid or Missing Signature", 400)
		return
	}
	if !signature.IsValid(p.secret, r) {
		p.logger.Debugf("environment: invalid signature in http.Request")
		http.Error(w, "Invalid Signature", 400)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		p.logger.Debugf("environment: cannot read http.Request body")
		w.WriteHeader(400)
		return
	}

	req := &Request{}
	err = json.Unmarshal(body, req)
	if err != nil {
		p.logger.Debugf("environment: cannot unmarshal http.Request body")
		http.Error(w, "Invalid Input", 400)
		return
	}

	auths, err := p.plugin.List(r.Context(), req)
	if err != nil {
		p.logger.Debugf("environment: cannot list registries: %s", err)
		http.Error(w, err.Error(), 404)
		return
	}
	out, _ := json.Marshal(auths)

	// If the client can optionally accept an encrypted
	// response, we encrypt the payload body using secretbox.
	if r.Header.Get("Accept-Encoding") == "aesgcm" {
		key, err := aesgcm.Key(p.secret)
		if err != nil {
			p.logger.Errorf("environment: invalid encryption key: %s", err)
			http.Error(w, err.Error(), 500)
			return
		}
		out, err = aesgcm.Encrypt(out, key)
		if err != nil {
			p.logger.Errorf("environment: cannot encrypt message: %s", err)
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Encoding", "aesgcm")
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	w.Write(out)
}
