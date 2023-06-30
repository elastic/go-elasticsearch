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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/decision"
)

// NodeAllocationExplanation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/cluster/allocation_explain/types.ts#L97-L106
type NodeAllocationExplanation struct {
	Deciders         []AllocationDecision `json:"deciders"`
	NodeAttributes   map[string]string    `json:"node_attributes"`
	NodeDecision     decision.Decision    `json:"node_decision"`
	NodeId           string               `json:"node_id"`
	NodeName         string               `json:"node_name"`
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
				return err
			}

		case "node_attributes":
			if s.NodeAttributes == nil {
				s.NodeAttributes = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.NodeAttributes); err != nil {
				return err
			}

		case "node_decision":
			if err := dec.Decode(&s.NodeDecision); err != nil {
				return err
			}

		case "node_id":
			if err := dec.Decode(&s.NodeId); err != nil {
				return err
			}

		case "node_name":
			if err := dec.Decode(&s.NodeName); err != nil {
				return err
			}

		case "store":
			if err := dec.Decode(&s.Store); err != nil {
				return err
			}

		case "transport_address":
			if err := dec.Decode(&s.TransportAddress); err != nil {
				return err
			}

		case "weight_ranking":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
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
		NodeAttributes: make(map[string]string, 0),
	}

	return r
}
