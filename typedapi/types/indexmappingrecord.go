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

// IndexMappingRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/get_mapping/IndicesGetMappingResponse.ts#L28-L31
type IndexMappingRecord struct {
	Item     *TypeMapping `json:"item,omitempty"`
	Mappings TypeMapping  `json:"mappings"`
}

// IndexMappingRecordBuilder holds IndexMappingRecord struct and provides a builder API.
type IndexMappingRecordBuilder struct {
	v *IndexMappingRecord
}

// NewIndexMappingRecord provides a builder for the IndexMappingRecord struct.
func NewIndexMappingRecordBuilder() *IndexMappingRecordBuilder {
	r := IndexMappingRecordBuilder{
		&IndexMappingRecord{},
	}

	return &r
}

// Build finalize the chain and returns the IndexMappingRecord struct
func (rb *IndexMappingRecordBuilder) Build() IndexMappingRecord {
	return *rb.v
}

func (rb *IndexMappingRecordBuilder) Item(item *TypeMappingBuilder) *IndexMappingRecordBuilder {
	v := item.Build()
	rb.v.Item = &v
	return rb
}

func (rb *IndexMappingRecordBuilder) Mappings(mappings *TypeMappingBuilder) *IndexMappingRecordBuilder {
	v := mappings.Build()
	rb.v.Mappings = v
	return rb
}
