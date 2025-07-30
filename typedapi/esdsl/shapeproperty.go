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
// https://github.com/elastic/elasticsearch-specification/tree/de4ff9ec1f716256f521d9e30011ad9c284b0dcc

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geoorientation"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _shapeProperty struct {
	v *types.ShapeProperty
}

func NewShapeProperty() *_shapeProperty {

	return &_shapeProperty{v: types.NewShapeProperty()}

}

func (s *_shapeProperty) Coerce(coerce bool) *_shapeProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_shapeProperty) CopyTo(fields ...string) *_shapeProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_shapeProperty) DocValues(docvalues bool) *_shapeProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_shapeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_shapeProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_shapeProperty) Fields(fields map[string]types.Property) *_shapeProperty {

	s.v.Fields = fields
	return s
}

func (s *_shapeProperty) AddField(key string, value types.PropertyVariant) *_shapeProperty {

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

func (s *_shapeProperty) IgnoreAbove(ignoreabove int) *_shapeProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_shapeProperty) IgnoreMalformed(ignoremalformed bool) *_shapeProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_shapeProperty) IgnoreZValue(ignorezvalue bool) *_shapeProperty {

	s.v.IgnoreZValue = &ignorezvalue

	return s
}

func (s *_shapeProperty) Meta(meta map[string]string) *_shapeProperty {

	s.v.Meta = meta
	return s
}

func (s *_shapeProperty) AddMeta(key string, value string) *_shapeProperty {

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

func (s *_shapeProperty) Orientation(orientation geoorientation.GeoOrientation) *_shapeProperty {

	s.v.Orientation = &orientation
	return s
}

func (s *_shapeProperty) Properties(properties map[string]types.Property) *_shapeProperty {

	s.v.Properties = properties
	return s
}

func (s *_shapeProperty) AddProperty(key string, value types.PropertyVariant) *_shapeProperty {

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

func (s *_shapeProperty) Store(store bool) *_shapeProperty {

	s.v.Store = &store

	return s
}

func (s *_shapeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_shapeProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_shapeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_shapeProperty) ShapePropertyCaster() *types.ShapeProperty {
	return s.v
}
