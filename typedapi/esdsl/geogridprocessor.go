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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geogridtargetformat"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geogridtiletype"
)

type _geoGridProcessor struct {
	v *types.GeoGridProcessor
}

// Converts geo-grid definitions of grid tiles or cells to regular bounding
// boxes or polygons which describe their shape.
// This is useful if there is a need to interact with the tile shapes as
// spatially indexable fields.
func NewGeoGridProcessor(field string, tiletype geogridtiletype.GeoGridTileType) *_geoGridProcessor {

	tmp := &_geoGridProcessor{v: types.NewGeoGridProcessor()}

	tmp.Field(field)

	tmp.TileType(tiletype)

	return tmp

}

func (s *_geoGridProcessor) ChildrenField(field string) *_geoGridProcessor {

	s.v.ChildrenField = &field

	return s
}

func (s *_geoGridProcessor) Description(description string) *_geoGridProcessor {

	s.v.Description = &description

	return s
}

func (s *_geoGridProcessor) Field(field string) *_geoGridProcessor {

	s.v.Field = field

	return s
}

func (s *_geoGridProcessor) If(if_ types.ScriptVariant) *_geoGridProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_geoGridProcessor) IgnoreFailure(ignorefailure bool) *_geoGridProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_geoGridProcessor) IgnoreMissing(ignoremissing bool) *_geoGridProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_geoGridProcessor) NonChildrenField(field string) *_geoGridProcessor {

	s.v.NonChildrenField = &field

	return s
}

func (s *_geoGridProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_geoGridProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_geoGridProcessor) ParentField(field string) *_geoGridProcessor {

	s.v.ParentField = &field

	return s
}

func (s *_geoGridProcessor) PrecisionField(field string) *_geoGridProcessor {

	s.v.PrecisionField = &field

	return s
}

func (s *_geoGridProcessor) Tag(tag string) *_geoGridProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_geoGridProcessor) TargetField(field string) *_geoGridProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_geoGridProcessor) TargetFormat(targetformat geogridtargetformat.GeoGridTargetFormat) *_geoGridProcessor {

	s.v.TargetFormat = &targetformat
	return s
}

func (s *_geoGridProcessor) TileType(tiletype geogridtiletype.GeoGridTileType) *_geoGridProcessor {

	s.v.TileType = tiletype
	return s
}

func (s *_geoGridProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.GeoGrid = s.v

	return container
}

func (s *_geoGridProcessor) GeoGridProcessorCaster() *types.GeoGridProcessor {
	return s.v
}
