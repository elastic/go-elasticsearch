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

// NodeAttributesRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/cat/nodeattrs/types.ts#L20-L55
type NodeAttributesRecord struct {
	// Attr attribute description
	Attr *string `json:"attr,omitempty"`
	// Host host name
	Host *string `json:"host,omitempty"`
	// Id unique node id
	Id *string `json:"id,omitempty"`
	// Ip ip address
	Ip *string `json:"ip,omitempty"`
	// Node node name
	Node *string `json:"node,omitempty"`
	// Pid process id
	Pid *string `json:"pid,omitempty"`
	// Port bound transport port
	Port *string `json:"port,omitempty"`
	// Value attribute value
	Value *string `json:"value,omitempty"`
}

// NewNodeAttributesRecord returns a NodeAttributesRecord.
func NewNodeAttributesRecord() *NodeAttributesRecord {
	r := &NodeAttributesRecord{}

	return r
}
