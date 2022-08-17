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

// DataframeAnalysisFeatureProcessorTargetMeanEncoding type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeAnalytics.ts#L295-L304
type DataframeAnalysisFeatureProcessorTargetMeanEncoding struct {
	// DefaultValue The default value if field value is not found in the target_map.
	DefaultValue int `json:"default_value"`
	// FeatureName The resulting feature name.
	FeatureName Name `json:"feature_name"`
	// Field The name of the field to encode.
	Field Field `json:"field"`
	// TargetMap The field value to target mean transition map.
	TargetMap map[string]interface{} `json:"target_map"`
}

// DataframeAnalysisFeatureProcessorTargetMeanEncodingBuilder holds DataframeAnalysisFeatureProcessorTargetMeanEncoding struct and provides a builder API.
type DataframeAnalysisFeatureProcessorTargetMeanEncodingBuilder struct {
	v *DataframeAnalysisFeatureProcessorTargetMeanEncoding
}

// NewDataframeAnalysisFeatureProcessorTargetMeanEncoding provides a builder for the DataframeAnalysisFeatureProcessorTargetMeanEncoding struct.
func NewDataframeAnalysisFeatureProcessorTargetMeanEncodingBuilder() *DataframeAnalysisFeatureProcessorTargetMeanEncodingBuilder {
	r := DataframeAnalysisFeatureProcessorTargetMeanEncodingBuilder{
		&DataframeAnalysisFeatureProcessorTargetMeanEncoding{
			TargetMap: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalysisFeatureProcessorTargetMeanEncoding struct
func (rb *DataframeAnalysisFeatureProcessorTargetMeanEncodingBuilder) Build() DataframeAnalysisFeatureProcessorTargetMeanEncoding {
	return *rb.v
}

// DefaultValue The default value if field value is not found in the target_map.

func (rb *DataframeAnalysisFeatureProcessorTargetMeanEncodingBuilder) DefaultValue(defaultvalue int) *DataframeAnalysisFeatureProcessorTargetMeanEncodingBuilder {
	rb.v.DefaultValue = defaultvalue
	return rb
}

// FeatureName The resulting feature name.

func (rb *DataframeAnalysisFeatureProcessorTargetMeanEncodingBuilder) FeatureName(featurename Name) *DataframeAnalysisFeatureProcessorTargetMeanEncodingBuilder {
	rb.v.FeatureName = featurename
	return rb
}

// Field The name of the field to encode.

func (rb *DataframeAnalysisFeatureProcessorTargetMeanEncodingBuilder) Field(field Field) *DataframeAnalysisFeatureProcessorTargetMeanEncodingBuilder {
	rb.v.Field = field
	return rb
}

// TargetMap The field value to target mean transition map.

func (rb *DataframeAnalysisFeatureProcessorTargetMeanEncodingBuilder) TargetMap(value map[string]interface{}) *DataframeAnalysisFeatureProcessorTargetMeanEncodingBuilder {
	rb.v.TargetMap = value
	return rb
}
