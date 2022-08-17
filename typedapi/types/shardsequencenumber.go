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

// ShardSequenceNumber type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/stats/types.ts#L164-L168
type ShardSequenceNumber struct {
	GlobalCheckpoint int64          `json:"global_checkpoint"`
	LocalCheckpoint  int64          `json:"local_checkpoint"`
	MaxSeqNo         SequenceNumber `json:"max_seq_no"`
}

// ShardSequenceNumberBuilder holds ShardSequenceNumber struct and provides a builder API.
type ShardSequenceNumberBuilder struct {
	v *ShardSequenceNumber
}

// NewShardSequenceNumber provides a builder for the ShardSequenceNumber struct.
func NewShardSequenceNumberBuilder() *ShardSequenceNumberBuilder {
	r := ShardSequenceNumberBuilder{
		&ShardSequenceNumber{},
	}

	return &r
}

// Build finalize the chain and returns the ShardSequenceNumber struct
func (rb *ShardSequenceNumberBuilder) Build() ShardSequenceNumber {
	return *rb.v
}

func (rb *ShardSequenceNumberBuilder) GlobalCheckpoint(globalcheckpoint int64) *ShardSequenceNumberBuilder {
	rb.v.GlobalCheckpoint = globalcheckpoint
	return rb
}

func (rb *ShardSequenceNumberBuilder) LocalCheckpoint(localcheckpoint int64) *ShardSequenceNumberBuilder {
	rb.v.LocalCheckpoint = localcheckpoint
	return rb
}

func (rb *ShardSequenceNumberBuilder) MaxSeqNo(maxseqno SequenceNumber) *ShardSequenceNumberBuilder {
	rb.v.MaxSeqNo = maxseqno
	return rb
}
