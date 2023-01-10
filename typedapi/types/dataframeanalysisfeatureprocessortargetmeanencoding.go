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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// DataframeAnalysisFeatureProcessorTargetMeanEncoding type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/DataframeAnalytics.ts#L295-L304
type DataframeAnalysisFeatureProcessorTargetMeanEncoding struct {
	// DefaultValue The default value if field value is not found in the target_map.
	DefaultValue int `json:"default_value"`
	// FeatureName The resulting feature name.
	FeatureName string `json:"feature_name"`
	// Field The name of the field to encode.
	Field string `json:"field"`
	// TargetMap The field value to target mean transition map.
	TargetMap map[string]interface{} `json:"target_map"`
}

// NewDataframeAnalysisFeatureProcessorTargetMeanEncoding returns a DataframeAnalysisFeatureProcessorTargetMeanEncoding.
func NewDataframeAnalysisFeatureProcessorTargetMeanEncoding() *DataframeAnalysisFeatureProcessorTargetMeanEncoding {
	r := &DataframeAnalysisFeatureProcessorTargetMeanEncoding{
		TargetMap: make(map[string]interface{}, 0),
	}

	return r
}
