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

// Process type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/nodes/_types/Stats.ts#L381-L387
type Process struct {
	Cpu                 *Cpu         `json:"cpu,omitempty"`
	MaxFileDescriptors  *int         `json:"max_file_descriptors,omitempty"`
	Mem                 *MemoryStats `json:"mem,omitempty"`
	OpenFileDescriptors *int         `json:"open_file_descriptors,omitempty"`
	Timestamp           *int64       `json:"timestamp,omitempty"`
}

func (s *Process) UnmarshalJSON(data []byte) error {

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

		case "cpu":
			if err := dec.Decode(&s.Cpu); err != nil {
				return err
			}

		case "max_file_descriptors":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxFileDescriptors = &value
			case float64:
				f := int(v)
				s.MaxFileDescriptors = &f
			}

		case "mem":
			if err := dec.Decode(&s.Mem); err != nil {
				return err
			}

		case "open_file_descriptors":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.OpenFileDescriptors = &value
			case float64:
				f := int(v)
				s.OpenFileDescriptors = &f
			}

		case "timestamp":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Timestamp = &value
			case float64:
				f := int64(v)
				s.Timestamp = &f
			}

		}
	}
	return nil
}

// NewProcess returns a Process.
func NewProcess() *Process {
	r := &Process{}

	return r
}
