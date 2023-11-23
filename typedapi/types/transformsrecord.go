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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// TransformsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cat/transforms/types.ts#L22-L197
type TransformsRecord struct {
	// ChangesLastDetectionTime The timestamp when changes were last detected in the source indices.
	ChangesLastDetectionTime string `json:"changes_last_detection_time,omitempty"`
	// Checkpoint The sequence number for the checkpoint.
	Checkpoint *string `json:"checkpoint,omitempty"`
	// CheckpointDurationTimeExpAvg The exponential moving average of the duration of the checkpoint, in
	// milliseconds.
	CheckpointDurationTimeExpAvg *string `json:"checkpoint_duration_time_exp_avg,omitempty"`
	// CheckpointProgress The progress of the next checkpoint that is currently in progress.
	CheckpointProgress string `json:"checkpoint_progress,omitempty"`
	// CreateTime The time the transform was created.
	CreateTime *string `json:"create_time,omitempty"`
	// DeleteTime The total time spent deleting documents, in milliseconds.
	DeleteTime *string `json:"delete_time,omitempty"`
	// Description The description of the transform.
	Description *string `json:"description,omitempty"`
	// DestIndex The destination index for the transform.
	DestIndex *string `json:"dest_index,omitempty"`
	// DocsPerSecond The number of input documents per second.
	DocsPerSecond *string `json:"docs_per_second,omitempty"`
	// DocumentsDeleted The number of documents deleted from the destination index due to the
	// retention policy for the transform.
	DocumentsDeleted *string `json:"documents_deleted,omitempty"`
	// DocumentsIndexed The number of documents that have been indexed into the destination index for
	// the transform.
	DocumentsIndexed *string `json:"documents_indexed,omitempty"`
	// DocumentsProcessed The number of documents that have been processed from the source index of the
	// transform.
	DocumentsProcessed *string `json:"documents_processed,omitempty"`
	// Frequency The interval between checks for changes in the source indices when the
	// transform is running continuously.
	Frequency *string `json:"frequency,omitempty"`
	// Id The transform identifier.
	Id *string `json:"id,omitempty"`
	// IndexFailure The total number of indexing failures.
	IndexFailure *string `json:"index_failure,omitempty"`
	// IndexTime The total time spent indexing documents, in milliseconds.
	IndexTime *string `json:"index_time,omitempty"`
	// IndexTotal The total number of index operations done by the transform.
	IndexTotal *string `json:"index_total,omitempty"`
	// IndexedDocumentsExpAvg The exponential moving average of the number of new documents that have been
	// indexed.
	IndexedDocumentsExpAvg *string `json:"indexed_documents_exp_avg,omitempty"`
	// LastSearchTime The timestamp of the last search in the source indices.
	// This field is shown only if the transform is running.
	LastSearchTime string `json:"last_search_time,omitempty"`
	// MaxPageSearchSize The initial page size that is used for the composite aggregation for each
	// checkpoint.
	MaxPageSearchSize *string `json:"max_page_search_size,omitempty"`
	// PagesProcessed The number of search or bulk index operations processed.
	// Documents are processed in batches instead of individually.
	PagesProcessed *string `json:"pages_processed,omitempty"`
	// Pipeline The unique identifier for the ingest pipeline.
	Pipeline *string `json:"pipeline,omitempty"`
	// ProcessedDocumentsExpAvg The exponential moving average of the number of documents that have been
	// processed.
	ProcessedDocumentsExpAvg *string `json:"processed_documents_exp_avg,omitempty"`
	// ProcessingTime The total time spent processing results, in milliseconds.
	ProcessingTime *string `json:"processing_time,omitempty"`
	// Reason If a transform has a `failed` state, these details describe the reason for
	// failure.
	Reason *string `json:"reason,omitempty"`
	// SearchFailure The total number of search failures.
	SearchFailure *string `json:"search_failure,omitempty"`
	// SearchTime The total amount of search time, in milliseconds.
	SearchTime *string `json:"search_time,omitempty"`
	// SearchTotal The total number of search operations on the source index for the transform.
	SearchTotal *string `json:"search_total,omitempty"`
	// SourceIndex The source indices for the transform.
	SourceIndex *string `json:"source_index,omitempty"`
	// State The status of the transform.
	// Returned values include:
	// `aborting`: The transform is aborting.
	// `failed: The transform failed. For more information about the failure, check
	// the `reason` field.
	// `indexing`: The transform is actively processing data and creating new
	// documents.
	// `started`: The transform is running but not actively indexing data.
	// `stopped`: The transform is stopped.
	// `stopping`: The transform is stopping.
	State *string `json:"state,omitempty"`
	// TransformType The type of transform: `batch` or `continuous`.
	TransformType *string `json:"transform_type,omitempty"`
	// TriggerCount The number of times the transform has been triggered by the scheduler.
	// For example, the scheduler triggers the transform indexer to check for
	// updates or ingest new data at an interval specified in the `frequency`
	// property.
	TriggerCount *string `json:"trigger_count,omitempty"`
	// Version The version of Elasticsearch that existed on the node when the transform was
	// created.
	Version *string `json:"version,omitempty"`
}

func (s *TransformsRecord) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "changes_last_detection_time", "cldt":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ChangesLastDetectionTime = o

		case "checkpoint", "c":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Checkpoint = &o

		case "checkpoint_duration_time_exp_avg", "cdtea", "checkpointTimeExpAvg":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CheckpointDurationTimeExpAvg = &o

		case "checkpoint_progress", "cp", "checkpointProgress":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CheckpointProgress = o

		case "create_time", "ct", "createTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CreateTime = &o

		case "delete_time", "dtime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DeleteTime = &o

		case "description", "d":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "dest_index", "di", "destIndex":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DestIndex = &o

		case "docs_per_second", "dps":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DocsPerSecond = &o

		case "documents_deleted", "docd":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DocumentsDeleted = &o

		case "documents_indexed", "doci":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DocumentsIndexed = &o

		case "documents_processed", "docp", "documentsProcessed":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DocumentsProcessed = &o

		case "frequency", "f":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Frequency = &o

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return err
			}

		case "index_failure", "if":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexFailure = &o

		case "index_time", "itime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexTime = &o

		case "index_total", "it":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexTotal = &o

		case "indexed_documents_exp_avg", "idea":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexedDocumentsExpAvg = &o

		case "last_search_time", "lst", "lastSearchTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.LastSearchTime = o

		case "max_page_search_size", "mpsz":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MaxPageSearchSize = &o

		case "pages_processed", "pp":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PagesProcessed = &o

		case "pipeline", "p":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Pipeline = &o

		case "processed_documents_exp_avg", "pdea":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ProcessedDocumentsExpAvg = &o

		case "processing_time", "pt":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ProcessingTime = &o

		case "reason", "r":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Reason = &o

		case "search_failure", "sf":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchFailure = &o

		case "search_time", "stime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchTime = &o

		case "search_total", "st":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchTotal = &o

		case "source_index", "si", "sourceIndex":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SourceIndex = &o

		case "state", "s":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.State = &o

		case "transform_type", "tt":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TransformType = &o

		case "trigger_count", "tc":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TriggerCount = &o

		case "version", "v":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewTransformsRecord returns a TransformsRecord.
func NewTransformsRecord() *TransformsRecord {
	r := &TransformsRecord{}

	return r
}
