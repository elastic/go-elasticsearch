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

// DataframeAnalyticsStatsMemoryUsage type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeAnalytics.ts#L350-L359
type DataframeAnalyticsStatsMemoryUsage struct {
	// MemoryReestimateBytes This value is present when the status is hard_limit and it is a new estimate
	// of how much memory the job needs.
	MemoryReestimateBytes *int64 `json:"memory_reestimate_bytes,omitempty"`
	// PeakUsageBytes The number of bytes used at the highest peak of memory usage.
	PeakUsageBytes int64 `json:"peak_usage_bytes"`
	// Status The memory usage status.
	Status string `json:"status"`
	// Timestamp The timestamp when memory usage was calculated.
	Timestamp *EpochTimeUnitMillis `json:"timestamp,omitempty"`
}

// DataframeAnalyticsStatsMemoryUsageBuilder holds DataframeAnalyticsStatsMemoryUsage struct and provides a builder API.
type DataframeAnalyticsStatsMemoryUsageBuilder struct {
	v *DataframeAnalyticsStatsMemoryUsage
}

// NewDataframeAnalyticsStatsMemoryUsage provides a builder for the DataframeAnalyticsStatsMemoryUsage struct.
func NewDataframeAnalyticsStatsMemoryUsageBuilder() *DataframeAnalyticsStatsMemoryUsageBuilder {
	r := DataframeAnalyticsStatsMemoryUsageBuilder{
		&DataframeAnalyticsStatsMemoryUsage{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalyticsStatsMemoryUsage struct
func (rb *DataframeAnalyticsStatsMemoryUsageBuilder) Build() DataframeAnalyticsStatsMemoryUsage {
	return *rb.v
}

// MemoryReestimateBytes This value is present when the status is hard_limit and it is a new estimate
// of how much memory the job needs.

func (rb *DataframeAnalyticsStatsMemoryUsageBuilder) MemoryReestimateBytes(memoryreestimatebytes int64) *DataframeAnalyticsStatsMemoryUsageBuilder {
	rb.v.MemoryReestimateBytes = &memoryreestimatebytes
	return rb
}

// PeakUsageBytes The number of bytes used at the highest peak of memory usage.

func (rb *DataframeAnalyticsStatsMemoryUsageBuilder) PeakUsageBytes(peakusagebytes int64) *DataframeAnalyticsStatsMemoryUsageBuilder {
	rb.v.PeakUsageBytes = peakusagebytes
	return rb
}

// Status The memory usage status.

func (rb *DataframeAnalyticsStatsMemoryUsageBuilder) Status(status string) *DataframeAnalyticsStatsMemoryUsageBuilder {
	rb.v.Status = status
	return rb
}

// Timestamp The timestamp when memory usage was calculated.

func (rb *DataframeAnalyticsStatsMemoryUsageBuilder) Timestamp(timestamp *EpochTimeUnitMillisBuilder) *DataframeAnalyticsStatsMemoryUsageBuilder {
	v := timestamp.Build()
	rb.v.Timestamp = &v
	return rb
}
