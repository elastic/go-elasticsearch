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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// TransformIndexerStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/transform/get_transform_stats/types.ts#L53-L71
type TransformIndexerStats struct {
	DeleteTimeInMs                     *int64  `json:"delete_time_in_ms,omitempty"`
	DocumentsDeleted                   *int64  `json:"documents_deleted,omitempty"`
	DocumentsIndexed                   int64   `json:"documents_indexed"`
	DocumentsProcessed                 int64   `json:"documents_processed"`
	ExponentialAvgCheckpointDurationMs Float64 `json:"exponential_avg_checkpoint_duration_ms"`
	ExponentialAvgDocumentsIndexed     Float64 `json:"exponential_avg_documents_indexed"`
	ExponentialAvgDocumentsProcessed   Float64 `json:"exponential_avg_documents_processed"`
	IndexFailures                      int64   `json:"index_failures"`
	IndexTimeInMs                      int64   `json:"index_time_in_ms"`
	IndexTotal                         int64   `json:"index_total"`
	PagesProcessed                     int64   `json:"pages_processed"`
	ProcessingTimeInMs                 int64   `json:"processing_time_in_ms"`
	ProcessingTotal                    int64   `json:"processing_total"`
	SearchFailures                     int64   `json:"search_failures"`
	SearchTimeInMs                     int64   `json:"search_time_in_ms"`
	SearchTotal                        int64   `json:"search_total"`
	TriggerCount                       int64   `json:"trigger_count"`
}

// NewTransformIndexerStats returns a TransformIndexerStats.
func NewTransformIndexerStats() *TransformIndexerStats {
	r := &TransformIndexerStats{}

	return r
}
