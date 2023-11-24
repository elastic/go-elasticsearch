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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package rollover

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package rollover
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/indices/rollover/IndicesRolloverRequest.ts#L29-L99
type Request struct {

	// Aliases Aliases for the target index.
	// Data streams do not support this parameter.
	Aliases map[string]types.Alias `json:"aliases,omitempty"`
	// Conditions Conditions for the rollover.
	// If specified, Elasticsearch only performs the rollover if the current index
	// satisfies these conditions.
	// If this parameter is not specified, Elasticsearch performs the rollover
	// unconditionally.
	// If conditions are specified, at least one of them must be a `max_*`
	// condition.
	// The index will rollover if any `max_*` condition is satisfied and all `min_*`
	// conditions are satisfied.
	Conditions *types.RolloverConditions `json:"conditions,omitempty"`
	// Mappings Mapping for fields in the index.
	// If specified, this mapping can include field names, field data types, and
	// mapping paramaters.
	Mappings *types.TypeMapping `json:"mappings,omitempty"`
	// Settings Configuration options for the index.
	// Data streams do not support this parameter.
	Settings map[string]json.RawMessage `json:"settings,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Aliases:  make(map[string]types.Alias, 0),
		Settings: make(map[string]json.RawMessage, 0),
	}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Rollover request: %w", err)
	}

	return &req, nil
}
