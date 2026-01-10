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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _longNumberProperty struct {
	v *types.LongNumberProperty
}

func NewLongNumberProperty() *_longNumberProperty {

	return &_longNumberProperty{v: types.NewLongNumberProperty()}

}

func (s *_longNumberProperty) NullValue(nullvalue int64) *_longNumberProperty {

	s.v.NullValue = &nullvalue

	return s
}

func (s *_longNumberProperty) Boost(boost types.Float64) *_longNumberProperty {

	s.v.Boost = &boost

	return s
}

func (s *_longNumberProperty) Coerce(coerce bool) *_longNumberProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_longNumberProperty) CopyTo(fields ...string) *_longNumberProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_longNumberProperty) DocValues(docvalues bool) *_longNumberProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_longNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_longNumberProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_longNumberProperty) Fields(fields map[string]types.Property) *_longNumberProperty {

	s.v.Fields = fields
	return s
}

func (s *_longNumberProperty) AddField(key string, value types.PropertyVariant) *_longNumberProperty {

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

func (s *_longNumberProperty) IgnoreAbove(ignoreabove int) *_longNumberProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_longNumberProperty) IgnoreMalformed(ignoremalformed bool) *_longNumberProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_longNumberProperty) Index(index bool) *_longNumberProperty {

	s.v.Index = &index

	return s
}

func (s *_longNumberProperty) Meta(meta map[string]string) *_longNumberProperty {

	s.v.Meta = meta
	return s
}

func (s *_longNumberProperty) AddMeta(key string, value string) *_longNumberProperty {

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

func (s *_longNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_longNumberProperty {

	s.v.OnScriptError = &onscripterror
	return s
}

func (s *_longNumberProperty) Properties(properties map[string]types.Property) *_longNumberProperty {

	s.v.Properties = properties
	return s
}

func (s *_longNumberProperty) AddProperty(key string, value types.PropertyVariant) *_longNumberProperty {

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

func (s *_longNumberProperty) Script(script types.ScriptVariant) *_longNumberProperty {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_longNumberProperty) Store(store bool) *_longNumberProperty {

	s.v.Store = &store

	return s
}

func (s *_longNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_longNumberProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_longNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_longNumberProperty {

	s.v.TimeSeriesDimension = &timeseriesdimension

	return s
}

func (s *_longNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_longNumberProperty {

	s.v.TimeSeriesMetric = &timeseriesmetric
	return s
}

func (s *_longNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_longNumberProperty) LongNumberPropertyCaster() *types.LongNumberProperty {
	return s.v
}
