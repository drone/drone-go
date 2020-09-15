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

package webhook

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/drone/drone-go/plugin/logger"

	"github.com/99designs/httpsignatures-go"
)

// Handler returns a http.Handler that accepts JSON-encoded
// HTTP requests for a webhook, invokes the underlying webhook
// plugin, and writes the JSON-encoded data to the HTTP response.
//
// The handler verifies the authenticity of the HTTP request
// using the http-signature, and returns a 400 Bad Request if
// the signature is missing or invalid.
//
// The handler can optionally encrypt the response body using
// aesgcm if the HTTP request includes the Accept-Encoding header
// set to aesgcm.
func Handler(plugin Plugin, secret string, logs logger.Logger) http.Handler {
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
		p.logger.Debugf("webhook: invalid or missing signature in http.Request")
		http.Error(w, "Invalid or Missing Signature", http.StatusBadRequest)
		return
	}
	if !signature.IsValid(p.secret, r) {
		p.logger.Debugf("webhook: invalid signature in http.Request")
		http.Error(w, "Invalid Signature", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		p.logger.Debugf("webhook: cannot read http.Request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req := &Request{}
	err = json.Unmarshal(body, req)
	if err != nil {
		p.logger.Debugf("webhook: cannot unmarshal http.Request body")
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	err = p.plugin.Deliver(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
