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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/valuetype"
)

type _compositeTermsAggregation struct {
	v *types.CompositeTermsAggregation
}

func NewCompositeTermsAggregation() *_compositeTermsAggregation {

	return &_compositeTermsAggregation{v: types.NewCompositeTermsAggregation()}

}

func (s *_compositeTermsAggregation) Field(field string) *_compositeTermsAggregation {

	s.v.Field = &field

	return s
}

func (s *_compositeTermsAggregation) MissingBucket(missingbucket bool) *_compositeTermsAggregation {

	s.v.MissingBucket = &missingbucket

	return s
}

func (s *_compositeTermsAggregation) MissingOrder(missingorder missingorder.MissingOrder) *_compositeTermsAggregation {

	s.v.MissingOrder = &missingorder
	return s
}

func (s *_compositeTermsAggregation) Order(order sortorder.SortOrder) *_compositeTermsAggregation {

	s.v.Order = &order
	return s
}

func (s *_compositeTermsAggregation) Script(script types.ScriptVariant) *_compositeTermsAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_compositeTermsAggregation) ValueType(valuetype valuetype.ValueType) *_compositeTermsAggregation {

	s.v.ValueType = &valuetype
	return s
}

func (s *_compositeTermsAggregation) CompositeTermsAggregationCaster() *types.CompositeTermsAggregation {
	return s.v
}
