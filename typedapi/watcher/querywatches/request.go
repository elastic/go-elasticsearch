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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package querywatches

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package querywatches
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/watcher/query_watches/WatcherQueryWatchesRequest.ts#L25-L49
type Request struct {

	// From The offset from the first result to fetch. Needs to be non-negative.
	From *int `json:"from,omitempty"`
	// Query Optional, query filter watches to be returned.
	Query *types.Query `json:"query,omitempty"`
	// SearchAfter Optional search After to do pagination using last hit’s sort values.
	SearchAfter []types.FieldValue `json:"search_after,omitempty"`
	// Size The number of hits to return. Needs to be non-negative.
	Size *int `json:"size,omitempty"`
	// Sort Optional sort definition.
	Sort []types.SortCombinations `json:"sort,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Querywatches request: %w", err)
	}

	return &req, nil
}
