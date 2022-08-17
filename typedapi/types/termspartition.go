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

// TermsPartition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L421-L424
type TermsPartition struct {
	NumPartitions int64 `json:"num_partitions"`
	Partition     int64 `json:"partition"`
}

// TermsPartitionBuilder holds TermsPartition struct and provides a builder API.
type TermsPartitionBuilder struct {
	v *TermsPartition
}

// NewTermsPartition provides a builder for the TermsPartition struct.
func NewTermsPartitionBuilder() *TermsPartitionBuilder {
	r := TermsPartitionBuilder{
		&TermsPartition{},
	}

	return &r
}

// Build finalize the chain and returns the TermsPartition struct
func (rb *TermsPartitionBuilder) Build() TermsPartition {
	return *rb.v
}

func (rb *TermsPartitionBuilder) NumPartitions(numpartitions int64) *TermsPartitionBuilder {
	rb.v.NumPartitions = numpartitions
	return rb
}

func (rb *TermsPartitionBuilder) Partition(partition int64) *TermsPartitionBuilder {
	rb.v.Partition = partition
	return rb
}
