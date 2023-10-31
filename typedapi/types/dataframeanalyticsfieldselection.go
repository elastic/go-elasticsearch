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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// DataframeAnalyticsFieldSelection type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/DataframeAnalytics.ts#L55-L68
type DataframeAnalyticsFieldSelection struct {
	// FeatureType The feature type of this field for the analysis. May be categorical or
	// numerical.
	FeatureType *string `json:"feature_type,omitempty"`
	// IsIncluded Whether the field is selected to be included in the analysis.
	IsIncluded bool `json:"is_included"`
	// IsRequired Whether the field is required.
	IsRequired bool `json:"is_required"`
	// MappingTypes The mapping types of the field.
	MappingTypes []string `json:"mapping_types"`
	// Name The field name.
	Name string `json:"name"`
	// Reason The reason a field is not selected to be included in the analysis.
	Reason *string `json:"reason,omitempty"`
}

func (s *DataframeAnalyticsFieldSelection) UnmarshalJSON(data []byte) error {

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

		case "feature_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FeatureType = &o

		case "is_included":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IsIncluded = value
			case bool:
				s.IsIncluded = v
			}

		case "is_required":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IsRequired = value
			case bool:
				s.IsRequired = v
			}

		case "mapping_types":
			if err := dec.Decode(&s.MappingTypes); err != nil {
				return err
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "reason":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Reason = &o

		}
	}
	return nil
}

// NewDataframeAnalyticsFieldSelection returns a DataframeAnalyticsFieldSelection.
func NewDataframeAnalyticsFieldSelection() *DataframeAnalyticsFieldSelection {
	r := &DataframeAnalyticsFieldSelection{}

	return r
}
