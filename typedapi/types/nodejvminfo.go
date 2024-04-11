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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// NodeJvmInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/nodes/info/types.ts#L356-L370
type NodeJvmInfo struct {
	GcCollectors                          []string          `json:"gc_collectors"`
	InputArguments                        []string          `json:"input_arguments"`
	Mem                                   NodeInfoJvmMemory `json:"mem"`
	MemoryPools                           []string          `json:"memory_pools"`
	Pid                                   int               `json:"pid"`
	StartTimeInMillis                     int64             `json:"start_time_in_millis"`
	UsingBundledJdk                       bool              `json:"using_bundled_jdk"`
	UsingCompressedOrdinaryObjectPointers string            `json:"using_compressed_ordinary_object_pointers,omitempty"`
	Version                               string            `json:"version"`
	VmName                                string            `json:"vm_name"`
	VmVendor                              string            `json:"vm_vendor"`
	VmVersion                             string            `json:"vm_version"`
}

func (s *NodeJvmInfo) UnmarshalJSON(data []byte) error {

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

		case "gc_collectors":
			if err := dec.Decode(&s.GcCollectors); err != nil {
				return fmt.Errorf("%s | %w", "GcCollectors", err)
			}

		case "input_arguments":
			if err := dec.Decode(&s.InputArguments); err != nil {
				return fmt.Errorf("%s | %w", "InputArguments", err)
			}

		case "mem":
			if err := dec.Decode(&s.Mem); err != nil {
				return fmt.Errorf("%s | %w", "Mem", err)
			}

		case "memory_pools":
			if err := dec.Decode(&s.MemoryPools); err != nil {
				return fmt.Errorf("%s | %w", "MemoryPools", err)
			}

		case "pid":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Pid", err)
				}
				s.Pid = value
			case float64:
				f := int(v)
				s.Pid = f
			}

		case "start_time_in_millis":
			if err := dec.Decode(&s.StartTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "StartTimeInMillis", err)
			}

		case "using_bundled_jdk", "bundled_jdk":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "UsingBundledJdk", err)
				}
				s.UsingBundledJdk = value
			case bool:
				s.UsingBundledJdk = v
			}

		case "using_compressed_ordinary_object_pointers":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "UsingCompressedOrdinaryObjectPointers", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.UsingCompressedOrdinaryObjectPointers = o

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		case "vm_name":
			if err := dec.Decode(&s.VmName); err != nil {
				return fmt.Errorf("%s | %w", "VmName", err)
			}

		case "vm_vendor":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "VmVendor", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.VmVendor = o

		case "vm_version":
			if err := dec.Decode(&s.VmVersion); err != nil {
				return fmt.Errorf("%s | %w", "VmVersion", err)
			}

		}
	}
	return nil
}

// NewNodeJvmInfo returns a NodeJvmInfo.
func NewNodeJvmInfo() *NodeJvmInfo {
	r := &NodeJvmInfo{}

	return r
}
