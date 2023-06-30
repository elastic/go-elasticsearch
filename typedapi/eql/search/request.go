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

package search

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/resultposition"
)

// Request holds the request body struct for the package search
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/eql/search/EqlSearchRequest.ts#L28-L115
type Request struct {
	CaseSensitive *bool `json:"case_sensitive,omitempty"`
	// EventCategoryField Field containing the event classification, such as process, file, or network.
	EventCategoryField *string `json:"event_category_field,omitempty"`
	// FetchSize Maximum number of events to search at a time for sequence queries.
	FetchSize *uint `json:"fetch_size,omitempty"`
	// Fields Array of wildcard (*) patterns. The response returns values for field names
	// matching these patterns in the fields property of each hit.
	Fields []types.FieldAndFormat `json:"fields,omitempty"`
	// Filter Query, written in Query DSL, used to filter the events on which the EQL query
	// runs.
	Filter           []types.Query  `json:"filter,omitempty"`
	KeepAlive        types.Duration `json:"keep_alive,omitempty"`
	KeepOnCompletion *bool          `json:"keep_on_completion,omitempty"`
	// Query EQL query you wish to run.
	Query           string                         `json:"query"`
	ResultPosition  *resultposition.ResultPosition `json:"result_position,omitempty"`
	RuntimeMappings types.RuntimeFields            `json:"runtime_mappings,omitempty"`
	// Size For basic queries, the maximum number of matching events to return. Defaults
	// to 10
	Size *uint `json:"size,omitempty"`
	// TiebreakerField Field used to sort hits with the same timestamp in ascending order
	TiebreakerField *string `json:"tiebreaker_field,omitempty"`
	// TimestampField Field containing event timestamp. Default "@timestamp"
	TimestampField           *string        `json:"timestamp_field,omitempty"`
	WaitForCompletionTimeout types.Duration `json:"wait_for_completion_timeout,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Search request: %w", err)
	}

	return &req, nil
}
