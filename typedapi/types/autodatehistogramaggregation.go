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
// https://github.com/elastic/elasticsearch-specification/tree/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/minimuminterval"
)

// AutoDateHistogramAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/_types/aggregations/bucket.ts#L65-L100
type AutoDateHistogramAggregation struct {
	// Buckets The target number of buckets.
	Buckets *int `json:"buckets,omitempty"`
	// Field The field on which to run the aggregation.
	Field *string `json:"field,omitempty"`
	// Format The date format used to format `key_as_string` in the response.
	// If no `format` is specified, the first date format specified in the field
	// mapping is used.
	Format *string  `json:"format,omitempty"`
	Meta   Metadata `json:"meta,omitempty"`
	// MinimumInterval The minimum rounding interval.
	// This can make the collection process more efficient, as the aggregation will
	// not attempt to round at any interval lower than `minimum_interval`.
	MinimumInterval *minimuminterval.MinimumInterval `json:"minimum_interval,omitempty"`
	// Missing The value to apply to documents that do not have a value.
	// By default, documents without a value are ignored.
	Missing DateTime `json:"missing,omitempty"`
	Name    *string  `json:"name,omitempty"`
	// Offset Time zone specified as a ISO 8601 UTC offset.
	Offset *string                    `json:"offset,omitempty"`
	Params map[string]json.RawMessage `json:"params,omitempty"`
	Script Script                     `json:"script,omitempty"`
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

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Buckets = &value
			case float64:
				f := int(v)
				s.Buckets = &f
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
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

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "minimum_interval":
			if err := dec.Decode(&s.MinimumInterval); err != nil {
				return err
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
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
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
				return err
			}

		case "script":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return err
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return err
				}

				switch t {

				case "lang", "options", "source":
					o := NewInlineScript()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Script = o

				case "id":
					o := NewStoredScriptId()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Script = o

				}
			}

		case "time_zone":
			if err := dec.Decode(&s.TimeZone); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewAutoDateHistogramAggregation returns a AutoDateHistogramAggregation.
func NewAutoDateHistogramAggregation() *AutoDateHistogramAggregation {
	r := &AutoDateHistogramAggregation{
		Params: make(map[string]json.RawMessage, 0),
	}

	return r
}
