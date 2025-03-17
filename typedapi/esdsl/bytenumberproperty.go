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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeseriesmetrictype"
)

type _byteNumberProperty struct {
	v *types.ByteNumberProperty
}

func NewByteNumberProperty() *_byteNumberProperty {

	return &_byteNumberProperty{v: types.NewByteNumberProperty()}

}

func (s *_byteNumberProperty) Boost(boost types.Float64) *_byteNumberProperty {

	s.v.Boost = &boost

	return s
}

func (s *_byteNumberProperty) Coerce(coerce bool) *_byteNumberProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_byteNumberProperty) CopyTo(fields ...string) *_byteNumberProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_byteNumberProperty) DocValues(docvalues bool) *_byteNumberProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_byteNumberProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_byteNumberProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_byteNumberProperty) Fields(fields map[string]types.Property) *_byteNumberProperty {

	s.v.Fields = fields
	return s
}

func (s *_byteNumberProperty) AddField(key string, value types.PropertyVariant) *_byteNumberProperty {

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

func (s *_byteNumberProperty) IgnoreAbove(ignoreabove int) *_byteNumberProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_byteNumberProperty) IgnoreMalformed(ignoremalformed bool) *_byteNumberProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_byteNumberProperty) Index(index bool) *_byteNumberProperty {

	s.v.Index = &index

	return s
}

// Metadata about the field.
func (s *_byteNumberProperty) Meta(meta map[string]string) *_byteNumberProperty {

	s.v.Meta = meta
	return s
}

func (s *_byteNumberProperty) AddMeta(key string, value string) *_byteNumberProperty {

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

func (s *_byteNumberProperty) NullValue(nullvalue byte) *_byteNumberProperty {

	s.v.NullValue = &nullvalue

	return s
}

func (s *_byteNumberProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_byteNumberProperty {

	s.v.OnScriptError = &onscripterror
	return s
}

func (s *_byteNumberProperty) Properties(properties map[string]types.Property) *_byteNumberProperty {

	s.v.Properties = properties
	return s
}

func (s *_byteNumberProperty) AddProperty(key string, value types.PropertyVariant) *_byteNumberProperty {

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

func (s *_byteNumberProperty) Script(script types.ScriptVariant) *_byteNumberProperty {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_byteNumberProperty) Store(store bool) *_byteNumberProperty {

	s.v.Store = &store

	return s
}

func (s *_byteNumberProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_byteNumberProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

// For internal use by Elastic only. Marks the field as a time series dimension.
// Defaults to false.
func (s *_byteNumberProperty) TimeSeriesDimension(timeseriesdimension bool) *_byteNumberProperty {

	s.v.TimeSeriesDimension = &timeseriesdimension

	return s
}

// For internal use by Elastic only. Marks the field as a time series dimension.
// Defaults to false.
func (s *_byteNumberProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_byteNumberProperty {

	s.v.TimeSeriesMetric = &timeseriesmetric
	return s
}

func (s *_byteNumberProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_byteNumberProperty) ByteNumberPropertyCaster() *types.ByteNumberProperty {
	return s.v
}
