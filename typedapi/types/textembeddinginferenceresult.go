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
// https://github.com/elastic/elasticsearch-specification/tree/3a94b6715915b1e9311724a2614c643368eece90

package types

// TextEmbeddingInferenceResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3a94b6715915b1e9311724a2614c643368eece90/specification/inference/_types/Results.ts#L67-L74
type TextEmbeddingInferenceResult struct {
	TextEmbedding      []TextEmbeddingResult     `json:"text_embedding,omitempty"`
	TextEmbeddingBytes []TextEmbeddingByteResult `json:"text_embedding_bytes,omitempty"`
}

// NewTextEmbeddingInferenceResult returns a TextEmbeddingInferenceResult.
func NewTextEmbeddingInferenceResult() *TextEmbeddingInferenceResult {
	r := &TextEmbeddingInferenceResult{}

	return r
}
