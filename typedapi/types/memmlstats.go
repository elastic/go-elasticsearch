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

// MemMlStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/get_memory_stats/types.ts#L90-L111
type MemMlStats struct {
	// AnomalyDetectors Amount of native memory set aside for anomaly detection jobs.
	AnomalyDetectors *ByteSize `json:"anomaly_detectors,omitempty"`
	// AnomalyDetectorsInBytes Amount of native memory, in bytes, set aside for anomaly detection jobs.
	AnomalyDetectorsInBytes int `json:"anomaly_detectors_in_bytes"`
	// DataFrameAnalytics Amount of native memory set aside for data frame analytics jobs.
	DataFrameAnalytics *ByteSize `json:"data_frame_analytics,omitempty"`
	// DataFrameAnalyticsInBytes Amount of native memory, in bytes, set aside for data frame analytics jobs.
	DataFrameAnalyticsInBytes int `json:"data_frame_analytics_in_bytes"`
	// Max Maximum amount of native memory (separate to the JVM heap) that may be used
	// by machine learning native processes.
	Max *ByteSize `json:"max,omitempty"`
	// MaxInBytes Maximum amount of native memory (separate to the JVM heap), in bytes, that
	// may be used by machine learning native processes.
	MaxInBytes int `json:"max_in_bytes"`
	// NativeCodeOverhead Amount of native memory set aside for loading machine learning native code
	// shared libraries.
	NativeCodeOverhead *ByteSize `json:"native_code_overhead,omitempty"`
	// NativeCodeOverheadInBytes Amount of native memory, in bytes, set aside for loading machine learning
	// native code shared libraries.
	NativeCodeOverheadInBytes int `json:"native_code_overhead_in_bytes"`
	// NativeInference Amount of native memory set aside for trained models that have a PyTorch
	// model_type.
	NativeInference *ByteSize `json:"native_inference,omitempty"`
	// NativeInferenceInBytes Amount of native memory, in bytes, set aside for trained models that have a
	// PyTorch model_type.
	NativeInferenceInBytes int `json:"native_inference_in_bytes"`
}

// NewMemMlStats returns a MemMlStats.
func NewMemMlStats() *MemMlStats {
	r := &MemMlStats{}

	return r
}
