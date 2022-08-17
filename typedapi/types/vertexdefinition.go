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

// VertexDefinition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/graph/_types/Vertex.ts#L30-L37
type VertexDefinition struct {
	Exclude          []string        `json:"exclude,omitempty"`
	Field            Field           `json:"field"`
	Include          []VertexInclude `json:"include,omitempty"`
	MinDocCount      *int64          `json:"min_doc_count,omitempty"`
	ShardMinDocCount *int64          `json:"shard_min_doc_count,omitempty"`
	Size             *int            `json:"size,omitempty"`
}

// VertexDefinitionBuilder holds VertexDefinition struct and provides a builder API.
type VertexDefinitionBuilder struct {
	v *VertexDefinition
}

// NewVertexDefinition provides a builder for the VertexDefinition struct.
func NewVertexDefinitionBuilder() *VertexDefinitionBuilder {
	r := VertexDefinitionBuilder{
		&VertexDefinition{},
	}

	return &r
}

// Build finalize the chain and returns the VertexDefinition struct
func (rb *VertexDefinitionBuilder) Build() VertexDefinition {
	return *rb.v
}

func (rb *VertexDefinitionBuilder) Exclude(exclude ...string) *VertexDefinitionBuilder {
	rb.v.Exclude = exclude
	return rb
}

func (rb *VertexDefinitionBuilder) Field(field Field) *VertexDefinitionBuilder {
	rb.v.Field = field
	return rb
}

func (rb *VertexDefinitionBuilder) Include(include []VertexIncludeBuilder) *VertexDefinitionBuilder {
	tmp := make([]VertexInclude, len(include))
	for _, value := range include {
		tmp = append(tmp, value.Build())
	}
	rb.v.Include = tmp
	return rb
}

func (rb *VertexDefinitionBuilder) MinDocCount(mindoccount int64) *VertexDefinitionBuilder {
	rb.v.MinDocCount = &mindoccount
	return rb
}

func (rb *VertexDefinitionBuilder) ShardMinDocCount(shardmindoccount int64) *VertexDefinitionBuilder {
	rb.v.ShardMinDocCount = &shardmindoccount
	return rb
}

func (rb *VertexDefinitionBuilder) Size(size int) *VertexDefinitionBuilder {
	rb.v.Size = &size
	return rb
}
