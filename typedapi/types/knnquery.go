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

// KnnQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Knn.ts#L24-L37
type KnnQuery struct {
	// Boost Boost value to apply to kNN scores
	Boost *float32 `json:"boost,omitempty"`
	// Field The name of the vector field to search against
	Field Field `json:"field"`
	// Filter Filters for the kNN search query
	Filter []QueryContainer `json:"filter,omitempty"`
	// K The final number of nearest neighbors to return as top hits
	K int64 `json:"k"`
	// NumCandidates The number of nearest neighbor candidates to consider per shard
	NumCandidates int64 `json:"num_candidates"`
	// QueryVector The query vector
	QueryVector []float64 `json:"query_vector"`
}

// KnnQueryBuilder holds KnnQuery struct and provides a builder API.
type KnnQueryBuilder struct {
	v *KnnQuery
}

// NewKnnQuery provides a builder for the KnnQuery struct.
func NewKnnQueryBuilder() *KnnQueryBuilder {
	r := KnnQueryBuilder{
		&KnnQuery{},
	}

	return &r
}

// Build finalize the chain and returns the KnnQuery struct
func (rb *KnnQueryBuilder) Build() KnnQuery {
	return *rb.v
}

// Boost Boost value to apply to kNN scores

func (rb *KnnQueryBuilder) Boost(boost float32) *KnnQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// Field The name of the vector field to search against

func (rb *KnnQueryBuilder) Field(field Field) *KnnQueryBuilder {
	rb.v.Field = field
	return rb
}

// Filter Filters for the kNN search query
func (rb *KnnQueryBuilder) Filter(arg []QueryContainer) *KnnQueryBuilder {
	rb.v.Filter = arg
	return rb
}

// K The final number of nearest neighbors to return as top hits

func (rb *KnnQueryBuilder) K(k int64) *KnnQueryBuilder {
	rb.v.K = k
	return rb
}

// NumCandidates The number of nearest neighbor candidates to consider per shard

func (rb *KnnQueryBuilder) NumCandidates(numcandidates int64) *KnnQueryBuilder {
	rb.v.NumCandidates = numcandidates
	return rb
}

// QueryVector The query vector

func (rb *KnnQueryBuilder) QueryVector(query_vector ...float64) *KnnQueryBuilder {
	rb.v.QueryVector = query_vector
	return rb
}
