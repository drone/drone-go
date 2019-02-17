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
	"context"

	"github.com/drone/drone-go/drone"
)

// V1 is version 1 of the admission API
const V1 = "application/vnd.drone.admission.v1+json"

// Admission event types.
const (
	EventLogin    = "login"
	EventRegister = "register"
)

type (
	// Request defines an admission request.
	Request struct {
		Event string     `json:"event,omitempty"`
		User  drone.User `json:"user,omitempty"`
	}

	// Plugin responds to a admission request.
	Plugin interface {
		Admit(context.Context, *Request) (*drone.User, error)
	}
)
