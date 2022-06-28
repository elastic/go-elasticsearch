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

// DataframeAnalyticsStatsOutlierDetection type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/ml/_types/DataframeAnalytics.ts#L384-L388
type DataframeAnalyticsStatsOutlierDetection struct {
	Parameters  OutlierDetectionParameters `json:"parameters"`
	Timestamp   DateString                 `json:"timestamp"`
	TimingStats TimingStats                `json:"timing_stats"`
}

// DataframeAnalyticsStatsOutlierDetectionBuilder holds DataframeAnalyticsStatsOutlierDetection struct and provides a builder API.
type DataframeAnalyticsStatsOutlierDetectionBuilder struct {
	v *DataframeAnalyticsStatsOutlierDetection
}

// NewDataframeAnalyticsStatsOutlierDetection provides a builder for the DataframeAnalyticsStatsOutlierDetection struct.
func NewDataframeAnalyticsStatsOutlierDetectionBuilder() *DataframeAnalyticsStatsOutlierDetectionBuilder {
	r := DataframeAnalyticsStatsOutlierDetectionBuilder{
		&DataframeAnalyticsStatsOutlierDetection{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalyticsStatsOutlierDetection struct
func (rb *DataframeAnalyticsStatsOutlierDetectionBuilder) Build() DataframeAnalyticsStatsOutlierDetection {
	return *rb.v
}

func (rb *DataframeAnalyticsStatsOutlierDetectionBuilder) Parameters(parameters *OutlierDetectionParametersBuilder) *DataframeAnalyticsStatsOutlierDetectionBuilder {
	v := parameters.Build()
	rb.v.Parameters = v
	return rb
}

func (rb *DataframeAnalyticsStatsOutlierDetectionBuilder) Timestamp(timestamp DateString) *DataframeAnalyticsStatsOutlierDetectionBuilder {
	rb.v.Timestamp = timestamp
	return rb
}

func (rb *DataframeAnalyticsStatsOutlierDetectionBuilder) TimingStats(timingstats *TimingStatsBuilder) *DataframeAnalyticsStatsOutlierDetectionBuilder {
	v := timingstats.Build()
	rb.v.TimingStats = v
	return rb
}
