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

// RecoveryIndexStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/recovery/types.ts#L64-L74
type RecoveryIndexStatus struct {
	Bytes                      *RecoveryBytes          `json:"bytes,omitempty"`
	Files                      RecoveryFiles           `json:"files"`
	Size                       RecoveryBytes           `json:"size"`
	SourceThrottleTime         *Duration               `json:"source_throttle_time,omitempty"`
	SourceThrottleTimeInMillis DurationValueUnitMillis `json:"source_throttle_time_in_millis"`
	TargetThrottleTime         *Duration               `json:"target_throttle_time,omitempty"`
	TargetThrottleTimeInMillis DurationValueUnitMillis `json:"target_throttle_time_in_millis"`
	TotalTime                  *Duration               `json:"total_time,omitempty"`
	TotalTimeInMillis          DurationValueUnitMillis `json:"total_time_in_millis"`
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

func (rb *RecoveryIndexStatusBuilder) SourceThrottleTime(sourcethrottletime *DurationBuilder) *RecoveryIndexStatusBuilder {
	v := sourcethrottletime.Build()
	rb.v.SourceThrottleTime = &v
	return rb
}

func (rb *RecoveryIndexStatusBuilder) SourceThrottleTimeInMillis(sourcethrottletimeinmillis *DurationValueUnitMillisBuilder) *RecoveryIndexStatusBuilder {
	v := sourcethrottletimeinmillis.Build()
	rb.v.SourceThrottleTimeInMillis = v
	return rb
}

func (rb *RecoveryIndexStatusBuilder) TargetThrottleTime(targetthrottletime *DurationBuilder) *RecoveryIndexStatusBuilder {
	v := targetthrottletime.Build()
	rb.v.TargetThrottleTime = &v
	return rb
}

func (rb *RecoveryIndexStatusBuilder) TargetThrottleTimeInMillis(targetthrottletimeinmillis *DurationValueUnitMillisBuilder) *RecoveryIndexStatusBuilder {
	v := targetthrottletimeinmillis.Build()
	rb.v.TargetThrottleTimeInMillis = v
	return rb
}

func (rb *RecoveryIndexStatusBuilder) TotalTime(totaltime *DurationBuilder) *RecoveryIndexStatusBuilder {
	v := totaltime.Build()
	rb.v.TotalTime = &v
	return rb
}

func (rb *RecoveryIndexStatusBuilder) TotalTimeInMillis(totaltimeinmillis *DurationValueUnitMillisBuilder) *RecoveryIndexStatusBuilder {
	v := totaltimeinmillis.Build()
	rb.v.TotalTimeInMillis = v
	return rb
}
