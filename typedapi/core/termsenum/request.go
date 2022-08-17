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


package termsenum

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package termsenum
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/terms_enum/TermsEnumRequest.ts#L26-L65
type Request struct {

	// CaseInsensitive When true the provided search string is matched against index terms without
	// case sensitivity.
	CaseInsensitive *bool `json:"case_insensitive,omitempty"`

	// Field The string to match at the start of indexed terms. If not provided, all terms
	// in the field are considered.
	Field types.Field `json:"field"`

	// IndexFilter Allows to filter an index shard if the provided query rewrites to match_none.
	IndexFilter *types.QueryContainer `json:"index_filter,omitempty"`

	SearchAfter *string `json:"search_after,omitempty"`

	// Size How many matching terms to return.
	Size *int `json:"size,omitempty"`

	// String The string after which terms in the index should be returned. Allows for a
	// form of pagination if the last result from one request is passed as the
	// search_after parameter for a subsequent request.
	String *string `json:"string,omitempty"`

	// Timeout The maximum length of time to spend collecting results. Defaults to "1s" (one
	// second). If the timeout is exceeded the complete flag set to false in the
	// response and the results may be partial or empty.
	Timeout *types.Duration `json:"timeout,omitempty"`
}

// RequestBuilder is the builder API for the termsenum.Request
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
		return nil, fmt.Errorf("could not deserialise json into Termsenum request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) CaseInsensitive(caseinsensitive bool) *RequestBuilder {
	rb.v.CaseInsensitive = &caseinsensitive
	return rb
}

func (rb *RequestBuilder) Field(field types.Field) *RequestBuilder {
	rb.v.Field = field
	return rb
}

func (rb *RequestBuilder) IndexFilter(indexfilter *types.QueryContainerBuilder) *RequestBuilder {
	v := indexfilter.Build()
	rb.v.IndexFilter = &v
	return rb
}

func (rb *RequestBuilder) SearchAfter(searchafter string) *RequestBuilder {
	rb.v.SearchAfter = &searchafter
	return rb
}

func (rb *RequestBuilder) Size(size int) *RequestBuilder {
	rb.v.Size = &size
	return rb
}

func (rb *RequestBuilder) String(string string) *RequestBuilder {
	rb.v.String = &string
	return rb
}

func (rb *RequestBuilder) Timeout(timeout *types.DurationBuilder) *RequestBuilder {
	v := timeout.Build()
	rb.v.Timeout = &v
	return rb
}
