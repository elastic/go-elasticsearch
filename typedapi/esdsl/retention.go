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

type _retention struct {
	v *types.Retention
}

func NewRetention(maxcount int, mincount int) *_retention {

	tmp := &_retention{v: types.NewRetention()}

	tmp.MaxCount(maxcount)

	tmp.MinCount(mincount)

	return tmp

}

// Time period after which a snapshot is considered expired and eligible for
// deletion. SLM deletes expired snapshots based on the slm.retention_schedule.
func (s *_retention) ExpireAfter(duration types.DurationVariant) *_retention {

	s.v.ExpireAfter = *duration.DurationCaster()

	return s
}

// Maximum number of snapshots to retain, even if the snapshots have not yet
// expired. If the number of snapshots in the repository exceeds this limit, the
// policy retains the most recent snapshots and deletes older snapshots.
func (s *_retention) MaxCount(maxcount int) *_retention {

	s.v.MaxCount = maxcount

	return s
}

// Minimum number of snapshots to retain, even if the snapshots have expired.
func (s *_retention) MinCount(mincount int) *_retention {

	s.v.MinCount = mincount

	return s
}

func (s *_retention) RetentionCaster() *types.Retention {
	return s.v
}
