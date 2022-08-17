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

// DatafeedTimingStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Datafeed.ts#L149-L156
type DatafeedTimingStats struct {
	AverageSearchTimePerBucketMs          *DurationValueUnitFloatMillis `json:"average_search_time_per_bucket_ms,omitempty"`
	BucketCount                           int64                         `json:"bucket_count"`
	ExponentialAverageSearchTimePerHourMs DurationValueUnitFloatMillis  `json:"exponential_average_search_time_per_hour_ms"`
	JobId                                 Id                            `json:"job_id"`
	SearchCount                           int64                         `json:"search_count"`
	TotalSearchTimeMs                     DurationValueUnitFloatMillis  `json:"total_search_time_ms"`
}

// DatafeedTimingStatsBuilder holds DatafeedTimingStats struct and provides a builder API.
type DatafeedTimingStatsBuilder struct {
	v *DatafeedTimingStats
}

// NewDatafeedTimingStats provides a builder for the DatafeedTimingStats struct.
func NewDatafeedTimingStatsBuilder() *DatafeedTimingStatsBuilder {
	r := DatafeedTimingStatsBuilder{
		&DatafeedTimingStats{},
	}

	return &r
}

// Build finalize the chain and returns the DatafeedTimingStats struct
func (rb *DatafeedTimingStatsBuilder) Build() DatafeedTimingStats {
	return *rb.v
}

func (rb *DatafeedTimingStatsBuilder) AverageSearchTimePerBucketMs(averagesearchtimeperbucketms *DurationValueUnitFloatMillisBuilder) *DatafeedTimingStatsBuilder {
	v := averagesearchtimeperbucketms.Build()
	rb.v.AverageSearchTimePerBucketMs = &v
	return rb
}

func (rb *DatafeedTimingStatsBuilder) BucketCount(bucketcount int64) *DatafeedTimingStatsBuilder {
	rb.v.BucketCount = bucketcount
	return rb
}

func (rb *DatafeedTimingStatsBuilder) ExponentialAverageSearchTimePerHourMs(exponentialaveragesearchtimeperhourms *DurationValueUnitFloatMillisBuilder) *DatafeedTimingStatsBuilder {
	v := exponentialaveragesearchtimeperhourms.Build()
	rb.v.ExponentialAverageSearchTimePerHourMs = v
	return rb
}

func (rb *DatafeedTimingStatsBuilder) JobId(jobid Id) *DatafeedTimingStatsBuilder {
	rb.v.JobId = jobid
	return rb
}

func (rb *DatafeedTimingStatsBuilder) SearchCount(searchcount int64) *DatafeedTimingStatsBuilder {
	rb.v.SearchCount = searchcount
	return rb
}

func (rb *DatafeedTimingStatsBuilder) TotalSearchTimeMs(totalsearchtimems *DurationValueUnitFloatMillisBuilder) *DatafeedTimingStatsBuilder {
	v := totalsearchtimems.Build()
	rb.v.TotalSearchTimeMs = v
	return rb
}
