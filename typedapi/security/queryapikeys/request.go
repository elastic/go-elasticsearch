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

package queryapikeys

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package queryapikeys
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/security/query_api_keys/QueryApiKeysRequest.ts#L25-L67
type Request struct {

	// From Starting document offset. By default, you cannot page through more than
	// 10,000
	// hits using the from and size parameters. To page through more hits, use the
	// search_after parameter.
	From *int `json:"from,omitempty"`
	// Query A query to filter which API keys to return.
	// The query supports a subset of query types, including match_all, bool, term,
	// terms, ids, prefix, wildcard, and range.
	// You can query all public information associated with an API key
	Query       *types.Query       `json:"query,omitempty"`
	SearchAfter []types.FieldValue `json:"search_after,omitempty"`
	// Size The number of hits to return. By default, you cannot page through more
	// than 10,000 hits using the from and size parameters. To page through more
	// hits, use the search_after parameter.
	Size *int                     `json:"size,omitempty"`
	Sort []types.SortCombinations `json:"sort,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Queryapikeys request: %w", err)
	}

	return &req, nil
}
