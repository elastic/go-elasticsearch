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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// KnnQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/Knn.ts#L26-L41
type KnnQuery struct {
	// Boost Boost value to apply to kNN scores
	Boost *float32 `json:"boost,omitempty"`
	// Field The name of the vector field to search against
	Field string `json:"field"`
	// Filter Filters for the kNN search query
	Filter []Query `json:"filter,omitempty"`
	// K The final number of nearest neighbors to return as top hits
	K int64 `json:"k"`
	// NumCandidates The number of nearest neighbor candidates to consider per shard
	NumCandidates int64 `json:"num_candidates"`
	// QueryVector The query vector
	QueryVector []float32 `json:"query_vector,omitempty"`
	// QueryVectorBuilder The query vector builder. You must provide a query_vector_builder or
	// query_vector, but not both.
	QueryVectorBuilder *QueryVectorBuilder `json:"query_vector_builder,omitempty"`
}

// NewKnnQuery returns a KnnQuery.
func NewKnnQuery() *KnnQuery {
	r := &KnnQuery{}

	return r
}
