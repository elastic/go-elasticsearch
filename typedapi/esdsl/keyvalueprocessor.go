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

type _keyValueProcessor struct {
	v *types.KeyValueProcessor
}

// This processor helps automatically parse messages (or specific event fields)
// which are of the `foo=bar` variety.
func NewKeyValueProcessor(fieldsplit string, valuesplit string) *_keyValueProcessor {

	tmp := &_keyValueProcessor{v: types.NewKeyValueProcessor()}

	tmp.FieldSplit(fieldsplit)

	tmp.ValueSplit(valuesplit)

	return tmp

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_keyValueProcessor) Description(description string) *_keyValueProcessor {

	s.v.Description = &description

	return s
}

// List of keys to exclude from document.
func (s *_keyValueProcessor) ExcludeKeys(excludekeys ...string) *_keyValueProcessor {

	for _, v := range excludekeys {

		s.v.ExcludeKeys = append(s.v.ExcludeKeys, v)

	}
	return s
}

// The field to be parsed.
// Supports template snippets.
func (s *_keyValueProcessor) Field(field string) *_keyValueProcessor {

	s.v.Field = field

	return s
}

// Regex pattern to use for splitting key-value pairs.
func (s *_keyValueProcessor) FieldSplit(fieldsplit string) *_keyValueProcessor {

	s.v.FieldSplit = fieldsplit

	return s
}

// Conditionally execute the processor.
func (s *_keyValueProcessor) If(if_ types.ScriptVariant) *_keyValueProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_keyValueProcessor) IgnoreFailure(ignorefailure bool) *_keyValueProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If `true` and `field` does not exist or is `null`, the processor quietly
// exits without modifying the document.
func (s *_keyValueProcessor) IgnoreMissing(ignoremissing bool) *_keyValueProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// List of keys to filter and insert into document.
// Defaults to including all keys.
func (s *_keyValueProcessor) IncludeKeys(includekeys ...string) *_keyValueProcessor {

	for _, v := range includekeys {

		s.v.IncludeKeys = append(s.v.IncludeKeys, v)

	}
	return s
}

// Handle failures for the processor.
func (s *_keyValueProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_keyValueProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Prefix to be added to extracted keys.
func (s *_keyValueProcessor) Prefix(prefix string) *_keyValueProcessor {

	s.v.Prefix = &prefix

	return s
}

// If `true`. strip brackets `()`, `<>`, `[]` as well as quotes `'` and `"` from
// extracted values.
func (s *_keyValueProcessor) StripBrackets(stripbrackets bool) *_keyValueProcessor {

	s.v.StripBrackets = &stripbrackets

	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_keyValueProcessor) Tag(tag string) *_keyValueProcessor {

	s.v.Tag = &tag

	return s
}

// The field to insert the extracted keys into.
// Defaults to the root of the document.
// Supports template snippets.
func (s *_keyValueProcessor) TargetField(field string) *_keyValueProcessor {

	s.v.TargetField = &field

	return s
}

// String of characters to trim from extracted keys.
func (s *_keyValueProcessor) TrimKey(trimkey string) *_keyValueProcessor {

	s.v.TrimKey = &trimkey

	return s
}

// String of characters to trim from extracted values.
func (s *_keyValueProcessor) TrimValue(trimvalue string) *_keyValueProcessor {

	s.v.TrimValue = &trimvalue

	return s
}

// Regex pattern to use for splitting the key from the value within a key-value
// pair.
func (s *_keyValueProcessor) ValueSplit(valuesplit string) *_keyValueProcessor {

	s.v.ValueSplit = valuesplit

	return s
}

func (s *_keyValueProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Kv = s.v

	return container
}

func (s *_keyValueProcessor) KeyValueProcessorCaster() *types.KeyValueProcessor {
	return s.v
}
