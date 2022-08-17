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

// MemMlStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/get_memory_stats/types.ts#L90-L111
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

// MemMlStatsBuilder holds MemMlStats struct and provides a builder API.
type MemMlStatsBuilder struct {
	v *MemMlStats
}

// NewMemMlStats provides a builder for the MemMlStats struct.
func NewMemMlStatsBuilder() *MemMlStatsBuilder {
	r := MemMlStatsBuilder{
		&MemMlStats{},
	}

	return &r
}

// Build finalize the chain and returns the MemMlStats struct
func (rb *MemMlStatsBuilder) Build() MemMlStats {
	return *rb.v
}

// AnomalyDetectors Amount of native memory set aside for anomaly detection jobs.

func (rb *MemMlStatsBuilder) AnomalyDetectors(anomalydetectors *ByteSizeBuilder) *MemMlStatsBuilder {
	v := anomalydetectors.Build()
	rb.v.AnomalyDetectors = &v
	return rb
}

// AnomalyDetectorsInBytes Amount of native memory, in bytes, set aside for anomaly detection jobs.

func (rb *MemMlStatsBuilder) AnomalyDetectorsInBytes(anomalydetectorsinbytes int) *MemMlStatsBuilder {
	rb.v.AnomalyDetectorsInBytes = anomalydetectorsinbytes
	return rb
}

// DataFrameAnalytics Amount of native memory set aside for data frame analytics jobs.

func (rb *MemMlStatsBuilder) DataFrameAnalytics(dataframeanalytics *ByteSizeBuilder) *MemMlStatsBuilder {
	v := dataframeanalytics.Build()
	rb.v.DataFrameAnalytics = &v
	return rb
}

// DataFrameAnalyticsInBytes Amount of native memory, in bytes, set aside for data frame analytics jobs.

func (rb *MemMlStatsBuilder) DataFrameAnalyticsInBytes(dataframeanalyticsinbytes int) *MemMlStatsBuilder {
	rb.v.DataFrameAnalyticsInBytes = dataframeanalyticsinbytes
	return rb
}

// Max Maximum amount of native memory (separate to the JVM heap) that may be used
// by machine learning native processes.

func (rb *MemMlStatsBuilder) Max(max *ByteSizeBuilder) *MemMlStatsBuilder {
	v := max.Build()
	rb.v.Max = &v
	return rb
}

// MaxInBytes Maximum amount of native memory (separate to the JVM heap), in bytes, that
// may be used by machine learning native processes.

func (rb *MemMlStatsBuilder) MaxInBytes(maxinbytes int) *MemMlStatsBuilder {
	rb.v.MaxInBytes = maxinbytes
	return rb
}

// NativeCodeOverhead Amount of native memory set aside for loading machine learning native code
// shared libraries.

func (rb *MemMlStatsBuilder) NativeCodeOverhead(nativecodeoverhead *ByteSizeBuilder) *MemMlStatsBuilder {
	v := nativecodeoverhead.Build()
	rb.v.NativeCodeOverhead = &v
	return rb
}

// NativeCodeOverheadInBytes Amount of native memory, in bytes, set aside for loading machine learning
// native code shared libraries.

func (rb *MemMlStatsBuilder) NativeCodeOverheadInBytes(nativecodeoverheadinbytes int) *MemMlStatsBuilder {
	rb.v.NativeCodeOverheadInBytes = nativecodeoverheadinbytes
	return rb
}

// NativeInference Amount of native memory set aside for trained models that have a PyTorch
// model_type.

func (rb *MemMlStatsBuilder) NativeInference(nativeinference *ByteSizeBuilder) *MemMlStatsBuilder {
	v := nativeinference.Build()
	rb.v.NativeInference = &v
	return rb
}

// NativeInferenceInBytes Amount of native memory, in bytes, set aside for trained models that have a
// PyTorch model_type.

func (rb *MemMlStatsBuilder) NativeInferenceInBytes(nativeinferenceinbytes int) *MemMlStatsBuilder {
	rb.v.NativeInferenceInBytes = nativeinferenceinbytes
	return rb
}
