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

// NodeOperatingSystemInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/nodes/info/types.ts#L367-L384
type NodeOperatingSystemInfo struct {
	// AllocatedProcessors The number of processors actually used to calculate thread pool size. This
	// number can be set with the node.processors setting of a node and defaults to
	// the number of processors reported by the OS.
	AllocatedProcessors *int `json:"allocated_processors,omitempty"`
	// Arch Name of the JVM architecture (ex: amd64, x86)
	Arch string `json:"arch"`
	// AvailableProcessors Number of processors available to the Java virtual machine
	AvailableProcessors int             `json:"available_processors"`
	Cpu                 *NodeInfoOSCPU  `json:"cpu,omitempty"`
	Mem                 *NodeInfoMemory `json:"mem,omitempty"`
	// Name Name of the operating system (ex: Linux, Windows, Mac OS X)
	Name       string `json:"name"`
	PrettyName string `json:"pretty_name"`
	// RefreshIntervalInMillis Refresh interval for the OS statistics
	RefreshIntervalInMillis int64           `json:"refresh_interval_in_millis"`
	Swap                    *NodeInfoMemory `json:"swap,omitempty"`
	// Version Version of the operating system
	Version string `json:"version"`
}

func (s *NodeOperatingSystemInfo) UnmarshalJSON(data []byte) error {

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
				s.AllocatedProcessors = &value
			case float64:
				f := int(v)
				s.AllocatedProcessors = &f
			}

		case "arch":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Arch = o

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

		case "cpu":
			if err := dec.Decode(&s.Cpu); err != nil {
				return err
			}

		case "mem":
			if err := dec.Decode(&s.Mem); err != nil {
				return err
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "pretty_name":
			if err := dec.Decode(&s.PrettyName); err != nil {
				return err
			}

		case "refresh_interval_in_millis":
			if err := dec.Decode(&s.RefreshIntervalInMillis); err != nil {
				return err
			}

		case "swap":
			if err := dec.Decode(&s.Swap); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewNodeOperatingSystemInfo returns a NodeOperatingSystemInfo.
func NewNodeOperatingSystemInfo() *NodeOperatingSystemInfo {
	r := &NodeOperatingSystemInfo{}

	return r
}
