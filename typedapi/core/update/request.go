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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package update

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package update
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_global/update/UpdateRequest.ts#L38-L151
type Request struct {

	// DetectNoop Set to false to disable setting 'result' in the response
	// to 'noop' if no change to the document occurred.
	DetectNoop *bool `json:"detect_noop,omitempty"`
	// Doc A partial update to an existing document.
	Doc json.RawMessage `json:"doc,omitempty"`
	// DocAsUpsert Set to true to use the contents of 'doc' as the value of 'upsert'
	DocAsUpsert *bool `json:"doc_as_upsert,omitempty"`
	// Script Script to execute to update the document.
	Script types.Script `json:"script,omitempty"`
	// ScriptedUpsert Set to true to execute the script whether or not the document exists.
	ScriptedUpsert *bool `json:"scripted_upsert,omitempty"`
	// Source_ Set to false to disable source retrieval. You can also specify a
	// comma-separated
	// list of the fields you want to retrieve.
	Source_ types.SourceConfig `json:"_source,omitempty"`
	// Upsert If the document does not already exist, the contents of 'upsert' are inserted
	// as a
	// new document. If the document exists, the 'script' is executed.
	Upsert json.RawMessage `json:"upsert,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Update request: %w", err)
	}

	return &req, nil
}
