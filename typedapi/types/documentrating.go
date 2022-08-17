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

// DocumentRating type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/rank_eval/types.ts#L116-L123
type DocumentRating struct {
	// Id_ The document ID.
	Id_ Id `json:"_id"`
	// Index_ The document’s index. For data streams, this should be the document’s backing
	// index.
	Index_ IndexName `json:"_index"`
	// Rating The document’s relevance with regard to this search request.
	Rating int `json:"rating"`
}

// DocumentRatingBuilder holds DocumentRating struct and provides a builder API.
type DocumentRatingBuilder struct {
	v *DocumentRating
}

// NewDocumentRating provides a builder for the DocumentRating struct.
func NewDocumentRatingBuilder() *DocumentRatingBuilder {
	r := DocumentRatingBuilder{
		&DocumentRating{},
	}

	return &r
}

// Build finalize the chain and returns the DocumentRating struct
func (rb *DocumentRatingBuilder) Build() DocumentRating {
	return *rb.v
}

// Id_ The document ID.

func (rb *DocumentRatingBuilder) Id_(id_ Id) *DocumentRatingBuilder {
	rb.v.Id_ = id_
	return rb
}

// Index_ The document’s index. For data streams, this should be the document’s backing
// index.

func (rb *DocumentRatingBuilder) Index_(index_ IndexName) *DocumentRatingBuilder {
	rb.v.Index_ = index_
	return rb
}

// Rating The document’s relevance with regard to this search request.

func (rb *DocumentRatingBuilder) Rating(rating int) *DocumentRatingBuilder {
	rb.v.Rating = rating
	return rb
}
