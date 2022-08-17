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

// ModelSnapshot type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Model.ts#L25-L46
type ModelSnapshot struct {
	// Description An optional description of the job.
	Description *string `json:"description,omitempty"`
	// JobId A numerical character string that uniquely identifies the job that the
	// snapshot was created for.
	JobId Id `json:"job_id"`
	// LatestRecordTimeStamp The timestamp of the latest processed record.
	LatestRecordTimeStamp *int `json:"latest_record_time_stamp,omitempty"`
	// LatestResultTimeStamp The timestamp of the latest bucket result.
	LatestResultTimeStamp *int `json:"latest_result_time_stamp,omitempty"`
	// MinVersion The minimum version required to be able to restore the model snapshot.
	MinVersion VersionString `json:"min_version"`
	// ModelSizeStats Summary information describing the model.
	ModelSizeStats *ModelSizeStats `json:"model_size_stats,omitempty"`
	// Retain If true, this snapshot will not be deleted during automatic cleanup of
	// snapshots older than model_snapshot_retention_days. However, this snapshot
	// will be deleted when the job is deleted. The default value is false.
	Retain bool `json:"retain"`
	// SnapshotDocCount For internal use only.
	SnapshotDocCount int64 `json:"snapshot_doc_count"`
	// SnapshotId A numerical character string that uniquely identifies the model snapshot.
	SnapshotId Id `json:"snapshot_id"`
	// Timestamp The creation timestamp for the snapshot.
	Timestamp int64 `json:"timestamp"`
}

// ModelSnapshotBuilder holds ModelSnapshot struct and provides a builder API.
type ModelSnapshotBuilder struct {
	v *ModelSnapshot
}

// NewModelSnapshot provides a builder for the ModelSnapshot struct.
func NewModelSnapshotBuilder() *ModelSnapshotBuilder {
	r := ModelSnapshotBuilder{
		&ModelSnapshot{},
	}

	return &r
}

// Build finalize the chain and returns the ModelSnapshot struct
func (rb *ModelSnapshotBuilder) Build() ModelSnapshot {
	return *rb.v
}

// Description An optional description of the job.

func (rb *ModelSnapshotBuilder) Description(description string) *ModelSnapshotBuilder {
	rb.v.Description = &description
	return rb
}

// JobId A numerical character string that uniquely identifies the job that the
// snapshot was created for.

func (rb *ModelSnapshotBuilder) JobId(jobid Id) *ModelSnapshotBuilder {
	rb.v.JobId = jobid
	return rb
}

// LatestRecordTimeStamp The timestamp of the latest processed record.

func (rb *ModelSnapshotBuilder) LatestRecordTimeStamp(latestrecordtimestamp int) *ModelSnapshotBuilder {
	rb.v.LatestRecordTimeStamp = &latestrecordtimestamp
	return rb
}

// LatestResultTimeStamp The timestamp of the latest bucket result.

func (rb *ModelSnapshotBuilder) LatestResultTimeStamp(latestresulttimestamp int) *ModelSnapshotBuilder {
	rb.v.LatestResultTimeStamp = &latestresulttimestamp
	return rb
}

// MinVersion The minimum version required to be able to restore the model snapshot.

func (rb *ModelSnapshotBuilder) MinVersion(minversion VersionString) *ModelSnapshotBuilder {
	rb.v.MinVersion = minversion
	return rb
}

// ModelSizeStats Summary information describing the model.

func (rb *ModelSnapshotBuilder) ModelSizeStats(modelsizestats *ModelSizeStatsBuilder) *ModelSnapshotBuilder {
	v := modelsizestats.Build()
	rb.v.ModelSizeStats = &v
	return rb
}

// Retain If true, this snapshot will not be deleted during automatic cleanup of
// snapshots older than model_snapshot_retention_days. However, this snapshot
// will be deleted when the job is deleted. The default value is false.

func (rb *ModelSnapshotBuilder) Retain(retain bool) *ModelSnapshotBuilder {
	rb.v.Retain = retain
	return rb
}

// SnapshotDocCount For internal use only.

func (rb *ModelSnapshotBuilder) SnapshotDocCount(snapshotdoccount int64) *ModelSnapshotBuilder {
	rb.v.SnapshotDocCount = snapshotdoccount
	return rb
}

// SnapshotId A numerical character string that uniquely identifies the model snapshot.

func (rb *ModelSnapshotBuilder) SnapshotId(snapshotid Id) *ModelSnapshotBuilder {
	rb.v.SnapshotId = snapshotid
	return rb
}

// Timestamp The creation timestamp for the snapshot.

func (rb *ModelSnapshotBuilder) Timestamp(timestamp int64) *ModelSnapshotBuilder {
	rb.v.Timestamp = timestamp
	return rb
}
