// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package invalidateapikey

import (
	"encoding/json"
	"fmt"
)

// Request holds the request body struct for the package invalidateapikey
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/security/invalidate_api_key/SecurityInvalidateApiKeyRequest.ts#L23-L66
type Request struct {
	Id *string `json:"id,omitempty"`
	// Ids A list of API key ids.
	// This parameter cannot be used with any of `name`, `realm_name`, or
	// `username`.
	Ids []string `json:"ids,omitempty"`
	// Name An API key name.
	// This parameter cannot be used with any of `ids`, `realm_name` or `username`.
	Name *string `json:"name,omitempty"`
	// Owner Can be used to query API keys owned by the currently authenticated user.
	// The `realm_name` or `username` parameters cannot be specified when this
	// parameter is set to `true` as they are assumed to be the currently
	// authenticated ones.
	Owner *bool `json:"owner,omitempty"`
	// RealmName The name of an authentication realm.
	// This parameter cannot be used with either `ids` or `name`, or when `owner`
	// flag is set to `true`.
	RealmName *string `json:"realm_name,omitempty"`
	// Username The username of a user.
	// This parameter cannot be used with either `ids` or `name`, or when `owner`
	// flag is set to `true`.
	Username *string `json:"username,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Invalidateapikey request: %w", err)
	}

	return &req, nil
}
