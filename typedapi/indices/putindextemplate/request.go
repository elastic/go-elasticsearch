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


package putindextemplate

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putindextemplate
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/put_index_template/IndicesPutIndexTemplateRequest.ts#L35-L58
type Request struct {
	ComposedOf []types.Name `json:"composed_of,omitempty"`

	DataStream *types.DataStreamVisibility `json:"data_stream,omitempty"`

	IndexPatterns *types.Indices `json:"index_patterns,omitempty"`

	Meta_ *types.Metadata `json:"_meta,omitempty"`

	Priority *int `json:"priority,omitempty"`

	Template *types.IndexTemplateMapping `json:"template,omitempty"`

	Version *types.VersionNumber `json:"version,omitempty"`
}

// RequestBuilder is the builder API for the putindextemplate.Request
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
		return nil, fmt.Errorf("could not deserialise json into Putindextemplate request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) ComposedOf(composed_of ...types.Name) *RequestBuilder {
	rb.v.ComposedOf = composed_of
	return rb
}

func (rb *RequestBuilder) DataStream(datastream *types.DataStreamVisibilityBuilder) *RequestBuilder {
	v := datastream.Build()
	rb.v.DataStream = &v
	return rb
}

func (rb *RequestBuilder) IndexPatterns(indexpatterns *types.IndicesBuilder) *RequestBuilder {
	v := indexpatterns.Build()
	rb.v.IndexPatterns = &v
	return rb
}

func (rb *RequestBuilder) Meta_(meta_ *types.MetadataBuilder) *RequestBuilder {
	v := meta_.Build()
	rb.v.Meta_ = &v
	return rb
}

func (rb *RequestBuilder) Priority(priority int) *RequestBuilder {
	rb.v.Priority = &priority
	return rb
}

func (rb *RequestBuilder) Template(template *types.IndexTemplateMappingBuilder) *RequestBuilder {
	v := template.Build()
	rb.v.Template = &v
	return rb
}

func (rb *RequestBuilder) Version(version types.VersionNumber) *RequestBuilder {
	rb.v.Version = &version
	return rb
}
