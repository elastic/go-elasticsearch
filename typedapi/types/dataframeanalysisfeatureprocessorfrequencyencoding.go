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

// DataframeAnalysisFeatureProcessorFrequencyEncoding type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeAnalytics.ts#L260-L267
type DataframeAnalysisFeatureProcessorFrequencyEncoding struct {
	// FeatureName The resulting feature name.
	FeatureName Name  `json:"feature_name"`
	Field       Field `json:"field"`
	// FrequencyMap The resulting frequency map for the field value. If the field value is
	// missing from the frequency_map, the resulting value is 0.
	FrequencyMap map[string]float64 `json:"frequency_map"`
}

// DataframeAnalysisFeatureProcessorFrequencyEncodingBuilder holds DataframeAnalysisFeatureProcessorFrequencyEncoding struct and provides a builder API.
type DataframeAnalysisFeatureProcessorFrequencyEncodingBuilder struct {
	v *DataframeAnalysisFeatureProcessorFrequencyEncoding
}

// NewDataframeAnalysisFeatureProcessorFrequencyEncoding provides a builder for the DataframeAnalysisFeatureProcessorFrequencyEncoding struct.
func NewDataframeAnalysisFeatureProcessorFrequencyEncodingBuilder() *DataframeAnalysisFeatureProcessorFrequencyEncodingBuilder {
	r := DataframeAnalysisFeatureProcessorFrequencyEncodingBuilder{
		&DataframeAnalysisFeatureProcessorFrequencyEncoding{
			FrequencyMap: make(map[string]float64, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalysisFeatureProcessorFrequencyEncoding struct
func (rb *DataframeAnalysisFeatureProcessorFrequencyEncodingBuilder) Build() DataframeAnalysisFeatureProcessorFrequencyEncoding {
	return *rb.v
}

// FeatureName The resulting feature name.

func (rb *DataframeAnalysisFeatureProcessorFrequencyEncodingBuilder) FeatureName(featurename Name) *DataframeAnalysisFeatureProcessorFrequencyEncodingBuilder {
	rb.v.FeatureName = featurename
	return rb
}

func (rb *DataframeAnalysisFeatureProcessorFrequencyEncodingBuilder) Field(field Field) *DataframeAnalysisFeatureProcessorFrequencyEncodingBuilder {
	rb.v.Field = field
	return rb
}

// FrequencyMap The resulting frequency map for the field value. If the field value is
// missing from the frequency_map, the resulting value is 0.

func (rb *DataframeAnalysisFeatureProcessorFrequencyEncodingBuilder) FrequencyMap(value map[string]float64) *DataframeAnalysisFeatureProcessorFrequencyEncodingBuilder {
	rb.v.FrequencyMap = value
	return rb
}
