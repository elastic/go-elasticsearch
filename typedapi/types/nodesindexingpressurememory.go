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

// NodesIndexingPressureMemory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/nodes/_types/Stats.ts#L123-L142
type NodesIndexingPressureMemory struct {
	// Current Contains statistics for current indexing load.
	Current *PressureMemory `json:"current,omitempty"`
	// Limit Configured memory limit for the indexing requests.
	// Replica requests have an automatic limit that is 1.5x this value.
	Limit ByteSize `json:"limit,omitempty"`
	// LimitInBytes Configured memory limit, in bytes, for the indexing requests.
	// Replica requests have an automatic limit that is 1.5x this value.
	LimitInBytes *int64 `json:"limit_in_bytes,omitempty"`
	// Total Contains statistics for the cumulative indexing load since the node started.
	Total *PressureMemory `json:"total,omitempty"`
}

func (s *NodesIndexingPressureMemory) UnmarshalJSON(data []byte) error {

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

		case "current":
			if err := dec.Decode(&s.Current); err != nil {
				return err
			}

		case "limit":
			if err := dec.Decode(&s.Limit); err != nil {
				return err
			}

		case "limit_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.LimitInBytes = &value
			case float64:
				f := int64(v)
				s.LimitInBytes = &f
			}

		case "total":
			if err := dec.Decode(&s.Total); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewNodesIndexingPressureMemory returns a NodesIndexingPressureMemory.
func NewNodesIndexingPressureMemory() *NodesIndexingPressureMemory {
	r := &NodesIndexingPressureMemory{}

	return r
}
