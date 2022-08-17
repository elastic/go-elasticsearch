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

// IndexVersioning type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexSettings.ts#L262-L265
type IndexVersioning struct {
	Created       VersionString  `json:"created"`
	CreatedString *VersionString `json:"created_string,omitempty"`
}

// IndexVersioningBuilder holds IndexVersioning struct and provides a builder API.
type IndexVersioningBuilder struct {
	v *IndexVersioning
}

// NewIndexVersioning provides a builder for the IndexVersioning struct.
func NewIndexVersioningBuilder() *IndexVersioningBuilder {
	r := IndexVersioningBuilder{
		&IndexVersioning{},
	}

	return &r
}

// Build finalize the chain and returns the IndexVersioning struct
func (rb *IndexVersioningBuilder) Build() IndexVersioning {
	return *rb.v
}

func (rb *IndexVersioningBuilder) Created(created VersionString) *IndexVersioningBuilder {
	rb.v.Created = created
	return rb
}

func (rb *IndexVersioningBuilder) CreatedString(createdstring VersionString) *IndexVersioningBuilder {
	rb.v.CreatedString = &createdstring
	return rb
}
