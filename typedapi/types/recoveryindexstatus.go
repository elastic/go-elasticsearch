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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// RecoveryIndexStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/indices/recovery/types.ts#L58-L68
type RecoveryIndexStatus struct {
	Bytes                      *RecoveryBytes `json:"bytes,omitempty"`
	Files                      RecoveryFiles  `json:"files"`
	Size                       RecoveryBytes  `json:"size"`
	SourceThrottleTime         *Time          `json:"source_throttle_time,omitempty"`
	SourceThrottleTimeInMillis EpochMillis    `json:"source_throttle_time_in_millis"`
	TargetThrottleTime         *Time          `json:"target_throttle_time,omitempty"`
	TargetThrottleTimeInMillis EpochMillis    `json:"target_throttle_time_in_millis"`
	TotalTime                  *Time          `json:"total_time,omitempty"`
	TotalTimeInMillis          EpochMillis    `json:"total_time_in_millis"`
}

// RecoveryIndexStatusBuilder holds RecoveryIndexStatus struct and provides a builder API.
type RecoveryIndexStatusBuilder struct {
	v *RecoveryIndexStatus
}

// NewRecoveryIndexStatus provides a builder for the RecoveryIndexStatus struct.
func NewRecoveryIndexStatusBuilder() *RecoveryIndexStatusBuilder {
	r := RecoveryIndexStatusBuilder{
		&RecoveryIndexStatus{},
	}

	return &r
}

// Build finalize the chain and returns the RecoveryIndexStatus struct
func (rb *RecoveryIndexStatusBuilder) Build() RecoveryIndexStatus {
	return *rb.v
}

func (rb *RecoveryIndexStatusBuilder) Bytes(bytes *RecoveryBytesBuilder) *RecoveryIndexStatusBuilder {
	v := bytes.Build()
	rb.v.Bytes = &v
	return rb
}

func (rb *RecoveryIndexStatusBuilder) Files(files *RecoveryFilesBuilder) *RecoveryIndexStatusBuilder {
	v := files.Build()
	rb.v.Files = v
	return rb
}

func (rb *RecoveryIndexStatusBuilder) Size(size *RecoveryBytesBuilder) *RecoveryIndexStatusBuilder {
	v := size.Build()
	rb.v.Size = v
	return rb
}

func (rb *RecoveryIndexStatusBuilder) SourceThrottleTime(sourcethrottletime *TimeBuilder) *RecoveryIndexStatusBuilder {
	v := sourcethrottletime.Build()
	rb.v.SourceThrottleTime = &v
	return rb
}

func (rb *RecoveryIndexStatusBuilder) SourceThrottleTimeInMillis(sourcethrottletimeinmillis *EpochMillisBuilder) *RecoveryIndexStatusBuilder {
	v := sourcethrottletimeinmillis.Build()
	rb.v.SourceThrottleTimeInMillis = v
	return rb
}

func (rb *RecoveryIndexStatusBuilder) TargetThrottleTime(targetthrottletime *TimeBuilder) *RecoveryIndexStatusBuilder {
	v := targetthrottletime.Build()
	rb.v.TargetThrottleTime = &v
	return rb
}

func (rb *RecoveryIndexStatusBuilder) TargetThrottleTimeInMillis(targetthrottletimeinmillis *EpochMillisBuilder) *RecoveryIndexStatusBuilder {
	v := targetthrottletimeinmillis.Build()
	rb.v.TargetThrottleTimeInMillis = v
	return rb
}

func (rb *RecoveryIndexStatusBuilder) TotalTime(totaltime *TimeBuilder) *RecoveryIndexStatusBuilder {
	v := totaltime.Build()
	rb.v.TotalTime = &v
	return rb
}

func (rb *RecoveryIndexStatusBuilder) TotalTimeInMillis(totaltimeinmillis *EpochMillisBuilder) *RecoveryIndexStatusBuilder {
	v := totaltimeinmillis.Build()
	rb.v.TotalTimeInMillis = v
	return rb
}
