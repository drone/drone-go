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
	"context"

	"github.com/drone/drone-go/plugin/internal/client"
)

// Client returns a new plugin client.
func Client(endpoint, secret string, skipverify bool) Plugin {
	client := client.New(endpoint, secret, skipverify)
	client.Accept = V1
	return &pluginClient{
		client: client,
	}
}

type pluginClient struct {
	client *client.Client
}

func (c *pluginClient) List(ctx context.Context, in *Request) (map[string]string, error) {
	res := map[string]string{}
	err := c.client.Do(in, &res)
	return res, err
}
