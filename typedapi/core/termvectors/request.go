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


package termvectors

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package termvectors
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/termvectors/TermVectorsRequest.ts#L33-L61
type Request struct {
	Doc interface{} `json:"doc,omitempty"`

	Filter *types.Filter `json:"filter,omitempty"`

	PerFieldAnalyzer map[types.Field]string `json:"per_field_analyzer,omitempty"`
}

// RequestBuilder is the builder API for the termvectors.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{
			PerFieldAnalyzer: make(map[types.Field]string, 0),
		},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Termvectors request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Doc(doc interface{}) *RequestBuilder {
	rb.v.Doc = doc
	return rb
}

func (rb *RequestBuilder) Filter(filter *types.FilterBuilder) *RequestBuilder {
	v := filter.Build()
	rb.v.Filter = &v
	return rb
}

func (rb *RequestBuilder) PerFieldAnalyzer(value map[types.Field]string) *RequestBuilder {
	rb.v.PerFieldAnalyzer = value
	return rb
}
