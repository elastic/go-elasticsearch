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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _dateProperty struct {
	v *types.DateProperty
}

func NewDateProperty() *_dateProperty {

	return &_dateProperty{v: types.NewDateProperty()}

}

func (s *_dateProperty) Boost(boost types.Float64) *_dateProperty {

	s.v.Boost = &boost

	return s
}

func (s *_dateProperty) Fielddata(fielddata types.NumericFielddataVariant) *_dateProperty {

	s.v.Fielddata = fielddata.NumericFielddataCaster()

	return s
}

func (s *_dateProperty) Format(format string) *_dateProperty {

	s.v.Format = &format

	return s
}

func (s *_dateProperty) IgnoreMalformed(ignoremalformed bool) *_dateProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_dateProperty) Index(index bool) *_dateProperty {

	s.v.Index = &index

	return s
}

func (s *_dateProperty) Locale(locale string) *_dateProperty {

	s.v.Locale = &locale

	return s
}

func (s *_dateProperty) NullValue(datetime types.DateTimeVariant) *_dateProperty {

	s.v.NullValue = *datetime.DateTimeCaster()

	return s
}

func (s *_dateProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_dateProperty {

	s.v.OnScriptError = &onscripterror
	return s
}

func (s *_dateProperty) PrecisionStep(precisionstep int) *_dateProperty {

	s.v.PrecisionStep = &precisionstep

	return s
}

func (s *_dateProperty) Script(script types.ScriptVariant) *_dateProperty {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_dateProperty) CopyTo(fields ...string) *_dateProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_dateProperty) DocValues(docvalues bool) *_dateProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_dateProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_dateProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_dateProperty) Fields(fields map[string]types.Property) *_dateProperty {

	s.v.Fields = fields
	return s
}

func (s *_dateProperty) AddField(key string, value types.PropertyVariant) *_dateProperty {

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

func (s *_dateProperty) IgnoreAbove(ignoreabove int) *_dateProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_dateProperty) Meta(meta map[string]string) *_dateProperty {

	s.v.Meta = meta
	return s
}

func (s *_dateProperty) AddMeta(key string, value string) *_dateProperty {

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

func (s *_dateProperty) Properties(properties map[string]types.Property) *_dateProperty {

	s.v.Properties = properties
	return s
}

func (s *_dateProperty) AddProperty(key string, value types.PropertyVariant) *_dateProperty {

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

func (s *_dateProperty) Store(store bool) *_dateProperty {

	s.v.Store = &store

	return s
}

func (s *_dateProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_dateProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_dateProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_dateProperty) DatePropertyCaster() *types.DateProperty {
	return s.v
}
