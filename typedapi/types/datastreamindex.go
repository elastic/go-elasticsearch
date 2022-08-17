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

// DataStreamIndex type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/DataStream.ts#L52-L55
type DataStreamIndex struct {
	IndexName IndexName `json:"index_name"`
	IndexUuid Uuid      `json:"index_uuid"`
}

// DataStreamIndexBuilder holds DataStreamIndex struct and provides a builder API.
type DataStreamIndexBuilder struct {
	v *DataStreamIndex
}

// NewDataStreamIndex provides a builder for the DataStreamIndex struct.
func NewDataStreamIndexBuilder() *DataStreamIndexBuilder {
	r := DataStreamIndexBuilder{
		&DataStreamIndex{},
	}

	return &r
}

// Build finalize the chain and returns the DataStreamIndex struct
func (rb *DataStreamIndexBuilder) Build() DataStreamIndex {
	return *rb.v
}

func (rb *DataStreamIndexBuilder) IndexName(indexname IndexName) *DataStreamIndexBuilder {
	rb.v.IndexName = indexname
	return rb
}

func (rb *DataStreamIndexBuilder) IndexUuid(indexuuid Uuid) *DataStreamIndexBuilder {
	rb.v.IndexUuid = indexuuid
	return rb
}
