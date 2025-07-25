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
// https://github.com/elastic/elasticsearch-specification/tree/a0b0db20330063a6d11f7997ff443fd2a1a827d1

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _scaledFloatNumberProperty struct {
	v *types.ScaledFloatNumberProperty
}

func NewScaledFloatNumberProperty() *_scaledFloatNumberProperty {

	return &_scaledFloatNumberProperty{v: types.NewScaledFloatNumberProperty()}

}

func (s *_scaledFloatNumberProperty) Boost(boost types.Float64) *_scaledFloatNumberProperty {

	s.v.Boost = &boost

	return s
}

func (s *_scaledFloatNumberProperty) Coerce(coerce bool) *_scaledFloatNumberProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_scaledFloatNumberProperty) CopyTo(fields ...string) *_scaledFloatNumberProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_scaledFloatNumberProperty) DocValues(docvalues bool) *_scaledFloatNumberProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_scaledFloatNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_scaledFloatNumberProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_scaledFloatNumberProperty) Fields(fields map[string]types.Property) *_scaledFloatNumberProperty {

	s.v.Fields = fields
	return s
}

func (s *_scaledFloatNumberProperty) AddField(key string, value types.PropertyVariant) *_scaledFloatNumberProperty {

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

func (s *_scaledFloatNumberProperty) IgnoreAbove(ignoreabove int) *_scaledFloatNumberProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_scaledFloatNumberProperty) IgnoreMalformed(ignoremalformed bool) *_scaledFloatNumberProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_scaledFloatNumberProperty) Index(index bool) *_scaledFloatNumberProperty {

	s.v.Index = &index

	return s
}

func (s *_scaledFloatNumberProperty) Meta(meta map[string]string) *_scaledFloatNumberProperty {

	s.v.Meta = meta
	return s
}

func (s *_scaledFloatNumberProperty) AddMeta(key string, value string) *_scaledFloatNumberProperty {

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

func (s *_scaledFloatNumberProperty) NullValue(nullvalue types.Float64) *_scaledFloatNumberProperty {

	s.v.NullValue = &nullvalue

	return s
}

func (s *_scaledFloatNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_scaledFloatNumberProperty {

	s.v.OnScriptError = &onscripterror
	return s
}

func (s *_scaledFloatNumberProperty) Properties(properties map[string]types.Property) *_scaledFloatNumberProperty {

	s.v.Properties = properties
	return s
}

func (s *_scaledFloatNumberProperty) AddProperty(key string, value types.PropertyVariant) *_scaledFloatNumberProperty {

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

func (s *_scaledFloatNumberProperty) ScalingFactor(scalingfactor types.Float64) *_scaledFloatNumberProperty {

	s.v.ScalingFactor = &scalingfactor

	return s
}

func (s *_scaledFloatNumberProperty) Script(script types.ScriptVariant) *_scaledFloatNumberProperty {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_scaledFloatNumberProperty) Store(store bool) *_scaledFloatNumberProperty {

	s.v.Store = &store

	return s
}

func (s *_scaledFloatNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_scaledFloatNumberProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_scaledFloatNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_scaledFloatNumberProperty {

	s.v.TimeSeriesDimension = &timeseriesdimension

	return s
}

func (s *_scaledFloatNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_scaledFloatNumberProperty {

	s.v.TimeSeriesMetric = &timeseriesmetric
	return s
}

func (s *_scaledFloatNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_scaledFloatNumberProperty) ScaledFloatNumberPropertyCaster() *types.ScaledFloatNumberProperty {
	return s.v
}
