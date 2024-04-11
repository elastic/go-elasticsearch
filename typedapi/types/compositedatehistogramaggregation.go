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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/valuetype"
)

// CompositeDateHistogramAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/_types/aggregations/bucket.ts#L172-L180
type CompositeDateHistogramAggregation struct {
	// CalendarInterval Either `calendar_interval` or `fixed_interval` must be present
	CalendarInterval *string `json:"calendar_interval,omitempty"`
	// Field Either `field` or `script` must be present
	Field *string `json:"field,omitempty"`
	// FixedInterval Either `calendar_interval` or `fixed_interval` must be present
	FixedInterval *string                    `json:"fixed_interval,omitempty"`
	Format        *string                    `json:"format,omitempty"`
	MissingBucket *bool                      `json:"missing_bucket,omitempty"`
	MissingOrder  *missingorder.MissingOrder `json:"missing_order,omitempty"`
	Offset        Duration                   `json:"offset,omitempty"`
	Order         *sortorder.SortOrder       `json:"order,omitempty"`
	// Script Either `field` or `script` must be present
	Script    Script               `json:"script,omitempty"`
	TimeZone  *string              `json:"time_zone,omitempty"`
	ValueType *valuetype.ValueType `json:"value_type,omitempty"`
}

func (s *CompositeDateHistogramAggregation) UnmarshalJSON(data []byte) error {

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
				return fmt.Errorf("%s | %w", "CalendarInterval", err)
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "fixed_interval":
			if err := dec.Decode(&s.FixedInterval); err != nil {
				return fmt.Errorf("%s | %w", "FixedInterval", err)
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

		case "missing_bucket":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MissingBucket", err)
				}
				s.MissingBucket = &value
			case bool:
				s.MissingBucket = &v
			}

		case "missing_order":
			if err := dec.Decode(&s.MissingOrder); err != nil {
				return fmt.Errorf("%s | %w", "MissingOrder", err)
			}

		case "offset":
			if err := dec.Decode(&s.Offset); err != nil {
				return fmt.Errorf("%s | %w", "Offset", err)
			}

		case "order":
			if err := dec.Decode(&s.Order); err != nil {
				return fmt.Errorf("%s | %w", "Order", err)
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

		case "time_zone":
			if err := dec.Decode(&s.TimeZone); err != nil {
				return fmt.Errorf("%s | %w", "TimeZone", err)
			}

		case "value_type":
			if err := dec.Decode(&s.ValueType); err != nil {
				return fmt.Errorf("%s | %w", "ValueType", err)
			}

		}
	}
	return nil
}

// NewCompositeDateHistogramAggregation returns a CompositeDateHistogramAggregation.
func NewCompositeDateHistogramAggregation() *CompositeDateHistogramAggregation {
	r := &CompositeDateHistogramAggregation{}

	return r
}
