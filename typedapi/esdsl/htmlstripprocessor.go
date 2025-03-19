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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _htmlStripProcessor struct {
	v *types.HtmlStripProcessor
}

// Removes HTML tags from the field.
// If the field is an array of strings, HTML tags will be removed from all
// members of the array.
func NewHtmlStripProcessor() *_htmlStripProcessor {

	return &_htmlStripProcessor{v: types.NewHtmlStripProcessor()}

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_htmlStripProcessor) Description(description string) *_htmlStripProcessor {

	s.v.Description = &description

	return s
}

// The string-valued field to remove HTML tags from.
func (s *_htmlStripProcessor) Field(field string) *_htmlStripProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_htmlStripProcessor) If(if_ types.ScriptVariant) *_htmlStripProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_htmlStripProcessor) IgnoreFailure(ignorefailure bool) *_htmlStripProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If `true` and `field` does not exist or is `null`, the processor quietly
// exits without modifying the document,
func (s *_htmlStripProcessor) IgnoreMissing(ignoremissing bool) *_htmlStripProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// Handle failures for the processor.
func (s *_htmlStripProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_htmlStripProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_htmlStripProcessor) Tag(tag string) *_htmlStripProcessor {

	s.v.Tag = &tag

	return s
}

// The field to assign the converted value to
// By default, the `field` is updated in-place.
func (s *_htmlStripProcessor) TargetField(field string) *_htmlStripProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_htmlStripProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.HtmlStrip = s.v

	return container
}

func (s *_htmlStripProcessor) HtmlStripProcessorCaster() *types.HtmlStripProcessor {
	return s.v
}
