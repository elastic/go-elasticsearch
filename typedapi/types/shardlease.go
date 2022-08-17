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

// ShardLease type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/stats/types.ts#L121-L126
type ShardLease struct {
	Id             Id             `json:"id"`
	RetainingSeqNo SequenceNumber `json:"retaining_seq_no"`
	Source         string         `json:"source"`
	Timestamp      int64          `json:"timestamp"`
}

// ShardLeaseBuilder holds ShardLease struct and provides a builder API.
type ShardLeaseBuilder struct {
	v *ShardLease
}

// NewShardLease provides a builder for the ShardLease struct.
func NewShardLeaseBuilder() *ShardLeaseBuilder {
	r := ShardLeaseBuilder{
		&ShardLease{},
	}

	return &r
}

// Build finalize the chain and returns the ShardLease struct
func (rb *ShardLeaseBuilder) Build() ShardLease {
	return *rb.v
}

func (rb *ShardLeaseBuilder) Id(id Id) *ShardLeaseBuilder {
	rb.v.Id = id
	return rb
}

func (rb *ShardLeaseBuilder) RetainingSeqNo(retainingseqno SequenceNumber) *ShardLeaseBuilder {
	rb.v.RetainingSeqNo = retainingseqno
	return rb
}

func (rb *ShardLeaseBuilder) Source(source string) *ShardLeaseBuilder {
	rb.v.Source = source
	return rb
}

func (rb *ShardLeaseBuilder) Timestamp(timestamp int64) *ShardLeaseBuilder {
	rb.v.Timestamp = timestamp
	return rb
}
