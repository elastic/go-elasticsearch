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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _doubleNumberProperty struct {
	v *types.DoubleNumberProperty
}

func NewDoubleNumberProperty() *_doubleNumberProperty {

	return &_doubleNumberProperty{v: types.NewDoubleNumberProperty()}

}

func (s *_doubleNumberProperty) NullValue(nullvalue types.Float64) *_doubleNumberProperty {

	s.v.NullValue = &nullvalue

	return s
}

func (s *_doubleNumberProperty) Boost(boost types.Float64) *_doubleNumberProperty {

	s.v.Boost = &boost

	return s
}

func (s *_doubleNumberProperty) Coerce(coerce bool) *_doubleNumberProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_doubleNumberProperty) CopyTo(fields ...string) *_doubleNumberProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_doubleNumberProperty) DocValues(docvalues bool) *_doubleNumberProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_doubleNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_doubleNumberProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_doubleNumberProperty) Fields(fields map[string]types.Property) *_doubleNumberProperty {

	s.v.Fields = fields
	return s
}

func (s *_doubleNumberProperty) AddField(key string, value types.PropertyVariant) *_doubleNumberProperty {

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

func (s *_doubleNumberProperty) IgnoreAbove(ignoreabove int) *_doubleNumberProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_doubleNumberProperty) IgnoreMalformed(ignoremalformed bool) *_doubleNumberProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_doubleNumberProperty) Index(index bool) *_doubleNumberProperty {

	s.v.Index = &index

	return s
}

func (s *_doubleNumberProperty) Meta(meta map[string]string) *_doubleNumberProperty {

	s.v.Meta = meta
	return s
}

func (s *_doubleNumberProperty) AddMeta(key string, value string) *_doubleNumberProperty {

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

func (s *_doubleNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_doubleNumberProperty {

	s.v.OnScriptError = &onscripterror
	return s
}

func (s *_doubleNumberProperty) Properties(properties map[string]types.Property) *_doubleNumberProperty {

	s.v.Properties = properties
	return s
}

func (s *_doubleNumberProperty) AddProperty(key string, value types.PropertyVariant) *_doubleNumberProperty {

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

func (s *_doubleNumberProperty) Script(script types.ScriptVariant) *_doubleNumberProperty {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_doubleNumberProperty) Store(store bool) *_doubleNumberProperty {

	s.v.Store = &store

	return s
}

func (s *_doubleNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_doubleNumberProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_doubleNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_doubleNumberProperty {

	s.v.TimeSeriesDimension = &timeseriesdimension

	return s
}

func (s *_doubleNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_doubleNumberProperty {

	s.v.TimeSeriesMetric = &timeseriesmetric
	return s
}

func (s *_doubleNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_doubleNumberProperty) DoubleNumberPropertyCaster() *types.DoubleNumberProperty {
	return s.v
}
