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

// IndexSegment type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/segments/types.ts#L24-L26
type IndexSegment struct {
	Shards map[string][]ShardsSegment `json:"shards"`
}

// IndexSegmentBuilder holds IndexSegment struct and provides a builder API.
type IndexSegmentBuilder struct {
	v *IndexSegment
}

// NewIndexSegment provides a builder for the IndexSegment struct.
func NewIndexSegmentBuilder() *IndexSegmentBuilder {
	r := IndexSegmentBuilder{
		&IndexSegment{
			Shards: make(map[string][]ShardsSegment, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the IndexSegment struct
func (rb *IndexSegmentBuilder) Build() IndexSegment {
	return *rb.v
}

func (rb *IndexSegmentBuilder) Shards(value map[string][]ShardsSegment) *IndexSegmentBuilder {
	rb.v.Shards = value
	return rb
}
