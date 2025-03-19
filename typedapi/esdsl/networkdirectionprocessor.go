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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _networkDirectionProcessor struct {
	v *types.NetworkDirectionProcessor
}

// Calculates the network direction given a source IP address, destination IP
// address, and a list of internal networks.
func NewNetworkDirectionProcessor() *_networkDirectionProcessor {

	return &_networkDirectionProcessor{v: types.NewNetworkDirectionProcessor()}

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_networkDirectionProcessor) Description(description string) *_networkDirectionProcessor {

	s.v.Description = &description

	return s
}

// Field containing the destination IP address.
func (s *_networkDirectionProcessor) DestinationIp(field string) *_networkDirectionProcessor {

	s.v.DestinationIp = &field

	return s
}

// Conditionally execute the processor.
func (s *_networkDirectionProcessor) If(if_ types.ScriptVariant) *_networkDirectionProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_networkDirectionProcessor) IgnoreFailure(ignorefailure bool) *_networkDirectionProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If true and any required fields are missing, the processor quietly exits
// without modifying the document.
func (s *_networkDirectionProcessor) IgnoreMissing(ignoremissing bool) *_networkDirectionProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// List of internal networks. Supports IPv4 and IPv6 addresses and ranges in
// CIDR notation. Also supports the named ranges listed below. These may be
// constructed with template snippets. Must specify only one of
// internal_networks or internal_networks_field.
func (s *_networkDirectionProcessor) InternalNetworks(internalnetworks ...string) *_networkDirectionProcessor {

	for _, v := range internalnetworks {

		s.v.InternalNetworks = append(s.v.InternalNetworks, v)

	}
	return s
}

// A field on the given document to read the internal_networks configuration
// from.
func (s *_networkDirectionProcessor) InternalNetworksField(field string) *_networkDirectionProcessor {

	s.v.InternalNetworksField = &field

	return s
}

// Handle failures for the processor.
func (s *_networkDirectionProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_networkDirectionProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Field containing the source IP address.
func (s *_networkDirectionProcessor) SourceIp(field string) *_networkDirectionProcessor {

	s.v.SourceIp = &field

	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_networkDirectionProcessor) Tag(tag string) *_networkDirectionProcessor {

	s.v.Tag = &tag

	return s
}

// Output field for the network direction.
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
