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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/shapetype"
)

type _circleProcessor struct {
	v *types.CircleProcessor
}

// Converts circle definitions of shapes to regular polygons which approximate
// them.
func NewCircleProcessor(errordistance types.Float64, shapetype shapetype.ShapeType) *_circleProcessor {

	tmp := &_circleProcessor{v: types.NewCircleProcessor()}

	tmp.ErrorDistance(errordistance)

	tmp.ShapeType(shapetype)

	return tmp

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_circleProcessor) Description(description string) *_circleProcessor {

	s.v.Description = &description

	return s
}

// The difference between the resulting inscribed distance from center to side
// and the circle’s radius (measured in meters for `geo_shape`, unit-less for
// `shape`).
func (s *_circleProcessor) ErrorDistance(errordistance types.Float64) *_circleProcessor {

	s.v.ErrorDistance = errordistance

	return s
}

// The field to interpret as a circle. Either a string in WKT format or a map
// for GeoJSON.
func (s *_circleProcessor) Field(field string) *_circleProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_circleProcessor) If(if_ types.ScriptVariant) *_circleProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_circleProcessor) IgnoreFailure(ignorefailure bool) *_circleProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If `true` and `field` does not exist, the processor quietly exits without
// modifying the document.
func (s *_circleProcessor) IgnoreMissing(ignoremissing bool) *_circleProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// Handle failures for the processor.
func (s *_circleProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_circleProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Which field mapping type is to be used when processing the circle:
// `geo_shape` or `shape`.
func (s *_circleProcessor) ShapeType(shapetype shapetype.ShapeType) *_circleProcessor {

	s.v.ShapeType = shapetype
	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_circleProcessor) Tag(tag string) *_circleProcessor {

	s.v.Tag = &tag

	return s
}

// The field to assign the polygon shape to
// By default, the field is updated in-place.
func (s *_circleProcessor) TargetField(field string) *_circleProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_circleProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Circle = s.v

	return container
}

func (s *_circleProcessor) CircleProcessorCaster() *types.CircleProcessor {
	return s.v
}
