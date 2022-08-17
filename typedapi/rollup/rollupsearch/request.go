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


package rollupsearch

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package rollupsearch
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/rollup/rollup_search/RollupSearchRequest.ts#L27-L47
type Request struct {
	Aggregations map[string]types.AggregationContainer `json:"aggregations,omitempty"`

	Query *types.QueryContainer `json:"query,omitempty"`

	// Size Must be zero if set, as rollups work on pre-aggregated data
	Size *int `json:"size,omitempty"`
}

// RequestBuilder is the builder API for the rollupsearch.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{
			Aggregations: make(map[string]types.AggregationContainer, 0),
		},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Rollupsearch request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Aggregations(values map[string]*types.AggregationContainerBuilder) *RequestBuilder {
	tmp := make(map[string]types.AggregationContainer, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *RequestBuilder) Query(query *types.QueryContainerBuilder) *RequestBuilder {
	v := query.Build()
	rb.v.Query = &v
	return rb
}

func (rb *RequestBuilder) Size(size int) *RequestBuilder {
	rb.v.Size = &size
	return rb
}
