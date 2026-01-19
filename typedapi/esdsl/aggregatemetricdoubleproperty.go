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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _aggregateMetricDoubleProperty struct {
	v *types.AggregateMetricDoubleProperty
}

func NewAggregateMetricDoubleProperty(defaultmetric string) *_aggregateMetricDoubleProperty {

	tmp := &_aggregateMetricDoubleProperty{v: types.NewAggregateMetricDoubleProperty()}

	tmp.DefaultMetric(defaultmetric)

	return tmp

}

func (s *_aggregateMetricDoubleProperty) DefaultMetric(defaultmetric string) *_aggregateMetricDoubleProperty {

	s.v.DefaultMetric = defaultmetric

	return s
}

func (s *_aggregateMetricDoubleProperty) IgnoreMalformed(ignoremalformed bool) *_aggregateMetricDoubleProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_aggregateMetricDoubleProperty) Metrics(metrics ...string) *_aggregateMetricDoubleProperty {

	for _, v := range metrics {

		s.v.Metrics = append(s.v.Metrics, v)

	}
	return s
}

func (s *_aggregateMetricDoubleProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_aggregateMetricDoubleProperty {

	s.v.TimeSeriesMetric = &timeseriesmetric
	return s
}

func (s *_aggregateMetricDoubleProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_aggregateMetricDoubleProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_aggregateMetricDoubleProperty) Fields(fields map[string]types.Property) *_aggregateMetricDoubleProperty {

	s.v.Fields = fields
	return s
}

func (s *_aggregateMetricDoubleProperty) AddField(key string, value types.PropertyVariant) *_aggregateMetricDoubleProperty {

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

func (s *_aggregateMetricDoubleProperty) IgnoreAbove(ignoreabove int) *_aggregateMetricDoubleProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_aggregateMetricDoubleProperty) Meta(meta map[string]string) *_aggregateMetricDoubleProperty {

	s.v.Meta = meta
	return s
}

func (s *_aggregateMetricDoubleProperty) AddMeta(key string, value string) *_aggregateMetricDoubleProperty {

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

func (s *_aggregateMetricDoubleProperty) Properties(properties map[string]types.Property) *_aggregateMetricDoubleProperty {

	s.v.Properties = properties
	return s
}

func (s *_aggregateMetricDoubleProperty) AddProperty(key string, value types.PropertyVariant) *_aggregateMetricDoubleProperty {

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

func (s *_aggregateMetricDoubleProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_aggregateMetricDoubleProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_aggregateMetricDoubleProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_aggregateMetricDoubleProperty) AggregateMetricDoublePropertyCaster() *types.AggregateMetricDoubleProperty {
	return s.v
}
