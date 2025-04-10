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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _termsQueryField struct {
	v types.TermsQueryField
}

func NewTermsQueryField() *_termsQueryField {
	return &_termsQueryField{v: nil}
}

func (u *_termsQueryField) FieldValues(fieldvalues ...types.FieldValueVariant) *_termsQueryField {

	u.v = make([]types.FieldValue, len(fieldvalues))
	for i, v := range fieldvalues {
		u.v.([]types.FieldValue)[i] = *v.FieldValueCaster()
	}

	return u
}

func (u *_termsQueryField) TermsLookup(termslookup types.TermsLookupVariant) *_termsQueryField {

	u.v = &termslookup

	return u
}

// Interface implementation for TermsLookup in TermsQueryField union
func (u *_termsLookup) TermsQueryFieldCaster() *types.TermsQueryField {
	t := types.TermsQueryField(u.v)
	return &t
}

func (u *_termsQueryField) TermsQueryFieldCaster() *types.TermsQueryField {
	return &u.v
}
