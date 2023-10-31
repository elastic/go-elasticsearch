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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// NodeInfoSettingsCluster type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/nodes/info/types.ts#L131-L141
type NodeInfoSettingsCluster struct {
	DeprecationIndexing *DeprecationIndexing            `json:"deprecation_indexing,omitempty"`
	Election            NodeInfoSettingsClusterElection `json:"election"`
	InitialMasterNodes  *string                         `json:"initial_master_nodes,omitempty"`
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
				return err
			}

		case "election":
			if err := dec.Decode(&s.Election); err != nil {
				return err
			}

		case "initial_master_nodes":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.InitialMasterNodes = &o

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return err
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
