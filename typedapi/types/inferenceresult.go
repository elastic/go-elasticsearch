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
// https://github.com/elastic/elasticsearch-specification/tree/07bf82537a186562d8699685e3704ea338b268ef

package types

// InferenceResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/07bf82537a186562d8699685e3704ea338b268ef/specification/inference/_types/Results.ts#L77-L87
type InferenceResult struct {
	Completion         []CompletionResult        `json:"completion,omitempty"`
	Rerank             []RankedDocument          `json:"rerank,omitempty"`
	SparseEmbedding    []SparseEmbeddingResult   `json:"sparse_embedding,omitempty"`
	TextEmbedding      []TextEmbeddingResult     `json:"text_embedding,omitempty"`
	TextEmbeddingBytes []TextEmbeddingByteResult `json:"text_embedding_bytes,omitempty"`
}

// NewInferenceResult returns a InferenceResult.
func NewInferenceResult() *InferenceResult {
	r := &InferenceResult{}

	return r
}
