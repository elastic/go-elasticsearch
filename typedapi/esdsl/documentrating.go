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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _documentRating struct {
	v *types.DocumentRating
}

func NewDocumentRating(rating int) *_documentRating {

	tmp := &_documentRating{v: types.NewDocumentRating()}

	tmp.Rating(rating)

	return tmp

}

// The document ID.
func (s *_documentRating) Id_(id string) *_documentRating {

	s.v.Id_ = id

	return s
}

// The document’s index. For data streams, this should be the document’s backing
// index.
func (s *_documentRating) Index_(indexname string) *_documentRating {

	s.v.Index_ = indexname

	return s
}

// The document’s relevance with regard to this search request.
func (s *_documentRating) Rating(rating int) *_documentRating {

	s.v.Rating = rating

	return s
}

func (s *_documentRating) DocumentRatingCaster() *types.DocumentRating {
	return s.v
}
