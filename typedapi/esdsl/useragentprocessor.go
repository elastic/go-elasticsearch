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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/useragentproperty"
)

type _userAgentProcessor struct {
	v *types.UserAgentProcessor
}

// The `user_agent` processor extracts details from the user agent string a
// browser sends with its web requests.
// This processor adds this information by default under the `user_agent` field.
func NewUserAgentProcessor() *_userAgentProcessor {

	return &_userAgentProcessor{v: types.NewUserAgentProcessor()}

}

func (s *_userAgentProcessor) Description(description string) *_userAgentProcessor {

	s.v.Description = &description

	return s
}

func (s *_userAgentProcessor) ExtractDeviceType(extractdevicetype bool) *_userAgentProcessor {

	s.v.ExtractDeviceType = &extractdevicetype

	return s
}

func (s *_userAgentProcessor) Field(field string) *_userAgentProcessor {

	s.v.Field = field

	return s
}

func (s *_userAgentProcessor) If(if_ types.ScriptVariant) *_userAgentProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_userAgentProcessor) IgnoreFailure(ignorefailure bool) *_userAgentProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_userAgentProcessor) IgnoreMissing(ignoremissing bool) *_userAgentProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_userAgentProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_userAgentProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_userAgentProcessor) Properties(properties ...useragentproperty.UserAgentProperty) *_userAgentProcessor {

	for _, v := range properties {

		s.v.Properties = append(s.v.Properties, v)

	}
	return s
}

func (s *_userAgentProcessor) RegexFile(regexfile string) *_userAgentProcessor {

	s.v.RegexFile = &regexfile

	return s
}

func (s *_userAgentProcessor) Tag(tag string) *_userAgentProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_userAgentProcessor) TargetField(field string) *_userAgentProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_userAgentProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.UserAgent = s.v

	return container
}

func (s *_userAgentProcessor) UserAgentProcessorCaster() *types.UserAgentProcessor {
	return s.v
}
