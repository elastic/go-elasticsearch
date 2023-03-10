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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package fieldcaps

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package fieldcaps
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_global/field_caps/FieldCapabilitiesRequest.ts#L25-L95
type Request struct {

	// Fields List of fields to retrieve capabilities for. Wildcard (`*`) expressions are
	// supported.
	Fields []string `json:"fields,omitempty"`
	// IndexFilter Allows to filter indices if the provided query rewrites to match_none on
	// every shard.
	IndexFilter *types.Query `json:"index_filter,omitempty"`
	// RuntimeMappings Defines ad-hoc runtime fields in the request similar to the way it is done in
	// search requests.
	// These fields exist only as part of the query and take precedence over fields
	// defined with the same name in the index mappings.
	RuntimeMappings map[string]types.RuntimeField `json:"runtime_mappings,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Fieldcaps request: %w", err)
	}

	return &req, nil
}
