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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _integerNumberProperty struct {
	v *types.IntegerNumberProperty
}

func NewIntegerNumberProperty() *_integerNumberProperty {

	return &_integerNumberProperty{v: types.NewIntegerNumberProperty()}

}

func (s *_integerNumberProperty) Boost(boost types.Float64) *_integerNumberProperty {

	s.v.Boost = &boost

	return s
}

func (s *_integerNumberProperty) Coerce(coerce bool) *_integerNumberProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_integerNumberProperty) CopyTo(fields ...string) *_integerNumberProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_integerNumberProperty) DocValues(docvalues bool) *_integerNumberProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_integerNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_integerNumberProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_integerNumberProperty) Fields(fields map[string]types.Property) *_integerNumberProperty {

	s.v.Fields = fields
	return s
}

func (s *_integerNumberProperty) AddField(key string, value types.PropertyVariant) *_integerNumberProperty {

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

func (s *_integerNumberProperty) IgnoreAbove(ignoreabove int) *_integerNumberProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_integerNumberProperty) IgnoreMalformed(ignoremalformed bool) *_integerNumberProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_integerNumberProperty) Index(index bool) *_integerNumberProperty {

	s.v.Index = &index

	return s
}

func (s *_integerNumberProperty) Meta(meta map[string]string) *_integerNumberProperty {

	s.v.Meta = meta
	return s
}

func (s *_integerNumberProperty) AddMeta(key string, value string) *_integerNumberProperty {

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

func (s *_integerNumberProperty) NullValue(nullvalue int) *_integerNumberProperty {

	s.v.NullValue = &nullvalue

	return s
}

func (s *_integerNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_integerNumberProperty {

	s.v.OnScriptError = &onscripterror
	return s
}

func (s *_integerNumberProperty) Properties(properties map[string]types.Property) *_integerNumberProperty {

	s.v.Properties = properties
	return s
}

func (s *_integerNumberProperty) AddProperty(key string, value types.PropertyVariant) *_integerNumberProperty {

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

func (s *_integerNumberProperty) Script(script types.ScriptVariant) *_integerNumberProperty {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_integerNumberProperty) Store(store bool) *_integerNumberProperty {

	s.v.Store = &store

	return s
}

func (s *_integerNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_integerNumberProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_integerNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_integerNumberProperty {

	s.v.TimeSeriesDimension = &timeseriesdimension

	return s
}

func (s *_integerNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_integerNumberProperty {

	s.v.TimeSeriesMetric = &timeseriesmetric
	return s
}

func (s *_integerNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_integerNumberProperty) IntegerNumberPropertyCaster() *types.IntegerNumberProperty {
	return s.v
}
