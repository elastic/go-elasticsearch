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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// CommunityIDProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ingest/_types/Processors.ts#L598-L659
type CommunityIDProcessor struct {
	// Description Description of the processor.
	// Useful for describing the purpose of the processor or its configuration.
	Description *string `json:"description,omitempty"`
	// DestinationIp Field containing the destination IP address.
	DestinationIp *string `json:"destination_ip,omitempty"`
	// DestinationPort Field containing the destination port.
	DestinationPort *string `json:"destination_port,omitempty"`
	// IanaNumber Field containing the IANA number.
	IanaNumber *string `json:"iana_number,omitempty"`
	// IcmpCode Field containing the ICMP code.
	IcmpCode *string `json:"icmp_code,omitempty"`
	// IcmpType Field containing the ICMP type.
	IcmpType *string `json:"icmp_type,omitempty"`
	// If Conditionally execute the processor.
	If *string `json:"if,omitempty"`
	// IgnoreFailure Ignore failures for the processor.
	IgnoreFailure *bool `json:"ignore_failure,omitempty"`
	// IgnoreMissing If true and any required fields are missing, the processor quietly exits
	// without modifying the document.
	IgnoreMissing *bool `json:"ignore_missing,omitempty"`
	// OnFailure Handle failures for the processor.
	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`
	// Seed Seed for the community ID hash. Must be between 0 and 65535 (inclusive). The
	// seed can prevent hash collisions between network domains, such as a staging
	// and production network that use the same addressing scheme.
	Seed *int `json:"seed,omitempty"`
	// SourceIp Field containing the source IP address.
	SourceIp *string `json:"source_ip,omitempty"`
	// SourcePort Field containing the source port.
	SourcePort *string `json:"source_port,omitempty"`
	// Tag Identifier for the processor.
	// Useful for debugging and metrics.
	Tag *string `json:"tag,omitempty"`
	// TargetField Output field for the community ID.
	TargetField *string `json:"target_field,omitempty"`
	// Transport Field containing the transport protocol name or number. Used only when the
	// iana_number field is not present. The following protocol names are currently
	// supported: eigrp, gre, icmp, icmpv6, igmp, ipv6-icmp, ospf, pim, sctp, tcp,
	// udp
	Transport *string `json:"transport,omitempty"`
}

func (s *CommunityIDProcessor) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "destination_ip":
			if err := dec.Decode(&s.DestinationIp); err != nil {
				return fmt.Errorf("%s | %w", "DestinationIp", err)
			}

		case "destination_port":
			if err := dec.Decode(&s.DestinationPort); err != nil {
				return fmt.Errorf("%s | %w", "DestinationPort", err)
			}

		case "iana_number":
			if err := dec.Decode(&s.IanaNumber); err != nil {
				return fmt.Errorf("%s | %w", "IanaNumber", err)
			}

		case "icmp_code":
			if err := dec.Decode(&s.IcmpCode); err != nil {
				return fmt.Errorf("%s | %w", "IcmpCode", err)
			}

		case "icmp_type":
			if err := dec.Decode(&s.IcmpType); err != nil {
				return fmt.Errorf("%s | %w", "IcmpType", err)
			}

		case "if":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "If", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.If = &o

		case "ignore_failure":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreFailure", err)
				}
				s.IgnoreFailure = &value
			case bool:
				s.IgnoreFailure = &v
			}

		case "ignore_missing":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreMissing", err)
				}
				s.IgnoreMissing = &value
			case bool:
				s.IgnoreMissing = &v
			}

		case "on_failure":
			if err := dec.Decode(&s.OnFailure); err != nil {
				return fmt.Errorf("%s | %w", "OnFailure", err)
			}

		case "seed":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Seed", err)
				}
				s.Seed = &value
			case float64:
				f := int(v)
				s.Seed = &f
			}

		case "source_ip":
			if err := dec.Decode(&s.SourceIp); err != nil {
				return fmt.Errorf("%s | %w", "SourceIp", err)
			}

		case "source_port":
			if err := dec.Decode(&s.SourcePort); err != nil {
				return fmt.Errorf("%s | %w", "SourcePort", err)
			}

		case "tag":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Tag", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Tag = &o

		case "target_field":
			if err := dec.Decode(&s.TargetField); err != nil {
				return fmt.Errorf("%s | %w", "TargetField", err)
			}

		case "transport":
			if err := dec.Decode(&s.Transport); err != nil {
				return fmt.Errorf("%s | %w", "Transport", err)
			}

		}
	}
	return nil
}

// NewCommunityIDProcessor returns a CommunityIDProcessor.
func NewCommunityIDProcessor() *CommunityIDProcessor {
	r := &CommunityIDProcessor{}

	return r
}
