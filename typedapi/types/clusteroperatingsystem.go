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

// ClusterOperatingSystem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cluster/stats/types.ts#L415-L442
type ClusterOperatingSystem struct {
	// AllocatedProcessors Number of processors used to calculate thread pool size across all selected
	// nodes.
	// This number can be set with the processors setting of a node and defaults to
	// the number of processors reported by the operating system.
	// In both cases, this number will never be larger than 32.
	AllocatedProcessors int `json:"allocated_processors"`
	// Architectures Contains statistics about processor architectures (for example, x86_64 or
	// aarch64) used by selected nodes.
	Architectures []ClusterOperatingSystemArchitecture `json:"architectures,omitempty"`
	// AvailableProcessors Number of processors available to JVM across all selected nodes.
	AvailableProcessors int `json:"available_processors"`
	// Mem Contains statistics about memory used by selected nodes.
	Mem OperatingSystemMemoryInfo `json:"mem"`
	// Names Contains statistics about operating systems used by selected nodes.
	Names []ClusterOperatingSystemName `json:"names"`
	// PrettyNames Contains statistics about operating systems used by selected nodes.
	PrettyNames []ClusterOperatingSystemPrettyName `json:"pretty_names"`
}

func (s *ClusterOperatingSystem) UnmarshalJSON(data []byte) error {

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

		case "allocated_processors":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.AllocatedProcessors = value
			case float64:
				f := int(v)
				s.AllocatedProcessors = f
			}

		case "architectures":
			if err := dec.Decode(&s.Architectures); err != nil {
				return err
			}

		case "available_processors":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.AvailableProcessors = value
			case float64:
				f := int(v)
				s.AvailableProcessors = f
			}

		case "mem":
			if err := dec.Decode(&s.Mem); err != nil {
				return err
			}

		case "names":
			if err := dec.Decode(&s.Names); err != nil {
				return err
			}

		case "pretty_names":
			if err := dec.Decode(&s.PrettyNames); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewClusterOperatingSystem returns a ClusterOperatingSystem.
func NewClusterOperatingSystem() *ClusterOperatingSystem {
	r := &ClusterOperatingSystem{}

	return r
}
