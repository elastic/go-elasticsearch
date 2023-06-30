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

// OutlierDetectionParameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/_types/DataframeAnalytics.ts#L412-L419
type OutlierDetectionParameters struct {
	ComputeFeatureInfluence   *bool    `json:"compute_feature_influence,omitempty"`
	FeatureInfluenceThreshold *Float64 `json:"feature_influence_threshold,omitempty"`
	Method                    *string  `json:"method,omitempty"`
	NNeighbors                *int     `json:"n_neighbors,omitempty"`
	OutlierFraction           *Float64 `json:"outlier_fraction,omitempty"`
	StandardizationEnabled    *bool    `json:"standardization_enabled,omitempty"`
}

func (s *OutlierDetectionParameters) UnmarshalJSON(data []byte) error {

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

		case "compute_feature_influence":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.ComputeFeatureInfluence = &value
			case bool:
				s.ComputeFeatureInfluence = &v
			}

		case "feature_influence_threshold":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.FeatureInfluenceThreshold = &f
			case float64:
				f := Float64(v)
				s.FeatureInfluenceThreshold = &f
			}

		case "method":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Method = &o

		case "n_neighbors":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.NNeighbors = &value
			case float64:
				f := int(v)
				s.NNeighbors = &f
			}

		case "outlier_fraction":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.OutlierFraction = &f
			case float64:
				f := Float64(v)
				s.OutlierFraction = &f
			}

		case "standardization_enabled":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.StandardizationEnabled = &value
			case bool:
				s.StandardizationEnabled = &v
			}

		}
	}
	return nil
}

// NewOutlierDetectionParameters returns a OutlierDetectionParameters.
func NewOutlierDetectionParameters() *OutlierDetectionParameters {
	r := &OutlierDetectionParameters{}

	return r
}
