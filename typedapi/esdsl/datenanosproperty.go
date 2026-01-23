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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _dateNanosProperty struct {
	v *types.DateNanosProperty
}

func NewDateNanosProperty() *_dateNanosProperty {

	return &_dateNanosProperty{v: types.NewDateNanosProperty()}

}

func (s *_dateNanosProperty) Boost(boost types.Float64) *_dateNanosProperty {

	s.v.Boost = &boost

	return s
}

func (s *_dateNanosProperty) Format(format string) *_dateNanosProperty {

	s.v.Format = &format

	return s
}

func (s *_dateNanosProperty) IgnoreMalformed(ignoremalformed bool) *_dateNanosProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_dateNanosProperty) Index(index bool) *_dateNanosProperty {

	s.v.Index = &index

	return s
}

func (s *_dateNanosProperty) NullValue(datetime types.DateTimeVariant) *_dateNanosProperty {

	s.v.NullValue = *datetime.DateTimeCaster()

	return s
}

func (s *_dateNanosProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_dateNanosProperty {

	s.v.OnScriptError = &onscripterror
	return s
}

func (s *_dateNanosProperty) PrecisionStep(precisionstep int) *_dateNanosProperty {

	s.v.PrecisionStep = &precisionstep

	return s
}

func (s *_dateNanosProperty) Script(script types.ScriptVariant) *_dateNanosProperty {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_dateNanosProperty) CopyTo(fields ...string) *_dateNanosProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_dateNanosProperty) DocValues(docvalues bool) *_dateNanosProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_dateNanosProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_dateNanosProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_dateNanosProperty) Fields(fields map[string]types.Property) *_dateNanosProperty {

	s.v.Fields = fields
	return s
}

func (s *_dateNanosProperty) AddField(key string, value types.PropertyVariant) *_dateNanosProperty {

	var tmp map[string]types.Property
	if s.v.Fields == nil {
		s.v.Fields = make(map[string]types.Property)
	} else {
		tmp = s.v.Fields
	}

	tmp[key] = *value.PropertyCaster()

	s.v.Fields = tmp
	return s
}

func (s *_dateNanosProperty) IgnoreAbove(ignoreabove int) *_dateNanosProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_dateNanosProperty) Meta(meta map[string]string) *_dateNanosProperty {

	s.v.Meta = meta
	return s
}

func (s *_dateNanosProperty) AddMeta(key string, value string) *_dateNanosProperty {

	var tmp map[string]string
	if s.v.Meta == nil {
		s.v.Meta = make(map[string]string)
	} else {
		tmp = s.v.Meta
	}

	tmp[key] = value

	s.v.Meta = tmp
	return s
}

func (s *_dateNanosProperty) Properties(properties map[string]types.Property) *_dateNanosProperty {

	s.v.Properties = properties
	return s
}

func (s *_dateNanosProperty) AddProperty(key string, value types.PropertyVariant) *_dateNanosProperty {

	var tmp map[string]types.Property
	if s.v.Properties == nil {
		s.v.Properties = make(map[string]types.Property)
	} else {
		tmp = s.v.Properties
	}

	tmp[key] = *value.PropertyCaster()

	s.v.Properties = tmp
	return s
}

func (s *_dateNanosProperty) Store(store bool) *_dateNanosProperty {

	s.v.Store = &store

	return s
}

func (s *_dateNanosProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_dateNanosProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_dateNanosProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_dateNanosProperty) DateNanosPropertyCaster() *types.DateNanosProperty {
	return s.v
}
