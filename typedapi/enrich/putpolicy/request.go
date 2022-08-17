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


package putpolicy

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putpolicy
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/enrich/put_policy/PutEnrichPolicyRequest.ts#L24-L38
type Request struct {
	GeoMatch *types.Policy `json:"geo_match,omitempty"`

	Match *types.Policy `json:"match,omitempty"`

	Range *types.Policy `json:"range,omitempty"`
}

// RequestBuilder is the builder API for the putpolicy.Request
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
		return nil, fmt.Errorf("could not deserialise json into Putpolicy request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) GeoMatch(geomatch *types.PolicyBuilder) *RequestBuilder {
	v := geomatch.Build()
	rb.v.GeoMatch = &v
	return rb
}

func (rb *RequestBuilder) Match(match *types.PolicyBuilder) *RequestBuilder {
	v := match.Build()
	rb.v.Match = &v
	return rb
}

func (rb *RequestBuilder) Range_(range_ *types.PolicyBuilder) *RequestBuilder {
	v := range_.Build()
	rb.v.Range = &v
	return rb
}
