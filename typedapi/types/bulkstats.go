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

// BulkStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/Stats.ts#L41-L51
type BulkStats struct {
	AvgSize           *ByteSize `json:"avg_size,omitempty"`
	AvgSizeInBytes    int64     `json:"avg_size_in_bytes"`
	AvgTime           *Duration `json:"avg_time,omitempty"`
	AvgTimeInMillis   int64     `json:"avg_time_in_millis"`
	TotalOperations   int64     `json:"total_operations"`
	TotalSize         *ByteSize `json:"total_size,omitempty"`
	TotalSizeInBytes  int64     `json:"total_size_in_bytes"`
	TotalTime         *Duration `json:"total_time,omitempty"`
	TotalTimeInMillis int64     `json:"total_time_in_millis"`
}

// NewBulkStats returns a BulkStats.
func NewBulkStats() *BulkStats {
	r := &BulkStats{}

	return r
}
