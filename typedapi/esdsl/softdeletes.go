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

type _softDeletes struct {
	v *types.SoftDeletes
}

func NewSoftDeletes() *_softDeletes {

	return &_softDeletes{v: types.NewSoftDeletes()}

}

// Indicates whether soft deletes are enabled on the index.
func (s *_softDeletes) Enabled(enabled bool) *_softDeletes {

	s.v.Enabled = &enabled

	return s
}

// The maximum period to retain a shard history retention lease before it is
// considered expired.
// Shard history retention leases ensure that soft deletes are retained during
// merges on the Lucene
// index. If a soft delete is merged away before it can be replicated to a
// follower the following
// process will fail due to incomplete history on the leader.
func (s *_softDeletes) RetentionLease(retentionlease types.RetentionLeaseVariant) *_softDeletes {

	s.v.RetentionLease = retentionlease.RetentionLeaseCaster()

	return s
}

func (s *_softDeletes) SoftDeletesCaster() *types.SoftDeletes {
	return s.v
}
