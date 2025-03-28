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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/jsonprocessorconflictstrategy"
)

type _jsonProcessor struct {
	v *types.JsonProcessor
}

// Converts a JSON string into a structured JSON object.
func NewJsonProcessor() *_jsonProcessor {

	return &_jsonProcessor{v: types.NewJsonProcessor()}

}

// Flag that forces the parsed JSON to be added at the top level of the
// document.
// `target_field` must not be set when this option is chosen.
func (s *_jsonProcessor) AddToRoot(addtoroot bool) *_jsonProcessor {

	s.v.AddToRoot = &addtoroot

	return s
}

// When set to `replace`, root fields that conflict with fields from the parsed
// JSON will be overridden.
// When set to `merge`, conflicting fields will be merged.
// Only applicable `if add_to_root` is set to true.
func (s *_jsonProcessor) AddToRootConflictStrategy(addtorootconflictstrategy jsonprocessorconflictstrategy.JsonProcessorConflictStrategy) *_jsonProcessor {

	s.v.AddToRootConflictStrategy = &addtorootconflictstrategy
	return s
}

// When set to `true`, the JSON parser will not fail if the JSON contains
// duplicate keys.
// Instead, the last encountered value for any duplicate key wins.
func (s *_jsonProcessor) AllowDuplicateKeys(allowduplicatekeys bool) *_jsonProcessor {

	s.v.AllowDuplicateKeys = &allowduplicatekeys

	return s
}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_jsonProcessor) Description(description string) *_jsonProcessor {

	s.v.Description = &description

	return s
}

// The field to be parsed.
func (s *_jsonProcessor) Field(field string) *_jsonProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_jsonProcessor) If(if_ types.ScriptVariant) *_jsonProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_jsonProcessor) IgnoreFailure(ignorefailure bool) *_jsonProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// Handle failures for the processor.
func (s *_jsonProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_jsonProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_jsonProcessor) Tag(tag string) *_jsonProcessor {

	s.v.Tag = &tag

	return s
}

// The field that the converted structured object will be written into.
// Any existing content in this field will be overwritten.
func (s *_jsonProcessor) TargetField(field string) *_jsonProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_jsonProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Json = s.v

	return container
}

func (s *_jsonProcessor) JsonProcessorCaster() *types.JsonProcessor {
	return s.v
}
