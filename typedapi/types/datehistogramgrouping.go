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

// DateHistogramGrouping type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/rollup/_types/Groupings.ts#L42-L73
type DateHistogramGrouping struct {
	// CalendarInterval The interval of time buckets to be generated when rolling up.
	CalendarInterval Duration `json:"calendar_interval,omitempty"`
	// Delay How long to wait before rolling up new documents.
	// By default, the indexer attempts to roll up all data that is available.
	// However, it is not uncommon for data to arrive out of order.
	// The indexer is unable to deal with data that arrives after a time-span has
	// been rolled up.
	// You need to specify a delay that matches the longest period of time you
	// expect out-of-order data to arrive.
	Delay Duration `json:"delay,omitempty"`
	// Field The date field that is to be rolled up.
	Field string `json:"field"`
	// FixedInterval The interval of time buckets to be generated when rolling up.
	FixedInterval Duration `json:"fixed_interval,omitempty"`
	Format        *string  `json:"format,omitempty"`
	Interval      Duration `json:"interval,omitempty"`
	// TimeZone Defines what `time_zone` the rollup documents are stored as.
	// Unlike raw data, which can shift timezones on the fly, rolled documents have
	// to be stored with a specific timezone.
	// By default, rollup documents are stored in `UTC`.
	TimeZone *string `json:"time_zone,omitempty"`
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
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
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
