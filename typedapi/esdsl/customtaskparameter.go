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
// https://github.com/elastic/elasticsearch-specification/tree/836fca874204ca4173ae5c36fb6b5107d28d2fc0

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _customTaskParameter struct {
	v types.CustomTaskParameter
}

func NewCustomTaskParameter() *_customTaskParameter {
	return &_customTaskParameter{v: nil}
}

func (u *_customTaskParameter) String(string string) *_customTaskParameter {

	u.v = &string

	return u
}

func (u *_customTaskParameter) Int(int int) *_customTaskParameter {

	u.v = &int

	return u
}

func (u *_customTaskParameter) Float64(float64 types.Float64) *_customTaskParameter {

	u.v = &float64

	return u
}

func (u *_customTaskParameter) Float32(float32 float32) *_customTaskParameter {

	u.v = &float32

	return u
}

func (u *_customTaskParameter) Bool(bool bool) *_customTaskParameter {

	u.v = &bool

	return u
}

func (u *_customTaskParameter) CustomTaskParameterCaster() *types.CustomTaskParameter {
	return &u.v
}
