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
// https://github.com/elastic/elasticsearch-specification/tree/cf6914e80d9c586e872b7d5e9e74ca34905dcf5f

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _unsignedLongNumberProperty struct {
	v *types.UnsignedLongNumberProperty
}

func NewUnsignedLongNumberProperty() *_unsignedLongNumberProperty {

	return &_unsignedLongNumberProperty{v: types.NewUnsignedLongNumberProperty()}

}

func (s *_unsignedLongNumberProperty) Boost(boost types.Float64) *_unsignedLongNumberProperty {

	s.v.Boost = &boost

	return s
}

func (s *_unsignedLongNumberProperty) Coerce(coerce bool) *_unsignedLongNumberProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_unsignedLongNumberProperty) CopyTo(fields ...string) *_unsignedLongNumberProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_unsignedLongNumberProperty) DocValues(docvalues bool) *_unsignedLongNumberProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_unsignedLongNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_unsignedLongNumberProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_unsignedLongNumberProperty) Fields(fields map[string]types.Property) *_unsignedLongNumberProperty {

	s.v.Fields = fields
	return s
}

func (s *_unsignedLongNumberProperty) AddField(key string, value types.PropertyVariant) *_unsignedLongNumberProperty {

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

func (s *_unsignedLongNumberProperty) IgnoreAbove(ignoreabove int) *_unsignedLongNumberProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_unsignedLongNumberProperty) IgnoreMalformed(ignoremalformed bool) *_unsignedLongNumberProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_unsignedLongNumberProperty) Index(index bool) *_unsignedLongNumberProperty {

	s.v.Index = &index

	return s
}

func (s *_unsignedLongNumberProperty) Meta(meta map[string]string) *_unsignedLongNumberProperty {

	s.v.Meta = meta
	return s
}

func (s *_unsignedLongNumberProperty) AddMeta(key string, value string) *_unsignedLongNumberProperty {

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

func (s *_unsignedLongNumberProperty) NullValue(nullvalue uint64) *_unsignedLongNumberProperty {

	s.v.NullValue = &nullvalue

	return s
}

func (s *_unsignedLongNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_unsignedLongNumberProperty {

	s.v.OnScriptError = &onscripterror
	return s
}

func (s *_unsignedLongNumberProperty) Properties(properties map[string]types.Property) *_unsignedLongNumberProperty {

	s.v.Properties = properties
	return s
}

func (s *_unsignedLongNumberProperty) AddProperty(key string, value types.PropertyVariant) *_unsignedLongNumberProperty {

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

func (s *_unsignedLongNumberProperty) Script(script types.ScriptVariant) *_unsignedLongNumberProperty {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_unsignedLongNumberProperty) Store(store bool) *_unsignedLongNumberProperty {

	s.v.Store = &store

	return s
}

func (s *_unsignedLongNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_unsignedLongNumberProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_unsignedLongNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_unsignedLongNumberProperty {

	s.v.TimeSeriesDimension = &timeseriesdimension

	return s
}

func (s *_unsignedLongNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_unsignedLongNumberProperty {

	s.v.TimeSeriesMetric = &timeseriesmetric
	return s
}

func (s *_unsignedLongNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_unsignedLongNumberProperty) UnsignedLongNumberPropertyCaster() *types.UnsignedLongNumberProperty {
	return s.v
}
