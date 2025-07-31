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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rareTermsAggregation struct {
	v *types.RareTermsAggregation
}

// A multi-bucket value source based aggregation which finds "rare" terms —
// terms that are at the long-tail of the distribution and are not frequent.
func NewRareTermsAggregation() *_rareTermsAggregation {

	return &_rareTermsAggregation{v: types.NewRareTermsAggregation()}

}

func (s *_rareTermsAggregation) Exclude(termsexcludes ...string) *_rareTermsAggregation {

	s.v.Exclude = termsexcludes

	return s
}

func (s *_rareTermsAggregation) Field(field string) *_rareTermsAggregation {

	s.v.Field = &field

	return s
}

func (s *_rareTermsAggregation) Include(termsinclude types.TermsIncludeVariant) *_rareTermsAggregation {

	s.v.Include = *termsinclude.TermsIncludeCaster()

	return s
}

func (s *_rareTermsAggregation) MaxDocCount(maxdoccount int64) *_rareTermsAggregation {

	s.v.MaxDocCount = &maxdoccount

	return s
}

func (s *_rareTermsAggregation) Missing(missing types.MissingVariant) *_rareTermsAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

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
