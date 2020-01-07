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
	"context"

	"github.com/drone/drone-go/drone"
)

// V1 is version 1 of the admission API
const V1 = "application/vnd.drone.webhook.v1+json"

// Webhook event types.
const (
	EventBuild = "build"
	EventRepo  = "repo"
	EventUser  = "user"
)

// Webhook action types.
const (
	ActionCreated  = "created"
	ActionUpdated  = "updated"
	ActionDeleted  = "deleted"
	ActionEnabled  = "enabled"
	ActionDisabled = "disabled"
)

type (
	// Request defines a webhook request.
	Request struct {
		Event  string        `json:"event"`
		Action string        `json:"action"`
		User   *drone.User   `json:"user,omitempty"`
		Repo   *drone.Repo   `json:"repo,omitempty"`
		Build  *drone.Build  `json:"build,omitempty"`
		System *drone.System `json:"system,omitempty"`
	}

	// Plugin responds to a webhook request.
	Plugin interface {
		Deliver(context.Context, *Request) error
	}
)
