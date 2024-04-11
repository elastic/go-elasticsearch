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
)

// NodeDiskUsage type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/cluster/allocation_explain/types.ts#L56-L60
type NodeDiskUsage struct {
	LeastAvailable DiskUsage `json:"least_available"`
	MostAvailable  DiskUsage `json:"most_available"`
	NodeName       string    `json:"node_name"`
}

func (s *NodeDiskUsage) UnmarshalJSON(data []byte) error {

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

		case "least_available":
			if err := dec.Decode(&s.LeastAvailable); err != nil {
				return fmt.Errorf("%s | %w", "LeastAvailable", err)
			}

		case "most_available":
			if err := dec.Decode(&s.MostAvailable); err != nil {
				return fmt.Errorf("%s | %w", "MostAvailable", err)
			}

		case "node_name":
			if err := dec.Decode(&s.NodeName); err != nil {
				return fmt.Errorf("%s | %w", "NodeName", err)
			}

		}
	}
	return nil
}

// NewNodeDiskUsage returns a NodeDiskUsage.
func NewNodeDiskUsage() *NodeDiskUsage {
	r := &NodeDiskUsage{}

	return r
}
