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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// TargetMeanEncodingPreprocessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/put_trained_model/types.ts#L49-L54
type TargetMeanEncodingPreprocessor struct {
	DefaultValue Float64            `json:"default_value"`
	FeatureName  string             `json:"feature_name"`
	Field        string             `json:"field"`
	TargetMap    map[string]Float64 `json:"target_map"`
}

func (s *TargetMeanEncodingPreprocessor) UnmarshalJSON(data []byte) error {

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

		case "default_value":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.DefaultValue = f
			case float64:
				f := Float64(v)
				s.DefaultValue = f
			}

		case "feature_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FeatureName = o

		case "field":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Field = o

		case "target_map":
			if s.TargetMap == nil {
				s.TargetMap = make(map[string]Float64, 0)
			}
			if err := dec.Decode(&s.TargetMap); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewTargetMeanEncodingPreprocessor returns a TargetMeanEncodingPreprocessor.
func NewTargetMeanEncodingPreprocessor() *TargetMeanEncodingPreprocessor {
	r := &TargetMeanEncodingPreprocessor{
		TargetMap: make(map[string]Float64, 0),
	}

	return r
}
