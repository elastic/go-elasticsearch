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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// ClusterIndicesShards type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cluster/stats/types.ts#L60-L72
type ClusterIndicesShards struct {
	// Index Contains statistics about shards assigned to selected nodes.
	Index *ClusterIndicesShardsIndex `json:"index,omitempty"`
	// Primaries Number of primary shards assigned to selected nodes.
	Primaries *Float64 `json:"primaries,omitempty"`
	// Replication Ratio of replica shards to primary shards across all selected nodes.
	Replication *Float64 `json:"replication,omitempty"`
	// Total Total number of shards assigned to selected nodes.
	Total *Float64 `json:"total,omitempty"`
}

func (s *ClusterIndicesShards) UnmarshalJSON(data []byte) error {

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

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "primaries":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Primaries = &f
			case float64:
				f := Float64(v)
				s.Primaries = &f
			}

		case "replication":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Replication = &f
			case float64:
				f := Float64(v)
				s.Replication = &f
			}

		case "total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Total = &f
			case float64:
				f := Float64(v)
				s.Total = &f
			}

		}
	}
	return nil
}

// NewClusterIndicesShards returns a ClusterIndicesShards.
func NewClusterIndicesShards() *ClusterIndicesShards {
	r := &ClusterIndicesShards{}

	return r
}
