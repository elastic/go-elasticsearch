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
// https://github.com/elastic/elasticsearch-specification/tree/836fca874204ca4173ae5c36fb6b5107d28d2fc0

package types

// QueryVectorBuilder type.
//
// https://github.com/elastic/elasticsearch-specification/blob/836fca874204ca4173ae5c36fb6b5107d28d2fc0/specification/_types/Knn.ts#L106-L116
type QueryVectorBuilder struct {
	// Lookup Lookup a vector from an existing document. Must reference a dense_vector
	// field and a single value.
	Lookup        *LookupQueryVectorBuilder `json:"lookup,omitempty"`
	TextEmbedding *TextEmbedding            `json:"text_embedding,omitempty"`
}

// NewQueryVectorBuilder returns a QueryVectorBuilder.
func NewQueryVectorBuilder() *QueryVectorBuilder {
	r := &QueryVectorBuilder{}

	return r
}

type QueryVectorBuilderVariant interface {
	QueryVectorBuilderCaster() *QueryVectorBuilder
}

func (s *QueryVectorBuilder) QueryVectorBuilderCaster() *QueryVectorBuilder {
	return s
}
