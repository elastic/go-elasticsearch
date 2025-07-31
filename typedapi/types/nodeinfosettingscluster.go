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
)

// NodeInfoSettingsCluster type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/nodes/info/types.ts#L142-L152
type NodeInfoSettingsCluster struct {
	DeprecationIndexing *DeprecationIndexing            `json:"deprecation_indexing,omitempty"`
	Election            NodeInfoSettingsClusterElection `json:"election"`
	InitialMasterNodes  []string                        `json:"initial_master_nodes,omitempty"`
	Name                string                          `json:"name"`
	Routing             *IndexRouting                   `json:"routing,omitempty"`
}

func (s *NodeInfoSettingsCluster) UnmarshalJSON(data []byte) error {

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

		case "deprecation_indexing":
			if err := dec.Decode(&s.DeprecationIndexing); err != nil {
				return fmt.Errorf("%s | %w", "DeprecationIndexing", err)
			}

		case "election":
			if err := dec.Decode(&s.Election); err != nil {
				return fmt.Errorf("%s | %w", "Election", err)
			}

		case "initial_master_nodes":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "InitialMasterNodes", err)
				}

				s.InitialMasterNodes = append(s.InitialMasterNodes, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.InitialMasterNodes); err != nil {
					return fmt.Errorf("%s | %w", "InitialMasterNodes", err)
				}
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return fmt.Errorf("%s | %w", "Routing", err)
			}

		}
	}
	return nil
}

// NewNodeInfoSettingsCluster returns a NodeInfoSettingsCluster.
func NewNodeInfoSettingsCluster() *NodeInfoSettingsCluster {
	r := &NodeInfoSettingsCluster{}

	return r
}
