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

// DataframeAnalyticsFieldSelection type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeAnalytics.ts#L55-L68
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
	Name Field `json:"name"`
	// Reason The reason a field is not selected to be included in the analysis.
	Reason *string `json:"reason,omitempty"`
}

// DataframeAnalyticsFieldSelectionBuilder holds DataframeAnalyticsFieldSelection struct and provides a builder API.
type DataframeAnalyticsFieldSelectionBuilder struct {
	v *DataframeAnalyticsFieldSelection
}

// NewDataframeAnalyticsFieldSelection provides a builder for the DataframeAnalyticsFieldSelection struct.
func NewDataframeAnalyticsFieldSelectionBuilder() *DataframeAnalyticsFieldSelectionBuilder {
	r := DataframeAnalyticsFieldSelectionBuilder{
		&DataframeAnalyticsFieldSelection{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalyticsFieldSelection struct
func (rb *DataframeAnalyticsFieldSelectionBuilder) Build() DataframeAnalyticsFieldSelection {
	return *rb.v
}

// FeatureType The feature type of this field for the analysis. May be categorical or
// numerical.

func (rb *DataframeAnalyticsFieldSelectionBuilder) FeatureType(featuretype string) *DataframeAnalyticsFieldSelectionBuilder {
	rb.v.FeatureType = &featuretype
	return rb
}

// IsIncluded Whether the field is selected to be included in the analysis.

func (rb *DataframeAnalyticsFieldSelectionBuilder) IsIncluded(isincluded bool) *DataframeAnalyticsFieldSelectionBuilder {
	rb.v.IsIncluded = isincluded
	return rb
}

// IsRequired Whether the field is required.

func (rb *DataframeAnalyticsFieldSelectionBuilder) IsRequired(isrequired bool) *DataframeAnalyticsFieldSelectionBuilder {
	rb.v.IsRequired = isrequired
	return rb
}

// MappingTypes The mapping types of the field.

func (rb *DataframeAnalyticsFieldSelectionBuilder) MappingTypes(mapping_types ...string) *DataframeAnalyticsFieldSelectionBuilder {
	rb.v.MappingTypes = mapping_types
	return rb
}

// Name The field name.

func (rb *DataframeAnalyticsFieldSelectionBuilder) Name(name Field) *DataframeAnalyticsFieldSelectionBuilder {
	rb.v.Name = name
	return rb
}

// Reason The reason a field is not selected to be included in the analysis.

func (rb *DataframeAnalyticsFieldSelectionBuilder) Reason(reason string) *DataframeAnalyticsFieldSelectionBuilder {
	rb.v.Reason = &reason
	return rb
}
