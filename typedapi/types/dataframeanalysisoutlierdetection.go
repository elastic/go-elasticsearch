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

// DataframeAnalysisOutlierDetection type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeAnalytics.ts#L103-L132
type DataframeAnalysisOutlierDetection struct {
	// ComputeFeatureInfluence Specifies whether the feature influence calculation is enabled.
	ComputeFeatureInfluence *bool `json:"compute_feature_influence,omitempty"`
	// FeatureInfluenceThreshold The minimum outlier score that a document needs to have in order to calculate
	// its feature influence score. Value range: 0-1.
	FeatureInfluenceThreshold *float64 `json:"feature_influence_threshold,omitempty"`
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
	OutlierFraction *float64 `json:"outlier_fraction,omitempty"`
	// StandardizationEnabled If true, the following operation is performed on the columns before computing
	// outlier scores: `(x_i - mean(x_i)) / sd(x_i)`.
	StandardizationEnabled *bool `json:"standardization_enabled,omitempty"`
}

// DataframeAnalysisOutlierDetectionBuilder holds DataframeAnalysisOutlierDetection struct and provides a builder API.
type DataframeAnalysisOutlierDetectionBuilder struct {
	v *DataframeAnalysisOutlierDetection
}

// NewDataframeAnalysisOutlierDetection provides a builder for the DataframeAnalysisOutlierDetection struct.
func NewDataframeAnalysisOutlierDetectionBuilder() *DataframeAnalysisOutlierDetectionBuilder {
	r := DataframeAnalysisOutlierDetectionBuilder{
		&DataframeAnalysisOutlierDetection{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalysisOutlierDetection struct
func (rb *DataframeAnalysisOutlierDetectionBuilder) Build() DataframeAnalysisOutlierDetection {
	return *rb.v
}

// ComputeFeatureInfluence Specifies whether the feature influence calculation is enabled.

func (rb *DataframeAnalysisOutlierDetectionBuilder) ComputeFeatureInfluence(computefeatureinfluence bool) *DataframeAnalysisOutlierDetectionBuilder {
	rb.v.ComputeFeatureInfluence = &computefeatureinfluence
	return rb
}

// FeatureInfluenceThreshold The minimum outlier score that a document needs to have in order to calculate
// its feature influence score. Value range: 0-1.

func (rb *DataframeAnalysisOutlierDetectionBuilder) FeatureInfluenceThreshold(featureinfluencethreshold float64) *DataframeAnalysisOutlierDetectionBuilder {
	rb.v.FeatureInfluenceThreshold = &featureinfluencethreshold
	return rb
}

// Method The method that outlier detection uses. Available methods are `lof`, `ldof`,
// `distance_kth_nn`, `distance_knn`, and `ensemble`. The default value is
// ensemble, which means that outlier detection uses an ensemble of different
// methods and normalises and combines their individual outlier scores to obtain
// the overall outlier score.

func (rb *DataframeAnalysisOutlierDetectionBuilder) Method(method string) *DataframeAnalysisOutlierDetectionBuilder {
	rb.v.Method = &method
	return rb
}

// NNeighbors Defines the value for how many nearest neighbors each method of outlier
// detection uses to calculate its outlier score. When the value is not set,
// different values are used for different ensemble members. This default
// behavior helps improve the diversity in the ensemble; only override it if you
// are confident that the value you choose is appropriate for the data set.

func (rb *DataframeAnalysisOutlierDetectionBuilder) NNeighbors(nneighbors int) *DataframeAnalysisOutlierDetectionBuilder {
	rb.v.NNeighbors = &nneighbors
	return rb
}

// OutlierFraction The proportion of the data set that is assumed to be outlying prior to
// outlier detection. For example, 0.05 means it is assumed that 5% of values
// are real outliers and 95% are inliers.

func (rb *DataframeAnalysisOutlierDetectionBuilder) OutlierFraction(outlierfraction float64) *DataframeAnalysisOutlierDetectionBuilder {
	rb.v.OutlierFraction = &outlierfraction
	return rb
}

// StandardizationEnabled If true, the following operation is performed on the columns before computing
// outlier scores: `(x_i - mean(x_i)) / sd(x_i)`.

func (rb *DataframeAnalysisOutlierDetectionBuilder) StandardizationEnabled(standardizationenabled bool) *DataframeAnalysisOutlierDetectionBuilder {
	rb.v.StandardizationEnabled = &standardizationenabled
	return rb
}
