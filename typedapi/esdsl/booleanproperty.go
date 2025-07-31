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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _booleanProperty struct {
	v *types.BooleanProperty
}

func NewBooleanProperty() *_booleanProperty {

	return &_booleanProperty{v: types.NewBooleanProperty()}

}

func (s *_booleanProperty) Boost(boost types.Float64) *_booleanProperty {

	s.v.Boost = &boost

	return s
}

func (s *_booleanProperty) CopyTo(fields ...string) *_booleanProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_booleanProperty) DocValues(docvalues bool) *_booleanProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_booleanProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_booleanProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_booleanProperty) Fielddata(fielddata types.NumericFielddataVariant) *_booleanProperty {

	s.v.Fielddata = fielddata.NumericFielddataCaster()

	return s
}

func (s *_booleanProperty) Fields(fields map[string]types.Property) *_booleanProperty {

	s.v.Fields = fields
	return s
}

func (s *_booleanProperty) AddField(key string, value types.PropertyVariant) *_booleanProperty {

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

func (s *_booleanProperty) IgnoreAbove(ignoreabove int) *_booleanProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_booleanProperty) IgnoreMalformed(ignoremalformed bool) *_booleanProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_booleanProperty) Index(index bool) *_booleanProperty {

	s.v.Index = &index

	return s
}

func (s *_booleanProperty) Meta(meta map[string]string) *_booleanProperty {

	s.v.Meta = meta
	return s
}

func (s *_booleanProperty) AddMeta(key string, value string) *_booleanProperty {

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

func (s *_booleanProperty) NullValue(nullvalue bool) *_booleanProperty {

	s.v.NullValue = &nullvalue

	return s
}

func (s *_booleanProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_booleanProperty {

	s.v.OnScriptError = &onscripterror
	return s
}

func (s *_booleanProperty) Properties(properties map[string]types.Property) *_booleanProperty {

	s.v.Properties = properties
	return s
}

func (s *_booleanProperty) AddProperty(key string, value types.PropertyVariant) *_booleanProperty {

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

func (s *_booleanProperty) Script(script types.ScriptVariant) *_booleanProperty {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_booleanProperty) Store(store bool) *_booleanProperty {

	s.v.Store = &store

	return s
}

func (s *_booleanProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_booleanProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_booleanProperty) TimeSeriesDimension(timeseriesdimension bool) *_booleanProperty {

	s.v.TimeSeriesDimension = &timeseriesdimension

	return s
}

func (s *_booleanProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_booleanProperty) BooleanPropertyCaster() *types.BooleanProperty {
	return s.v
}
