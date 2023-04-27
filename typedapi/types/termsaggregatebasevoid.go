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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// TermsAggregateBaseVoid type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/aggregations/Aggregate.ts#L376-L381
type TermsAggregateBaseVoid struct {
	Buckets                 BucketsVoid `json:"buckets"`
	DocCountErrorUpperBound *int64      `json:"doc_count_error_upper_bound,omitempty"`
	Meta                    Metadata    `json:"meta,omitempty"`
	SumOtherDocCount        *int64      `json:"sum_other_doc_count,omitempty"`
}

func (s *TermsAggregateBaseVoid) UnmarshalJSON(data []byte) error {

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

		case "buckets":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := make(map[string]interface{}, 0)
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Buckets = o
			case '[':
				o := []interface{}{}
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Buckets = o
			}

		case "doc_count_error_upper_bound":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DocCountErrorUpperBound = &value
			case float64:
				f := int64(v)
				s.DocCountErrorUpperBound = &f
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "sum_other_doc_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.SumOtherDocCount = &value
			case float64:
				f := int64(v)
				s.SumOtherDocCount = &f
			}

		}
	}
	return nil
}

// NewTermsAggregateBaseVoid returns a TermsAggregateBaseVoid.
func NewTermsAggregateBaseVoid() *TermsAggregateBaseVoid {
	r := &TermsAggregateBaseVoid{}

	return r
}
