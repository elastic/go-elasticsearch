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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/runtimefieldtype"
)

type _runtimeField struct {
	v *types.RuntimeField
}

func NewRuntimeField(type_ runtimefieldtype.RuntimeFieldType) *_runtimeField {

	tmp := &_runtimeField{v: types.NewRuntimeField()}

	tmp.Type(type_)

	return tmp

}

func (s *_runtimeField) FetchFields(fetchfields ...types.RuntimeFieldFetchFieldsVariant) *_runtimeField {

	for _, v := range fetchfields {

		s.v.FetchFields = append(s.v.FetchFields, *v.RuntimeFieldFetchFieldsCaster())

	}
	return s
}

func (s *_runtimeField) Fields(fields map[string]types.CompositeSubField) *_runtimeField {

	s.v.Fields = fields
	return s
}

func (s *_runtimeField) AddField(key string, value types.CompositeSubFieldVariant) *_runtimeField {

	var tmp map[string]types.CompositeSubField
	if s.v.Fields == nil {
		s.v.Fields = make(map[string]types.CompositeSubField)
	} else {
		tmp = s.v.Fields
	}

	tmp[key] = *value.CompositeSubFieldCaster()

	s.v.Fields = tmp
	return s
}

func (s *_runtimeField) Format(format string) *_runtimeField {

	s.v.Format = &format

	return s
}

func (s *_runtimeField) InputField(field string) *_runtimeField {

	s.v.InputField = &field

	return s
}

func (s *_runtimeField) Script(script types.ScriptVariant) *_runtimeField {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_runtimeField) TargetField(field string) *_runtimeField {

	s.v.TargetField = &field

	return s
}

func (s *_runtimeField) TargetIndex(indexname string) *_runtimeField {

	s.v.TargetIndex = &indexname

	return s
}

func (s *_runtimeField) Type(type_ runtimefieldtype.RuntimeFieldType) *_runtimeField {

	s.v.Type = type_
	return s
}

func (s *_runtimeField) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Runtime = s.v

	return container
}

func (s *_runtimeField) RuntimeFieldCaster() *types.RuntimeField {
	return s.v
}
