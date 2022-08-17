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

// DataframeAnalyticsStatsProgress type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeAnalytics.ts#L343-L348
type DataframeAnalyticsStatsProgress struct {
	// Phase Defines the phase of the data frame analytics job.
	Phase string `json:"phase"`
	// ProgressPercent The progress that the data frame analytics job has made expressed in
	// percentage.
	ProgressPercent int `json:"progress_percent"`
}

// DataframeAnalyticsStatsProgressBuilder holds DataframeAnalyticsStatsProgress struct and provides a builder API.
type DataframeAnalyticsStatsProgressBuilder struct {
	v *DataframeAnalyticsStatsProgress
}

// NewDataframeAnalyticsStatsProgress provides a builder for the DataframeAnalyticsStatsProgress struct.
func NewDataframeAnalyticsStatsProgressBuilder() *DataframeAnalyticsStatsProgressBuilder {
	r := DataframeAnalyticsStatsProgressBuilder{
		&DataframeAnalyticsStatsProgress{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalyticsStatsProgress struct
func (rb *DataframeAnalyticsStatsProgressBuilder) Build() DataframeAnalyticsStatsProgress {
	return *rb.v
}

// Phase Defines the phase of the data frame analytics job.

func (rb *DataframeAnalyticsStatsProgressBuilder) Phase(phase string) *DataframeAnalyticsStatsProgressBuilder {
	rb.v.Phase = phase
	return rb
}

// ProgressPercent The progress that the data frame analytics job has made expressed in
// percentage.

func (rb *DataframeAnalyticsStatsProgressBuilder) ProgressPercent(progresspercent int) *DataframeAnalyticsStatsProgressBuilder {
	rb.v.ProgressPercent = progresspercent
	return rb
}
