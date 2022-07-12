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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// KnnQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/query_dsl/term.ts#L40-L47
type KnnQuery struct {
	Boost         *float32  `json:"boost,omitempty"`
	Field         Field     `json:"field"`
	NumCandidates int       `json:"num_candidates"`
	QueryName_    *string   `json:"_name,omitempty"`
	QueryVector   []float64 `json:"query_vector"`
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

func (rb *KnnQueryBuilder) Boost(boost float32) *KnnQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *KnnQueryBuilder) Field(field Field) *KnnQueryBuilder {
	rb.v.Field = field
	return rb
}

func (rb *KnnQueryBuilder) NumCandidates(numcandidates int) *KnnQueryBuilder {
	rb.v.NumCandidates = numcandidates
	return rb
}

func (rb *KnnQueryBuilder) QueryName_(queryname_ string) *KnnQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *KnnQueryBuilder) QueryVector(query_vector ...float64) *KnnQueryBuilder {
	rb.v.QueryVector = query_vector
	return rb
}
