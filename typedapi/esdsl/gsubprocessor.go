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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _gsubProcessor struct {
	v *types.GsubProcessor
}

// Converts a string field by applying a regular expression and a replacement.
// If the field is an array of string, all members of the array will be
// converted.
// If any non-string values are encountered, the processor will throw an
// exception.
func NewGsubProcessor(pattern string, replacement string) *_gsubProcessor {

	tmp := &_gsubProcessor{v: types.NewGsubProcessor()}

	tmp.Pattern(pattern)

	tmp.Replacement(replacement)

	return tmp

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_gsubProcessor) Description(description string) *_gsubProcessor {

	s.v.Description = &description

	return s
}

// The field to apply the replacement to.
func (s *_gsubProcessor) Field(field string) *_gsubProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_gsubProcessor) If(if_ types.ScriptVariant) *_gsubProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_gsubProcessor) IgnoreFailure(ignorefailure bool) *_gsubProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If `true` and `field` does not exist or is `null`, the processor quietly
// exits without modifying the document.
func (s *_gsubProcessor) IgnoreMissing(ignoremissing bool) *_gsubProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// Handle failures for the processor.
func (s *_gsubProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_gsubProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// The pattern to be replaced.
func (s *_gsubProcessor) Pattern(pattern string) *_gsubProcessor {

	s.v.Pattern = pattern

	return s
}

// The string to replace the matching patterns with.
func (s *_gsubProcessor) Replacement(replacement string) *_gsubProcessor {

	s.v.Replacement = replacement

	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_gsubProcessor) Tag(tag string) *_gsubProcessor {

	s.v.Tag = &tag

	return s
}

// The field to assign the converted value to
// By default, the `field` is updated in-place.
func (s *_gsubProcessor) TargetField(field string) *_gsubProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_gsubProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Gsub = s.v

	return container
}

func (s *_gsubProcessor) GsubProcessorCaster() *types.GsubProcessor {
	return s.v
}
