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

// Recording type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L88-L93
type Recording struct {
	CumulativeExecutionCount      *int64                   `json:"cumulative_execution_count,omitempty"`
	CumulativeExecutionTime       *Duration                `json:"cumulative_execution_time,omitempty"`
	CumulativeExecutionTimeMillis *DurationValueUnitMillis `json:"cumulative_execution_time_millis,omitempty"`
	Name                          *string                  `json:"name,omitempty"`
}

// RecordingBuilder holds Recording struct and provides a builder API.
type RecordingBuilder struct {
	v *Recording
}

// NewRecording provides a builder for the Recording struct.
func NewRecordingBuilder() *RecordingBuilder {
	r := RecordingBuilder{
		&Recording{},
	}

	return &r
}

// Build finalize the chain and returns the Recording struct
func (rb *RecordingBuilder) Build() Recording {
	return *rb.v
}

func (rb *RecordingBuilder) CumulativeExecutionCount(cumulativeexecutioncount int64) *RecordingBuilder {
	rb.v.CumulativeExecutionCount = &cumulativeexecutioncount
	return rb
}

func (rb *RecordingBuilder) CumulativeExecutionTime(cumulativeexecutiontime *DurationBuilder) *RecordingBuilder {
	v := cumulativeexecutiontime.Build()
	rb.v.CumulativeExecutionTime = &v
	return rb
}

func (rb *RecordingBuilder) CumulativeExecutionTimeMillis(cumulativeexecutiontimemillis *DurationValueUnitMillisBuilder) *RecordingBuilder {
	v := cumulativeexecutiontimemillis.Build()
	rb.v.CumulativeExecutionTimeMillis = &v
	return rb
}

func (rb *RecordingBuilder) Name(name string) *RecordingBuilder {
	rb.v.Name = &name
	return rb
}
