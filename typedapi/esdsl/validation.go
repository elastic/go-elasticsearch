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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _validation struct {
	v types.Validation
}

func NewValidation() *_validation {
	return &_validation{v: nil}
}

func (u *_validation) LessThanValidation(lessthanvalidation types.LessThanValidationVariant) *_validation {

	u.v = lessthanvalidation.LessThanValidationCaster()

	return u
}

// Interface implementation for LessThanValidation in Validation union
func (u *_lessThanValidation) ValidationCaster() *types.Validation {
	t := types.Validation(u.v)
	return &t
}

func (u *_validation) GreaterThanValidation(greaterthanvalidation types.GreaterThanValidationVariant) *_validation {

	u.v = greaterthanvalidation.GreaterThanValidationCaster()

	return u
}

// Interface implementation for GreaterThanValidation in Validation union
func (u *_greaterThanValidation) ValidationCaster() *types.Validation {
	t := types.Validation(u.v)
	return &t
}

func (u *_validation) ListTypeValidation(listtypevalidation types.ListTypeValidationVariant) *_validation {

	u.v = listtypevalidation.ListTypeValidationCaster()

	return u
}

// Interface implementation for ListTypeValidation in Validation union
func (u *_listTypeValidation) ValidationCaster() *types.Validation {
	t := types.Validation(u.v)
	return &t
}

func (u *_validation) IncludedInValidation(includedinvalidation types.IncludedInValidationVariant) *_validation {

	u.v = includedinvalidation.IncludedInValidationCaster()

	return u
}

// Interface implementation for IncludedInValidation in Validation union
func (u *_includedInValidation) ValidationCaster() *types.Validation {
	t := types.Validation(u.v)
	return &t
}

func (u *_validation) RegexValidation(regexvalidation types.RegexValidationVariant) *_validation {

	u.v = regexvalidation.RegexValidationCaster()

	return u
}

// Interface implementation for RegexValidation in Validation union
func (u *_regexValidation) ValidationCaster() *types.Validation {
	t := types.Validation(u.v)
	return &t
}

func (u *_validation) ValidationCaster() *types.Validation {
	return &u.v
}
