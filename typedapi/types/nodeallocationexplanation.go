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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/decision"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/noderole"
)

// NodeAllocationExplanation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/allocation_explain/types.ts#L103-L117
type NodeAllocationExplanation struct {
	Deciders         []AllocationDecision `json:"deciders"`
	NodeAttributes   map[string]string    `json:"node_attributes"`
	NodeDecision     decision.Decision    `json:"node_decision"`
	NodeId           string               `json:"node_id"`
	NodeName         string               `json:"node_name"`
	Roles            []noderole.NodeRole  `json:"roles"`
	Store            *AllocationStore     `json:"store,omitempty"`
	TransportAddress string               `json:"transport_address"`
	WeightRanking    int                  `json:"weight_ranking"`
}

func (s *NodeAllocationExplanation) UnmarshalJSON(data []byte) error {

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

		case "deciders":
			if err := dec.Decode(&s.Deciders); err != nil {
				return fmt.Errorf("%s | %w", "Deciders", err)
			}

		case "node_attributes":
			if s.NodeAttributes == nil {
				s.NodeAttributes = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.NodeAttributes); err != nil {
				return fmt.Errorf("%s | %w", "NodeAttributes", err)
			}

		case "node_decision":
			if err := dec.Decode(&s.NodeDecision); err != nil {
				return fmt.Errorf("%s | %w", "NodeDecision", err)
			}

		case "node_id":
			if err := dec.Decode(&s.NodeId); err != nil {
				return fmt.Errorf("%s | %w", "NodeId", err)
			}

		case "node_name":
			if err := dec.Decode(&s.NodeName); err != nil {
				return fmt.Errorf("%s | %w", "NodeName", err)
			}

		case "roles":
			if err := dec.Decode(&s.Roles); err != nil {
				return fmt.Errorf("%s | %w", "Roles", err)
			}

		case "store":
			if err := dec.Decode(&s.Store); err != nil {
				return fmt.Errorf("%s | %w", "Store", err)
			}

		case "transport_address":
			if err := dec.Decode(&s.TransportAddress); err != nil {
				return fmt.Errorf("%s | %w", "TransportAddress", err)
			}

		case "weight_ranking":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "WeightRanking", err)
				}
				s.WeightRanking = value
			case float64:
				f := int(v)
				s.WeightRanking = f
			}

		}
	}
	return nil
}

// NewNodeAllocationExplanation returns a NodeAllocationExplanation.
func NewNodeAllocationExplanation() *NodeAllocationExplanation {
	r := &NodeAllocationExplanation{
		NodeAttributes: make(map[string]string),
	}

	return r
}
