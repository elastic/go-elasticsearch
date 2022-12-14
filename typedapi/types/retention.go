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
// https://github.com/elastic/elasticsearch-specification/tree/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33


package types

// Retention type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/slm/_types/SnapshotLifecycle.ts#L84-L97
type Retention struct {
	// ExpireAfter Time period after which a snapshot is considered expired and eligible for
	// deletion. SLM deletes expired snapshots based on the slm.retention_schedule.
	ExpireAfter Duration `json:"expire_after"`
	// MaxCount Maximum number of snapshots to retain, even if the snapshots have not yet
	// expired. If the number of snapshots in the repository exceeds this limit, the
	// policy retains the most recent snapshots and deletes older snapshots.
	MaxCount int `json:"max_count"`
	// MinCount Minimum number of snapshots to retain, even if the snapshots have expired.
	MinCount int `json:"min_count"`
}

// NewRetention returns a Retention.
func NewRetention() *Retention {
	r := &Retention{}

	return r
}
