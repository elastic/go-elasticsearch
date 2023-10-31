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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/shardroutingstate"
)

// NodeShard type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/Node.ts#L60-L71
type NodeShard struct {
	AllocationId          map[string]string                   `json:"allocation_id,omitempty"`
	Index                 string                              `json:"index"`
	Node                  *string                             `json:"node,omitempty"`
	Primary               bool                                `json:"primary"`
	RecoverySource        map[string]string                   `json:"recovery_source,omitempty"`
	RelocatingNode        string                              `json:"relocating_node,omitempty"`
	RelocationFailureInfo *RelocationFailureInfo              `json:"relocation_failure_info,omitempty"`
	Shard                 int                                 `json:"shard"`
	State                 shardroutingstate.ShardRoutingState `json:"state"`
	UnassignedInfo        *UnassignedInformation              `json:"unassigned_info,omitempty"`
}

func (s *NodeShard) UnmarshalJSON(data []byte) error {

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

		case "allocation_id":
			if s.AllocationId == nil {
				s.AllocationId = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.AllocationId); err != nil {
				return err
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "node":
			if err := dec.Decode(&s.Node); err != nil {
				return err
			}

		case "primary":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Primary = value
			case bool:
				s.Primary = v
			}

		case "recovery_source":
			if s.RecoverySource == nil {
				s.RecoverySource = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.RecoverySource); err != nil {
				return err
			}

		case "relocating_node":
			if err := dec.Decode(&s.RelocatingNode); err != nil {
				return err
			}

		case "relocation_failure_info":
			if err := dec.Decode(&s.RelocationFailureInfo); err != nil {
				return err
			}

		case "shard":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Shard = value
			case float64:
				f := int(v)
				s.Shard = f
			}

		case "state":
			if err := dec.Decode(&s.State); err != nil {
				return err
			}

		case "unassigned_info":
			if err := dec.Decode(&s.UnassignedInfo); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewNodeShard returns a NodeShard.
func NewNodeShard() *NodeShard {
	r := &NodeShard{
		AllocationId:   make(map[string]string, 0),
		RecoverySource: make(map[string]string, 0),
	}

	return r
}
