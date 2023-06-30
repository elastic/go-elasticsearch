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

// MappingLimitSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/indices/_types/IndexSettings.ts#L402-L415
type MappingLimitSettings struct {
	Coerce          *bool                                `json:"coerce,omitempty"`
	Depth           *MappingLimitSettingsDepth           `json:"depth,omitempty"`
	DimensionFields *MappingLimitSettingsDimensionFields `json:"dimension_fields,omitempty"`
	FieldNameLength *MappingLimitSettingsFieldNameLength `json:"field_name_length,omitempty"`
	IgnoreMalformed *bool                                `json:"ignore_malformed,omitempty"`
	NestedFields    *MappingLimitSettingsNestedFields    `json:"nested_fields,omitempty"`
	NestedObjects   *MappingLimitSettingsNestedObjects   `json:"nested_objects,omitempty"`
	TotalFields     *MappingLimitSettingsTotalFields     `json:"total_fields,omitempty"`
}

func (s *MappingLimitSettings) UnmarshalJSON(data []byte) error {

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

		case "coerce":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Coerce = &value
			case bool:
				s.Coerce = &v
			}

		case "depth":
			if err := dec.Decode(&s.Depth); err != nil {
				return err
			}

		case "dimension_fields":
			if err := dec.Decode(&s.DimensionFields); err != nil {
				return err
			}

		case "field_name_length":
			if err := dec.Decode(&s.FieldNameLength); err != nil {
				return err
			}

		case "ignore_malformed":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IgnoreMalformed = &value
			case bool:
				s.IgnoreMalformed = &v
			}

		case "nested_fields":
			if err := dec.Decode(&s.NestedFields); err != nil {
				return err
			}

		case "nested_objects":
			if err := dec.Decode(&s.NestedObjects); err != nil {
				return err
			}

		case "total_fields":
			if err := dec.Decode(&s.TotalFields); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewMappingLimitSettings returns a MappingLimitSettings.
func NewMappingLimitSettings() *MappingLimitSettings {
	r := &MappingLimitSettings{}

	return r
}
