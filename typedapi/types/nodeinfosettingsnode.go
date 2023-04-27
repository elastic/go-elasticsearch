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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// NodeInfoSettingsNode type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/nodes/info/types.ts#L148-L152
type NodeInfoSettingsNode struct {
	Attr                 map[string]json.RawMessage `json:"attr"`
	MaxLocalStorageNodes *string                    `json:"max_local_storage_nodes,omitempty"`
	Name                 string                     `json:"name"`
}

func (s *NodeInfoSettingsNode) UnmarshalJSON(data []byte) error {

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

		case "attr":
			if s.Attr == nil {
				s.Attr = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Attr); err != nil {
				return err
			}

		case "max_local_storage_nodes":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.MaxLocalStorageNodes = &o

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewNodeInfoSettingsNode returns a NodeInfoSettingsNode.
func NewNodeInfoSettingsNode() *NodeInfoSettingsNode {
	r := &NodeInfoSettingsNode{
		Attr: make(map[string]json.RawMessage, 0),
	}

	return r
}
