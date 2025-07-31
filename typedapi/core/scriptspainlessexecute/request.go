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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package scriptspainlessexecute

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/painlesscontext"
)

// Request holds the request body struct for the package scriptspainlessexecute
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/scripts_painless_execute/ExecutePainlessScriptRequest.ts#L24-L64
type Request struct {

	// Context The context that the script should run in.
	// NOTE: Result ordering in the field contexts is not guaranteed.
	Context *painlesscontext.PainlessContext `json:"context,omitempty"`
	// ContextSetup Additional parameters for the `context`.
	// NOTE: This parameter is required for all contexts except `painless_test`,
	// which is the default if no value is provided for `context`.
	ContextSetup *types.PainlessContextSetup `json:"context_setup,omitempty"`
	// Script The Painless script to run.
	Script *types.Script `json:"script,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Scriptspainlessexecute request: %w", err)
	}

	return &req, nil
}
