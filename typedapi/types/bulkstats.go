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

// BulkStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Stats.ts#L41-L51
type BulkStats struct {
	AvgSize           *ByteSize               `json:"avg_size,omitempty"`
	AvgSizeInBytes    int64                   `json:"avg_size_in_bytes"`
	AvgTime           *Duration               `json:"avg_time,omitempty"`
	AvgTimeInMillis   DurationValueUnitMillis `json:"avg_time_in_millis"`
	TotalOperations   int64                   `json:"total_operations"`
	TotalSize         *ByteSize               `json:"total_size,omitempty"`
	TotalSizeInBytes  int64                   `json:"total_size_in_bytes"`
	TotalTime         *Duration               `json:"total_time,omitempty"`
	TotalTimeInMillis DurationValueUnitMillis `json:"total_time_in_millis"`
}

// BulkStatsBuilder holds BulkStats struct and provides a builder API.
type BulkStatsBuilder struct {
	v *BulkStats
}

// NewBulkStats provides a builder for the BulkStats struct.
func NewBulkStatsBuilder() *BulkStatsBuilder {
	r := BulkStatsBuilder{
		&BulkStats{},
	}

	return &r
}

// Build finalize the chain and returns the BulkStats struct
func (rb *BulkStatsBuilder) Build() BulkStats {
	return *rb.v
}

func (rb *BulkStatsBuilder) AvgSize(avgsize *ByteSizeBuilder) *BulkStatsBuilder {
	v := avgsize.Build()
	rb.v.AvgSize = &v
	return rb
}

func (rb *BulkStatsBuilder) AvgSizeInBytes(avgsizeinbytes int64) *BulkStatsBuilder {
	rb.v.AvgSizeInBytes = avgsizeinbytes
	return rb
}

func (rb *BulkStatsBuilder) AvgTime(avgtime *DurationBuilder) *BulkStatsBuilder {
	v := avgtime.Build()
	rb.v.AvgTime = &v
	return rb
}

func (rb *BulkStatsBuilder) AvgTimeInMillis(avgtimeinmillis *DurationValueUnitMillisBuilder) *BulkStatsBuilder {
	v := avgtimeinmillis.Build()
	rb.v.AvgTimeInMillis = v
	return rb
}

func (rb *BulkStatsBuilder) TotalOperations(totaloperations int64) *BulkStatsBuilder {
	rb.v.TotalOperations = totaloperations
	return rb
}

func (rb *BulkStatsBuilder) TotalSize(totalsize *ByteSizeBuilder) *BulkStatsBuilder {
	v := totalsize.Build()
	rb.v.TotalSize = &v
	return rb
}

func (rb *BulkStatsBuilder) TotalSizeInBytes(totalsizeinbytes int64) *BulkStatsBuilder {
	rb.v.TotalSizeInBytes = totalsizeinbytes
	return rb
}

func (rb *BulkStatsBuilder) TotalTime(totaltime *DurationBuilder) *BulkStatsBuilder {
	v := totaltime.Build()
	rb.v.TotalTime = &v
	return rb
}

func (rb *BulkStatsBuilder) TotalTimeInMillis(totaltimeinmillis *DurationValueUnitMillisBuilder) *BulkStatsBuilder {
	v := totaltimeinmillis.Build()
	rb.v.TotalTimeInMillis = v
	return rb
}
