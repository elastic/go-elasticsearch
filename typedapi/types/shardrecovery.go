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

// ShardRecovery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/recovery/types.ts#L118-L135
type ShardRecovery struct {
	Id                int64                   `json:"id"`
	Index             RecoveryIndexStatus     `json:"index"`
	Primary           bool                    `json:"primary"`
	Source            RecoveryOrigin          `json:"source"`
	Stage             string                  `json:"stage"`
	Start             *RecoveryStartStatus    `json:"start,omitempty"`
	StartTime         *DateTime               `json:"start_time,omitempty"`
	StartTimeInMillis EpochTimeUnitMillis     `json:"start_time_in_millis"`
	StopTime          *DateTime               `json:"stop_time,omitempty"`
	StopTimeInMillis  *EpochTimeUnitMillis    `json:"stop_time_in_millis,omitempty"`
	Target            RecoveryOrigin          `json:"target"`
	TotalTime         *Duration               `json:"total_time,omitempty"`
	TotalTimeInMillis DurationValueUnitMillis `json:"total_time_in_millis"`
	Translog          TranslogStatus          `json:"translog"`
	Type              string                  `json:"type"`
	VerifyIndex       VerifyIndex             `json:"verify_index"`
}

// ShardRecoveryBuilder holds ShardRecovery struct and provides a builder API.
type ShardRecoveryBuilder struct {
	v *ShardRecovery
}

// NewShardRecovery provides a builder for the ShardRecovery struct.
func NewShardRecoveryBuilder() *ShardRecoveryBuilder {
	r := ShardRecoveryBuilder{
		&ShardRecovery{},
	}

	return &r
}

// Build finalize the chain and returns the ShardRecovery struct
func (rb *ShardRecoveryBuilder) Build() ShardRecovery {
	return *rb.v
}

func (rb *ShardRecoveryBuilder) Id(id int64) *ShardRecoveryBuilder {
	rb.v.Id = id
	return rb
}

func (rb *ShardRecoveryBuilder) Index(index *RecoveryIndexStatusBuilder) *ShardRecoveryBuilder {
	v := index.Build()
	rb.v.Index = v
	return rb
}

func (rb *ShardRecoveryBuilder) Primary(primary bool) *ShardRecoveryBuilder {
	rb.v.Primary = primary
	return rb
}

func (rb *ShardRecoveryBuilder) Source(source *RecoveryOriginBuilder) *ShardRecoveryBuilder {
	v := source.Build()
	rb.v.Source = v
	return rb
}

func (rb *ShardRecoveryBuilder) Stage(stage string) *ShardRecoveryBuilder {
	rb.v.Stage = stage
	return rb
}

func (rb *ShardRecoveryBuilder) Start(start *RecoveryStartStatusBuilder) *ShardRecoveryBuilder {
	v := start.Build()
	rb.v.Start = &v
	return rb
}

func (rb *ShardRecoveryBuilder) StartTime(starttime *DateTimeBuilder) *ShardRecoveryBuilder {
	v := starttime.Build()
	rb.v.StartTime = &v
	return rb
}

func (rb *ShardRecoveryBuilder) StartTimeInMillis(starttimeinmillis *EpochTimeUnitMillisBuilder) *ShardRecoveryBuilder {
	v := starttimeinmillis.Build()
	rb.v.StartTimeInMillis = v
	return rb
}

func (rb *ShardRecoveryBuilder) StopTime(stoptime *DateTimeBuilder) *ShardRecoveryBuilder {
	v := stoptime.Build()
	rb.v.StopTime = &v
	return rb
}

func (rb *ShardRecoveryBuilder) StopTimeInMillis(stoptimeinmillis *EpochTimeUnitMillisBuilder) *ShardRecoveryBuilder {
	v := stoptimeinmillis.Build()
	rb.v.StopTimeInMillis = &v
	return rb
}

func (rb *ShardRecoveryBuilder) Target(target *RecoveryOriginBuilder) *ShardRecoveryBuilder {
	v := target.Build()
	rb.v.Target = v
	return rb
}

func (rb *ShardRecoveryBuilder) TotalTime(totaltime *DurationBuilder) *ShardRecoveryBuilder {
	v := totaltime.Build()
	rb.v.TotalTime = &v
	return rb
}

func (rb *ShardRecoveryBuilder) TotalTimeInMillis(totaltimeinmillis *DurationValueUnitMillisBuilder) *ShardRecoveryBuilder {
	v := totaltimeinmillis.Build()
	rb.v.TotalTimeInMillis = v
	return rb
}

func (rb *ShardRecoveryBuilder) Translog(translog *TranslogStatusBuilder) *ShardRecoveryBuilder {
	v := translog.Build()
	rb.v.Translog = v
	return rb
}

func (rb *ShardRecoveryBuilder) Type_(type_ string) *ShardRecoveryBuilder {
	rb.v.Type = type_
	return rb
}

func (rb *ShardRecoveryBuilder) VerifyIndex(verifyindex *VerifyIndexBuilder) *ShardRecoveryBuilder {
	v := verifyindex.Build()
	rb.v.VerifyIndex = v
	return rb
}
