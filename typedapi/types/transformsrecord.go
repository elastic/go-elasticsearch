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

// TransformsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/transforms/types.ts#L22-L187
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
	Id *Id `json:"id,omitempty"`
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
	Version *VersionString `json:"version,omitempty"`
}

// TransformsRecordBuilder holds TransformsRecord struct and provides a builder API.
type TransformsRecordBuilder struct {
	v *TransformsRecord
}

// NewTransformsRecord provides a builder for the TransformsRecord struct.
func NewTransformsRecordBuilder() *TransformsRecordBuilder {
	r := TransformsRecordBuilder{
		&TransformsRecord{},
	}

	return &r
}

// Build finalize the chain and returns the TransformsRecord struct
func (rb *TransformsRecordBuilder) Build() TransformsRecord {
	return *rb.v
}

// ChangesLastDetectionTime changes last detected time

func (rb *TransformsRecordBuilder) ChangesLastDetectionTime(changeslastdetectiontime string) *TransformsRecordBuilder {
	rb.v.ChangesLastDetectionTime = changeslastdetectiontime
	return rb
}

// Checkpoint checkpoint

func (rb *TransformsRecordBuilder) Checkpoint(checkpoint string) *TransformsRecordBuilder {
	rb.v.Checkpoint = &checkpoint
	return rb
}

// CheckpointDurationTimeExpAvg exponential average checkpoint processing time (milliseconds)

func (rb *TransformsRecordBuilder) CheckpointDurationTimeExpAvg(checkpointdurationtimeexpavg string) *TransformsRecordBuilder {
	rb.v.CheckpointDurationTimeExpAvg = &checkpointdurationtimeexpavg
	return rb
}

// CheckpointProgress progress of the checkpoint

func (rb *TransformsRecordBuilder) CheckpointProgress(checkpointprogress string) *TransformsRecordBuilder {
	rb.v.CheckpointProgress = checkpointprogress
	return rb
}

// CreateTime transform creation time

func (rb *TransformsRecordBuilder) CreateTime(createtime string) *TransformsRecordBuilder {
	rb.v.CreateTime = &createtime
	return rb
}

// DeleteTime total time spent deleting documents

func (rb *TransformsRecordBuilder) DeleteTime(deletetime string) *TransformsRecordBuilder {
	rb.v.DeleteTime = &deletetime
	return rb
}

// Description description

func (rb *TransformsRecordBuilder) Description(description string) *TransformsRecordBuilder {
	rb.v.Description = &description
	return rb
}

// DestIndex destination index

func (rb *TransformsRecordBuilder) DestIndex(destindex string) *TransformsRecordBuilder {
	rb.v.DestIndex = &destindex
	return rb
}

// DocsPerSecond docs per second

func (rb *TransformsRecordBuilder) DocsPerSecond(docspersecond string) *TransformsRecordBuilder {
	rb.v.DocsPerSecond = &docspersecond
	return rb
}

// DocumentsDeleted the number of documents deleted from the destination index

func (rb *TransformsRecordBuilder) DocumentsDeleted(documentsdeleted string) *TransformsRecordBuilder {
	rb.v.DocumentsDeleted = &documentsdeleted
	return rb
}

// DocumentsIndexed the number of documents written to the destination index

func (rb *TransformsRecordBuilder) DocumentsIndexed(documentsindexed string) *TransformsRecordBuilder {
	rb.v.DocumentsIndexed = &documentsindexed
	return rb
}

// DocumentsProcessed the number of documents read from source indices and processed

func (rb *TransformsRecordBuilder) DocumentsProcessed(documentsprocessed string) *TransformsRecordBuilder {
	rb.v.DocumentsProcessed = &documentsprocessed
	return rb
}

// Frequency frequency of transform

func (rb *TransformsRecordBuilder) Frequency(frequency string) *TransformsRecordBuilder {
	rb.v.Frequency = &frequency
	return rb
}

// Id the id

func (rb *TransformsRecordBuilder) Id(id Id) *TransformsRecordBuilder {
	rb.v.Id = &id
	return rb
}

// IndexFailure total number of index failures

