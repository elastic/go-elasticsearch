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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// DataframeAnalyticsStatsMemoryUsage type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/DataframeAnalytics.ts#L350-L359
type DataframeAnalyticsStatsMemoryUsage struct {
	// MemoryReestimateBytes This value is present when the status is hard_limit and it is a new estimate
	// of how much memory the job needs.
	MemoryReestimateBytes *int64 `json:"memory_reestimate_bytes,omitempty"`
	// PeakUsageBytes The number of bytes used at the highest peak of memory usage.
	PeakUsageBytes int64 `json:"peak_usage_bytes"`
	// Status The memory usage status.
	Status string `json:"status"`
	// Timestamp The timestamp when memory usage was calculated.
	Timestamp *int64 `json:"timestamp,omitempty"`
}

// NewDataframeAnalyticsStatsMemoryUsage returns a DataframeAnalyticsStatsMemoryUsage.
func NewDataframeAnalyticsStatsMemoryUsage() *DataframeAnalyticsStatsMemoryUsage {
	r := &DataframeAnalyticsStatsMemoryUsage{}

	return r
}
