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

// ClusterFileSystem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cluster/stats/types.ts#L34-L49
type ClusterFileSystem struct {
	// AvailableInBytes Total number of bytes available to JVM in file stores across all selected
	// nodes.
	// Depending on operating system or process-level restrictions, this number may
	// be less than `nodes.fs.free_in_byes`.
	// This is the actual amount of free disk space the selected Elasticsearch nodes
	// can use.
	AvailableInBytes int64 `json:"available_in_bytes"`
	// FreeInBytes Total number of unallocated bytes in file stores across all selected nodes.
	FreeInBytes int64 `json:"free_in_bytes"`
	// TotalInBytes Total size, in bytes, of all file stores across all selected nodes.
	TotalInBytes int64 `json:"total_in_bytes"`
}

func (s *ClusterFileSystem) UnmarshalJSON(data []byte) error {

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

		case "available_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.AvailableInBytes = value
			case float64:
				f := int64(v)
				s.AvailableInBytes = f
			}

		case "free_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.FreeInBytes = value
			case float64:
				f := int64(v)
				s.FreeInBytes = f
			}

		case "total_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalInBytes = value
			case float64:
				f := int64(v)
				s.TotalInBytes = f
			}

		}
	}
	return nil
}

// NewClusterFileSystem returns a ClusterFileSystem.
func NewClusterFileSystem() *ClusterFileSystem {
	r := &ClusterFileSystem{}

	return r
}
