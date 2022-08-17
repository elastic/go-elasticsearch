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

// RecoveryRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/recovery/types.ts#L24-L155
type RecoveryRecord struct {
	// Bytes number of bytes to recover
	Bytes *string `json:"bytes,omitempty"`
	// BytesPercent percent of bytes recovered
	BytesPercent *Percentage `json:"bytes_percent,omitempty"`
	// BytesRecovered bytes recovered
	BytesRecovered *string `json:"bytes_recovered,omitempty"`
	// BytesTotal total number of bytes
	BytesTotal *string `json:"bytes_total,omitempty"`
	// Files number of files to recover
	Files *string `json:"files,omitempty"`
	// FilesPercent percent of files recovered
	FilesPercent *Percentage `json:"files_percent,omitempty"`
	// FilesRecovered files recovered
	FilesRecovered *string `json:"files_recovered,omitempty"`
	// FilesTotal total number of files
	FilesTotal *string `json:"files_total,omitempty"`
	// Index index name
	Index *IndexName `json:"index,omitempty"`
	// Repository repository
	Repository *string `json:"repository,omitempty"`
	// Shard shard name
	Shard *string `json:"shard,omitempty"`
	// Snapshot snapshot
	Snapshot *string `json:"snapshot,omitempty"`
	// SourceHost source host
	SourceHost *string `json:"source_host,omitempty"`
	// SourceNode source node name
	SourceNode *string `json:"source_node,omitempty"`
	// Stage recovery stage
	Stage *string `json:"stage,omitempty"`
	// StartTime recovery start time
	StartTime *DateTime `json:"start_time,omitempty"`
	// StartTimeMillis recovery start time in epoch milliseconds
	StartTimeMillis *EpochTimeUnitMillis `json:"start_time_millis,omitempty"`
	// StopTime recovery stop time
	StopTime *DateTime `json:"stop_time,omitempty"`
	// StopTimeMillis recovery stop time in epoch milliseconds
	StopTimeMillis *EpochTimeUnitMillis `json:"stop_time_millis,omitempty"`
	// TargetHost target host
	TargetHost *string `json:"target_host,omitempty"`
	// TargetNode target node name
	TargetNode *string `json:"target_node,omitempty"`
	// Time recovery time
	Time *Duration `json:"time,omitempty"`
	// TranslogOps number of translog ops to recover
	TranslogOps *string `json:"translog_ops,omitempty"`
	// TranslogOpsPercent percent of translog ops recovered
	TranslogOpsPercent *Percentage `json:"translog_ops_percent,omitempty"`
	// TranslogOpsRecovered translog ops recovered
	TranslogOpsRecovered *string `json:"translog_ops_recovered,omitempty"`
	// Type recovery type
	Type *string `json:"type,omitempty"`
}

// RecoveryRecordBuilder holds RecoveryRecord struct and provides a builder API.
type RecoveryRecordBuilder struct {
	v *RecoveryRecord
}

// NewRecoveryRecord provides a builder for the RecoveryRecord struct.
func NewRecoveryRecordBuilder() *RecoveryRecordBuilder {
	r := RecoveryRecordBuilder{
		&RecoveryRecord{},
	}

	return &r
}

// Build finalize the chain and returns the RecoveryRecord struct
func (rb *RecoveryRecordBuilder) Build() RecoveryRecord {
	return *rb.v
}

// Bytes number of bytes to recover

func (rb *RecoveryRecordBuilder) Bytes(bytes string) *RecoveryRecordBuilder {
	rb.v.Bytes = &bytes
	return rb
}

// BytesPercent percent of bytes recovered

func (rb *RecoveryRecordBuilder) BytesPercent(bytespercent *PercentageBuilder) *RecoveryRecordBuilder {
	v := bytespercent.Build()
	rb.v.BytesPercent = &v
	return rb
}

// BytesRecovered bytes recovered

func (rb *RecoveryRecordBuilder) BytesRecovered(bytesrecovered string) *RecoveryRecordBuilder {
	rb.v.BytesRecovered = &bytesrecovered
	return rb
}

// BytesTotal total number of bytes

func (rb *RecoveryRecordBuilder) BytesTotal(bytestotal string) *RecoveryRecordBuilder {
	rb.v.BytesTotal = &bytestotal
	return rb
}

// Files number of files to recover

func (rb *RecoveryRecordBuilder) Files(files string) *RecoveryRecordBuilder {
	rb.v.Files = &files
	return rb
}

// FilesPercent percent of files recovered

