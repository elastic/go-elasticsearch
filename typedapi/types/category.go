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

// Category type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/_types/Category.ts#L23-L49
type Category struct {
	// CategoryId A unique identifier for the category. category_id is unique at the job level,
	// even when per-partition categorization is enabled.
	CategoryId uint64 `json:"category_id"`
	// Examples A list of examples of actual values that matched the category.
	Examples []string `json:"examples"`
	// GrokPattern [experimental] A Grok pattern that could be used in Logstash or an ingest
	// pipeline to extract fields from messages that match the category. This field
	// is experimental and may be changed or removed in a future release. The Grok
	// patterns that are found are not optimal, but are often a good starting point
	// for manual tweaking.
	GrokPattern *string `json:"grok_pattern,omitempty"`
	// JobId Identifier for the anomaly detection job.
	JobId string `json:"job_id"`
	// MaxMatchingLength The maximum length of the fields that matched the category. The value is
	// increased by 10% to enable matching for similar fields that have not been
	// analyzed.
	MaxMatchingLength uint64 `json:"max_matching_length"`
	Mlcategory        string `json:"mlcategory"`
	// NumMatches The number of messages that have been matched by this category. This is only
	// guaranteed to have the latest accurate count after a job _flush or _close
	NumMatches *int64  `json:"num_matches,omitempty"`
	P          *string `json:"p,omitempty"`
	// PartitionFieldName If per-partition categorization is enabled, this property identifies the
	// field used to segment the categorization. It is not present when
	// per-partition categorization is disabled.
	PartitionFieldName *string `json:"partition_field_name,omitempty"`
	// PartitionFieldValue If per-partition categorization is enabled, this property identifies the
	// value of the partition_field_name for the category. It is not present when
	// per-partition categorization is disabled.
	PartitionFieldValue *string `json:"partition_field_value,omitempty"`
	// PreferredToCategories A list of category_id entries that this current category encompasses. Any new
	// message that is processed by the categorizer will match against this category
	// and not any of the categories in this list. This is only guaranteed to have
	// the latest accurate list of categories after a job _flush or _close
	PreferredToCategories []string `json:"preferred_to_categories,omitempty"`
	// Regex A regular expression that is used to search for values that match the
	// category.
	Regex      string `json:"regex"`
	ResultType string `json:"result_type"`
	// Terms A space separated list of the common tokens that are matched in values of the
	// category.
	Terms string `json:"terms"`
}

func (s *Category) UnmarshalJSON(data []byte) error {

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

		case "category_id":
			if err := dec.Decode(&s.CategoryId); err != nil {
				return err
			}

		case "examples":
			if err := dec.Decode(&s.Examples); err != nil {
				return err
			}

		case "grok_pattern":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.GrokPattern = &o

		case "job_id":
			if err := dec.Decode(&s.JobId); err != nil {
				return err
			}

		case "max_matching_length":
			if err := dec.Decode(&s.MaxMatchingLength); err != nil {
				return err
			}

		case "mlcategory":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Mlcategory = o

		case "num_matches":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.NumMatches = &value
			case float64:
				f := int64(v)
				s.NumMatches = &f
			}

		case "p":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.P = &o

		case "partition_field_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PartitionFieldName = &o

		case "partition_field_value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PartitionFieldValue = &o

		case "preferred_to_categories":
			if err := dec.Decode(&s.PreferredToCategories); err != nil {
				return err
			}

		case "regex":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Regex = o

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

		case "terms":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Terms = o

		}
	}
	return nil
}

// NewCategory returns a Category.
func NewCategory() *Category {
	r := &Category{}

	return r
}
