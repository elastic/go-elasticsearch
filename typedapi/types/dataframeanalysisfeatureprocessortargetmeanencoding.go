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

import (
	"encoding/json"
)

// DataframeAnalysisFeatureProcessorTargetMeanEncoding type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/DataframeAnalytics.ts#L295-L304
type DataframeAnalysisFeatureProcessorTargetMeanEncoding struct {
	// DefaultValue The default value if field value is not found in the target_map.
	DefaultValue int `json:"default_value"`
	// FeatureName The resulting feature name.
	FeatureName string `json:"feature_name"`
	// Field The name of the field to encode.
	Field string `json:"field"`
	// TargetMap The field value to target mean transition map.
	TargetMap map[string]json.RawMessage `json:"target_map"`
}

// NewDataframeAnalysisFeatureProcessorTargetMeanEncoding returns a DataframeAnalysisFeatureProcessorTargetMeanEncoding.
func NewDataframeAnalysisFeatureProcessorTargetMeanEncoding() *DataframeAnalysisFeatureProcessorTargetMeanEncoding {
	r := &DataframeAnalysisFeatureProcessorTargetMeanEncoding{
		TargetMap: make(map[string]json.RawMessage, 0),
	}

	return r
}
