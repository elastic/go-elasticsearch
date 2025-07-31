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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _networkDirectionProcessor struct {
	v *types.NetworkDirectionProcessor
}

// Calculates the network direction given a source IP address, destination IP
// address, and a list of internal networks.
func NewNetworkDirectionProcessor() *_networkDirectionProcessor {

	return &_networkDirectionProcessor{v: types.NewNetworkDirectionProcessor()}

}

func (s *_networkDirectionProcessor) Description(description string) *_networkDirectionProcessor {

	s.v.Description = &description

	return s
}

func (s *_networkDirectionProcessor) DestinationIp(field string) *_networkDirectionProcessor {

	s.v.DestinationIp = &field

	return s
}

func (s *_networkDirectionProcessor) If(if_ types.ScriptVariant) *_networkDirectionProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_networkDirectionProcessor) IgnoreFailure(ignorefailure bool) *_networkDirectionProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_networkDirectionProcessor) IgnoreMissing(ignoremissing bool) *_networkDirectionProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_networkDirectionProcessor) InternalNetworks(internalnetworks ...string) *_networkDirectionProcessor {

	for _, v := range internalnetworks {

		s.v.InternalNetworks = append(s.v.InternalNetworks, v)

	}
	return s
}

func (s *_networkDirectionProcessor) InternalNetworksField(field string) *_networkDirectionProcessor {

	s.v.InternalNetworksField = &field

	return s
}

func (s *_networkDirectionProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_networkDirectionProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_networkDirectionProcessor) SourceIp(field string) *_networkDirectionProcessor {

	s.v.SourceIp = &field

	return s
}

func (s *_networkDirectionProcessor) Tag(tag string) *_networkDirectionProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_networkDirectionProcessor) TargetField(field string) *_networkDirectionProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_networkDirectionProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.NetworkDirection = s.v

	return container
}

func (s *_networkDirectionProcessor) NetworkDirectionProcessorCaster() *types.NetworkDirectionProcessor {
	return s.v
}
