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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// OutlierDetectionParameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeAnalytics.ts#L412-L419
type OutlierDetectionParameters struct {
	ComputeFeatureInfluence   *bool    `json:"compute_feature_influence,omitempty"`
	FeatureInfluenceThreshold *float64 `json:"feature_influence_threshold,omitempty"`
	Method                    *string  `json:"method,omitempty"`
	NNeighbors                *int     `json:"n_neighbors,omitempty"`
	OutlierFraction           *float64 `json:"outlier_fraction,omitempty"`
	StandardizationEnabled    *bool    `json:"standardization_enabled,omitempty"`
}

// OutlierDetectionParametersBuilder holds OutlierDetectionParameters struct and provides a builder API.
type OutlierDetectionParametersBuilder struct {
	v *OutlierDetectionParameters
}

// NewOutlierDetectionParameters provides a builder for the OutlierDetectionParameters struct.
func NewOutlierDetectionParametersBuilder() *OutlierDetectionParametersBuilder {
	r := OutlierDetectionParametersBuilder{
		&OutlierDetectionParameters{},
	}

	return &r
}

// Build finalize the chain and returns the OutlierDetectionParameters struct
func (rb *OutlierDetectionParametersBuilder) Build() OutlierDetectionParameters {
	return *rb.v
}

func (rb *OutlierDetectionParametersBuilder) ComputeFeatureInfluence(computefeatureinfluence bool) *OutlierDetectionParametersBuilder {
	rb.v.ComputeFeatureInfluence = &computefeatureinfluence
	return rb
}

func (rb *OutlierDetectionParametersBuilder) FeatureInfluenceThreshold(featureinfluencethreshold float64) *OutlierDetectionParametersBuilder {
	rb.v.FeatureInfluenceThreshold = &featureinfluencethreshold
	return rb
}

func (rb *OutlierDetectionParametersBuilder) Method(method string) *OutlierDetectionParametersBuilder {
	rb.v.Method = &method
	return rb
}

func (rb *OutlierDetectionParametersBuilder) NNeighbors(nneighbors int) *OutlierDetectionParametersBuilder {
	rb.v.NNeighbors = &nneighbors
	return rb
}

func (rb *OutlierDetectionParametersBuilder) OutlierFraction(outlierfraction float64) *OutlierDetectionParametersBuilder {
	rb.v.OutlierFraction = &outlierfraction
	return rb
}

func (rb *OutlierDetectionParametersBuilder) StandardizationEnabled(standardizationenabled bool) *OutlierDetectionParametersBuilder {
	rb.v.StandardizationEnabled = &standardizationenabled
	return rb
}
