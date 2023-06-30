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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/holtwinterstype"
)

// HoltWintersModelSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/aggregations/pipeline.ts#L235-L242
type HoltWintersModelSettings struct {
	Alpha  *float32                         `json:"alpha,omitempty"`
	Beta   *float32                         `json:"beta,omitempty"`
	Gamma  *float32                         `json:"gamma,omitempty"`
	Pad    *bool                            `json:"pad,omitempty"`
	Period *int                             `json:"period,omitempty"`
	Type   *holtwinterstype.HoltWintersType `json:"type,omitempty"`
}

func (s *HoltWintersModelSettings) UnmarshalJSON(data []byte) error {

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

		case "alpha":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Alpha = &f
			case float64:
				f := float32(v)
				s.Alpha = &f
			}

		case "beta":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Beta = &f
			case float64:
				f := float32(v)
				s.Beta = &f
			}

		case "gamma":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Gamma = &f
			case float64:
				f := float32(v)
				s.Gamma = &f
			}

		case "pad":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Pad = &value
			case bool:
				s.Pad = &v
			}

		case "period":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Period = &value
			case float64:
				f := int(v)
				s.Period = &f
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewHoltWintersModelSettings returns a HoltWintersModelSettings.
func NewHoltWintersModelSettings() *HoltWintersModelSettings {
	r := &HoltWintersModelSettings{}

	return r
}
