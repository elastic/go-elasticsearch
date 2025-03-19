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

type _dateIndexNameProcessor struct {
	v *types.DateIndexNameProcessor
}

// The purpose of this processor is to point documents to the right time based
// index based on a date or timestamp field in a document by using the date math
// index name support.
func NewDateIndexNameProcessor(daterounding string) *_dateIndexNameProcessor {

	tmp := &_dateIndexNameProcessor{v: types.NewDateIndexNameProcessor()}

	tmp.DateRounding(daterounding)

	return tmp

}

// An array of the expected date formats for parsing dates / timestamps in the
// document being preprocessed.
// Can be a java time pattern or one of the following formats: ISO8601, UNIX,
// UNIX_MS, or TAI64N.
func (s *_dateIndexNameProcessor) DateFormats(dateformats ...string) *_dateIndexNameProcessor {

	for _, v := range dateformats {

		s.v.DateFormats = append(s.v.DateFormats, v)

	}
	return s
}

// How to round the date when formatting the date into the index name. Valid
// values are:
// `y` (year), `M` (month), `w` (week), `d` (day), `h` (hour), `m` (minute) and
// `s` (second).
// Supports template snippets.
func (s *_dateIndexNameProcessor) DateRounding(daterounding string) *_dateIndexNameProcessor {

	s.v.DateRounding = daterounding

	return s
}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_dateIndexNameProcessor) Description(description string) *_dateIndexNameProcessor {

	s.v.Description = &description

	return s
}

// The field to get the date or timestamp from.
func (s *_dateIndexNameProcessor) Field(field string) *_dateIndexNameProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_dateIndexNameProcessor) If(if_ types.ScriptVariant) *_dateIndexNameProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_dateIndexNameProcessor) IgnoreFailure(ignorefailure bool) *_dateIndexNameProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// The format to be used when printing the parsed date into the index name.
// A valid java time pattern is expected here.
// Supports template snippets.
func (s *_dateIndexNameProcessor) IndexNameFormat(indexnameformat string) *_dateIndexNameProcessor {

	s.v.IndexNameFormat = &indexnameformat

	return s
}

// A prefix of the index name to be prepended before the printed date.
// Supports template snippets.
func (s *_dateIndexNameProcessor) IndexNamePrefix(indexnameprefix string) *_dateIndexNameProcessor {

	s.v.IndexNamePrefix = &indexnameprefix

	return s
}

// The locale to use when parsing the date from the document being preprocessed,
// relevant when parsing month names or week days.
func (s *_dateIndexNameProcessor) Locale(locale string) *_dateIndexNameProcessor {

	s.v.Locale = &locale

	return s
}

// Handle failures for the processor.
func (s *_dateIndexNameProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_dateIndexNameProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_dateIndexNameProcessor) Tag(tag string) *_dateIndexNameProcessor {

	s.v.Tag = &tag

	return s
}

// The timezone to use when parsing the date and when date math index supports
// resolves expressions into concrete index names.
func (s *_dateIndexNameProcessor) Timezone(timezone string) *_dateIndexNameProcessor {

	s.v.Timezone = &timezone

	return s
}

func (s *_dateIndexNameProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.DateIndexName = s.v

	return container
}

func (s *_dateIndexNameProcessor) DateIndexNameProcessorCaster() *types.DateIndexNameProcessor {
	return s.v
}
