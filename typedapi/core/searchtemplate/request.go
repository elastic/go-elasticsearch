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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package searchtemplate

import (
	"encoding/json"
	"fmt"
)

// Request holds the request body struct for the package searchtemplate
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_global/search_template/SearchTemplateRequest.ts#L32-L96
type Request struct {
	Explain *bool `json:"explain,omitempty"`
	// Id ID of the search template to use. If no source is specified,
	// this parameter is required.
	Id      *string                    `json:"id,omitempty"`
	Params  map[string]json.RawMessage `json:"params,omitempty"`
	Profile *bool                      `json:"profile,omitempty"`
	// Source An inline search template. Supports the same parameters as the search API's
	// request body. Also supports Mustache variables. If no id is specified, this
	// parameter is required.
	Source *string `json:"source,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Params: make(map[string]json.RawMessage, 0),
	}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Searchtemplate request: %w", err)
	}

	return &req, nil
}
