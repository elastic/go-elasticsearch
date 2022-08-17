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

// FuzzyQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/term.ts#L40-L51
type FuzzyQuery struct {
	Boost          *float32               `json:"boost,omitempty"`
	Fuzziness      *Fuzziness             `json:"fuzziness,omitempty"`
	MaxExpansions  *int                   `json:"max_expansions,omitempty"`
	PrefixLength   *int                   `json:"prefix_length,omitempty"`
	QueryName_     *string                `json:"_name,omitempty"`
	Rewrite        *MultiTermQueryRewrite `json:"rewrite,omitempty"`
	Transpositions *bool                  `json:"transpositions,omitempty"`
	Value          string                 `json:"value"`
}

// FuzzyQueryBuilder holds FuzzyQuery struct and provides a builder API.
type FuzzyQueryBuilder struct {
	v *FuzzyQuery
}

// NewFuzzyQuery provides a builder for the FuzzyQuery struct.
func NewFuzzyQueryBuilder() *FuzzyQueryBuilder {
	r := FuzzyQueryBuilder{
		&FuzzyQuery{},
	}

	return &r
}

// Build finalize the chain and returns the FuzzyQuery struct
func (rb *FuzzyQueryBuilder) Build() FuzzyQuery {
	return *rb.v
}

func (rb *FuzzyQueryBuilder) Boost(boost float32) *FuzzyQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *FuzzyQueryBuilder) Fuzziness(fuzziness *FuzzinessBuilder) *FuzzyQueryBuilder {
	v := fuzziness.Build()
	rb.v.Fuzziness = &v
	return rb
}

func (rb *FuzzyQueryBuilder) MaxExpansions(maxexpansions int) *FuzzyQueryBuilder {
	rb.v.MaxExpansions = &maxexpansions
	return rb
}

func (rb *FuzzyQueryBuilder) PrefixLength(prefixlength int) *FuzzyQueryBuilder {
	rb.v.PrefixLength = &prefixlength
	return rb
}

func (rb *FuzzyQueryBuilder) QueryName_(queryname_ string) *FuzzyQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *FuzzyQueryBuilder) Rewrite(rewrite MultiTermQueryRewrite) *FuzzyQueryBuilder {
	rb.v.Rewrite = &rewrite
	return rb
}

func (rb *FuzzyQueryBuilder) Transpositions(transpositions bool) *FuzzyQueryBuilder {
	rb.v.Transpositions = &transpositions
	return rb
}

func (rb *FuzzyQueryBuilder) Value(arg string) *FuzzyQueryBuilder {
	rb.v.Value = arg
	return rb
}
