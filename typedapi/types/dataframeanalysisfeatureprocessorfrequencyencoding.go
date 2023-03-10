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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// DataframeAnalysisFeatureProcessorFrequencyEncoding type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/DataframeAnalytics.ts#L260-L267
type DataframeAnalysisFeatureProcessorFrequencyEncoding struct {
	// FeatureName The resulting feature name.
	FeatureName string `json:"feature_name"`
	Field       string `json:"field"`
	// FrequencyMap The resulting frequency map for the field value. If the field value is
	// missing from the frequency_map, the resulting value is 0.
	FrequencyMap map[string]Float64 `json:"frequency_map"`
}

// NewDataframeAnalysisFeatureProcessorFrequencyEncoding returns a DataframeAnalysisFeatureProcessorFrequencyEncoding.
func NewDataframeAnalysisFeatureProcessorFrequencyEncoding() *DataframeAnalysisFeatureProcessorFrequencyEncoding {
	r := &DataframeAnalysisFeatureProcessorFrequencyEncoding{
		FrequencyMap: make(map[string]Float64, 0),
	}

	return r
}
