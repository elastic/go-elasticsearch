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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// NodeInfoXpackMl type.
//
// https://github.com/elastic/elasticsearch-specification/blob/60a81659be928bfe6cec53708c7f7613555a5eaf/specification/nodes/info/types.ts#L253-L255
type NodeInfoXpackMl struct {
	UseAutoMachineMemoryPercent *bool `json:"use_auto_machine_memory_percent,omitempty"`
}

func (s *NodeInfoXpackMl) UnmarshalJSON(data []byte) error {

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

		case "use_auto_machine_memory_percent":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "UseAutoMachineMemoryPercent", err)
				}
				s.UseAutoMachineMemoryPercent = &value
			case bool:
				s.UseAutoMachineMemoryPercent = &v
			}

		}
	}
	return nil
}

// NewNodeInfoXpackMl returns a NodeInfoXpackMl.
func NewNodeInfoXpackMl() *NodeInfoXpackMl {
	r := &NodeInfoXpackMl{}

	return r
}

// false
