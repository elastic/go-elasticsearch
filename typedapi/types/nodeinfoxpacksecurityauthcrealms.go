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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

// NodeInfoXpackSecurityAuthcRealms type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/nodes/info/types.ts#L281-L285
type NodeInfoXpackSecurityAuthcRealms struct {
	File   map[string]NodeInfoXpackSecurityAuthcRealmsStatus `json:"file,omitempty"`
	Native map[string]NodeInfoXpackSecurityAuthcRealmsStatus `json:"native,omitempty"`
	Pki    map[string]NodeInfoXpackSecurityAuthcRealmsStatus `json:"pki,omitempty"`
}

// NewNodeInfoXpackSecurityAuthcRealms returns a NodeInfoXpackSecurityAuthcRealms.
func NewNodeInfoXpackSecurityAuthcRealms() *NodeInfoXpackSecurityAuthcRealms {
	r := &NodeInfoXpackSecurityAuthcRealms{
		File:   make(map[string]NodeInfoXpackSecurityAuthcRealmsStatus),
		Native: make(map[string]NodeInfoXpackSecurityAuthcRealmsStatus),
		Pki:    make(map[string]NodeInfoXpackSecurityAuthcRealmsStatus),
	}

	return r
}
