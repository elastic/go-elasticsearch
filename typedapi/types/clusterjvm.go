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

// ClusterJvm type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cluster/stats/types.ts#L275-L292
type ClusterJvm struct {
	// MaxUptimeInMillis Uptime duration, in milliseconds, since JVM last started.
	MaxUptimeInMillis int64 `json:"max_uptime_in_millis"`
	// Mem Contains statistics about memory used by selected nodes.
	Mem ClusterJvmMemory `json:"mem"`
	// Threads Number of active threads in use by JVM across all selected nodes.
	Threads int64 `json:"threads"`
	// Versions Contains statistics about the JVM versions used by selected nodes.
	Versions []ClusterJvmVersion `json:"versions"`
}

func (s *ClusterJvm) UnmarshalJSON(data []byte) error {

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

		case "max_uptime_in_millis":
			if err := dec.Decode(&s.MaxUptimeInMillis); err != nil {
				return err
			}

		case "mem":
			if err := dec.Decode(&s.Mem); err != nil {
				return err
			}

		case "threads":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Threads = value
			case float64:
				f := int64(v)
				s.Threads = f
			}

		case "versions":
			if err := dec.Decode(&s.Versions); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewClusterJvm returns a ClusterJvm.
func NewClusterJvm() *ClusterJvm {
	r := &ClusterJvm{}

	return r
}
