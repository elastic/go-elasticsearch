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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _termsPartition struct {
	v *types.TermsPartition
}

func NewTermsPartition(numpartitions int64, partition int64) *_termsPartition {

	tmp := &_termsPartition{v: types.NewTermsPartition()}

	tmp.NumPartitions(numpartitions)

	tmp.Partition(partition)

	return tmp

}

// The number of partitions.
func (s *_termsPartition) NumPartitions(numpartitions int64) *_termsPartition {

	s.v.NumPartitions = numpartitions

	return s
}

// The partition number for this request.
func (s *_termsPartition) Partition(partition int64) *_termsPartition {

	s.v.Partition = partition

	return s
}

func (s *_termsPartition) TermsPartitionCaster() *types.TermsPartition {
	return s.v
}
