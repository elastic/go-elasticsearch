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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// RoleDescriptorRead type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/security/_types/RoleDescriptor.ts#L38-L47
type RoleDescriptorRead struct {
	Applications      []ApplicationPrivileges  `json:"applications,omitempty"`
	Cluster           []string                 `json:"cluster"`
	Global            []GlobalPrivilege        `json:"global,omitempty"`
	Indices           []IndicesPrivileges      `json:"indices"`
	Metadata          Metadata                 `json:"metadata,omitempty"`
	RunAs             []string                 `json:"run_as,omitempty"`
	TransientMetadata *TransientMetadataConfig `json:"transient_metadata,omitempty"`
}

func (s *RoleDescriptorRead) UnmarshalJSON(data []byte) error {

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

		case "applications":
			if err := dec.Decode(&s.Applications); err != nil {
				return err
			}

		case "cluster":
			if err := dec.Decode(&s.Cluster); err != nil {
				return err
			}

		case "global":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewGlobalPrivilege()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Global = append(s.Global, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Global); err != nil {
					return err
				}
			}

		case "indices", "index":
			if err := dec.Decode(&s.Indices); err != nil {
				return err
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return err
			}

		case "run_as":
			if err := dec.Decode(&s.RunAs); err != nil {
				return err
			}

		case "transient_metadata":
			if err := dec.Decode(&s.TransientMetadata); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewRoleDescriptorRead returns a RoleDescriptorRead.
func NewRoleDescriptorRead() *RoleDescriptorRead {
	r := &RoleDescriptorRead{}

	return r
}
