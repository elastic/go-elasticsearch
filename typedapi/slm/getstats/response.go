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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package getstats

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package getstats
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/slm/get_stats/GetSnapshotLifecycleStatsResponse.ts#L23-L36

type Response struct {
	PolicyStats                   []string       `json:"policy_stats"`
	RetentionDeletionTime         types.Duration `json:"retention_deletion_time"`
	RetentionDeletionTimeMillis   int64          `json:"retention_deletion_time_millis"`
	RetentionFailed               int64          `json:"retention_failed"`
	RetentionRuns                 int64          `json:"retention_runs"`
	RetentionTimedOut             int64          `json:"retention_timed_out"`
	TotalSnapshotDeletionFailures int64          `json:"total_snapshot_deletion_failures"`
	TotalSnapshotsDeleted         int64          `json:"total_snapshots_deleted"`
	TotalSnapshotsFailed          int64          `json:"total_snapshots_failed"`
	TotalSnapshotsTaken           int64          `json:"total_snapshots_taken"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
