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

package search

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package search
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/eql/search/EqlSearchResponse.ts#L22-L24

type Response struct {

	// Hits Contains matching events and sequences. Also contains related metadata.
	Hits types.EqlHits `json:"hits"`
	// Id Identifier for the search.
	Id *string `json:"id,omitempty"`
	// IsPartial If true, the response does not contain complete search results.
	IsPartial *bool `json:"is_partial,omitempty"`
	// IsRunning If true, the search request is still executing.
	IsRunning *bool `json:"is_running,omitempty"`
	// TimedOut If true, the request timed out before completion.
	TimedOut *bool `json:"timed_out,omitempty"`
	// Took Milliseconds it took Elasticsearch to execute the request.
	Took *int64 `json:"took,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
