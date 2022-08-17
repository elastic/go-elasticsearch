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

// Statistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/slm/_types/SnapshotLifecycle.ts#L51-L74
type Statistics struct {
	Policy                        *Id                      `json:"policy,omitempty"`
	RetentionDeletionTime         *Duration                `json:"retention_deletion_time,omitempty"`
	RetentionDeletionTimeMillis   *DurationValueUnitMillis `json:"retention_deletion_time_millis,omitempty"`
	RetentionFailed               *int64                   `json:"retention_failed,omitempty"`
	RetentionRuns                 *int64                   `json:"retention_runs,omitempty"`
	RetentionTimedOut             *int64                   `json:"retention_timed_out,omitempty"`
	TotalSnapshotDeletionFailures *int64                   `json:"total_snapshot_deletion_failures,omitempty"`
	TotalSnapshotsDeleted         *int64                   `json:"total_snapshots_deleted,omitempty"`
	TotalSnapshotsFailed          *int64                   `json:"total_snapshots_failed,omitempty"`
	TotalSnapshotsTaken           *int64                   `json:"total_snapshots_taken,omitempty"`
}

// StatisticsBuilder holds Statistics struct and provides a builder API.
type StatisticsBuilder struct {
	v *Statistics
}

// NewStatistics provides a builder for the Statistics struct.
func NewStatisticsBuilder() *StatisticsBuilder {
	r := StatisticsBuilder{
		&Statistics{},
	}

	return &r
}

// Build finalize the chain and returns the Statistics struct
func (rb *StatisticsBuilder) Build() Statistics {
	return *rb.v
}

func (rb *StatisticsBuilder) Policy(policy Id) *StatisticsBuilder {
	rb.v.Policy = &policy
	return rb
}

func (rb *StatisticsBuilder) RetentionDeletionTime(retentiondeletiontime *DurationBuilder) *StatisticsBuilder {
	v := retentiondeletiontime.Build()
	rb.v.RetentionDeletionTime = &v
	return rb
}

func (rb *StatisticsBuilder) RetentionDeletionTimeMillis(retentiondeletiontimemillis *DurationValueUnitMillisBuilder) *StatisticsBuilder {
	v := retentiondeletiontimemillis.Build()
	rb.v.RetentionDeletionTimeMillis = &v
	return rb
}

func (rb *StatisticsBuilder) RetentionFailed(retentionfailed int64) *StatisticsBuilder {
	rb.v.RetentionFailed = &retentionfailed
	return rb
}

func (rb *StatisticsBuilder) RetentionRuns(retentionruns int64) *StatisticsBuilder {
	rb.v.RetentionRuns = &retentionruns
	return rb
}

func (rb *StatisticsBuilder) RetentionTimedOut(retentiontimedout int64) *StatisticsBuilder {
	rb.v.RetentionTimedOut = &retentiontimedout
	return rb
}

func (rb *StatisticsBuilder) TotalSnapshotDeletionFailures(totalsnapshotdeletionfailures int64) *StatisticsBuilder {
	rb.v.TotalSnapshotDeletionFailures = &totalsnapshotdeletionfailures
	return rb
}

func (rb *StatisticsBuilder) TotalSnapshotsDeleted(totalsnapshotsdeleted int64) *StatisticsBuilder {
	rb.v.TotalSnapshotsDeleted = &totalsnapshotsdeleted
	return rb
}

func (rb *StatisticsBuilder) TotalSnapshotsFailed(totalsnapshotsfailed int64) *StatisticsBuilder {
	rb.v.TotalSnapshotsFailed = &totalsnapshotsfailed
	return rb
}

func (rb *StatisticsBuilder) TotalSnapshotsTaken(totalsnapshotstaken int64) *StatisticsBuilder {
	rb.v.TotalSnapshotsTaken = &totalsnapshotstaken
	return rb
}
