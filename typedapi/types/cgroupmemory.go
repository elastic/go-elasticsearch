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

// CgroupMemory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/nodes/_types/Stats.ts#L521-L537
type CgroupMemory struct {
	// ControlGroup The `memory` control group to which the Elasticsearch process belongs.
	ControlGroup *string `json:"control_group,omitempty"`
	// LimitInBytes The maximum amount of user memory (including file cache) allowed for all
	// tasks in the same cgroup as the Elasticsearch process.
	// This value can be too big to store in a `long`, so is returned as a string so
	// that the value returned can exactly match what the underlying operating
	// system interface returns.
	// Any value that is too large to parse into a `long` almost certainly means no
	// limit has been set for the cgroup.
	LimitInBytes *string `json:"limit_in_bytes,omitempty"`
	// UsageInBytes The total current memory usage by processes in the cgroup, in bytes, by all
	// tasks in the same cgroup as the Elasticsearch process.
	// This value is stored as a string for consistency with `limit_in_bytes`.
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
