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

package query

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package query
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/sql/query/QuerySqlResponse.ts#L23-L60

type Response struct {

	// Columns Column headings for the search results. Each object is a column.
	Columns []types.Column `json:"columns,omitempty"`
	// Cursor Cursor for the next set of paginated results. For CSV, TSV, and
	// TXT responses, this value is returned in the `Cursor` HTTP header.
	Cursor *string `json:"cursor,omitempty"`
	// Id Identifier for the search. This value is only returned for async and saved
	// synchronous searches. For CSV, TSV, and TXT responses, this value is returned
	// in the `Async-ID` HTTP header.
	Id *string `json:"id,omitempty"`
	// IsPartial If `true`, the response does not contain complete search results. If
	// `is_partial`
	// is `true` and `is_running` is `true`, the search is still running. If
	// `is_partial`
	// is `true` but `is_running` is `false`, the results are partial due to a
	// failure or
	// timeout. This value is only returned for async and saved synchronous
	// searches.
	// For CSV, TSV, and TXT responses, this value is returned in the
	// `Async-partial` HTTP header.
	IsPartial *bool `json:"is_partial,omitempty"`
	// IsRunning If `true`, the search is still running. If false, the search has finished.
	// This value is only returned for async and saved synchronous searches. For
	// CSV, TSV, and TXT responses, this value is returned in the `Async-partial`
	// HTTP header.
	IsRunning *bool `json:"is_running,omitempty"`
	// Rows Values for the search results.
	Rows [][]json.RawMessage `json:"rows"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
