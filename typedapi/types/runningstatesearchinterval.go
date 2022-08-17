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

// RunningStateSearchInterval type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Datafeed.ts#L164-L169
type RunningStateSearchInterval struct {
	End     *Duration               `json:"end,omitempty"`
	EndMs   DurationValueUnitMillis `json:"end_ms"`
	Start   *Duration               `json:"start,omitempty"`
	StartMs DurationValueUnitMillis `json:"start_ms"`
}

// RunningStateSearchIntervalBuilder holds RunningStateSearchInterval struct and provides a builder API.
type RunningStateSearchIntervalBuilder struct {
	v *RunningStateSearchInterval
}

// NewRunningStateSearchInterval provides a builder for the RunningStateSearchInterval struct.
func NewRunningStateSearchIntervalBuilder() *RunningStateSearchIntervalBuilder {
	r := RunningStateSearchIntervalBuilder{
		&RunningStateSearchInterval{},
	}

	return &r
}

// Build finalize the chain and returns the RunningStateSearchInterval struct
func (rb *RunningStateSearchIntervalBuilder) Build() RunningStateSearchInterval {
	return *rb.v
}

func (rb *RunningStateSearchIntervalBuilder) End(end *DurationBuilder) *RunningStateSearchIntervalBuilder {
	v := end.Build()
	rb.v.End = &v
	return rb
}

func (rb *RunningStateSearchIntervalBuilder) EndMs(endms *DurationValueUnitMillisBuilder) *RunningStateSearchIntervalBuilder {
	v := endms.Build()
	rb.v.EndMs = v
	return rb
}

func (rb *RunningStateSearchIntervalBuilder) Start(start *DurationBuilder) *RunningStateSearchIntervalBuilder {
	v := start.Build()
	rb.v.Start = &v
	return rb
}

func (rb *RunningStateSearchIntervalBuilder) StartMs(startms *DurationValueUnitMillisBuilder) *RunningStateSearchIntervalBuilder {
	v := startms.Build()
	rb.v.StartMs = v
	return rb
}
