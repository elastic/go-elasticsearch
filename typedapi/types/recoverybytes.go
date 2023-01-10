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

// RecoveryBytes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/indices/recovery/types.ts#L38-L48
type RecoveryBytes struct {
	Percent                      Percentage `json:"percent"`
	Recovered                    *ByteSize  `json:"recovered,omitempty"`
	RecoveredFromSnapshot        *ByteSize  `json:"recovered_from_snapshot,omitempty"`
	RecoveredFromSnapshotInBytes *ByteSize  `json:"recovered_from_snapshot_in_bytes,omitempty"`
	RecoveredInBytes             ByteSize   `json:"recovered_in_bytes"`
	Reused                       *ByteSize  `json:"reused,omitempty"`
	ReusedInBytes                ByteSize   `json:"reused_in_bytes"`
	Total                        *ByteSize  `json:"total,omitempty"`
	TotalInBytes                 ByteSize   `json:"total_in_bytes"`
}

// NewRecoveryBytes returns a RecoveryBytes.
func NewRecoveryBytes() *RecoveryBytes {
	r := &RecoveryBytes{}

	return r
}
