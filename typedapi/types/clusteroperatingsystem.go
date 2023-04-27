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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// ClusterOperatingSystem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/cluster/stats/types.ts#L235-L242
type ClusterOperatingSystem struct {
	AllocatedProcessors int                                  `json:"allocated_processors"`
	Architectures       []ClusterOperatingSystemArchitecture `json:"architectures,omitempty"`
	AvailableProcessors int                                  `json:"available_processors"`
	Mem                 OperatingSystemMemoryInfo            `json:"mem"`
	Names               []ClusterOperatingSystemName         `json:"names"`
	PrettyNames         []ClusterOperatingSystemPrettyName   `json:"pretty_names"`
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
