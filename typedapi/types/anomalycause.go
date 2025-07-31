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

// AnomalyCause type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/_types/Anomaly.ts#L123-L139
type AnomalyCause struct {
	Actual                 []Float64   `json:"actual,omitempty"`
	ByFieldName            *string     `json:"by_field_name,omitempty"`
	ByFieldValue           *string     `json:"by_field_value,omitempty"`
	CorrelatedByFieldValue *string     `json:"correlated_by_field_value,omitempty"`
	FieldName              *string     `json:"field_name,omitempty"`
	Function               *string     `json:"function,omitempty"`
	FunctionDescription    *string     `json:"function_description,omitempty"`
	GeoResults             *GeoResults `json:"geo_results,omitempty"`
	Influencers            []Influence `json:"influencers,omitempty"`
	OverFieldName          *string     `json:"over_field_name,omitempty"`
	OverFieldValue         *string     `json:"over_field_value,omitempty"`
	PartitionFieldName     *string     `json:"partition_field_name,omitempty"`
	PartitionFieldValue    *string     `json:"partition_field_value,omitempty"`
	Probability            Float64     `json:"probability"`
	Typical                []Float64   `json:"typical,omitempty"`
}

func (s *AnomalyCause) UnmarshalJSON(data []byte) error {

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

		case "actual":
			if err := dec.Decode(&s.Actual); err != nil {
				return fmt.Errorf("%s | %w", "Actual", err)
			}

		case "by_field_name":
			if err := dec.Decode(&s.ByFieldName); err != nil {
				return fmt.Errorf("%s | %w", "ByFieldName", err)
			}

		case "by_field_value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ByFieldValue", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ByFieldValue = &o

		case "correlated_by_field_value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "CorrelatedByFieldValue", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CorrelatedByFieldValue = &o

		case "field_name":
			if err := dec.Decode(&s.FieldName); err != nil {
				return fmt.Errorf("%s | %w", "FieldName", err)
			}

		case "function":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Function", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Function = &o

		case "function_description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "FunctionDescription", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FunctionDescription = &o

		case "geo_results":
			if err := dec.Decode(&s.GeoResults); err != nil {
				return fmt.Errorf("%s | %w", "GeoResults", err)
			}

		case "influencers":
			if err := dec.Decode(&s.Influencers); err != nil {
				return fmt.Errorf("%s | %w", "Influencers", err)
			}

		case "over_field_name":
			if err := dec.Decode(&s.OverFieldName); err != nil {
				return fmt.Errorf("%s | %w", "OverFieldName", err)
			}

		case "over_field_value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "OverFieldValue", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.OverFieldValue = &o

		case "partition_field_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "PartitionFieldName", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PartitionFieldName = &o

		case "partition_field_value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "PartitionFieldValue", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PartitionFieldValue = &o

		case "probability":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Probability", err)
				}
				f := Float64(value)
				s.Probability = f
			case float64:
				f := Float64(v)
				s.Probability = f
			}

		case "typical":
			if err := dec.Decode(&s.Typical); err != nil {
				return fmt.Errorf("%s | %w", "Typical", err)
			}

		}
	}
	return nil
}

// NewAnomalyCause returns a AnomalyCause.
func NewAnomalyCause() *AnomalyCause {
	r := &AnomalyCause{}

	return r
}
