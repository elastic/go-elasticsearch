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

// IntervalsAllOf type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/query_dsl/fulltext.ts#L50-L70
type IntervalsAllOf struct {
	// Filter Rule used to filter returned intervals.
	Filter *IntervalsFilter `json:"filter,omitempty"`
	// Intervals An array of rules to combine. All rules must produce a match in a document
	// for the overall source to match.
	Intervals []Intervals `json:"intervals"`
	// MaxGaps Maximum number of positions between the matching terms.
	// Intervals produced by the rules further apart than this are not considered
	// matches.
	MaxGaps *int `json:"max_gaps,omitempty"`
	// Ordered If `true`, intervals produced by the rules should appear in the order in
	// which they are specified.
	Ordered *bool `json:"ordered,omitempty"`
}

func (s *IntervalsAllOf) UnmarshalJSON(data []byte) error {

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

		case "filter":
			if err := dec.Decode(&s.Filter); err != nil {
				return err
			}

		case "intervals":
			if err := dec.Decode(&s.Intervals); err != nil {
				return err
			}

		case "max_gaps":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxGaps = &value
			case float64:
				f := int(v)
				s.MaxGaps = &f
			}

		case "ordered":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Ordered = &value
			case bool:
				s.Ordered = &v
			}

		}
	}
	return nil
}

// NewIntervalsAllOf returns a IntervalsAllOf.
func NewIntervalsAllOf() *IntervalsAllOf {
	r := &IntervalsAllOf{}

	return r
}
