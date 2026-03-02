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
// https://github.com/elastic/elasticsearch-specification/tree/224e96968e3ab27c2d1d33f015783b44ed183c1f

// Package profilingoperationmode
package profilingoperationmode

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/224e96968e3ab27c2d1d33f015783b44ed183c1f/specification/profiling/status/types.ts#L20-L24
type ProfilingOperationMode struct {
	Name string
}

var (
	RUNNING = ProfilingOperationMode{"RUNNING"}

	STOPPING = ProfilingOperationMode{"STOPPING"}

	STOPPED = ProfilingOperationMode{"STOPPED"}
)

func (p ProfilingOperationMode) MarshalText() (text []byte, err error) {
	return []byte(p.String()), nil
}

func (p *ProfilingOperationMode) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "running":
		*p = RUNNING
	case "stopping":
		*p = STOPPING
	case "stopped":
		*p = STOPPED
	default:
		*p = ProfilingOperationMode{string(text)}
	}

	return nil
}

func (p ProfilingOperationMode) String() string {
	return p.Name
}
