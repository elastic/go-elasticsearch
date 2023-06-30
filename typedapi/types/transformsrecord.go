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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

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
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/cat/transforms/types.ts#L22-L187
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
