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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geoorientation"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geostrategy"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _geoShapeProperty struct {
	v *types.GeoShapeProperty
}

func NewGeoShapeProperty() *_geoShapeProperty {

	return &_geoShapeProperty{v: types.NewGeoShapeProperty()}

}

func (s *_geoShapeProperty) Coerce(coerce bool) *_geoShapeProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_geoShapeProperty) IgnoreMalformed(ignoremalformed bool) *_geoShapeProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_geoShapeProperty) IgnoreZValue(ignorezvalue bool) *_geoShapeProperty {

	s.v.IgnoreZValue = &ignorezvalue

	return s
}

func (s *_geoShapeProperty) Index(index bool) *_geoShapeProperty {

	s.v.Index = &index

	return s
}

func (s *_geoShapeProperty) Orientation(orientation geoorientation.GeoOrientation) *_geoShapeProperty {

	s.v.Orientation = &orientation
	return s
}

func (s *_geoShapeProperty) Strategy(strategy geostrategy.GeoStrategy) *_geoShapeProperty {

	s.v.Strategy = &strategy
	return s
}

func (s *_geoShapeProperty) CopyTo(fields ...string) *_geoShapeProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_geoShapeProperty) DocValues(docvalues bool) *_geoShapeProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_geoShapeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_geoShapeProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_geoShapeProperty) Fields(fields map[string]types.Property) *_geoShapeProperty {

	s.v.Fields = fields
	return s
}

func (s *_geoShapeProperty) AddField(key string, value types.PropertyVariant) *_geoShapeProperty {

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

func (s *_geoShapeProperty) IgnoreAbove(ignoreabove int) *_geoShapeProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_geoShapeProperty) Meta(meta map[string]string) *_geoShapeProperty {

	s.v.Meta = meta
	return s
}

func (s *_geoShapeProperty) AddMeta(key string, value string) *_geoShapeProperty {

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

func (s *_geoShapeProperty) Properties(properties map[string]types.Property) *_geoShapeProperty {

	s.v.Properties = properties
	return s
}

func (s *_geoShapeProperty) AddProperty(key string, value types.PropertyVariant) *_geoShapeProperty {

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

func (s *_geoShapeProperty) Store(store bool) *_geoShapeProperty {

	s.v.Store = &store

	return s
}

func (s *_geoShapeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_geoShapeProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_geoShapeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_geoShapeProperty) GeoShapePropertyCaster() *types.GeoShapeProperty {
	return s.v
}
