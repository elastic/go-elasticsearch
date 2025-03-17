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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _rareTermsAggregation struct {
	v *types.RareTermsAggregation
}

// A multi-bucket value source based aggregation which finds "rare" terms —
// terms that are at the long-tail of the distribution and are not frequent.
func NewRareTermsAggregation() *_rareTermsAggregation {

	return &_rareTermsAggregation{v: types.NewRareTermsAggregation()}

}

// Terms that should be excluded from the aggregation.
func (s *_rareTermsAggregation) Exclude(termsexcludes ...string) *_rareTermsAggregation {

	s.v.Exclude = termsexcludes

	return s
}

// The field from which to return rare terms.
func (s *_rareTermsAggregation) Field(field string) *_rareTermsAggregation {

	s.v.Field = &field

	return s
}

// Terms that should be included in the aggregation.
func (s *_rareTermsAggregation) Include(termsinclude types.TermsIncludeVariant) *_rareTermsAggregation {

	s.v.Include = *termsinclude.TermsIncludeCaster()

	return s
}

// The maximum number of documents a term should appear in.
func (s *_rareTermsAggregation) MaxDocCount(maxdoccount int64) *_rareTermsAggregation {

	s.v.MaxDocCount = &maxdoccount

	return s
}

// The value to apply to documents that do not have a value.
// By default, documents without a value are ignored.
func (s *_rareTermsAggregation) Missing(missing types.MissingVariant) *_rareTermsAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

// The precision of the internal CuckooFilters.
// Smaller precision leads to better approximation, but higher memory usage.
func (s *_rareTermsAggregation) Precision(precision types.Float64) *_rareTermsAggregation {

	s.v.Precision = &precision

	return s
}

func (s *_rareTermsAggregation) ValueType(valuetype string) *_rareTermsAggregation {

	s.v.ValueType = &valuetype

	return s
}

func (s *_rareTermsAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.RareTerms = s.v

	return container
}

func (s *_rareTermsAggregation) RareTermsAggregationCaster() *types.RareTermsAggregation {
	return s.v
}
