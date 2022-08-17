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


package types

// Page type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Page.ts#L22-L33
type Page struct {
	// From Skips the specified number of items.
	From *int `json:"from,omitempty"`
	// Size Specifies the maximum number of items to obtain.
	Size *int `json:"size,omitempty"`
}

// PageBuilder holds Page struct and provides a builder API.
type PageBuilder struct {
	v *Page
}

// NewPage provides a builder for the Page struct.
func NewPageBuilder() *PageBuilder {
	r := PageBuilder{
		&Page{},
	}

	return &r
}

// Build finalize the chain and returns the Page struct
func (rb *PageBuilder) Build() Page {
	return *rb.v
}

// From Skips the specified number of items.

func (rb *PageBuilder) From(from int) *PageBuilder {
	rb.v.From = &from
	return rb
}

// Size Specifies the maximum number of items to obtain.

func (rb *PageBuilder) Size(size int) *PageBuilder {
	rb.v.Size = &size
	return rb
}
