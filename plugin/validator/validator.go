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
	"context"
	"errors"

	"github.com/drone/drone-go/drone"
)

// V1 is version 1 of the validator API
const V1 = "application/vnd.drone.validate.v1+json"

// ErrSkip is returned when the build should be skipped
// instead of throwing an error.
var ErrSkip = errors.New("skip")

// ErrBlock is returned when the build should be blocked
// instead of throwing an error.
var ErrBlock = errors.New("block")

type (
	// Request defines a validator request.
	Request struct {
		Build  drone.Build  `json:"build,omitempty"`
		Config drone.Config `json:"config,omitempty"`
		Repo   drone.Repo   `json:"repo,omitempty"`
		Token  drone.Token  `json:"token,omitempty"` // not implemented
	}

	// Plugin responds to a validator request.
	Plugin interface {
		Validate(context.Context, *Request) error
	}
)
