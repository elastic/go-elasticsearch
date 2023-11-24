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

// RollupJobStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/rollup/get_jobs/types.ts#L45-L58
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

func (s *RollupJobStats) UnmarshalJSON(data []byte) error {

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

		case "documents_processed":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DocumentsProcessed = value
			case float64:
				f := int64(v)
				s.DocumentsProcessed = f
			}

		case "index_failures":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.IndexFailures = value
			case float64:
				f := int64(v)
				s.IndexFailures = f
			}

		case "index_time_in_ms":
			if err := dec.Decode(&s.IndexTimeInMs); err != nil {
				return err
			}

		case "index_total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
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
					return err
				}
				s.PagesProcessed = value
			case float64:
				f := int64(v)
				s.PagesProcessed = f
			}

		case "processing_time_in_ms":
			if err := dec.Decode(&s.ProcessingTimeInMs); err != nil {
				return err
			}

		case "processing_total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ProcessingTotal = value
			case float64:
				f := int64(v)
				s.ProcessingTotal = f
			}

		case "rollups_indexed":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.RollupsIndexed = value
			case float64:
				f := int64(v)
				s.RollupsIndexed = f
			}

		case "search_failures":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.SearchFailures = value
			case float64:
				f := int64(v)
				s.SearchFailures = f
			}

		case "search_time_in_ms":
			if err := dec.Decode(&s.SearchTimeInMs); err != nil {
				return err
			}

		case "search_total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
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
					return err
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

// NewRollupJobStats returns a RollupJobStats.
func NewRollupJobStats() *RollupJobStats {
	r := &RollupJobStats{}

	return r
}
