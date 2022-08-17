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

// TrainedModelInferenceStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/TrainedModel.ts#L98-L118
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

// TrainedModelInferenceStatsBuilder holds TrainedModelInferenceStats struct and provides a builder API.
type TrainedModelInferenceStatsBuilder struct {
	v *TrainedModelInferenceStats
}

// NewTrainedModelInferenceStats provides a builder for the TrainedModelInferenceStats struct.
func NewTrainedModelInferenceStatsBuilder() *TrainedModelInferenceStatsBuilder {
	r := TrainedModelInferenceStatsBuilder{
		&TrainedModelInferenceStats{},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelInferenceStats struct
func (rb *TrainedModelInferenceStatsBuilder) Build() TrainedModelInferenceStats {
	return *rb.v
}

// CacheMissCount The number of times the model was loaded for inference and was not retrieved
// from the cache.
// If this number is close to the `inference_count`, the cache is not being
// appropriately used.
// This can be solved by increasing the cache size or its time-to-live (TTL).
// Refer to general machine learning settings for the appropriate settings.

func (rb *TrainedModelInferenceStatsBuilder) CacheMissCount(cachemisscount int) *TrainedModelInferenceStatsBuilder {
	rb.v.CacheMissCount = cachemisscount
	return rb
}

// FailureCount The number of failures when using the model for inference.

func (rb *TrainedModelInferenceStatsBuilder) FailureCount(failurecount int) *TrainedModelInferenceStatsBuilder {
	rb.v.FailureCount = failurecount
	return rb
}

// InferenceCount The total number of times the model has been called for inference.
// This is across all inference contexts, including all pipelines.

func (rb *TrainedModelInferenceStatsBuilder) InferenceCount(inferencecount int) *TrainedModelInferenceStatsBuilder {
	rb.v.InferenceCount = inferencecount
	return rb
}

// MissingAllFieldsCount The number of inference calls where all the training features for the model
// were missing.

func (rb *TrainedModelInferenceStatsBuilder) MissingAllFieldsCount(missingallfieldscount int) *TrainedModelInferenceStatsBuilder {
	rb.v.MissingAllFieldsCount = missingallfieldscount
	return rb
}

// Timestamp The time when the statistics were last updated.

func (rb *TrainedModelInferenceStatsBuilder) Timestamp(timestamp *DateTimeBuilder) *TrainedModelInferenceStatsBuilder {
	v := timestamp.Build()
	rb.v.Timestamp = v
	return rb
}
