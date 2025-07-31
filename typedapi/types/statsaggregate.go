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

// StatsAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/Aggregate.ts#L257-L273
type StatsAggregate struct {
	Avg         *Float64 `json:"avg,omitempty"`
	AvgAsString *string  `json:"avg_as_string,omitempty"`
	Count       int64    `json:"count"`
	Max         *Float64 `json:"max,omitempty"`
	MaxAsString *string  `json:"max_as_string,omitempty"`
	Meta        Metadata `json:"meta,omitempty"`
	Min         *Float64 `json:"min,omitempty"`
	MinAsString *string  `json:"min_as_string,omitempty"`
	Sum         Float64  `json:"sum"`
	SumAsString *string  `json:"sum_as_string,omitempty"`
}

func (s *StatsAggregate) UnmarshalJSON(data []byte) error {

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

		case "avg":
			if err := dec.Decode(&s.Avg); err != nil {
				return fmt.Errorf("%s | %w", "Avg", err)
			}

		case "avg_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "AvgAsString", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.AvgAsString = &o

		case "count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Count", err)
				}
				s.Count = value
			case float64:
				f := int64(v)
				s.Count = f
			}

		case "max":
			if err := dec.Decode(&s.Max); err != nil {
				return fmt.Errorf("%s | %w", "Max", err)
			}

		case "max_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MaxAsString", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MaxAsString = &o

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return fmt.Errorf("%s | %w", "Meta", err)
			}

		case "min":
			if err := dec.Decode(&s.Min); err != nil {
				return fmt.Errorf("%s | %w", "Min", err)
			}

		case "min_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MinAsString", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MinAsString = &o

		case "sum":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Sum", err)
				}
				f := Float64(value)
				s.Sum = f
			case float64:
				f := Float64(v)
				s.Sum = f
			}

		case "sum_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "SumAsString", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SumAsString = &o

		}
	}
	return nil
}

// NewStatsAggregate returns a StatsAggregate.
func NewStatsAggregate() *StatsAggregate {
	r := &StatsAggregate{}

	return r
}
