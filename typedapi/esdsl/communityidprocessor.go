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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _communityIDProcessor struct {
	v *types.CommunityIDProcessor
}

// Computes the Community ID for network flow data as defined in the
// Community ID Specification. You can use a community ID to correlate network
// events related to a single flow.
func NewCommunityIDProcessor() *_communityIDProcessor {

	return &_communityIDProcessor{v: types.NewCommunityIDProcessor()}

}

func (s *_communityIDProcessor) Description(description string) *_communityIDProcessor {

	s.v.Description = &description

	return s
}

func (s *_communityIDProcessor) DestinationIp(field string) *_communityIDProcessor {

	s.v.DestinationIp = &field

	return s
}

func (s *_communityIDProcessor) DestinationPort(field string) *_communityIDProcessor {

	s.v.DestinationPort = &field

	return s
}

func (s *_communityIDProcessor) IanaNumber(field string) *_communityIDProcessor {

	s.v.IanaNumber = &field

	return s
}

func (s *_communityIDProcessor) IcmpCode(field string) *_communityIDProcessor {

	s.v.IcmpCode = &field

	return s
}

func (s *_communityIDProcessor) IcmpType(field string) *_communityIDProcessor {

	s.v.IcmpType = &field

	return s
}

func (s *_communityIDProcessor) If(if_ types.ScriptVariant) *_communityIDProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_communityIDProcessor) IgnoreFailure(ignorefailure bool) *_communityIDProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_communityIDProcessor) IgnoreMissing(ignoremissing bool) *_communityIDProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_communityIDProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_communityIDProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_communityIDProcessor) Seed(seed int) *_communityIDProcessor {

	s.v.Seed = &seed

	return s
}

func (s *_communityIDProcessor) SourceIp(field string) *_communityIDProcessor {

	s.v.SourceIp = &field

	return s
}

func (s *_communityIDProcessor) SourcePort(field string) *_communityIDProcessor {

	s.v.SourcePort = &field

	return s
}

func (s *_communityIDProcessor) Tag(tag string) *_communityIDProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_communityIDProcessor) TargetField(field string) *_communityIDProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_communityIDProcessor) Transport(field string) *_communityIDProcessor {

	s.v.Transport = &field

	return s
}

func (s *_communityIDProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.CommunityId = s.v

	return container
}

func (s *_communityIDProcessor) CommunityIDProcessorCaster() *types.CommunityIDProcessor {
	return s.v
}
