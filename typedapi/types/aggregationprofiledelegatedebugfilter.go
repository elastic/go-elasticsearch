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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// AggregationProfileDelegateDebugFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_global/search/_types/profile.ts#L70-L75
type AggregationProfileDelegateDebugFilter struct {
	Query                         *string `json:"query,omitempty"`
	ResultsFromMetadata           *int    `json:"results_from_metadata,omitempty"`
	SegmentsCountedInConstantTime *int    `json:"segments_counted_in_constant_time,omitempty"`
	SpecializedFor                *string `json:"specialized_for,omitempty"`
}

func (s *AggregationProfileDelegateDebugFilter) UnmarshalJSON(data []byte) error {

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

		case "query":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Query = &o

		case "results_from_metadata":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ResultsFromMetadata = &value
			case float64:
				f := int(v)
				s.ResultsFromMetadata = &f
			}

		case "segments_counted_in_constant_time":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.SegmentsCountedInConstantTime = &value
			case float64:
				f := int(v)
				s.SegmentsCountedInConstantTime = &f
			}

		case "specialized_for":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SpecializedFor = &o

		}
	}
	return nil
}

// NewAggregationProfileDelegateDebugFilter returns a AggregationProfileDelegateDebugFilter.
func NewAggregationProfileDelegateDebugFilter() *AggregationProfileDelegateDebugFilter {
	r := &AggregationProfileDelegateDebugFilter{}

	return r
}
