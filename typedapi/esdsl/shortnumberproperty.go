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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _shortNumberProperty struct {
	v *types.ShortNumberProperty
}

func NewShortNumberProperty() *_shortNumberProperty {

	return &_shortNumberProperty{v: types.NewShortNumberProperty()}

}

func (s *_shortNumberProperty) Boost(boost types.Float64) *_shortNumberProperty {

	s.v.Boost = &boost

	return s
}

func (s *_shortNumberProperty) Coerce(coerce bool) *_shortNumberProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_shortNumberProperty) CopyTo(fields ...string) *_shortNumberProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_shortNumberProperty) DocValues(docvalues bool) *_shortNumberProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_shortNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_shortNumberProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_shortNumberProperty) Fields(fields map[string]types.Property) *_shortNumberProperty {

	s.v.Fields = fields
	return s
}

func (s *_shortNumberProperty) AddField(key string, value types.PropertyVariant) *_shortNumberProperty {

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

func (s *_shortNumberProperty) IgnoreAbove(ignoreabove int) *_shortNumberProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_shortNumberProperty) IgnoreMalformed(ignoremalformed bool) *_shortNumberProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_shortNumberProperty) Index(index bool) *_shortNumberProperty {

	s.v.Index = &index

	return s
}

func (s *_shortNumberProperty) Meta(meta map[string]string) *_shortNumberProperty {

	s.v.Meta = meta
	return s
}

func (s *_shortNumberProperty) AddMeta(key string, value string) *_shortNumberProperty {

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

func (s *_shortNumberProperty) NullValue(nullvalue int) *_shortNumberProperty {

	s.v.NullValue = &nullvalue

	return s
}

func (s *_shortNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_shortNumberProperty {

	s.v.OnScriptError = &onscripterror
	return s
}

func (s *_shortNumberProperty) Properties(properties map[string]types.Property) *_shortNumberProperty {

	s.v.Properties = properties
	return s
}

func (s *_shortNumberProperty) AddProperty(key string, value types.PropertyVariant) *_shortNumberProperty {

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

func (s *_shortNumberProperty) Script(script types.ScriptVariant) *_shortNumberProperty {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_shortNumberProperty) Store(store bool) *_shortNumberProperty {

	s.v.Store = &store

	return s
}

func (s *_shortNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_shortNumberProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_shortNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_shortNumberProperty {

	s.v.TimeSeriesDimension = &timeseriesdimension

	return s
}

func (s *_shortNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_shortNumberProperty {

	s.v.TimeSeriesMetric = &timeseriesmetric
	return s
}

func (s *_shortNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_shortNumberProperty) ShortNumberPropertyCaster() *types.ShortNumberProperty {
	return s.v
}
