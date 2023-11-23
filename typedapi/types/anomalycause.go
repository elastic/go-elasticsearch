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

// AnomalyCause type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/Anomaly.ts#L123-L138
type AnomalyCause struct {
	Actual                 []Float64   `json:"actual"`
	ByFieldName            string      `json:"by_field_name"`
	ByFieldValue           string      `json:"by_field_value"`
	CorrelatedByFieldValue string      `json:"correlated_by_field_value"`
	FieldName              string      `json:"field_name"`
	Function               string      `json:"function"`
	FunctionDescription    string      `json:"function_description"`
	Influencers            []Influence `json:"influencers"`
	OverFieldName          string      `json:"over_field_name"`
	OverFieldValue         string      `json:"over_field_value"`
	PartitionFieldName     string      `json:"partition_field_name"`
	PartitionFieldValue    string      `json:"partition_field_value"`
	Probability            Float64     `json:"probability"`
	Typical                []Float64   `json:"typical"`
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
				return err
			}

		case "by_field_name":
			if err := dec.Decode(&s.ByFieldName); err != nil {
				return err
			}

		case "by_field_value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ByFieldValue = o

		case "correlated_by_field_value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CorrelatedByFieldValue = o

		case "field_name":
			if err := dec.Decode(&s.FieldName); err != nil {
				return err
			}

		case "function":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Function = o

		case "function_description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FunctionDescription = o

		case "influencers":
			if err := dec.Decode(&s.Influencers); err != nil {
				return err
			}

		case "over_field_name":
			if err := dec.Decode(&s.OverFieldName); err != nil {
				return err
			}

		case "over_field_value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.OverFieldValue = o

		case "partition_field_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PartitionFieldName = o

		case "partition_field_value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PartitionFieldValue = o

		case "probability":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Probability = f
			case float64:
				f := Float64(v)
				s.Probability = f
			}

		case "typical":
			if err := dec.Decode(&s.Typical); err != nil {
				return err
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
