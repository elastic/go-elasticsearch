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

// DataframeAnalysisOutlierDetection type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/DataframeAnalytics.ts#L103-L132
type DataframeAnalysisOutlierDetection struct {
	// ComputeFeatureInfluence Specifies whether the feature influence calculation is enabled.
	ComputeFeatureInfluence *bool `json:"compute_feature_influence,omitempty"`
	// FeatureInfluenceThreshold The minimum outlier score that a document needs to have in order to calculate
	// its feature influence score. Value range: 0-1.
	FeatureInfluenceThreshold *Float64 `json:"feature_influence_threshold,omitempty"`
	// Method The method that outlier detection uses. Available methods are `lof`, `ldof`,
	// `distance_kth_nn`, `distance_knn`, and `ensemble`. The default value is
	// ensemble, which means that outlier detection uses an ensemble of different
	// methods and normalises and combines their individual outlier scores to obtain
	// the overall outlier score.
	Method *string `json:"method,omitempty"`
	// NNeighbors Defines the value for how many nearest neighbors each method of outlier
	// detection uses to calculate its outlier score. When the value is not set,
	// different values are used for different ensemble members. This default
	// behavior helps improve the diversity in the ensemble; only override it if you
	// are confident that the value you choose is appropriate for the data set.
	NNeighbors *int `json:"n_neighbors,omitempty"`
	// OutlierFraction The proportion of the data set that is assumed to be outlying prior to
	// outlier detection. For example, 0.05 means it is assumed that 5% of values
	// are real outliers and 95% are inliers.
	OutlierFraction *Float64 `json:"outlier_fraction,omitempty"`
	// StandardizationEnabled If true, the following operation is performed on the columns before computing
	// outlier scores: `(x_i - mean(x_i)) / sd(x_i)`.
	StandardizationEnabled *bool `json:"standardization_enabled,omitempty"`
}

func (s *DataframeAnalysisOutlierDetection) UnmarshalJSON(data []byte) error {

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

// NewDataframeAnalysisOutlierDetection returns a DataframeAnalysisOutlierDetection.
func NewDataframeAnalysisOutlierDetection() *DataframeAnalysisOutlierDetection {
	r := &DataframeAnalysisOutlierDetection{}

	return r
}
