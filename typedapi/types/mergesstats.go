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

// MergesStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Stats.ts#L118-L135
type MergesStats struct {
	Current                    int64                   `json:"current"`
	CurrentDocs                int64                   `json:"current_docs"`
	CurrentSize                *string                 `json:"current_size,omitempty"`
	CurrentSizeInBytes         int64                   `json:"current_size_in_bytes"`
	Total                      int64                   `json:"total"`
	TotalAutoThrottle          *string                 `json:"total_auto_throttle,omitempty"`
	TotalAutoThrottleInBytes   int64                   `json:"total_auto_throttle_in_bytes"`
	TotalDocs                  int64                   `json:"total_docs"`
	TotalSize                  *string                 `json:"total_size,omitempty"`
	TotalSizeInBytes           int64                   `json:"total_size_in_bytes"`
	TotalStoppedTime           *Duration               `json:"total_stopped_time,omitempty"`
	TotalStoppedTimeInMillis   DurationValueUnitMillis `json:"total_stopped_time_in_millis"`
	TotalThrottledTime         *Duration               `json:"total_throttled_time,omitempty"`
	TotalThrottledTimeInMillis DurationValueUnitMillis `json:"total_throttled_time_in_millis"`
	TotalTime                  *Duration               `json:"total_time,omitempty"`
	TotalTimeInMillis          DurationValueUnitMillis `json:"total_time_in_millis"`
}

// MergesStatsBuilder holds MergesStats struct and provides a builder API.
type MergesStatsBuilder struct {
	v *MergesStats
}

// NewMergesStats provides a builder for the MergesStats struct.
func NewMergesStatsBuilder() *MergesStatsBuilder {
	r := MergesStatsBuilder{
		&MergesStats{},
	}

	return &r
}

// Build finalize the chain and returns the MergesStats struct
func (rb *MergesStatsBuilder) Build() MergesStats {
	return *rb.v
}

func (rb *MergesStatsBuilder) Current(current int64) *MergesStatsBuilder {
	rb.v.Current = current
	return rb
}

func (rb *MergesStatsBuilder) CurrentDocs(currentdocs int64) *MergesStatsBuilder {
	rb.v.CurrentDocs = currentdocs
	return rb
}

func (rb *MergesStatsBuilder) CurrentSize(currentsize string) *MergesStatsBuilder {
	rb.v.CurrentSize = &currentsize
	return rb
}

func (rb *MergesStatsBuilder) CurrentSizeInBytes(currentsizeinbytes int64) *MergesStatsBuilder {
	rb.v.CurrentSizeInBytes = currentsizeinbytes
	return rb
}

func (rb *MergesStatsBuilder) Total(total int64) *MergesStatsBuilder {
	rb.v.Total = total
	return rb
}

func (rb *MergesStatsBuilder) TotalAutoThrottle(totalautothrottle string) *MergesStatsBuilder {
	rb.v.TotalAutoThrottle = &totalautothrottle
	return rb
}

func (rb *MergesStatsBuilder) TotalAutoThrottleInBytes(totalautothrottleinbytes int64) *MergesStatsBuilder {
	rb.v.TotalAutoThrottleInBytes = totalautothrottleinbytes
	return rb
}

func (rb *MergesStatsBuilder) TotalDocs(totaldocs int64) *MergesStatsBuilder {
	rb.v.TotalDocs = totaldocs
	return rb
}

func (rb *MergesStatsBuilder) TotalSize(totalsize string) *MergesStatsBuilder {
	rb.v.TotalSize = &totalsize
	return rb
}

func (rb *MergesStatsBuilder) TotalSizeInBytes(totalsizeinbytes int64) *MergesStatsBuilder {
	rb.v.TotalSizeInBytes = totalsizeinbytes
	return rb
}

func (rb *MergesStatsBuilder) TotalStoppedTime(totalstoppedtime *DurationBuilder) *MergesStatsBuilder {
	v := totalstoppedtime.Build()
	rb.v.TotalStoppedTime = &v
	return rb
}

func (rb *MergesStatsBuilder) TotalStoppedTimeInMillis(totalstoppedtimeinmillis *DurationValueUnitMillisBuilder) *MergesStatsBuilder {
	v := totalstoppedtimeinmillis.Build()
	rb.v.TotalStoppedTimeInMillis = v
	return rb
}

func (rb *MergesStatsBuilder) TotalThrottledTime(totalthrottledtime *DurationBuilder) *MergesStatsBuilder {
	v := totalthrottledtime.Build()
	rb.v.TotalThrottledTime = &v
	return rb
}

func (rb *MergesStatsBuilder) TotalThrottledTimeInMillis(totalthrottledtimeinmillis *DurationValueUnitMillisBuilder) *MergesStatsBuilder {
	v := totalthrottledtimeinmillis.Build()
	rb.v.TotalThrottledTimeInMillis = v
	return rb
}

func (rb *MergesStatsBuilder) TotalTime(totaltime *DurationBuilder) *MergesStatsBuilder {
	v := totaltime.Build()
	rb.v.TotalTime = &v
	return rb
}

func (rb *MergesStatsBuilder) TotalTimeInMillis(totaltimeinmillis *DurationValueUnitMillisBuilder) *MergesStatsBuilder {
	v := totaltimeinmillis.Build()
	rb.v.TotalTimeInMillis = v
	return rb
}
