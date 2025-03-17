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

type _communityIDProcessor struct {
	v *types.CommunityIDProcessor
}

// Computes the Community ID for network flow data as defined in the
// Community ID Specification. You can use a community ID to correlate network
// events related to a single flow.
func NewCommunityIDProcessor() *_communityIDProcessor {

	return &_communityIDProcessor{v: types.NewCommunityIDProcessor()}

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_communityIDProcessor) Description(description string) *_communityIDProcessor {

	s.v.Description = &description

	return s
}

// Field containing the destination IP address.
func (s *_communityIDProcessor) DestinationIp(field string) *_communityIDProcessor {

	s.v.DestinationIp = &field

	return s
}

// Field containing the destination port.
func (s *_communityIDProcessor) DestinationPort(field string) *_communityIDProcessor {

	s.v.DestinationPort = &field

	return s
}

// Field containing the IANA number.
func (s *_communityIDProcessor) IanaNumber(field string) *_communityIDProcessor {

	s.v.IanaNumber = &field

	return s
}

// Field containing the ICMP code.
func (s *_communityIDProcessor) IcmpCode(field string) *_communityIDProcessor {

	s.v.IcmpCode = &field

	return s
}

// Field containing the ICMP type.
func (s *_communityIDProcessor) IcmpType(field string) *_communityIDProcessor {

	s.v.IcmpType = &field

	return s
}

// Conditionally execute the processor.
func (s *_communityIDProcessor) If(if_ types.ScriptVariant) *_communityIDProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_communityIDProcessor) IgnoreFailure(ignorefailure bool) *_communityIDProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If true and any required fields are missing, the processor quietly exits
// without modifying the document.
func (s *_communityIDProcessor) IgnoreMissing(ignoremissing bool) *_communityIDProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// Handle failures for the processor.
func (s *_communityIDProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_communityIDProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Seed for the community ID hash. Must be between 0 and 65535 (inclusive). The
// seed can prevent hash collisions between network domains, such as a staging
// and production network that use the same addressing scheme.
func (s *_communityIDProcessor) Seed(seed int) *_communityIDProcessor {

	s.v.Seed = &seed

	return s
}

// Field containing the source IP address.
func (s *_communityIDProcessor) SourceIp(field string) *_communityIDProcessor {

	s.v.SourceIp = &field

	return s
}

// Field containing the source port.
func (s *_communityIDProcessor) SourcePort(field string) *_communityIDProcessor {

	s.v.SourcePort = &field

	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_communityIDProcessor) Tag(tag string) *_communityIDProcessor {

	s.v.Tag = &tag

	return s
}

// Output field for the community ID.
func (s *_communityIDProcessor) TargetField(field string) *_communityIDProcessor {

	s.v.TargetField = &field

	return s
}

// Field containing the transport protocol name or number. Used only when the
// iana_number field is not present. The following protocol names are currently
// supported: eigrp, gre, icmp, icmpv6, igmp, ipv6-icmp, ospf, pim, sctp, tcp,
// udp
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
