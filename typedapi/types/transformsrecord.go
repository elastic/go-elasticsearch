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

// TransformsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/cat/transforms/types.ts#L22-L187
type TransformsRecord struct {
	// ChangesLastDetectionTime changes last detected time
	ChangesLastDetectionTime string `json:"changes_last_detection_time,omitempty"`
	// Checkpoint checkpoint
	Checkpoint *string `json:"checkpoint,omitempty"`
	// CheckpointDurationTimeExpAvg exponential average checkpoint processing time (milliseconds)
	CheckpointDurationTimeExpAvg *string `json:"checkpoint_duration_time_exp_avg,omitempty"`
	// CheckpointProgress progress of the checkpoint
	CheckpointProgress string `json:"checkpoint_progress,omitempty"`
	// CreateTime transform creation time
	CreateTime *string `json:"create_time,omitempty"`
	// DeleteTime total time spent deleting documents
	DeleteTime *string `json:"delete_time,omitempty"`
	// Description description
	Description *string `json:"description,omitempty"`
	// DestIndex destination index
	DestIndex *string `json:"dest_index,omitempty"`
	// DocsPerSecond docs per second
	DocsPerSecond *string `json:"docs_per_second,omitempty"`
	// DocumentsDeleted the number of documents deleted from the destination index
	DocumentsDeleted *string `json:"documents_deleted,omitempty"`
	// DocumentsIndexed the number of documents written to the destination index
	DocumentsIndexed *string `json:"documents_indexed,omitempty"`
	// DocumentsProcessed the number of documents read from source indices and processed
	DocumentsProcessed *string `json:"documents_processed,omitempty"`
	// Frequency frequency of transform
	Frequency *string `json:"frequency,omitempty"`
	// Id the id
	Id *string `json:"id,omitempty"`
	// IndexFailure total number of index failures
	IndexFailure *string `json:"index_failure,omitempty"`
	// IndexTime total time spent indexing documents
	IndexTime *string `json:"index_time,omitempty"`
	// IndexTotal total number of index phases done by the transform
	IndexTotal *string `json:"index_total,omitempty"`
	// IndexedDocumentsExpAvg exponential average number of documents indexed
	IndexedDocumentsExpAvg *string `json:"indexed_documents_exp_avg,omitempty"`
	// LastSearchTime last time transform searched for updates
	LastSearchTime string `json:"last_search_time,omitempty"`
	// MaxPageSearchSize max page search size
	MaxPageSearchSize *string `json:"max_page_search_size,omitempty"`
	// PagesProcessed the number of pages processed
	PagesProcessed *string `json:"pages_processed,omitempty"`
	// Pipeline transform pipeline
	Pipeline *string `json:"pipeline,omitempty"`
	// ProcessedDocumentsExpAvg exponential average number of documents processed
	ProcessedDocumentsExpAvg *string `json:"processed_documents_exp_avg,omitempty"`
	// ProcessingTime the total time spent processing documents
	ProcessingTime *string `json:"processing_time,omitempty"`
	// Reason reason for the current state
	Reason *string `json:"reason,omitempty"`
	// SearchFailure total number of search failures
	SearchFailure *string `json:"search_failure,omitempty"`
	// SearchTime total search time
	SearchTime *string `json:"search_time,omitempty"`
	// SearchTotal total number of search phases
	SearchTotal *string `json:"search_total,omitempty"`
	// SourceIndex source index
	SourceIndex *string `json:"source_index,omitempty"`
	// State transform state
	State *string `json:"state,omitempty"`
	// TransformType batch or continuous transform
	TransformType *string `json:"transform_type,omitempty"`
	// TriggerCount the number of times the transform has been triggered
	TriggerCount *string `json:"trigger_count,omitempty"`
	// Version the version of Elasticsearch when the transform was created
	Version *string `json:"version,omitempty"`
}

// NewTransformsRecord returns a TransformsRecord.
func NewTransformsRecord() *TransformsRecord {
	r := &TransformsRecord{}

	return r
}
