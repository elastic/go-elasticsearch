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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/minimuminterval"
)

// AutoDateHistogramAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/bucket.ts#L72-L110
type AutoDateHistogramAggregation struct {
	// Buckets The target number of buckets.
	Buckets *int `json:"buckets,omitempty"`
	// Field The field on which to run the aggregation.
	Field *string `json:"field,omitempty"`
	// Format The date format used to format `key_as_string` in the response.
	// If no `format` is specified, the first date format specified in the field
	// mapping is used.
	Format *string `json:"format,omitempty"`
	// MinimumInterval The minimum rounding interval.
	// This can make the collection process more efficient, as the aggregation will
	// not attempt to round at any interval lower than `minimum_interval`.
	MinimumInterval *minimuminterval.MinimumInterval `json:"minimum_interval,omitempty"`
	// Missing The value to apply to documents that do not have a value.
	// By default, documents without a value are ignored.
	Missing DateTime `json:"missing,omitempty"`
	// Offset Time zone specified as a ISO 8601 UTC offset.
	Offset *string                    `json:"offset,omitempty"`
	Params map[string]json.RawMessage `json:"params,omitempty"`
	Script *Script                    `json:"script,omitempty"`
	// TimeZone Time zone ID.
	TimeZone *string `json:"time_zone,omitempty"`
}

func (s *AutoDateHistogramAggregation) UnmarshalJSON(data []byte) error {

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

		case "buckets":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Buckets", err)
				}
				s.Buckets = &value
			case float64:
				f := int(v)
				s.Buckets = &f
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Format", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Format = &o

		case "minimum_interval":
			if err := dec.Decode(&s.MinimumInterval); err != nil {
				return fmt.Errorf("%s | %w", "MinimumInterval", err)
			}

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return fmt.Errorf("%s | %w", "Missing", err)
			}

		case "offset":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Offset", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Offset = &o

		case "params":
			if s.Params == nil {
				s.Params = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Params); err != nil {
				return fmt.Errorf("%s | %w", "Params", err)
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}

		case "time_zone":
			if err := dec.Decode(&s.TimeZone); err != nil {
				return fmt.Errorf("%s | %w", "TimeZone", err)
			}

		}
	}
	return nil
}

// NewAutoDateHistogramAggregation returns a AutoDateHistogramAggregation.
func NewAutoDateHistogramAggregation() *AutoDateHistogramAggregation {
	r := &AutoDateHistogramAggregation{
		Params: make(map[string]json.RawMessage),
	}

	return r
}
