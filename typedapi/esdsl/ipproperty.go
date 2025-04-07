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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syntheticsourcekeepenum"
)

type _ipProperty struct {
	v *types.IpProperty
}

func NewIpProperty() *_ipProperty {

	return &_ipProperty{v: types.NewIpProperty()}

}

func (s *_ipProperty) Boost(boost types.Float64) *_ipProperty {

	s.v.Boost = &boost

	return s
}

func (s *_ipProperty) CopyTo(fields ...string) *_ipProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_ipProperty) DocValues(docvalues bool) *_ipProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_ipProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_ipProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_ipProperty) Fields(fields map[string]types.Property) *_ipProperty {

	s.v.Fields = fields
	return s
}

func (s *_ipProperty) AddField(key string, value types.PropertyVariant) *_ipProperty {

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

func (s *_ipProperty) IgnoreAbove(ignoreabove int) *_ipProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_ipProperty) IgnoreMalformed(ignoremalformed bool) *_ipProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_ipProperty) Index(index bool) *_ipProperty {

	s.v.Index = &index

	return s
}

func (s *_ipProperty) Meta(meta map[string]string) *_ipProperty {

	s.v.Meta = meta
	return s
}

func (s *_ipProperty) AddMeta(key string, value string) *_ipProperty {

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

func (s *_ipProperty) NullValue(nullvalue string) *_ipProperty {

	s.v.NullValue = &nullvalue

	return s
}

func (s *_ipProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_ipProperty {

	s.v.OnScriptError = &onscripterror
	return s
}

func (s *_ipProperty) Properties(properties map[string]types.Property) *_ipProperty {

	s.v.Properties = properties
	return s
}

func (s *_ipProperty) AddProperty(key string, value types.PropertyVariant) *_ipProperty {

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

func (s *_ipProperty) Script(script types.ScriptVariant) *_ipProperty {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_ipProperty) Store(store bool) *_ipProperty {

	s.v.Store = &store

	return s
}

func (s *_ipProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_ipProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_ipProperty) TimeSeriesDimension(timeseriesdimension bool) *_ipProperty {

	s.v.TimeSeriesDimension = &timeseriesdimension

	return s
}

func (s *_ipProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_ipProperty) IpPropertyCaster() *types.IpProperty {
	return s.v
}
