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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/calendarinterval"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/ratemode"
)

// RateAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/_types/aggregations/metric.ts#L230-L241
type RateAggregation struct {
	// Field The field on which to run the aggregation.
	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`
	// Missing The value to apply to documents that do not have a value.
	// By default, documents without a value are ignored.
	Missing Missing `json:"missing,omitempty"`
	// Mode How the rate is calculated.
	Mode   *ratemode.RateMode `json:"mode,omitempty"`
	Script Script             `json:"script,omitempty"`
	// Unit The interval used to calculate the rate.
	// By default, the interval of the `date_histogram` is used.
	Unit *calendarinterval.CalendarInterval `json:"unit,omitempty"`
}

func (s *RateAggregation) UnmarshalJSON(data []byte) error {

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

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return fmt.Errorf("%s | %w", "Missing", err)
			}

		case "mode":
			if err := dec.Decode(&s.Mode); err != nil {
				return fmt.Errorf("%s | %w", "Mode", err)
			}

		case "script":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Script", err)
				}

				switch t {

				case "lang", "options", "source":
					o := NewInlineScript()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Script", err)
					}
					s.Script = o

				case "id":
					o := NewStoredScriptId()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Script", err)
					}
					s.Script = o

				}
			}

		case "unit":
			if err := dec.Decode(&s.Unit); err != nil {
				return fmt.Errorf("%s | %w", "Unit", err)
			}

		}
	}
	return nil
}

// NewRateAggregation returns a RateAggregation.
func NewRateAggregation() *RateAggregation {
	r := &RateAggregation{}

	return r
}
