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
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _csvProcessor struct {
	v *types.CsvProcessor
}

// Extracts fields from CSV line out of a single text field within a document.
// Any empty field in CSV will be skipped.
func NewCsvProcessor() *_csvProcessor {

	return &_csvProcessor{v: types.NewCsvProcessor()}

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_csvProcessor) Description(description string) *_csvProcessor {

	s.v.Description = &description

	return s
}

// Value used to fill empty fields.
// Empty fields are skipped if this is not provided.
// An empty field is one with no value (2 consecutive separators) or empty
// quotes (`""`).
func (s *_csvProcessor) EmptyValue(emptyvalue json.RawMessage) *_csvProcessor {

	s.v.EmptyValue = emptyvalue

	return s
}

// The field to extract data from.
func (s *_csvProcessor) Field(field string) *_csvProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_csvProcessor) If(if_ types.ScriptVariant) *_csvProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_csvProcessor) IgnoreFailure(ignorefailure bool) *_csvProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If `true` and `field` does not exist, the processor quietly exits without
// modifying the document.
func (s *_csvProcessor) IgnoreMissing(ignoremissing bool) *_csvProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// Handle failures for the processor.
func (s *_csvProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_csvProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Quote used in CSV, has to be single character string.
func (s *_csvProcessor) Quote(quote string) *_csvProcessor {

	s.v.Quote = &quote

	return s
}

// Separator used in CSV, has to be single character string.
func (s *_csvProcessor) Separator(separator string) *_csvProcessor {

	s.v.Separator = &separator

	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_csvProcessor) Tag(tag string) *_csvProcessor {

	s.v.Tag = &tag

	return s
}

// The array of fields to assign extracted values to.
func (s *_csvProcessor) TargetFields(fields ...string) *_csvProcessor {

	s.v.TargetFields = fields

	return s
}

// Trim whitespaces in unquoted fields.
func (s *_csvProcessor) Trim(trim bool) *_csvProcessor {

	s.v.Trim = &trim

	return s
}

func (s *_csvProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Csv = s.v

	return container
}

func (s *_csvProcessor) CsvProcessorCaster() *types.CsvProcessor {
	return s.v
}
