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

// IndexingPressure type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L55-L57
type IndexingPressure struct {
	Memory *IndexingPressureMemory `json:"memory,omitempty"`
}

// IndexingPressureBuilder holds IndexingPressure struct and provides a builder API.
type IndexingPressureBuilder struct {
	v *IndexingPressure
}

// NewIndexingPressure provides a builder for the IndexingPressure struct.
func NewIndexingPressureBuilder() *IndexingPressureBuilder {
	r := IndexingPressureBuilder{
		&IndexingPressure{},
	}

	return &r
}

// Build finalize the chain and returns the IndexingPressure struct
func (rb *IndexingPressureBuilder) Build() IndexingPressure {
	return *rb.v
}

func (rb *IndexingPressureBuilder) Memory(memory *IndexingPressureMemoryBuilder) *IndexingPressureBuilder {
	v := memory.Build()
	rb.v.Memory = &v
	return rb
}
