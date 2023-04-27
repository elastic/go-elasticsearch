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

	"encoding/json"
)

// ClusterNetworkTypes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/cluster/stats/types.ts#L178-L181
type ClusterNetworkTypes struct {
	HttpTypes      map[string]int `json:"http_types"`
	TransportTypes map[string]int `json:"transport_types"`
}

func (s *ClusterNetworkTypes) UnmarshalJSON(data []byte) error {

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

		case "http_types":
			if s.HttpTypes == nil {
				s.HttpTypes = make(map[string]int, 0)
			}
			if err := dec.Decode(&s.HttpTypes); err != nil {
				return err
			}

		case "transport_types":
			if s.TransportTypes == nil {
				s.TransportTypes = make(map[string]int, 0)
			}
			if err := dec.Decode(&s.TransportTypes); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewClusterNetworkTypes returns a ClusterNetworkTypes.
func NewClusterNetworkTypes() *ClusterNetworkTypes {
	r := &ClusterNetworkTypes{
		HttpTypes:      make(map[string]int, 0),
		TransportTypes: make(map[string]int, 0),
	}

	return r
}
