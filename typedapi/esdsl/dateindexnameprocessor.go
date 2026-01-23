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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

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

func (s *_dateIndexNameProcessor) DateFormats(dateformats ...string) *_dateIndexNameProcessor {

	for _, v := range dateformats {

		s.v.DateFormats = append(s.v.DateFormats, v)

	}
	return s
}

func (s *_dateIndexNameProcessor) DateRounding(daterounding string) *_dateIndexNameProcessor {

	s.v.DateRounding = daterounding

	return s
}

func (s *_dateIndexNameProcessor) Field(field string) *_dateIndexNameProcessor {

	s.v.Field = field

	return s
}

func (s *_dateIndexNameProcessor) IndexNameFormat(indexnameformat string) *_dateIndexNameProcessor {

	s.v.IndexNameFormat = &indexnameformat

	return s
}

func (s *_dateIndexNameProcessor) IndexNamePrefix(indexnameprefix string) *_dateIndexNameProcessor {

	s.v.IndexNamePrefix = &indexnameprefix

	return s
}

func (s *_dateIndexNameProcessor) Locale(locale string) *_dateIndexNameProcessor {

	s.v.Locale = &locale

	return s
}

func (s *_dateIndexNameProcessor) Timezone(timezone string) *_dateIndexNameProcessor {

	s.v.Timezone = &timezone

	return s
}

func (s *_dateIndexNameProcessor) Description(description string) *_dateIndexNameProcessor {

	s.v.Description = &description

	return s
}

func (s *_dateIndexNameProcessor) If(if_ types.ScriptVariant) *_dateIndexNameProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_dateIndexNameProcessor) IgnoreFailure(ignorefailure bool) *_dateIndexNameProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_dateIndexNameProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_dateIndexNameProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_dateIndexNameProcessor) Tag(tag string) *_dateIndexNameProcessor {

	s.v.Tag = &tag

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
