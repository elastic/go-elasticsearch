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

// TransformIndexerStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/transform/get_transform_stats/types.ts#L48-L66
type TransformIndexerStats struct {
	DeleteTimeInMs                     *EpochTimeUnitMillis         `json:"delete_time_in_ms,omitempty"`
	DocumentsDeleted                   *int64                       `json:"documents_deleted,omitempty"`
	DocumentsIndexed                   int64                        `json:"documents_indexed"`
	DocumentsProcessed                 int64                        `json:"documents_processed"`
	ExponentialAvgCheckpointDurationMs DurationValueUnitFloatMillis `json:"exponential_avg_checkpoint_duration_ms"`
	ExponentialAvgDocumentsIndexed     float64                      `json:"exponential_avg_documents_indexed"`
	ExponentialAvgDocumentsProcessed   float64                      `json:"exponential_avg_documents_processed"`
	IndexFailures                      int64                        `json:"index_failures"`
	IndexTimeInMs                      DurationValueUnitMillis      `json:"index_time_in_ms"`
	IndexTotal                         int64                        `json:"index_total"`
	PagesProcessed                     int64                        `json:"pages_processed"`
	ProcessingTimeInMs                 DurationValueUnitMillis      `json:"processing_time_in_ms"`
	ProcessingTotal                    int64                        `json:"processing_total"`
	SearchFailures                     int64                        `json:"search_failures"`
	SearchTimeInMs                     DurationValueUnitMillis      `json:"search_time_in_ms"`
	SearchTotal                        int64                        `json:"search_total"`
	TriggerCount                       int64                        `json:"trigger_count"`
}

// TransformIndexerStatsBuilder holds TransformIndexerStats struct and provides a builder API.
type TransformIndexerStatsBuilder struct {
	v *TransformIndexerStats
}

// NewTransformIndexerStats provides a builder for the TransformIndexerStats struct.
func NewTransformIndexerStatsBuilder() *TransformIndexerStatsBuilder {
	r := TransformIndexerStatsBuilder{
		&TransformIndexerStats{},
	}

	return &r
}

// Build finalize the chain and returns the TransformIndexerStats struct
func (rb *TransformIndexerStatsBuilder) Build() TransformIndexerStats {
	return *rb.v
}

func (rb *TransformIndexerStatsBuilder) DeleteTimeInMs(deletetimeinms *EpochTimeUnitMillisBuilder) *TransformIndexerStatsBuilder {
	v := deletetimeinms.Build()
	rb.v.DeleteTimeInMs = &v
	return rb
}

func (rb *TransformIndexerStatsBuilder) DocumentsDeleted(documentsdeleted int64) *TransformIndexerStatsBuilder {
	rb.v.DocumentsDeleted = &documentsdeleted
	return rb
}

func (rb *TransformIndexerStatsBuilder) DocumentsIndexed(documentsindexed int64) *TransformIndexerStatsBuilder {
	rb.v.DocumentsIndexed = documentsindexed
	return rb
}

func (rb *TransformIndexerStatsBuilder) DocumentsProcessed(documentsprocessed int64) *TransformIndexerStatsBuilder {
	rb.v.DocumentsProcessed = documentsprocessed
	return rb
}

func (rb *TransformIndexerStatsBuilder) ExponentialAvgCheckpointDurationMs(exponentialavgcheckpointdurationms *DurationValueUnitFloatMillisBuilder) *TransformIndexerStatsBuilder {
	v := exponentialavgcheckpointdurationms.Build()
	rb.v.ExponentialAvgCheckpointDurationMs = v
	return rb
}

func (rb *TransformIndexerStatsBuilder) ExponentialAvgDocumentsIndexed(exponentialavgdocumentsindexed float64) *TransformIndexerStatsBuilder {
	rb.v.ExponentialAvgDocumentsIndexed = exponentialavgdocumentsindexed
	return rb
}

func (rb *TransformIndexerStatsBuilder) ExponentialAvgDocumentsProcessed(exponentialavgdocumentsprocessed float64) *TransformIndexerStatsBuilder {
	rb.v.ExponentialAvgDocumentsProcessed = exponentialavgdocumentsprocessed
	return rb
}

func (rb *TransformIndexerStatsBuilder) IndexFailures(indexfailures int64) *TransformIndexerStatsBuilder {
	rb.v.IndexFailures = indexfailures
	return rb
}

func (rb *TransformIndexerStatsBuilder) IndexTimeInMs(indextimeinms *DurationValueUnitMillisBuilder) *TransformIndexerStatsBuilder {
	v := indextimeinms.Build()
	rb.v.IndexTimeInMs = v
	return rb
}

func (rb *TransformIndexerStatsBuilder) IndexTotal(indextotal int64) *TransformIndexerStatsBuilder {
	rb.v.IndexTotal = indextotal
	return rb
}

func (rb *TransformIndexerStatsBuilder) PagesProcessed(pagesprocessed int64) *TransformIndexerStatsBuilder {
	rb.v.PagesProcessed = pagesprocessed
	return rb
}

func (rb *TransformIndexerStatsBuilder) ProcessingTimeInMs(processingtimeinms *DurationValueUnitMillisBuilder) *TransformIndexerStatsBuilder {
	v := processingtimeinms.Build()
	rb.v.ProcessingTimeInMs = v
	return rb
}

func (rb *TransformIndexerStatsBuilder) ProcessingTotal(processingtotal int64) *TransformIndexerStatsBuilder {
	rb.v.ProcessingTotal = processingtotal
	return rb
}

func (rb *TransformIndexerStatsBuilder) SearchFailures(searchfailures int64) *TransformIndexerStatsBuilder {
	rb.v.SearchFailures = searchfailures
	return rb
}

func (rb *TransformIndexerStatsBuilder) SearchTimeInMs(searchtimeinms *DurationValueUnitMillisBuilder) *TransformIndexerStatsBuilder {
	v := searchtimeinms.Build()
	rb.v.SearchTimeInMs = v
	return rb
}

func (rb *TransformIndexerStatsBuilder) SearchTotal(searchtotal int64) *TransformIndexerStatsBuilder {
	rb.v.SearchTotal = searchtotal
	return rb
}

func (rb *TransformIndexerStatsBuilder) TriggerCount(triggercount int64) *TransformIndexerStatsBuilder {
	rb.v.TriggerCount = triggercount
	return rb
}
