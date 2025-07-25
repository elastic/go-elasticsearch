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
// https://github.com/elastic/elasticsearch-specification/tree/cf6914e80d9c586e872b7d5e9e74ca34905dcf5f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _stringifiedboolean struct {
	v types.Stringifiedboolean
}

func NewStringifiedboolean() *_stringifiedboolean {
	return &_stringifiedboolean{v: nil}
}

func (u *_stringifiedboolean) Bool(bool bool) *_stringifiedboolean {

	u.v = &bool

	return u
}

func (u *_stringifiedboolean) String(string string) *_stringifiedboolean {

	u.v = &string

	return u
}

func (u *_stringifiedboolean) StringifiedbooleanCaster() *types.Stringifiedboolean {
	return &u.v
}
