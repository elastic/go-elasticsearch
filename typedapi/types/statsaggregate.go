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

// StatsAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/aggregations/Aggregate.ts#L240-L255
type StatsAggregate struct {
	Avg         Float64  `json:"avg,omitempty"`
	AvgAsString *string  `json:"avg_as_string,omitempty"`
	Count       int64    `json:"count"`
	Max         Float64  `json:"max,omitempty"`
	MaxAsString *string  `json:"max_as_string,omitempty"`
	Meta        Metadata `json:"meta,omitempty"`
	Min         Float64  `json:"min,omitempty"`
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
				return err
			}

		case "avg_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.AvgAsString = &o

		case "count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Count = value
			case float64:
				f := int64(v)
				s.Count = f
			}

		case "max":
			if err := dec.Decode(&s.Max); err != nil {
				return err
			}

		case "max_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MaxAsString = &o

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "min":
			if err := dec.Decode(&s.Min); err != nil {
				return err
			}

		case "min_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MinAsString = &o

		case "sum":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
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
				return err
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
