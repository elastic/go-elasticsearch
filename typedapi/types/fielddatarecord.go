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

// FielddataRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/cat/fielddata/types.ts#L20-L48
type FielddataRecord struct {
	// Field field name
	Field *string `json:"field,omitempty"`
	// Host host name
	Host *string `json:"host,omitempty"`
	// Id node id
	Id *string `json:"id,omitempty"`
	// Ip ip address
	Ip *string `json:"ip,omitempty"`
	// Node node name
	Node *string `json:"node,omitempty"`
	// Size field data usage
	Size *string `json:"size,omitempty"`
}

// NewFielddataRecord returns a FielddataRecord.
func NewFielddataRecord() *FielddataRecord {
	r := &FielddataRecord{}

	return r
}
