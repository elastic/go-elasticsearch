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

// IntervalsAllOf type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/query_dsl/fulltext.ts#L49-L56
type IntervalsAllOf struct {
	Filter    *IntervalsFilter `json:"filter,omitempty"`
	Intervals []Intervals      `json:"intervals"`
	MaxGaps   *int             `json:"max_gaps,omitempty"`
	Ordered   *bool            `json:"ordered,omitempty"`
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
