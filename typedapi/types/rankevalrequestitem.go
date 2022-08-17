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

// RankEvalRequestItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/rank_eval/types.ts#L98-L109
type RankEvalRequestItem struct {
	// Id The search request’s ID, used to group result details later.
	Id Id `json:"id"`
	// Params The search template parameters.
	Params map[string]interface{} `json:"params,omitempty"`
	// Ratings List of document ratings
	Ratings []DocumentRating `json:"ratings"`
	// Request The query being evaluated.
	Request *RankEvalQuery `json:"request,omitempty"`
	// TemplateId The search template Id
	TemplateId *Id `json:"template_id,omitempty"`
}

// RankEvalRequestItemBuilder holds RankEvalRequestItem struct and provides a builder API.
type RankEvalRequestItemBuilder struct {
	v *RankEvalRequestItem
}

// NewRankEvalRequestItem provides a builder for the RankEvalRequestItem struct.
func NewRankEvalRequestItemBuilder() *RankEvalRequestItemBuilder {
	r := RankEvalRequestItemBuilder{
		&RankEvalRequestItem{
			Params: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the RankEvalRequestItem struct
func (rb *RankEvalRequestItemBuilder) Build() RankEvalRequestItem {
	return *rb.v
}

// Id The search request’s ID, used to group result details later.

func (rb *RankEvalRequestItemBuilder) Id(id Id) *RankEvalRequestItemBuilder {
	rb.v.Id = id
	return rb
}

// Params The search template parameters.

func (rb *RankEvalRequestItemBuilder) Params(value map[string]interface{}) *RankEvalRequestItemBuilder {
	rb.v.Params = value
	return rb
}

// Ratings List of document ratings

func (rb *RankEvalRequestItemBuilder) Ratings(ratings []DocumentRatingBuilder) *RankEvalRequestItemBuilder {
	tmp := make([]DocumentRating, len(ratings))
	for _, value := range ratings {
		tmp = append(tmp, value.Build())
	}
	rb.v.Ratings = tmp
	return rb
}

// Request The query being evaluated.

func (rb *RankEvalRequestItemBuilder) Request(request *RankEvalQueryBuilder) *RankEvalRequestItemBuilder {
	v := request.Build()
	rb.v.Request = &v
	return rb
}

// TemplateId The search template Id

func (rb *RankEvalRequestItemBuilder) TemplateId(templateid Id) *RankEvalRequestItemBuilder {
	rb.v.TemplateId = &templateid
	return rb
}
