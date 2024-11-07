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
// https://github.com/elastic/elasticsearch-specification/tree/4fcf747dfafc951e1dcf3077327e3dcee9107db3

package types

// RetrieverContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4fcf747dfafc951e1dcf3077327e3dcee9107db3/specification/_types/Retriever.ts#L26-L38
type RetrieverContainer struct {
	// Knn A retriever that replaces the functionality  of a knn search.
	Knn *KnnRetriever `json:"knn,omitempty"`
	// Rrf A retriever that produces top documents from reciprocal rank fusion (RRF).
	Rrf *RRFRetriever `json:"rrf,omitempty"`
	// Standard A retriever that replaces the functionality of a traditional query.
	Standard *StandardRetriever `json:"standard,omitempty"`
	// TextSimilarityReranker A retriever that reranks the top documents based on a reranking model using
	// the InferenceAPI
	TextSimilarityReranker *TextSimilarityReranker `json:"text_similarity_reranker,omitempty"`
}

// NewRetrieverContainer returns a RetrieverContainer.
func NewRetrieverContainer() *RetrieverContainer {
	r := &RetrieverContainer{}

	return r
}
