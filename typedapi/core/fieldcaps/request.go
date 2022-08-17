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


package fieldcaps

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package fieldcaps
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/field_caps/FieldCapabilitiesRequest.ts#L25-L90
type Request struct {

	// IndexFilter Allows to filter indices if the provided query rewrites to match_none on
	// every shard.
	IndexFilter *types.QueryContainer `json:"index_filter,omitempty"`

	// RuntimeMappings Defines ad-hoc runtime fields in the request similar to the way it is done in
	// search requests.
	// These fields exist only as part of the query and take precedence over fields
	// defined with the same name in the index mappings.
	RuntimeMappings *types.RuntimeFields `json:"runtime_mappings,omitempty"`
}

// RequestBuilder is the builder API for the fieldcaps.Request
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
		return nil, fmt.Errorf("could not deserialise json into Fieldcaps request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) IndexFilter(indexfilter *types.QueryContainerBuilder) *RequestBuilder {
	v := indexfilter.Build()
	rb.v.IndexFilter = &v
	return rb
}

func (rb *RequestBuilder) RuntimeMappings(runtimemappings *types.RuntimeFieldsBuilder) *RequestBuilder {
	v := runtimemappings.Build()
	rb.v.RuntimeMappings = &v
	return rb
}
