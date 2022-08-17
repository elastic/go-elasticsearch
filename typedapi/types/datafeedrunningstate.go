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

// DatafeedRunningState type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Datafeed.ts#L158-L162
type DatafeedRunningState struct {
	RealTimeConfigured bool                        `json:"real_time_configured"`
	RealTimeRunning    bool                        `json:"real_time_running"`
	SearchInterval     *RunningStateSearchInterval `json:"search_interval,omitempty"`
}

// DatafeedRunningStateBuilder holds DatafeedRunningState struct and provides a builder API.
type DatafeedRunningStateBuilder struct {
	v *DatafeedRunningState
}

// NewDatafeedRunningState provides a builder for the DatafeedRunningState struct.
func NewDatafeedRunningStateBuilder() *DatafeedRunningStateBuilder {
	r := DatafeedRunningStateBuilder{
		&DatafeedRunningState{},
	}

	return &r
}

// Build finalize the chain and returns the DatafeedRunningState struct
func (rb *DatafeedRunningStateBuilder) Build() DatafeedRunningState {
	return *rb.v
}

func (rb *DatafeedRunningStateBuilder) RealTimeConfigured(realtimeconfigured bool) *DatafeedRunningStateBuilder {
	rb.v.RealTimeConfigured = realtimeconfigured
	return rb
}

func (rb *DatafeedRunningStateBuilder) RealTimeRunning(realtimerunning bool) *DatafeedRunningStateBuilder {
	rb.v.RealTimeRunning = realtimerunning
	return rb
}

func (rb *DatafeedRunningStateBuilder) SearchInterval(searchinterval *RunningStateSearchIntervalBuilder) *DatafeedRunningStateBuilder {
	v := searchinterval.Build()
	rb.v.SearchInterval = &v
	return rb
}
