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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/totalhitsrelation"
)

// TotalHits type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/hits.ts#L90-L93
type TotalHits struct {
	Relation totalhitsrelation.TotalHitsRelation `json:"relation"`
	Value    int64                               `json:"value"`
}

// TotalHitsBuilder holds TotalHits struct and provides a builder API.
type TotalHitsBuilder struct {
	v *TotalHits
}

// NewTotalHits provides a builder for the TotalHits struct.
func NewTotalHitsBuilder() *TotalHitsBuilder {
	r := TotalHitsBuilder{
		&TotalHits{},
	}

	return &r
}

// Build finalize the chain and returns the TotalHits struct
func (rb *TotalHitsBuilder) Build() TotalHits {
	return *rb.v
}

func (rb *TotalHitsBuilder) Relation(relation totalhitsrelation.TotalHitsRelation) *TotalHitsBuilder {
	rb.v.Relation = relation
	return rb
}

func (rb *TotalHitsBuilder) Value(value int64) *TotalHitsBuilder {
	rb.v.Value = value
	return rb
}
