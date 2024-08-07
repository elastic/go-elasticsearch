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
// https://github.com/elastic/elasticsearch-specification/tree/19027dbdd366978ccae41842a040a636730e7c10

package inference

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package inference
//
// https://github.com/elastic/elasticsearch-specification/blob/19027dbdd366978ccae41842a040a636730e7c10/specification/inference/inference/InferenceResponse.ts#L22-L24
type Response struct {
	Completion         []types.CompletionResult        `json:"completion,omitempty"`
	Rerank             []types.RankedDocument          `json:"rerank,omitempty"`
	SparseEmbedding    []types.SparseEmbeddingResult   `json:"sparse_embedding,omitempty"`
	TextEmbedding      []types.TextEmbeddingResult     `json:"text_embedding,omitempty"`
	TextEmbeddingBytes []types.TextEmbeddingByteResult `json:"text_embedding_bytes,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
