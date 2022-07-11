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

// RollupJobStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/rollup/get_jobs/types.ts#L45-L58
type RollupJobStats struct {
	DocumentsProcessed int64 `json:"documents_processed"`
	IndexFailures      int64 `json:"index_failures"`
	IndexTimeInMs      int64 `json:"index_time_in_ms"`
	IndexTotal         int64 `json:"index_total"`
	PagesProcessed     int64 `json:"pages_processed"`
	ProcessingTimeInMs int64 `json:"processing_time_in_ms"`
	ProcessingTotal    int64 `json:"processing_total"`
	RollupsIndexed     int64 `json:"rollups_indexed"`
	SearchFailures     int64 `json:"search_failures"`
	SearchTimeInMs     int64 `json:"search_time_in_ms"`
	SearchTotal        int64 `json:"search_total"`
	TriggerCount       int64 `json:"trigger_count"`
}

// RollupJobStatsBuilder holds RollupJobStats struct and provides a builder API.
type RollupJobStatsBuilder struct {
	v *RollupJobStats
}

// NewRollupJobStats provides a builder for the RollupJobStats struct.
func NewRollupJobStatsBuilder() *RollupJobStatsBuilder {
	r := RollupJobStatsBuilder{
		&RollupJobStats{},
	}

	return &r
}

// Build finalize the chain and returns the RollupJobStats struct
func (rb *RollupJobStatsBuilder) Build() RollupJobStats {
	return *rb.v
}

func (rb *RollupJobStatsBuilder) DocumentsProcessed(documentsprocessed int64) *RollupJobStatsBuilder {
	rb.v.DocumentsProcessed = documentsprocessed
	return rb
}

func (rb *RollupJobStatsBuilder) IndexFailures(indexfailures int64) *RollupJobStatsBuilder {
	rb.v.IndexFailures = indexfailures
	return rb
}

func (rb *RollupJobStatsBuilder) IndexTimeInMs(indextimeinms int64) *RollupJobStatsBuilder {
	rb.v.IndexTimeInMs = indextimeinms
	return rb
}

func (rb *RollupJobStatsBuilder) IndexTotal(indextotal int64) *RollupJobStatsBuilder {
	rb.v.IndexTotal = indextotal
	return rb
}

func (rb *RollupJobStatsBuilder) PagesProcessed(pagesprocessed int64) *RollupJobStatsBuilder {
	rb.v.PagesProcessed = pagesprocessed
	return rb
}

func (rb *RollupJobStatsBuilder) ProcessingTimeInMs(processingtimeinms int64) *RollupJobStatsBuilder {
	rb.v.ProcessingTimeInMs = processingtimeinms
	return rb
}

func (rb *RollupJobStatsBuilder) ProcessingTotal(processingtotal int64) *RollupJobStatsBuilder {
	rb.v.ProcessingTotal = processingtotal
	return rb
}

func (rb *RollupJobStatsBuilder) RollupsIndexed(rollupsindexed int64) *RollupJobStatsBuilder {
	rb.v.RollupsIndexed = rollupsindexed
	return rb
}

func (rb *RollupJobStatsBuilder) SearchFailures(searchfailures int64) *RollupJobStatsBuilder {
	rb.v.SearchFailures = searchfailures
	return rb
}

func (rb *RollupJobStatsBuilder) SearchTimeInMs(searchtimeinms int64) *RollupJobStatsBuilder {
	rb.v.SearchTimeInMs = searchtimeinms
	return rb
}

func (rb *RollupJobStatsBuilder) SearchTotal(searchtotal int64) *RollupJobStatsBuilder {
	rb.v.SearchTotal = searchtotal
	return rb
}

func (rb *RollupJobStatsBuilder) TriggerCount(triggercount int64) *RollupJobStatsBuilder {
	rb.v.TriggerCount = triggercount
	return rb
}
