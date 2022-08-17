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

// RecoveryStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/recovery/types.ts#L98-L100
type RecoveryStatus struct {
	Shards []ShardRecovery `json:"shards"`
}

// RecoveryStatusBuilder holds RecoveryStatus struct and provides a builder API.
type RecoveryStatusBuilder struct {
	v *RecoveryStatus
}

// NewRecoveryStatus provides a builder for the RecoveryStatus struct.
func NewRecoveryStatusBuilder() *RecoveryStatusBuilder {
	r := RecoveryStatusBuilder{
		&RecoveryStatus{},
	}

	return &r
}

// Build finalize the chain and returns the RecoveryStatus struct
func (rb *RecoveryStatusBuilder) Build() RecoveryStatus {
	return *rb.v
}

func (rb *RecoveryStatusBuilder) Shards(shards []ShardRecoveryBuilder) *RecoveryStatusBuilder {
	tmp := make([]ShardRecovery, len(shards))
	for _, value := range shards {
		tmp = append(tmp, value.Build())
	}
	rb.v.Shards = tmp
	return rb
}