func (rb *RecoveryRecordBuilder) FilesPercent(filespercent *PercentageBuilder) *RecoveryRecordBuilder {
	v := filespercent.Build()
	rb.v.FilesPercent = &v
	return rb
}

// FilesRecovered files recovered

func (rb *RecoveryRecordBuilder) FilesRecovered(filesrecovered string) *RecoveryRecordBuilder {
	rb.v.FilesRecovered = &filesrecovered
	return rb
}

// FilesTotal total number of files

func (rb *RecoveryRecordBuilder) FilesTotal(filestotal string) *RecoveryRecordBuilder {
	rb.v.FilesTotal = &filestotal
	return rb
}

// Index index name

func (rb *RecoveryRecordBuilder) Index(index IndexName) *RecoveryRecordBuilder {
	rb.v.Index = &index
	return rb
}

// Repository repository

func (rb *RecoveryRecordBuilder) Repository(repository string) *RecoveryRecordBuilder {
	rb.v.Repository = &repository
	return rb
}

// Shard shard name

func (rb *RecoveryRecordBuilder) Shard(shard string) *RecoveryRecordBuilder {
	rb.v.Shard = &shard
	return rb
}

// Snapshot snapshot

func (rb *RecoveryRecordBuilder) Snapshot(snapshot string) *RecoveryRecordBuilder {
	rb.v.Snapshot = &snapshot
	return rb
}

// SourceHost source host

func (rb *RecoveryRecordBuilder) SourceHost(sourcehost string) *RecoveryRecordBuilder {
	rb.v.SourceHost = &sourcehost
	return rb
}

// SourceNode source node name

func (rb *RecoveryRecordBuilder) SourceNode(sourcenode string) *RecoveryRecordBuilder {
	rb.v.SourceNode = &sourcenode
	return rb
}

// Stage recovery stage

func (rb *RecoveryRecordBuilder) Stage(stage string) *RecoveryRecordBuilder {
	rb.v.Stage = &stage
	return rb
}

// StartTime recovery start time

func (rb *RecoveryRecordBuilder) StartTime(starttime *DateTimeBuilder) *RecoveryRecordBuilder {
	v := starttime.Build()
	rb.v.StartTime = &v
	return rb
}

// StartTimeMillis recovery start time in epoch milliseconds

func (rb *RecoveryRecordBuilder) StartTimeMillis(starttimemillis *EpochTimeUnitMillisBuilder) *RecoveryRecordBuilder {
	v := starttimemillis.Build()
	rb.v.StartTimeMillis = &v
	return rb
}

// StopTime recovery stop time

func (rb *RecoveryRecordBuilder) StopTime(stoptime *DateTimeBuilder) *RecoveryRecordBuilder {
	v := stoptime.Build()
	rb.v.StopTime = &v
	return rb
}

// StopTimeMillis recovery stop time in epoch milliseconds

func (rb *RecoveryRecordBuilder) StopTimeMillis(stoptimemillis *EpochTimeUnitMillisBuilder) *RecoveryRecordBuilder {
	v := stoptimemillis.Build()
	rb.v.StopTimeMillis = &v
	return rb
}

// TargetHost target host

func (rb *RecoveryRecordBuilder) TargetHost(targethost string) *RecoveryRecordBuilder {
	rb.v.TargetHost = &targethost
	return rb
}

// TargetNode target node name

func (rb *RecoveryRecordBuilder) TargetNode(targetnode string) *RecoveryRecordBuilder {
	rb.v.TargetNode = &targetnode
	return rb
}

// Time recovery time

func (rb *RecoveryRecordBuilder) Time(time *DurationBuilder) *RecoveryRecordBuilder {
	v := time.Build()
	rb.v.Time = &v
	return rb
}

// TranslogOps number of translog ops to recover

func (rb *RecoveryRecordBuilder) TranslogOps(translogops string) *RecoveryRecordBuilder {
	rb.v.TranslogOps = &translogops
	return rb
}

// TranslogOpsPercent percent of translog ops recovered

func (rb *RecoveryRecordBuilder) TranslogOpsPercent(translogopspercent *PercentageBuilder) *RecoveryRecordBuilder {
	v := translogopspercent.Build()
	rb.v.TranslogOpsPercent = &v
	return rb
}

// TranslogOpsRecovered translog ops recovered

func (rb *RecoveryRecordBuilder) TranslogOpsRecovered(translogopsrecovered string) *RecoveryRecordBuilder {
	rb.v.TranslogOpsRecovered = &translogopsrecovered
	return rb
}

// Type recovery type

func (rb *RecoveryRecordBuilder) Type_(type_ string) *RecoveryRecordBuilder {
	rb.v.Type = &type_
	return rb
}
