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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package queryapikeys

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package queryapikeys
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/query_api_keys/QueryApiKeysRequest.ts#L25-L56
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
	Query *types.QueryContainer `json:"query,omitempty"`

	SearchAfter *types.SortResults `json:"search_after,omitempty"`

	// Size The number of hits to return. By default, you cannot page through more
	// than 10,000 hits using the from and size parameters. To page through more
	// hits, use the search_after parameter.
	Size *int `json:"size,omitempty"`

	Sort *types.Sort `json:"sort,omitempty"`
}

// RequestBuilder is the builder API for the queryapikeys.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Queryapikeys request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) From(from int) *RequestBuilder {
	rb.v.From = &from
	return rb
}

func (rb *RequestBuilder) Query(query *types.QueryContainerBuilder) *RequestBuilder {
	v := query.Build()
	rb.v.Query = &v
	return rb
}

func (rb *RequestBuilder) SearchAfter(searchafter *types.SortResultsBuilder) *RequestBuilder {
	v := searchafter.Build()
	rb.v.SearchAfter = &v
	return rb
}

func (rb *RequestBuilder) Size(size int) *RequestBuilder {
	rb.v.Size = &size
	return rb
}

func (rb *RequestBuilder) Sort(sort *types.SortBuilder) *RequestBuilder {
	v := sort.Build()
	rb.v.Sort = &v
	return rb
}
