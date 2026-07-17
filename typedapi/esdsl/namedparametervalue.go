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
// https://github.com/elastic/elasticsearch-specification/tree/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _namedParameterValue struct {
	v types.NamedParameterValue
}

func NewNamedParameterValue() *_namedParameterValue {
	return &_namedParameterValue{v: nil}
}

func (u *_namedParameterValue) SingleOrMultiValue(singleormultivalues ...types.FieldValueVariant) *_namedParameterValue {

	convertedItems := make([]types.FieldValue, 0, len(singleormultivalues))
	for _, v := range singleormultivalues {
		convertedItems = append(convertedItems, *v.FieldValueCaster())
	}
	u.v = convertedItems

	return u
}

func (u *_namedParameterValue) SingleOrMultiValueValues(singleormultivaluevalues []types.FieldValue) *_namedParameterValue {

	u.v = singleormultivaluevalues
	return u
}

// Interface implementation for SingleOrMultiValue in NamedParameterValue union
func (u *_singleOrMultiValue) NamedParameterValueCaster() *types.NamedParameterValue {
	t := types.NamedParameterValue(u.v)
	return &t
}

func (u *_namedParameterValue) ClassifiedNamedParameter(classifiednamedparameter types.ClassifiedNamedParameterVariant) *_namedParameterValue {

	u.v = classifiednamedparameter.ClassifiedNamedParameterCaster()

	return u
}

// Interface implementation for ClassifiedNamedParameter in NamedParameterValue union
func (u *_classifiedNamedParameter) NamedParameterValueCaster() *types.NamedParameterValue {
	t := types.NamedParameterValue(u.v)
	return &t
}

func (u *_namedParameterValue) NamedParameterValueCaster() *types.NamedParameterValue {
	return &u.v
}
