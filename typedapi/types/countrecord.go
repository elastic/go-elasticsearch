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

// CountRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/cat/count/types.ts#L23-L39
type CountRecord struct {
	// Count the document count
	Count *string `json:"count,omitempty"`
	// Epoch seconds since 1970-01-01 00:00:00
	Epoch StringifiedEpochTimeUnitSeconds `json:"epoch,omitempty"`
	// Timestamp time in HH:MM:SS
	Timestamp *string `json:"timestamp,omitempty"`
}

// NewCountRecord returns a CountRecord.
func NewCountRecord() *CountRecord {
	r := &CountRecord{}

	return r
}
