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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/shapetype"
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

func (s *_circleProcessor) Description(description string) *_circleProcessor {

	s.v.Description = &description

	return s
}

func (s *_circleProcessor) ErrorDistance(errordistance types.Float64) *_circleProcessor {

	s.v.ErrorDistance = errordistance

	return s
}

func (s *_circleProcessor) Field(field string) *_circleProcessor {

	s.v.Field = field

	return s
}

func (s *_circleProcessor) If(if_ types.ScriptVariant) *_circleProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_circleProcessor) IgnoreFailure(ignorefailure bool) *_circleProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_circleProcessor) IgnoreMissing(ignoremissing bool) *_circleProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_circleProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_circleProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_circleProcessor) ShapeType(shapetype shapetype.ShapeType) *_circleProcessor {

	s.v.ShapeType = shapetype
	return s
}

func (s *_circleProcessor) Tag(tag string) *_circleProcessor {

	s.v.Tag = &tag

	return s
}

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
