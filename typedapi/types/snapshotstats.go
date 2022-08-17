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

// SnapshotStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/snapshot/_types/SnapshotStats.ts#L23-L29
type SnapshotStats struct {
	Incremental       FileCountSnapshotStats  `json:"incremental"`
	StartTimeInMillis EpochTimeUnitMillis     `json:"start_time_in_millis"`
	Time              *Duration               `json:"time,omitempty"`
	TimeInMillis      DurationValueUnitMillis `json:"time_in_millis"`
	Total             FileCountSnapshotStats  `json:"total"`
}

// SnapshotStatsBuilder holds SnapshotStats struct and provides a builder API.
type SnapshotStatsBuilder struct {
	v *SnapshotStats
}

// NewSnapshotStats provides a builder for the SnapshotStats struct.
func NewSnapshotStatsBuilder() *SnapshotStatsBuilder {
	r := SnapshotStatsBuilder{
		&SnapshotStats{},
	}

	return &r
}

// Build finalize the chain and returns the SnapshotStats struct
func (rb *SnapshotStatsBuilder) Build() SnapshotStats {
	return *rb.v
}

func (rb *SnapshotStatsBuilder) Incremental(incremental *FileCountSnapshotStatsBuilder) *SnapshotStatsBuilder {
	v := incremental.Build()
	rb.v.Incremental = v
	return rb
}

func (rb *SnapshotStatsBuilder) StartTimeInMillis(starttimeinmillis *EpochTimeUnitMillisBuilder) *SnapshotStatsBuilder {
	v := starttimeinmillis.Build()
	rb.v.StartTimeInMillis = v
	return rb
}

func (rb *SnapshotStatsBuilder) Time(time *DurationBuilder) *SnapshotStatsBuilder {
	v := time.Build()
	rb.v.Time = &v
	return rb
}

func (rb *SnapshotStatsBuilder) TimeInMillis(timeinmillis *DurationValueUnitMillisBuilder) *SnapshotStatsBuilder {
	v := timeinmillis.Build()
	rb.v.TimeInMillis = v
	return rb
}

func (rb *SnapshotStatsBuilder) Total(total *FileCountSnapshotStatsBuilder) *SnapshotStatsBuilder {
	v := total.Build()
	rb.v.Total = v
	return rb
}
