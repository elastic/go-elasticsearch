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

// FrequentItemSetsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/aggregations/bucket.ts#L1159-L1183
type FrequentItemSetsAggregation struct {
	// Fields Fields to analyze.
	Fields []FrequentItemSetsField `json:"fields"`
	// Filter Query that filters documents from analysis.
	Filter *Query `json:"filter,omitempty"`
	// MinimumSetSize The minimum size of one item set.
	MinimumSetSize *int `json:"minimum_set_size,omitempty"`
	// MinimumSupport The minimum support of one item set.
	MinimumSupport *Float64 `json:"minimum_support,omitempty"`
	// Size The number of top item sets to return.
	Size *int `json:"size,omitempty"`
}

func (s *FrequentItemSetsAggregation) UnmarshalJSON(data []byte) error {

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

		case "fields":
			if err := dec.Decode(&s.Fields); err != nil {
				return err
			}

		case "filter":
			if err := dec.Decode(&s.Filter); err != nil {
				return err
			}

		case "minimum_set_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MinimumSetSize = &value
			case float64:
				f := int(v)
				s.MinimumSetSize = &f
			}

		case "minimum_support":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.MinimumSupport = &f
			case float64:
				f := Float64(v)
				s.MinimumSupport = &f
			}

		case "size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Size = &value
			case float64:
				f := int(v)
				s.Size = &f
			}

		}
	}
	return nil
}

// NewFrequentItemSetsAggregation returns a FrequentItemSetsAggregation.
func NewFrequentItemSetsAggregation() *FrequentItemSetsAggregation {
	r := &FrequentItemSetsAggregation{}

	return r
}
