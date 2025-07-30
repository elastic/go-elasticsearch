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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geopointmetrictype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _geoPointProperty struct {
	v *types.GeoPointProperty
}

func NewGeoPointProperty() *_geoPointProperty {

	return &_geoPointProperty{v: types.NewGeoPointProperty()}

}

func (s *_geoPointProperty) CopyTo(fields ...string) *_geoPointProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_geoPointProperty) DocValues(docvalues bool) *_geoPointProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_geoPointProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_geoPointProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_geoPointProperty) Fields(fields map[string]types.Property) *_geoPointProperty {

	s.v.Fields = fields
	return s
}

func (s *_geoPointProperty) AddField(key string, value types.PropertyVariant) *_geoPointProperty {

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

func (s *_geoPointProperty) IgnoreAbove(ignoreabove int) *_geoPointProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_geoPointProperty) IgnoreMalformed(ignoremalformed bool) *_geoPointProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_geoPointProperty) IgnoreZValue(ignorezvalue bool) *_geoPointProperty {

	s.v.IgnoreZValue = &ignorezvalue

	return s
}

func (s *_geoPointProperty) Index(index bool) *_geoPointProperty {

	s.v.Index = &index

	return s
}

func (s *_geoPointProperty) Meta(meta map[string]string) *_geoPointProperty {

	s.v.Meta = meta
	return s
}

func (s *_geoPointProperty) AddMeta(key string, value string) *_geoPointProperty {

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

func (s *_geoPointProperty) NullValue(geolocation types.GeoLocationVariant) *_geoPointProperty {

	s.v.NullValue = *geolocation.GeoLocationCaster()

	return s
}

func (s *_geoPointProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_geoPointProperty {

	s.v.OnScriptError = &onscripterror
	return s
}

func (s *_geoPointProperty) Properties(properties map[string]types.Property) *_geoPointProperty {

	s.v.Properties = properties
	return s
}

func (s *_geoPointProperty) AddProperty(key string, value types.PropertyVariant) *_geoPointProperty {

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

func (s *_geoPointProperty) Script(script types.ScriptVariant) *_geoPointProperty {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_geoPointProperty) Store(store bool) *_geoPointProperty {

	s.v.Store = &store

	return s
}

func (s *_geoPointProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_geoPointProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_geoPointProperty) TimeSeriesMetric(timeseriesmetric geopointmetrictype.GeoPointMetricType) *_geoPointProperty {

	s.v.TimeSeriesMetric = &timeseriesmetric
	return s
}

func (s *_geoPointProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_geoPointProperty) GeoPointPropertyCaster() *types.GeoPointProperty {
	return s.v
}
