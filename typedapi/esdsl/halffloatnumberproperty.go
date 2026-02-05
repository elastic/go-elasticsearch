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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _halfFloatNumberProperty struct {
	v *types.HalfFloatNumberProperty
}

func NewHalfFloatNumberProperty() *_halfFloatNumberProperty {

	return &_halfFloatNumberProperty{v: types.NewHalfFloatNumberProperty()}

}

func (s *_halfFloatNumberProperty) NullValue(nullvalue float32) *_halfFloatNumberProperty {

	s.v.NullValue = &nullvalue

	return s
}

func (s *_halfFloatNumberProperty) Boost(boost types.Float64) *_halfFloatNumberProperty {

	s.v.Boost = &boost

	return s
}

func (s *_halfFloatNumberProperty) Coerce(coerce bool) *_halfFloatNumberProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_halfFloatNumberProperty) CopyTo(fields ...string) *_halfFloatNumberProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_halfFloatNumberProperty) DocValues(docvalues bool) *_halfFloatNumberProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_halfFloatNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_halfFloatNumberProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_halfFloatNumberProperty) Fields(fields map[string]types.Property) *_halfFloatNumberProperty {

	s.v.Fields = fields
	return s
}

func (s *_halfFloatNumberProperty) AddField(key string, value types.PropertyVariant) *_halfFloatNumberProperty {

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

func (s *_halfFloatNumberProperty) IgnoreAbove(ignoreabove int) *_halfFloatNumberProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_halfFloatNumberProperty) IgnoreMalformed(ignoremalformed bool) *_halfFloatNumberProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_halfFloatNumberProperty) Index(index bool) *_halfFloatNumberProperty {

	s.v.Index = &index

	return s
}

func (s *_halfFloatNumberProperty) Meta(meta map[string]string) *_halfFloatNumberProperty {

	s.v.Meta = meta
	return s
}

func (s *_halfFloatNumberProperty) AddMeta(key string, value string) *_halfFloatNumberProperty {

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

func (s *_halfFloatNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_halfFloatNumberProperty {

	s.v.OnScriptError = &onscripterror
	return s
}

func (s *_halfFloatNumberProperty) Properties(properties map[string]types.Property) *_halfFloatNumberProperty {

	s.v.Properties = properties
	return s
}

func (s *_halfFloatNumberProperty) AddProperty(key string, value types.PropertyVariant) *_halfFloatNumberProperty {

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

func (s *_halfFloatNumberProperty) Script(script types.ScriptVariant) *_halfFloatNumberProperty {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_halfFloatNumberProperty) Store(store bool) *_halfFloatNumberProperty {

	s.v.Store = &store

	return s
}

func (s *_halfFloatNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_halfFloatNumberProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_halfFloatNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_halfFloatNumberProperty {

	s.v.TimeSeriesDimension = &timeseriesdimension

	return s
}

func (s *_halfFloatNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_halfFloatNumberProperty {

	s.v.TimeSeriesMetric = &timeseriesmetric
	return s
}

func (s *_halfFloatNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_halfFloatNumberProperty) HalfFloatNumberPropertyCaster() *types.HalfFloatNumberProperty {
	return s.v
}
