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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/runtimefieldtype"
)

type _compositeSubField struct {
	v *types.CompositeSubField
}

func NewCompositeSubField(type_ runtimefieldtype.RuntimeFieldType) *_compositeSubField {

	tmp := &_compositeSubField{v: types.NewCompositeSubField()}

	tmp.Type(type_)

	return tmp

}

func (s *_compositeSubField) Type(type_ runtimefieldtype.RuntimeFieldType) *_compositeSubField {

	s.v.Type = type_
	return s
}

func (s *_compositeSubField) CompositeSubFieldCaster() *types.CompositeSubField {
	return s.v
}
