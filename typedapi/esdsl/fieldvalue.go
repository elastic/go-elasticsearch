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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _fieldValue struct {
	v types.FieldValue
}

func NewFieldValue() *_fieldValue {
	return &_fieldValue{v: nil}
}

func (u *_fieldValue) Int64(int64 int64) *_fieldValue {

	u.v = &int64

	return u
}

func (u *_fieldValue) Float64(float64 types.Float64) *_fieldValue {

	u.v = &float64

	return u
}

func (u *_fieldValue) String(string string) *_fieldValue {

	u.v = &string

	return u
}

func (u *_fieldValue) Bool(bool bool) *_fieldValue {

	u.v = &bool

	return u
}

func (u *_fieldValue) Nil() *_fieldValue {
	u.v = types.NullValue{}
	return u
}

func (u *_fieldValue) FieldValueCaster() *types.FieldValue {
	return &u.v
}
