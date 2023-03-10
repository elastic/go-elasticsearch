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

// TrainedModelInferenceStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/TrainedModel.ts#L99-L119
type TrainedModelInferenceStats struct {
	// CacheMissCount The number of times the model was loaded for inference and was not retrieved
	// from the cache.
	// If this number is close to the `inference_count`, the cache is not being
	// appropriately used.
	// This can be solved by increasing the cache size or its time-to-live (TTL).
	// Refer to general machine learning settings for the appropriate settings.
	CacheMissCount int `json:"cache_miss_count"`
	// FailureCount The number of failures when using the model for inference.
	FailureCount int `json:"failure_count"`
	// InferenceCount The total number of times the model has been called for inference.
	// This is across all inference contexts, including all pipelines.
	InferenceCount int `json:"inference_count"`
	// MissingAllFieldsCount The number of inference calls where all the training features for the model
	// were missing.
	MissingAllFieldsCount int `json:"missing_all_fields_count"`
	// Timestamp The time when the statistics were last updated.
	Timestamp DateTime `json:"timestamp"`
}

// NewTrainedModelInferenceStats returns a TrainedModelInferenceStats.
func NewTrainedModelInferenceStats() *TrainedModelInferenceStats {
	r := &TrainedModelInferenceStats{}

	return r
}
