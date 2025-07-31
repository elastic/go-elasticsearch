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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package textembedding

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package textembedding
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/text_embedding/TextEmbeddingResponse.ts#L22-L24
type Response struct {
	AdditionalTextEmbeddingInferenceResultProperty map[string]json.RawMessage      `json:"-"`
	TextEmbedding                                  []types.TextEmbeddingResult     `json:"text_embedding,omitempty"`
	TextEmbeddingBits                              []types.TextEmbeddingByteResult `json:"text_embedding_bits,omitempty"`
	TextEmbeddingBytes                             []types.TextEmbeddingByteResult `json:"text_embedding_bytes,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		AdditionalTextEmbeddingInferenceResultProperty: make(map[string]json.RawMessage, 0),
	}
	return r
}
