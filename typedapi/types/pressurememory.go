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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// PressureMemory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/nodes/_types/Stats.ts#L65-L74
type PressureMemory struct {
	AllInBytes                            *int64 `json:"all_in_bytes,omitempty"`
	CombinedCoordinatingAndPrimaryInBytes *int64 `json:"combined_coordinating_and_primary_in_bytes,omitempty"`
	CoordinatingInBytes                   *int64 `json:"coordinating_in_bytes,omitempty"`
	CoordinatingRejections                *int64 `json:"coordinating_rejections,omitempty"`
	PrimaryInBytes                        *int64 `json:"primary_in_bytes,omitempty"`
	PrimaryRejections                     *int64 `json:"primary_rejections,omitempty"`
	ReplicaInBytes                        *int64 `json:"replica_in_bytes,omitempty"`
	ReplicaRejections                     *int64 `json:"replica_rejections,omitempty"`
}

// NewPressureMemory returns a PressureMemory.
func NewPressureMemory() *PressureMemory {
	r := &PressureMemory{}

	return r
}
