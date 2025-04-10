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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geoshaperelation"
)

type _enrichProcessor struct {
	v *types.EnrichProcessor
}

// The `enrich` processor can enrich documents with data from another index.
func NewEnrichProcessor(policyname string) *_enrichProcessor {

	tmp := &_enrichProcessor{v: types.NewEnrichProcessor()}

	tmp.PolicyName(policyname)

	return tmp

}

func (s *_enrichProcessor) Description(description string) *_enrichProcessor {

	s.v.Description = &description

	return s
}

func (s *_enrichProcessor) Field(field string) *_enrichProcessor {

	s.v.Field = field

	return s
}

func (s *_enrichProcessor) If(if_ types.ScriptVariant) *_enrichProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_enrichProcessor) IgnoreFailure(ignorefailure bool) *_enrichProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_enrichProcessor) IgnoreMissing(ignoremissing bool) *_enrichProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_enrichProcessor) MaxMatches(maxmatches int) *_enrichProcessor {

	s.v.MaxMatches = &maxmatches

	return s
}

func (s *_enrichProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_enrichProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_enrichProcessor) Override(override bool) *_enrichProcessor {

	s.v.Override = &override

	return s
}

func (s *_enrichProcessor) PolicyName(policyname string) *_enrichProcessor {

	s.v.PolicyName = policyname

	return s
}

func (s *_enrichProcessor) ShapeRelation(shaperelation geoshaperelation.GeoShapeRelation) *_enrichProcessor {

	s.v.ShapeRelation = &shaperelation
	return s
}

func (s *_enrichProcessor) Tag(tag string) *_enrichProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_enrichProcessor) TargetField(field string) *_enrichProcessor {

	s.v.TargetField = field

	return s
}

func (s *_enrichProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Enrich = s.v

	return container
}

func (s *_enrichProcessor) EnrichProcessorCaster() *types.EnrichProcessor {
	return s.v
}
