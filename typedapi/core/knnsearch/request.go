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


package knnsearch

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package knnsearch
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/knn_search/KnnSearchRequest.ts#L27-L79
type Request struct {

	// DocvalueFields The request returns doc values for field names matching these patterns
	// in the hits.fields property of the response. Accepts wildcard (*) patterns.
	DocvalueFields []types.FieldAndFormat `json:"docvalue_fields,omitempty"`

	// Fields The request returns values for field names matching these patterns
	// in the hits.fields property of the response. Accepts wildcard (*) patterns.
	Fields *types.Fields `json:"fields,omitempty"`

	// Filter Query to filter the documents that can match. The kNN search will return the
	// top
	// `k` documents that also match this filter. The value can be a single query or
	// a
	// list of queries. If `filter` isn't provided, all documents are allowed to
	// match.
	Filter []types.QueryContainer `json:"filter,omitempty"`

	// Knn kNN query to execute
	Knn types.Query `json:"knn"`

	// Source_ Indicates which source fields are returned for matching documents. These
	// fields are returned in the hits._source property of the search response.
	Source_ *types.SourceConfig `json:"_source,omitempty"`

	// StoredFields List of stored fields to return as part of a hit. If no fields are specified,
	// no stored fields are included in the response. If this field is specified,
	// the _source
	// parameter defaults to false. You can pass _source: true to return both source
	// fields
	// and stored fields in the search response.
	StoredFields *types.Fields `json:"stored_fields,omitempty"`
}

// RequestBuilder is the builder API for the knnsearch.Request
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
		return nil, fmt.Errorf("could not deserialise json into Knnsearch request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) DocvalueFields(docvalue_fields []types.FieldAndFormatBuilder) *RequestBuilder {
	tmp := make([]types.FieldAndFormat, len(docvalue_fields))
	for _, value := range docvalue_fields {
		tmp = append(tmp, value.Build())
	}
	rb.v.DocvalueFields = tmp
	return rb
}

func (rb *RequestBuilder) Fields(fields *types.FieldsBuilder) *RequestBuilder {
	v := fields.Build()
	rb.v.Fields = &v
	return rb
}

func (rb *RequestBuilder) Filter(arg []types.QueryContainer) *RequestBuilder {
	rb.v.Filter = arg
	return rb
}

func (rb *RequestBuilder) Knn(knn *types.QueryBuilder) *RequestBuilder {
	v := knn.Build()
	rb.v.Knn = v
	return rb
}

func (rb *RequestBuilder) Source_(source_ *types.SourceConfigBuilder) *RequestBuilder {
	v := source_.Build()
	rb.v.Source_ = &v
	return rb
}

func (rb *RequestBuilder) StoredFields(storedfields *types.FieldsBuilder) *RequestBuilder {
	v := storedfields.Build()
	rb.v.StoredFields = &v
	return rb
}