func (rb *TransformsRecordBuilder) IndexFailure(indexfailure string) *TransformsRecordBuilder {
	rb.v.IndexFailure = &indexfailure
	return rb
}

// IndexTime total time spent indexing documents

func (rb *TransformsRecordBuilder) IndexTime(indextime string) *TransformsRecordBuilder {
	rb.v.IndexTime = &indextime
	return rb
}

// IndexTotal total number of index phases done by the transform

func (rb *TransformsRecordBuilder) IndexTotal(indextotal string) *TransformsRecordBuilder {
	rb.v.IndexTotal = &indextotal
	return rb
}

// IndexedDocumentsExpAvg exponential average number of documents indexed

func (rb *TransformsRecordBuilder) IndexedDocumentsExpAvg(indexeddocumentsexpavg string) *TransformsRecordBuilder {
	rb.v.IndexedDocumentsExpAvg = &indexeddocumentsexpavg
	return rb
}

// LastSearchTime last time transform searched for updates

func (rb *TransformsRecordBuilder) LastSearchTime(lastsearchtime string) *TransformsRecordBuilder {
	rb.v.LastSearchTime = lastsearchtime
	return rb
}

// MaxPageSearchSize max page search size

func (rb *TransformsRecordBuilder) MaxPageSearchSize(maxpagesearchsize string) *TransformsRecordBuilder {
	rb.v.MaxPageSearchSize = &maxpagesearchsize
	return rb
}

// PagesProcessed the number of pages processed

func (rb *TransformsRecordBuilder) PagesProcessed(pagesprocessed string) *TransformsRecordBuilder {
	rb.v.PagesProcessed = &pagesprocessed
	return rb
}

// Pipeline transform pipeline

func (rb *TransformsRecordBuilder) Pipeline(pipeline string) *TransformsRecordBuilder {
	rb.v.Pipeline = &pipeline
	return rb
}

// ProcessedDocumentsExpAvg exponential average number of documents processed

func (rb *TransformsRecordBuilder) ProcessedDocumentsExpAvg(processeddocumentsexpavg string) *TransformsRecordBuilder {
	rb.v.ProcessedDocumentsExpAvg = &processeddocumentsexpavg
	return rb
}

// ProcessingTime the total time spent processing documents

func (rb *TransformsRecordBuilder) ProcessingTime(processingtime string) *TransformsRecordBuilder {
	rb.v.ProcessingTime = &processingtime
	return rb
}

// Reason reason for the current state

func (rb *TransformsRecordBuilder) Reason(reason string) *TransformsRecordBuilder {
	rb.v.Reason = &reason
	return rb
}

// SearchFailure total number of search failures

func (rb *TransformsRecordBuilder) SearchFailure(searchfailure string) *TransformsRecordBuilder {
	rb.v.SearchFailure = &searchfailure
	return rb
}

// SearchTime total search time

func (rb *TransformsRecordBuilder) SearchTime(searchtime string) *TransformsRecordBuilder {
	rb.v.SearchTime = &searchtime
	return rb
}

// SearchTotal total number of search phases

func (rb *TransformsRecordBuilder) SearchTotal(searchtotal string) *TransformsRecordBuilder {
	rb.v.SearchTotal = &searchtotal
	return rb
}

// SourceIndex source index

func (rb *TransformsRecordBuilder) SourceIndex(sourceindex string) *TransformsRecordBuilder {
	rb.v.SourceIndex = &sourceindex
	return rb
}

// State transform state

func (rb *TransformsRecordBuilder) State(state string) *TransformsRecordBuilder {
	rb.v.State = &state
	return rb
}

// TransformType batch or continuous transform

func (rb *TransformsRecordBuilder) TransformType(transformtype string) *TransformsRecordBuilder {
	rb.v.TransformType = &transformtype
	return rb
}

// TriggerCount the number of times the transform has been triggered

func (rb *TransformsRecordBuilder) TriggerCount(triggercount string) *TransformsRecordBuilder {
	rb.v.TriggerCount = &triggercount
	return rb
}

// Version the version of Elasticsearch when the transform was created

func (rb *TransformsRecordBuilder) Version(version VersionString) *TransformsRecordBuilder {
	rb.v.Version = &version
	return rb
}
