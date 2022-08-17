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


package allocationexplain

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package allocationexplain
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/allocation_explain/ClusterAllocationExplainRequest.ts#L24-L61
type Request struct {

	// CurrentNode Specifies the node ID or the name of the node to only explain a shard that is
	// currently located on the specified node.
	CurrentNode *string `json:"current_node,omitempty"`

	// Index Specifies the name of the index that you would like an explanation for.
	Index *types.IndexName `json:"index,omitempty"`

	// Primary If true, returns explanation for the primary shard for the given shard ID.
	Primary *bool `json:"primary,omitempty"`

	// Shard Specifies the ID of the shard that you would like an explanation for.
	Shard *int `json:"shard,omitempty"`
}

// RequestBuilder is the builder API for the allocationexplain.Request
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
		return nil, fmt.Errorf("could not deserialise json into Allocationexplain request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) CurrentNode(currentnode string) *RequestBuilder {
	rb.v.CurrentNode = &currentnode
	return rb
}

func (rb *RequestBuilder) Index(index types.IndexName) *RequestBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *RequestBuilder) Primary(primary bool) *RequestBuilder {
	rb.v.Primary = &primary
	return rb
}

func (rb *RequestBuilder) Shard(shard int) *RequestBuilder {
	rb.v.Shard = &shard
	return rb
}
