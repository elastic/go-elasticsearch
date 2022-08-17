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

// SnapshotInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/snapshot/_types/SnapshotInfo.ts#L41-L65
type SnapshotInfo struct {
	DataStreams        []string                   `json:"data_streams"`
	Duration           *Duration                  `json:"duration,omitempty"`
	DurationInMillis   *DurationValueUnitMillis   `json:"duration_in_millis,omitempty"`
	EndTime            *DateTime                  `json:"end_time,omitempty"`
	EndTimeInMillis    *EpochTimeUnitMillis       `json:"end_time_in_millis,omitempty"`
	Failures           []SnapshotShardFailure     `json:"failures,omitempty"`
	FeatureStates      []InfoFeatureState         `json:"feature_states,omitempty"`
	IncludeGlobalState *bool                      `json:"include_global_state,omitempty"`
	IndexDetails       map[IndexName]IndexDetails `json:"index_details,omitempty"`
	Indices            []IndexName                `json:"indices,omitempty"`
	Metadata           *Metadata                  `json:"metadata,omitempty"`
	Reason             *string                    `json:"reason,omitempty"`
	Repository         *Name                      `json:"repository,omitempty"`
	Shards             *ShardStatistics           `json:"shards,omitempty"`
	Snapshot           Name                       `json:"snapshot"`
	StartTime          *DateTime                  `json:"start_time,omitempty"`
	StartTimeInMillis  *EpochTimeUnitMillis       `json:"start_time_in_millis,omitempty"`
	State              *string                    `json:"state,omitempty"`
	Uuid               Uuid                       `json:"uuid"`
	Version            *VersionString             `json:"version,omitempty"`
	VersionId          *VersionNumber             `json:"version_id,omitempty"`
}

// SnapshotInfoBuilder holds SnapshotInfo struct and provides a builder API.
type SnapshotInfoBuilder struct {
	v *SnapshotInfo
}

// NewSnapshotInfo provides a builder for the SnapshotInfo struct.
func NewSnapshotInfoBuilder() *SnapshotInfoBuilder {
	r := SnapshotInfoBuilder{
		&SnapshotInfo{
			IndexDetails: make(map[IndexName]IndexDetails, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the SnapshotInfo struct
func (rb *SnapshotInfoBuilder) Build() SnapshotInfo {
	return *rb.v
}

func (rb *SnapshotInfoBuilder) DataStreams(data_streams ...string) *SnapshotInfoBuilder {
	rb.v.DataStreams = data_streams
	return rb
}

func (rb *SnapshotInfoBuilder) Duration(duration *DurationBuilder) *SnapshotInfoBuilder {
	v := duration.Build()
	rb.v.Duration = &v
	return rb
}

func (rb *SnapshotInfoBuilder) DurationInMillis(durationinmillis *DurationValueUnitMillisBuilder) *SnapshotInfoBuilder {
	v := durationinmillis.Build()
	rb.v.DurationInMillis = &v
	return rb
}

func (rb *SnapshotInfoBuilder) EndTime(endtime *DateTimeBuilder) *SnapshotInfoBuilder {
	v := endtime.Build()
	rb.v.EndTime = &v
	return rb
}

func (rb *SnapshotInfoBuilder) EndTimeInMillis(endtimeinmillis *EpochTimeUnitMillisBuilder) *SnapshotInfoBuilder {
	v := endtimeinmillis.Build()
	rb.v.EndTimeInMillis = &v
	return rb
}

func (rb *SnapshotInfoBuilder) Failures(failures []SnapshotShardFailureBuilder) *SnapshotInfoBuilder {
	tmp := make([]SnapshotShardFailure, len(failures))
	for _, value := range failures {
		tmp = append(tmp, value.Build())
	}
	rb.v.Failures = tmp
	return rb
}

func (rb *SnapshotInfoBuilder) FeatureStates(feature_states []InfoFeatureStateBuilder) *SnapshotInfoBuilder {
	tmp := make([]InfoFeatureState, len(feature_states))
	for _, value := range feature_states {
		tmp = append(tmp, value.Build())
	}
	rb.v.FeatureStates = tmp
	return rb
}

func (rb *SnapshotInfoBuilder) IncludeGlobalState(includeglobalstate bool) *SnapshotInfoBuilder {
	rb.v.IncludeGlobalState = &includeglobalstate
	return rb
}

func (rb *SnapshotInfoBuilder) IndexDetails(values map[IndexName]*IndexDetailsBuilder) *SnapshotInfoBuilder {
	tmp := make(map[IndexName]IndexDetails, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.IndexDetails = tmp
	return rb
}

func (rb *SnapshotInfoBuilder) Indices(indices ...IndexName) *SnapshotInfoBuilder {
	rb.v.Indices = indices
	return rb
}

func (rb *SnapshotInfoBuilder) Metadata(metadata *MetadataBuilder) *SnapshotInfoBuilder {
	v := metadata.Build()
	rb.v.Metadata = &v
	return rb
}

func (rb *SnapshotInfoBuilder) Reason(reason string) *SnapshotInfoBuilder {
	rb.v.Reason = &reason
	return rb
}

func (rb *SnapshotInfoBuilder) Repository(repository Name) *SnapshotInfoBuilder {
	rb.v.Repository = &repository
	return rb
}

func (rb *SnapshotInfoBuilder) Shards(shards *ShardStatisticsBuilder) *SnapshotInfoBuilder {
	v := shards.Build()
	rb.v.Shards = &v
	return rb
}

func (rb *SnapshotInfoBuilder) Snapshot(snapshot Name) *SnapshotInfoBuilder {
	rb.v.Snapshot = snapshot
	return rb
}

func (rb *SnapshotInfoBuilder) StartTime(starttime *DateTimeBuilder) *SnapshotInfoBuilder {
	v := starttime.Build()
	rb.v.StartTime = &v
	return rb
}

func (rb *SnapshotInfoBuilder) StartTimeInMillis(starttimeinmillis *EpochTimeUnitMillisBuilder) *SnapshotInfoBuilder {
	v := starttimeinmillis.Build()
	rb.v.StartTimeInMillis = &v
	return rb
}

func (rb *SnapshotInfoBuilder) State(state string) *SnapshotInfoBuilder {
	rb.v.State = &state
	return rb
}

func (rb *SnapshotInfoBuilder) Uuid(uuid Uuid) *SnapshotInfoBuilder {
	rb.v.Uuid = uuid
	return rb
}

func (rb *SnapshotInfoBuilder) Version(version VersionString) *SnapshotInfoBuilder {
	rb.v.Version = &version
	return rb
}

func (rb *SnapshotInfoBuilder) VersionId(versionid VersionNumber) *SnapshotInfoBuilder {
	rb.v.VersionId = &versionid
	return rb
}
