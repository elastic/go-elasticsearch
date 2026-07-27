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
// https://github.com/elastic/elasticsearch-specification/tree/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c

package putdatasource

import (
	"encoding/json"
	"fmt"
)

// Request holds the request body struct for the package putdatasource
//
// https://github.com/elastic/elasticsearch-specification/blob/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c/specification/esql/put_data_source/PutDataSourceRequest.ts#L26-L72
type Request struct {
	// Description A free-text description of the data source.
	Description *string `json:"description,omitempty"`
	// Settings Type-specific settings. The accepted keys depend on the data source type's
	// validator.
	Settings map[string]json.RawMessage `json:"settings,omitempty"`
	// Type The data source type. Must be lowercase and contain no whitespace.
	Type string `json:"type"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Settings: make(map[string]json.RawMessage, 0),
	}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putdatasource request: %w", err)
	}

	return &req, nil
}
