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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _floatNumberProperty struct {
	v *types.FloatNumberProperty
}

func NewFloatNumberProperty() *_floatNumberProperty {

	return &_floatNumberProperty{v: types.NewFloatNumberProperty()}

}

func (s *_floatNumberProperty) NullValue(nullvalue float32) *_floatNumberProperty {

	s.v.NullValue = &nullvalue

	return s
}

func (s *_floatNumberProperty) Boost(boost types.Float64) *_floatNumberProperty {

	s.v.Boost = &boost

	return s
}

func (s *_floatNumberProperty) Coerce(coerce bool) *_floatNumberProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_floatNumberProperty) CopyTo(fields ...string) *_floatNumberProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_floatNumberProperty) DocValues(docvalues bool) *_floatNumberProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_floatNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_floatNumberProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_floatNumberProperty) Fields(fields map[string]types.Property) *_floatNumberProperty {

	s.v.Fields = fields
	return s
}

func (s *_floatNumberProperty) AddField(key string, value types.PropertyVariant) *_floatNumberProperty {

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

func (s *_floatNumberProperty) IgnoreAbove(ignoreabove int) *_floatNumberProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_floatNumberProperty) IgnoreMalformed(ignoremalformed bool) *_floatNumberProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_floatNumberProperty) Index(index bool) *_floatNumberProperty {

	s.v.Index = &index

	return s
}

func (s *_floatNumberProperty) Meta(meta map[string]string) *_floatNumberProperty {

	s.v.Meta = meta
	return s
}

func (s *_floatNumberProperty) AddMeta(key string, value string) *_floatNumberProperty {

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

func (s *_floatNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_floatNumberProperty {

	s.v.OnScriptError = &onscripterror
	return s
}

func (s *_floatNumberProperty) Properties(properties map[string]types.Property) *_floatNumberProperty {

	s.v.Properties = properties
	return s
}

func (s *_floatNumberProperty) AddProperty(key string, value types.PropertyVariant) *_floatNumberProperty {

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

func (s *_floatNumberProperty) Script(script types.ScriptVariant) *_floatNumberProperty {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_floatNumberProperty) Store(store bool) *_floatNumberProperty {

	s.v.Store = &store

	return s
}

func (s *_floatNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_floatNumberProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_floatNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_floatNumberProperty {

	s.v.TimeSeriesDimension = &timeseriesdimension

	return s
}

func (s *_floatNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_floatNumberProperty {

	s.v.TimeSeriesMetric = &timeseriesmetric
	return s
}

func (s *_floatNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_floatNumberProperty) FloatNumberPropertyCaster() *types.FloatNumberProperty {
	return s.v
}
