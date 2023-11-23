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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/calendarinterval"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

// DateHistogramAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/aggregations/bucket.ts#L189-L247
type DateHistogramAggregation struct {
	// CalendarInterval Calendar-aware interval.
	// Can be specified using the unit name, such as `month`, or as a single unit
	// quantity, such as `1M`.
	CalendarInterval *calendarinterval.CalendarInterval `json:"calendar_interval,omitempty"`
	// ExtendedBounds Enables extending the bounds of the histogram beyond the data itself.
	ExtendedBounds *ExtendedBoundsFieldDateMath `json:"extended_bounds,omitempty"`
	// Field The date field whose values are use to build a histogram.
	Field *string `json:"field,omitempty"`
	// FixedInterval Fixed intervals: a fixed number of SI units and never deviate, regardless of
	// where they fall on the calendar.
	FixedInterval Duration `json:"fixed_interval,omitempty"`
	// Format The date format used to format `key_as_string` in the response.
	// If no `format` is specified, the first date format specified in the field
	// mapping is used.
	Format *string `json:"format,omitempty"`
	// HardBounds Limits the histogram to specified bounds.
	HardBounds *ExtendedBoundsFieldDateMath `json:"hard_bounds,omitempty"`
	Interval   Duration                     `json:"interval,omitempty"`
	// Keyed Set to `true` to associate a unique string key with each bucket and return
	// the ranges as a hash rather than an array.
	Keyed *bool    `json:"keyed,omitempty"`
	Meta  Metadata `json:"meta,omitempty"`
	// MinDocCount Only returns buckets that have `min_doc_count` number of documents.
	// By default, all buckets between the first bucket that matches documents and
	// the last one are returned.
	MinDocCount *int `json:"min_doc_count,omitempty"`
	// Missing The value to apply to documents that do not have a value.
	// By default, documents without a value are ignored.
	Missing DateTime `json:"missing,omitempty"`
	Name    *string  `json:"name,omitempty"`
	// Offset Changes the start value of each bucket by the specified positive (`+`) or
	// negative offset (`-`) duration.
	Offset Duration `json:"offset,omitempty"`
	// Order The sort order of the returned buckets.
	Order  AggregateOrder             `json:"order,omitempty"`
	Params map[string]json.RawMessage `json:"params,omitempty"`
	Script Script                     `json:"script,omitempty"`
	// TimeZone Time zone used for bucketing and rounding.
	// Defaults to Coordinated Universal Time (UTC).
	TimeZone *string `json:"time_zone,omitempty"`
}

func (s *DateHistogramAggregation) UnmarshalJSON(data []byte) error {

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

		case "extended_bounds":
			if err := dec.Decode(&s.ExtendedBounds); err != nil {
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

		case "hard_bounds":
			if err := dec.Decode(&s.HardBounds); err != nil {
				return err
			}

		case "interval":
			if err := dec.Decode(&s.Interval); err != nil {
				return err
			}

		case "keyed":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Keyed = &value
			case bool:
				s.Keyed = &v
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "min_doc_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MinDocCount = &value
			case float64:
				f := int(v)
				s.MinDocCount = &f
			}

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return err
			}

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = &o

		case "offset":
			if err := dec.Decode(&s.Offset); err != nil {
				return err
			}

		case "order":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := make(map[string]sortorder.SortOrder, 0)
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Order = o
			case '[':
				o := make([]map[string]sortorder.SortOrder, 0)
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Order = o
			}

		case "params":
			if s.Params == nil {
				s.Params = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Params); err != nil {
				return err
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
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

// NewDateHistogramAggregation returns a DateHistogramAggregation.
func NewDateHistogramAggregation() *DateHistogramAggregation {
	r := &DateHistogramAggregation{
		Params: make(map[string]json.RawMessage, 0),
	}

	return r
}
