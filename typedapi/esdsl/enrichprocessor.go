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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geoshaperelation"
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

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_enrichProcessor) Description(description string) *_enrichProcessor {

	s.v.Description = &description

	return s
}

// The field in the input document that matches the policies match_field used to
// retrieve the enrichment data.
// Supports template snippets.
func (s *_enrichProcessor) Field(field string) *_enrichProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_enrichProcessor) If(if_ types.ScriptVariant) *_enrichProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_enrichProcessor) IgnoreFailure(ignorefailure bool) *_enrichProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If `true` and `field` does not exist, the processor quietly exits without
// modifying the document.
func (s *_enrichProcessor) IgnoreMissing(ignoremissing bool) *_enrichProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// The maximum number of matched documents to include under the configured
// target field.
// The `target_field` will be turned into a json array if `max_matches` is
// higher than 1, otherwise `target_field` will become a json object.
// In order to avoid documents getting too large, the maximum allowed value is
// 128.
func (s *_enrichProcessor) MaxMatches(maxmatches int) *_enrichProcessor {

	s.v.MaxMatches = &maxmatches

	return s
}

// Handle failures for the processor.
func (s *_enrichProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_enrichProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// If processor will update fields with pre-existing non-null-valued field.
// When set to `false`, such fields will not be touched.
func (s *_enrichProcessor) Override(override bool) *_enrichProcessor {

	s.v.Override = &override

	return s
}

// The name of the enrich policy to use.
func (s *_enrichProcessor) PolicyName(policyname string) *_enrichProcessor {

	s.v.PolicyName = policyname

	return s
}

// A spatial relation operator used to match the geoshape of incoming documents
// to documents in the enrich index.
// This option is only used for `geo_match` enrich policy types.
func (s *_enrichProcessor) ShapeRelation(shaperelation geoshaperelation.GeoShapeRelation) *_enrichProcessor {

	s.v.ShapeRelation = &shaperelation
	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_enrichProcessor) Tag(tag string) *_enrichProcessor {

	s.v.Tag = &tag

	return s
}

// Field added to incoming documents to contain enrich data. This field contains
// both the `match_field` and `enrich_fields` specified in the enrich policy.
// Supports template snippets.
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
