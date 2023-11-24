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

package putalias

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putalias
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/indices/put_alias/IndicesPutAliasRequest.ts#L25-L91
type Request struct {

	// Filter Query used to limit documents the alias can access.
	Filter *types.Query `json:"filter,omitempty"`
	// IndexRouting Value used to route indexing operations to a specific shard.
	// If specified, this overwrites the `routing` value for indexing operations.
	// Data stream aliases don’t support this parameter.
	IndexRouting *string `json:"index_routing,omitempty"`
	// IsWriteIndex If `true`, sets the write index or data stream for the alias.
	// If an alias points to multiple indices or data streams and `is_write_index`
	// isn’t set, the alias rejects write requests.
	// If an index alias points to one index and `is_write_index` isn’t set, the
	// index automatically acts as the write index.
	// Data stream aliases don’t automatically set a write data stream, even if the
	// alias points to one data stream.
	IsWriteIndex *bool `json:"is_write_index,omitempty"`
	// Routing Value used to route indexing and search operations to a specific shard.
	// Data stream aliases don’t support this parameter.
	Routing *string `json:"routing,omitempty"`
	// SearchRouting Value used to route search operations to a specific shard.
	// If specified, this overwrites the `routing` value for search operations.
	// Data stream aliases don’t support this parameter.
	SearchRouting *string `json:"search_routing,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Putalias request: %w", err)
	}

	return &req, nil
}
