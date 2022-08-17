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

// ResolveIndexDataStreamsItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/resolve_index/ResolveIndexResponse.ts#L42-L46
type ResolveIndexDataStreamsItem struct {
	BackingIndices Indices        `json:"backing_indices"`
	Name           DataStreamName `json:"name"`
	TimestampField Field          `json:"timestamp_field"`
}

// ResolveIndexDataStreamsItemBuilder holds ResolveIndexDataStreamsItem struct and provides a builder API.
type ResolveIndexDataStreamsItemBuilder struct {
	v *ResolveIndexDataStreamsItem
}

// NewResolveIndexDataStreamsItem provides a builder for the ResolveIndexDataStreamsItem struct.
func NewResolveIndexDataStreamsItemBuilder() *ResolveIndexDataStreamsItemBuilder {
	r := ResolveIndexDataStreamsItemBuilder{
		&ResolveIndexDataStreamsItem{},
	}

	return &r
}

// Build finalize the chain and returns the ResolveIndexDataStreamsItem struct
func (rb *ResolveIndexDataStreamsItemBuilder) Build() ResolveIndexDataStreamsItem {
	return *rb.v
}

func (rb *ResolveIndexDataStreamsItemBuilder) BackingIndices(backingindices *IndicesBuilder) *ResolveIndexDataStreamsItemBuilder {
	v := backingindices.Build()
	rb.v.BackingIndices = v
	return rb
}

func (rb *ResolveIndexDataStreamsItemBuilder) Name(name DataStreamName) *ResolveIndexDataStreamsItemBuilder {
	rb.v.Name = name
	return rb
}

func (rb *ResolveIndexDataStreamsItemBuilder) TimestampField(timestampfield Field) *ResolveIndexDataStreamsItemBuilder {
	rb.v.TimestampField = timestampfield
	return rb
}
