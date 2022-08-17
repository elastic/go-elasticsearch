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


package reindex

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/conflicts"
)

// Request holds the request body struct for the package reindex
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/reindex/ReindexRequest.ts#L27-L51
type Request struct {
	Conflicts *conflicts.Conflicts `json:"conflicts,omitempty"`

	Dest types.Destination `json:"dest"`

	MaxDocs *int64 `json:"max_docs,omitempty"`

	Script *types.Script `json:"script,omitempty"`

	Size *int64 `json:"size,omitempty"`

	Source types.Source `json:"source"`
}

// RequestBuilder is the builder API for the reindex.Request
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
		return nil, fmt.Errorf("could not deserialise json into Reindex request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Conflicts(conflicts conflicts.Conflicts) *RequestBuilder {
	rb.v.Conflicts = &conflicts
	return rb
}

func (rb *RequestBuilder) Dest(dest *types.DestinationBuilder) *RequestBuilder {
	v := dest.Build()
	rb.v.Dest = v
	return rb
}

func (rb *RequestBuilder) MaxDocs(maxdocs int64) *RequestBuilder {
	rb.v.MaxDocs = &maxdocs
	return rb
}

func (rb *RequestBuilder) Script(script *types.ScriptBuilder) *RequestBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *RequestBuilder) Size(size int64) *RequestBuilder {
	rb.v.Size = &size
	return rb
}

func (rb *RequestBuilder) Source(source *types.SourceBuilder) *RequestBuilder {
	v := source.Build()
	rb.v.Source = v
	return rb
}
