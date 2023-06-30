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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/categorizationstatus"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/memorystatus"
)

// ModelSizeStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/_types/Model.ts#L56-L78
type ModelSizeStats struct {
	AssignmentMemoryBasis         *string                                   `json:"assignment_memory_basis,omitempty"`
	BucketAllocationFailuresCount int64                                     `json:"bucket_allocation_failures_count"`
	CategorizationStatus          categorizationstatus.CategorizationStatus `json:"categorization_status"`
	CategorizedDocCount           int                                       `json:"categorized_doc_count"`
	DeadCategoryCount             int                                       `json:"dead_category_count"`
	FailedCategoryCount           int                                       `json:"failed_category_count"`
	FrequentCategoryCount         int                                       `json:"frequent_category_count"`
	JobId                         string                                    `json:"job_id"`
	LogTime                       DateTime                                  `json:"log_time"`
	MemoryStatus                  memorystatus.MemoryStatus                 `json:"memory_status"`
	ModelBytes                    ByteSize                                  `json:"model_bytes"`
	ModelBytesExceeded            ByteSize                                  `json:"model_bytes_exceeded,omitempty"`
	ModelBytesMemoryLimit         ByteSize                                  `json:"model_bytes_memory_limit,omitempty"`
	PeakModelBytes                ByteSize                                  `json:"peak_model_bytes,omitempty"`
	RareCategoryCount             int                                       `json:"rare_category_count"`
	ResultType                    string                                    `json:"result_type"`
	Timestamp                     *int64                                    `json:"timestamp,omitempty"`
	TotalByFieldCount             int64                                     `json:"total_by_field_count"`
	TotalCategoryCount            int                                       `json:"total_category_count"`
	TotalOverFieldCount           int64                                     `json:"total_over_field_count"`
	TotalPartitionFieldCount      int64                                     `json:"total_partition_field_count"`
}

func (s *ModelSizeStats) UnmarshalJSON(data []byte) error {

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

		case "assignment_memory_basis":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.AssignmentMemoryBasis = &o

		case "bucket_allocation_failures_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.BucketAllocationFailuresCount = value
			case float64:
				f := int64(v)
				s.BucketAllocationFailuresCount = f
			}

		case "categorization_status":
			if err := dec.Decode(&s.CategorizationStatus); err != nil {
				return err
			}

		case "categorized_doc_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.CategorizedDocCount = value
			case float64:
				f := int(v)
				s.CategorizedDocCount = f
			}

		case "dead_category_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.DeadCategoryCount = value
			case float64:
				f := int(v)
				s.DeadCategoryCount = f
			}

		case "failed_category_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.FailedCategoryCount = value
			case float64:
				f := int(v)
				s.FailedCategoryCount = f
			}

		case "frequent_category_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.FrequentCategoryCount = value
			case float64:
				f := int(v)
				s.FrequentCategoryCount = f
			}

		case "job_id":
			if err := dec.Decode(&s.JobId); err != nil {
				return err
			}

		case "log_time":
			if err := dec.Decode(&s.LogTime); err != nil {
				return err
			}

		case "memory_status":
			if err := dec.Decode(&s.MemoryStatus); err != nil {
				return err
			}

		case "model_bytes":
			if err := dec.Decode(&s.ModelBytes); err != nil {
				return err
			}

		case "model_bytes_exceeded":
			if err := dec.Decode(&s.ModelBytesExceeded); err != nil {
				return err
			}

		case "model_bytes_memory_limit":
			if err := dec.Decode(&s.ModelBytesMemoryLimit); err != nil {
				return err
			}

		case "peak_model_bytes":
			if err := dec.Decode(&s.PeakModelBytes); err != nil {
				return err
			}

		case "rare_category_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.RareCategoryCount = value
			case float64:
				f := int(v)
				s.RareCategoryCount = f
			}

		case "result_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ResultType = o

		case "timestamp":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Timestamp = &value
			case float64:
				f := int64(v)
				s.Timestamp = &f
			}

		case "total_by_field_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalByFieldCount = value
			case float64:
				f := int64(v)
				s.TotalByFieldCount = f
			}

		case "total_category_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.TotalCategoryCount = value
			case float64:
				f := int(v)
				s.TotalCategoryCount = f
			}

		case "total_over_field_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalOverFieldCount = value
			case float64:
				f := int64(v)
				s.TotalOverFieldCount = f
			}

		case "total_partition_field_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalPartitionFieldCount = value
			case float64:
				f := int64(v)
				s.TotalPartitionFieldCount = f
			}

		}
	}
	return nil
}

// NewModelSizeStats returns a ModelSizeStats.
func NewModelSizeStats() *ModelSizeStats {
	r := &ModelSizeStats{}

	return r
}
