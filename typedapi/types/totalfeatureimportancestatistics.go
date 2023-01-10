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

// TotalFeatureImportanceStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/TrainedModel.ts#L237-L244
type TotalFeatureImportanceStatistics struct {
	// Max The maximum importance value across all the training data for this feature.
	Max int `json:"max"`
	// MeanMagnitude The average magnitude of this feature across all the training data. This
	// value is the average of the absolute values of the importance for this
	// feature.
	MeanMagnitude float64 `json:"mean_magnitude"`
	// Min The minimum importance value across all the training data for this feature.
	Min int `json:"min"`
}

// NewTotalFeatureImportanceStatistics returns a TotalFeatureImportanceStatistics.
func NewTotalFeatureImportanceStatistics() *TotalFeatureImportanceStatistics {
	r := &TotalFeatureImportanceStatistics{}

	return r
}
