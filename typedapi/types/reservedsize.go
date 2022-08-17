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

// ReservedSize type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/allocation_explain/types.ts#L71-L76
type ReservedSize struct {
	NodeId Id       `json:"node_id"`
	Path   string   `json:"path"`
	Shards []string `json:"shards"`
	Total  int64    `json:"total"`
}

// ReservedSizeBuilder holds ReservedSize struct and provides a builder API.
type ReservedSizeBuilder struct {
	v *ReservedSize
}

// NewReservedSize provides a builder for the ReservedSize struct.
func NewReservedSizeBuilder() *ReservedSizeBuilder {
	r := ReservedSizeBuilder{
		&ReservedSize{},
	}

	return &r
}

// Build finalize the chain and returns the ReservedSize struct
func (rb *ReservedSizeBuilder) Build() ReservedSize {
	return *rb.v
}

func (rb *ReservedSizeBuilder) NodeId(nodeid Id) *ReservedSizeBuilder {
	rb.v.NodeId = nodeid
	return rb
}

func (rb *ReservedSizeBuilder) Path(path string) *ReservedSizeBuilder {
	rb.v.Path = path
	return rb
}

func (rb *ReservedSizeBuilder) Shards(shards ...string) *ReservedSizeBuilder {
	rb.v.Shards = shards
	return rb
}

func (rb *ReservedSizeBuilder) Total(total int64) *ReservedSizeBuilder {
	rb.v.Total = total
	return rb
}
