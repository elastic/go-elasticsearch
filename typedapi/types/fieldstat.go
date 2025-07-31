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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// FieldStat type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/text_structure/_types/Structure.ts#L23-L33
type FieldStat struct {
	Cardinality int      `json:"cardinality"`
	Count       int      `json:"count"`
	Earliest    *string  `json:"earliest,omitempty"`
	Latest      *string  `json:"latest,omitempty"`
	MaxValue    *int     `json:"max_value,omitempty"`
	MeanValue   *int     `json:"mean_value,omitempty"`
	MedianValue *int     `json:"median_value,omitempty"`
	MinValue    *int     `json:"min_value,omitempty"`
	TopHits     []TopHit `json:"top_hits"`
}

func (s *FieldStat) UnmarshalJSON(data []byte) error {

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

		case "cardinality":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Cardinality", err)
				}
				s.Cardinality = value
			case float64:
				f := int(v)
				s.Cardinality = f
			}

		case "count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Count", err)
				}
				s.Count = value
			case float64:
				f := int(v)
				s.Count = f
			}

		case "earliest":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Earliest", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Earliest = &o

		case "latest":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Latest", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Latest = &o

		case "max_value":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxValue", err)
				}
				s.MaxValue = &value
			case float64:
				f := int(v)
				s.MaxValue = &f
			}

		case "mean_value":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MeanValue", err)
				}
				s.MeanValue = &value
			case float64:
				f := int(v)
				s.MeanValue = &f
			}

		case "median_value":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MedianValue", err)
				}
				s.MedianValue = &value
			case float64:
				f := int(v)
				s.MedianValue = &f
			}

		case "min_value":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinValue", err)
				}
				s.MinValue = &value
			case float64:
				f := int(v)
				s.MinValue = &f
			}

		case "top_hits":
			if err := dec.Decode(&s.TopHits); err != nil {
				return fmt.Errorf("%s | %w", "TopHits", err)
			}

		}
	}
	return nil
}

// NewFieldStat returns a FieldStat.
func NewFieldStat() *FieldStat {
	r := &FieldStat{}

	return r
}
