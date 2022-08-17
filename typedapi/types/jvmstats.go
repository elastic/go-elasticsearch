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

// JvmStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/get_memory_stats/types.ts#L50-L63
type JvmStats struct {
	// HeapMax Maximum amount of memory available for use by the heap.
	HeapMax *ByteSize `json:"heap_max,omitempty"`
	// HeapMaxInBytes Maximum amount of memory, in bytes, available for use by the heap.
	HeapMaxInBytes int `json:"heap_max_in_bytes"`
	// JavaInference Amount of Java heap currently being used for caching inference models.
	JavaInference *ByteSize `json:"java_inference,omitempty"`
	// JavaInferenceInBytes Amount of Java heap, in bytes, currently being used for caching inference
	// models.
	JavaInferenceInBytes int `json:"java_inference_in_bytes"`
	// JavaInferenceMax Maximum amount of Java heap to be used for caching inference models.
	JavaInferenceMax *ByteSize `json:"java_inference_max,omitempty"`
	// JavaInferenceMaxInBytes Maximum amount of Java heap, in bytes, to be used for caching inference
	// models.
	JavaInferenceMaxInBytes int `json:"java_inference_max_in_bytes"`
}

// JvmStatsBuilder holds JvmStats struct and provides a builder API.
type JvmStatsBuilder struct {
	v *JvmStats
}

// NewJvmStats provides a builder for the JvmStats struct.
func NewJvmStatsBuilder() *JvmStatsBuilder {
	r := JvmStatsBuilder{
		&JvmStats{},
	}

	return &r
}

// Build finalize the chain and returns the JvmStats struct
func (rb *JvmStatsBuilder) Build() JvmStats {
	return *rb.v
}

// HeapMax Maximum amount of memory available for use by the heap.

func (rb *JvmStatsBuilder) HeapMax(heapmax *ByteSizeBuilder) *JvmStatsBuilder {
	v := heapmax.Build()
	rb.v.HeapMax = &v
	return rb
}

// HeapMaxInBytes Maximum amount of memory, in bytes, available for use by the heap.

func (rb *JvmStatsBuilder) HeapMaxInBytes(heapmaxinbytes int) *JvmStatsBuilder {
	rb.v.HeapMaxInBytes = heapmaxinbytes
	return rb
}

// JavaInference Amount of Java heap currently being used for caching inference models.

func (rb *JvmStatsBuilder) JavaInference(javainference *ByteSizeBuilder) *JvmStatsBuilder {
	v := javainference.Build()
	rb.v.JavaInference = &v
	return rb
}

// JavaInferenceInBytes Amount of Java heap, in bytes, currently being used for caching inference
// models.

func (rb *JvmStatsBuilder) JavaInferenceInBytes(javainferenceinbytes int) *JvmStatsBuilder {
	rb.v.JavaInferenceInBytes = javainferenceinbytes
	return rb
}

// JavaInferenceMax Maximum amount of Java heap to be used for caching inference models.

func (rb *JvmStatsBuilder) JavaInferenceMax(javainferencemax *ByteSizeBuilder) *JvmStatsBuilder {
	v := javainferencemax.Build()
	rb.v.JavaInferenceMax = &v
	return rb
}

// JavaInferenceMaxInBytes Maximum amount of Java heap, in bytes, to be used for caching inference
// models.

func (rb *JvmStatsBuilder) JavaInferenceMaxInBytes(javainferencemaxinbytes int) *JvmStatsBuilder {
	rb.v.JavaInferenceMaxInBytes = javainferencemaxinbytes
	return rb
}
