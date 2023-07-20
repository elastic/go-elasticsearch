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
// https://github.com/elastic/elasticsearch-specification/tree/656080dee792f93a849cd7fbf566f528f858a5ea

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// CgroupMemory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/656080dee792f93a849cd7fbf566f528f858a5ea/specification/nodes/_types/Stats.ts#L212-L216
type CgroupMemory struct {
	ControlGroup *string `json:"control_group,omitempty"`
	LimitInBytes *string `json:"limit_in_bytes,omitempty"`
	UsageInBytes *string `json:"usage_in_bytes,omitempty"`
}

func (s *CgroupMemory) UnmarshalJSON(data []byte) error {

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

		case "control_group":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ControlGroup = &o

		case "limit_in_bytes":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.LimitInBytes = &o

		case "usage_in_bytes":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.UsageInBytes = &o

		}
	}
	return nil
}

// NewCgroupMemory returns a CgroupMemory.
func NewCgroupMemory() *CgroupMemory {
	r := &CgroupMemory{}

	return r
}
