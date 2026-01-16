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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _histogramProperty struct {
	v *types.HistogramProperty
}

func NewHistogramProperty() *_histogramProperty {

	return &_histogramProperty{v: types.NewHistogramProperty()}

}

func (s *_histogramProperty) IgnoreMalformed(ignoremalformed bool) *_histogramProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_histogramProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_histogramProperty {

	s.v.TimeSeriesMetric = &timeseriesmetric
	return s
}

func (s *_histogramProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_histogramProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_histogramProperty) Fields(fields map[string]types.Property) *_histogramProperty {

	s.v.Fields = fields
	return s
}

func (s *_histogramProperty) AddField(key string, value types.PropertyVariant) *_histogramProperty {

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

func (s *_histogramProperty) IgnoreAbove(ignoreabove int) *_histogramProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_histogramProperty) Meta(meta map[string]string) *_histogramProperty {

	s.v.Meta = meta
	return s
}

func (s *_histogramProperty) AddMeta(key string, value string) *_histogramProperty {

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

func (s *_histogramProperty) Properties(properties map[string]types.Property) *_histogramProperty {

	s.v.Properties = properties
	return s
}

func (s *_histogramProperty) AddProperty(key string, value types.PropertyVariant) *_histogramProperty {

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

func (s *_histogramProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_histogramProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_histogramProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_histogramProperty) HistogramPropertyCaster() *types.HistogramProperty {
	return s.v
}
