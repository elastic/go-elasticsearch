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
// https://github.com/elastic/elasticsearch-specification/tree/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33


package types

// NodeProcessInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/nodes/info/types.ts#L383-L390
type NodeProcessInfo struct {
	// Id Process identifier (PID)
	Id int64 `json:"id"`
	// Mlockall Indicates if the process address space has been successfully locked in memory
	Mlockall bool `json:"mlockall"`
	// RefreshIntervalInMillis Refresh interval for the process statistics
	RefreshIntervalInMillis int64 `json:"refresh_interval_in_millis"`
}

// NewNodeProcessInfo returns a NodeProcessInfo.
func NewNodeProcessInfo() *NodeProcessInfo {
	r := &NodeProcessInfo{}

	return r
}
