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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// NodeInfoSettingsNode type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/nodes/info/types.ts#L152-L156
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
				return fmt.Errorf("%s | %w", "Attr", err)
			}

		case "max_local_storage_nodes":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MaxLocalStorageNodes", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MaxLocalStorageNodes = &o

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
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
