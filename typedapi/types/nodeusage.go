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
)

// NodeUsage type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/nodes/usage/types.ts#L25-L30
type NodeUsage struct {
	Aggregations map[string]json.RawMessage `json:"aggregations"`
	RestActions  map[string]int             `json:"rest_actions"`
	Since        int64                      `json:"since"`
	Timestamp    int64                      `json:"timestamp"`
}

func (s *NodeUsage) UnmarshalJSON(data []byte) error {

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

		case "aggregations":
			if s.Aggregations == nil {
				s.Aggregations = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Aggregations); err != nil {
				return err
			}

		case "rest_actions":
			if s.RestActions == nil {
				s.RestActions = make(map[string]int, 0)
			}
			if err := dec.Decode(&s.RestActions); err != nil {
				return err
			}

		case "since":
			if err := dec.Decode(&s.Since); err != nil {
				return err
			}

		case "timestamp":
			if err := dec.Decode(&s.Timestamp); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewNodeUsage returns a NodeUsage.
func NewNodeUsage() *NodeUsage {
	r := &NodeUsage{
		Aggregations: make(map[string]json.RawMessage, 0),
		RestActions:  make(map[string]int, 0),
	}

	return r
}
