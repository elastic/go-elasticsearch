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

type _splitProcessor struct {
	v *types.SplitProcessor
}

// Splits a field into an array using a separator character.
// Only works on string fields.
func NewSplitProcessor(separator string) *_splitProcessor {

	tmp := &_splitProcessor{v: types.NewSplitProcessor()}

	tmp.Separator(separator)

	return tmp

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_splitProcessor) Description(description string) *_splitProcessor {

	s.v.Description = &description

	return s
}

// The field to split.
func (s *_splitProcessor) Field(field string) *_splitProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_splitProcessor) If(if_ types.ScriptVariant) *_splitProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_splitProcessor) IgnoreFailure(ignorefailure bool) *_splitProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If `true` and `field` does not exist, the processor quietly exits without
// modifying the document.
func (s *_splitProcessor) IgnoreMissing(ignoremissing bool) *_splitProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// Handle failures for the processor.
func (s *_splitProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_splitProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Preserves empty trailing fields, if any.
func (s *_splitProcessor) PreserveTrailing(preservetrailing bool) *_splitProcessor {

	s.v.PreserveTrailing = &preservetrailing

	return s
}

// A regex which matches the separator, for example, `,` or `\s+`.
func (s *_splitProcessor) Separator(separator string) *_splitProcessor {

	s.v.Separator = separator

	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_splitProcessor) Tag(tag string) *_splitProcessor {

	s.v.Tag = &tag

	return s
}

// The field to assign the split value to.
// By default, the field is updated in-place.
func (s *_splitProcessor) TargetField(field string) *_splitProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_splitProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Split = s.v

	return container
}

func (s *_splitProcessor) SplitProcessorCaster() *types.SplitProcessor {
	return s.v
}
