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

// BulkStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/Stats.ts#L40-L50
type BulkStats struct {
	AvgSize           *ByteSize `json:"avg_size,omitempty"`
	AvgSizeInBytes    int64     `json:"avg_size_in_bytes"`
	AvgTime           *string   `json:"avg_time,omitempty"`
	AvgTimeInMillis   int64     `json:"avg_time_in_millis"`
	TotalOperations   int64     `json:"total_operations"`
	TotalSize         *ByteSize `json:"total_size,omitempty"`
	TotalSizeInBytes  int64     `json:"total_size_in_bytes"`
	TotalTime         *string   `json:"total_time,omitempty"`
	TotalTimeInMillis int64     `json:"total_time_in_millis"`
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

func (rb *BulkStatsBuilder) AvgTime(avgtime string) *BulkStatsBuilder {
	rb.v.AvgTime = &avgtime
	return rb
}

func (rb *BulkStatsBuilder) AvgTimeInMillis(avgtimeinmillis int64) *BulkStatsBuilder {
	rb.v.AvgTimeInMillis = avgtimeinmillis
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

func (rb *BulkStatsBuilder) TotalTime(totaltime string) *BulkStatsBuilder {
	rb.v.TotalTime = &totaltime
	return rb
}

func (rb *BulkStatsBuilder) TotalTimeInMillis(totaltimeinmillis int64) *BulkStatsBuilder {
	rb.v.TotalTimeInMillis = totaltimeinmillis
	return rb
}
