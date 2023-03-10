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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// Memory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/get_memory_stats/types.ts#L25-L48
type Memory struct {
	Attributes  map[string]string `json:"attributes"`
	EphemeralId string            `json:"ephemeral_id"`
	// Jvm Contains Java Virtual Machine (JVM) statistics for the node.
	Jvm JvmStats `json:"jvm"`
	// Mem Contains statistics about memory usage for the node.
	Mem MemStats `json:"mem"`
	// Name Human-readable identifier for the node. Based on the Node name setting
	// setting.
	Name string `json:"name"`
	// Roles Roles assigned to the node.
	Roles []string `json:"roles"`
	// TransportAddress The host and port where transport HTTP connections are accepted.
	TransportAddress string `json:"transport_address"`
}

// NewMemory returns a Memory.
func NewMemory() *Memory {
	r := &Memory{
		Attributes: make(map[string]string, 0),
	}

	return r
}
