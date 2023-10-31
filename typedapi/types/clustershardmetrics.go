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

// ClusterShardMetrics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cluster/stats/types.ts#L511-L524
type ClusterShardMetrics struct {
	// Avg Mean number of shards in an index, counting only shards assigned to selected
	// nodes.
	Avg Float64 `json:"avg"`
	// Max Maximum number of shards in an index, counting only shards assigned to
	// selected nodes.
	Max Float64 `json:"max"`
	// Min Minimum number of shards in an index, counting only shards assigned to
	// selected nodes.
	Min Float64 `json:"min"`
}

func (s *ClusterShardMetrics) UnmarshalJSON(data []byte) error {

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

		case "avg":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Avg = f
			case float64:
				f := Float64(v)
				s.Avg = f
			}

		case "max":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Max = f
			case float64:
				f := Float64(v)
				s.Max = f
			}

		case "min":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Min = f
			case float64:
				f := Float64(v)
				s.Min = f
			}

		}
	}
	return nil
}

// NewClusterShardMetrics returns a ClusterShardMetrics.
func NewClusterShardMetrics() *ClusterShardMetrics {
	r := &ClusterShardMetrics{}

	return r
}
