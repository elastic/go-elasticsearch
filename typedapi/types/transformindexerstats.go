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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// TransformIndexerStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/transform/get_transform_stats/types.ts#L56-L74
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

func (s *TransformIndexerStats) UnmarshalJSON(data []byte) error {

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

		case "delete_time_in_ms":
			if err := dec.Decode(&s.DeleteTimeInMs); err != nil {
				return fmt.Errorf("%s | %w", "DeleteTimeInMs", err)
			}

		case "documents_deleted":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DocumentsDeleted", err)
				}
				s.DocumentsDeleted = &value
			case float64:
				f := int64(v)
				s.DocumentsDeleted = &f
			}

		case "documents_indexed":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DocumentsIndexed", err)
				}
				s.DocumentsIndexed = value
			case float64:
				f := int64(v)
				s.DocumentsIndexed = f
			}

		case "documents_processed":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DocumentsProcessed", err)
				}
				s.DocumentsProcessed = value
			case float64:
				f := int64(v)
				s.DocumentsProcessed = f
			}

		case "exponential_avg_checkpoint_duration_ms":
			if err := dec.Decode(&s.ExponentialAvgCheckpointDurationMs); err != nil {
				return fmt.Errorf("%s | %w", "ExponentialAvgCheckpointDurationMs", err)
			}

		case "exponential_avg_documents_indexed":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ExponentialAvgDocumentsIndexed", err)
				}
				f := Float64(value)
				s.ExponentialAvgDocumentsIndexed = f
			case float64:
				f := Float64(v)
				s.ExponentialAvgDocumentsIndexed = f
			}

		case "exponential_avg_documents_processed":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ExponentialAvgDocumentsProcessed", err)
				}
				f := Float64(value)
				s.ExponentialAvgDocumentsProcessed = f
			case float64:
				f := Float64(v)
				s.ExponentialAvgDocumentsProcessed = f
			}

		case "index_failures":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndexFailures", err)
				}
				s.IndexFailures = value
			case float64:
				f := int64(v)
				s.IndexFailures = f
			}

		case "index_time_in_ms":
			if err := dec.Decode(&s.IndexTimeInMs); err != nil {
				return fmt.Errorf("%s | %w", "IndexTimeInMs", err)
			}

		case "index_total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndexTotal", err)
				}
				s.IndexTotal = value
			case float64:
				f := int64(v)
				s.IndexTotal = f
			}

		case "pages_processed":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PagesProcessed", err)
				}
				s.PagesProcessed = value
			case float64:
				f := int64(v)
				s.PagesProcessed = f
			}

		case "processing_time_in_ms":
			if err := dec.Decode(&s.ProcessingTimeInMs); err != nil {
				return fmt.Errorf("%s | %w", "ProcessingTimeInMs", err)
			}

		case "processing_total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ProcessingTotal", err)
				}
				s.ProcessingTotal = value
			case float64:
				f := int64(v)
				s.ProcessingTotal = f
			}

		case "search_failures":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SearchFailures", err)
				}
				s.SearchFailures = value
			case float64:
				f := int64(v)
				s.SearchFailures = f
			}

		case "search_time_in_ms":
			if err := dec.Decode(&s.SearchTimeInMs); err != nil {
				return fmt.Errorf("%s | %w", "SearchTimeInMs", err)
			}

		case "search_total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SearchTotal", err)
				}
				s.SearchTotal = value
			case float64:
				f := int64(v)
				s.SearchTotal = f
			}

		case "trigger_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TriggerCount", err)
				}
				s.TriggerCount = value
			case float64:
				f := int64(v)
				s.TriggerCount = f
			}

		}
	}
	return nil
}

// NewTransformIndexerStats returns a TransformIndexerStats.
func NewTransformIndexerStats() *TransformIndexerStats {
	r := &TransformIndexerStats{}

	return r
}
