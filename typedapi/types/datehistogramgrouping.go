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

	"encoding/json"
)

// DateHistogramGrouping type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/rollup/_types/Groupings.ts#L30-L38
type DateHistogramGrouping struct {
	CalendarInterval Duration `json:"calendar_interval,omitempty"`
	Delay            Duration `json:"delay,omitempty"`
	Field            string   `json:"field"`
	FixedInterval    Duration `json:"fixed_interval,omitempty"`
	Format           *string  `json:"format,omitempty"`
	Interval         Duration `json:"interval,omitempty"`
	TimeZone         *string  `json:"time_zone,omitempty"`
}

func (s *DateHistogramGrouping) UnmarshalJSON(data []byte) error {

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

		case "calendar_interval":
			if err := dec.Decode(&s.CalendarInterval); err != nil {
				return err
			}

		case "delay":
			if err := dec.Decode(&s.Delay); err != nil {
				return err
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "fixed_interval":
			if err := dec.Decode(&s.FixedInterval); err != nil {
				return err
			}

		case "format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Format = &o

		case "interval":
			if err := dec.Decode(&s.Interval); err != nil {
				return err
			}

		case "time_zone":
			if err := dec.Decode(&s.TimeZone); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDateHistogramGrouping returns a DateHistogramGrouping.
func NewDateHistogramGrouping() *DateHistogramGrouping {
	r := &DateHistogramGrouping{}

	return r
}
