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
// https://github.com/elastic/elasticsearch-specification/tree/836fca874204ca4173ae5c36fb6b5107d28d2fc0

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// GPU vector indexing usage statistics.
//
// https://github.com/elastic/elasticsearch-specification/blob/836fca874204ca4173ae5c36fb6b5107d28d2fc0/specification/xpack/usage/types.ts#L524-L534
type GpuVectorIndexing struct {
	Available bool `json:"available"`
	Enabled   bool `json:"enabled"`
	// IndexBuildCount Total GPU index builds across the cluster.
	IndexBuildCount int64 `json:"index_build_count"`
	// Nodes Per-node GPU details including type, memory, enabled status, and build count.
	Nodes []GpuNodeStats `json:"nodes"`
	// NodesWithGpu Count of data nodes with GPU support.
	NodesWithGpu int `json:"nodes_with_gpu"`
}

func (s *GpuVectorIndexing) UnmarshalJSON(data []byte) error {

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

		case "available":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Available", err)
				}
				s.Available = value
			case bool:
				s.Available = v
			}

		case "enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Enabled", err)
				}
				s.Enabled = value
			case bool:
				s.Enabled = v
			}

		case "index_build_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndexBuildCount", err)
				}
				s.IndexBuildCount = value
			case float64:
				f := int64(v)
				s.IndexBuildCount = f
			}

		case "nodes":
			if err := dec.Decode(&s.Nodes); err != nil {
				return fmt.Errorf("%s | %w", "Nodes", err)
			}

		case "nodes_with_gpu":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NodesWithGpu", err)
				}
				s.NodesWithGpu = value
			case float64:
				f := int(v)
				s.NodesWithGpu = f
			}

		}
	}
	return nil
}

// NewGpuVectorIndexing returns a GpuVectorIndexing.
func NewGpuVectorIndexing() *GpuVectorIndexing {
	r := &GpuVectorIndexing{}

	return r
}
