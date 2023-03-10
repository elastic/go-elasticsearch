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

// TimeSync type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/transform/_types/Transform.ts#L175-L187
type TimeSync struct {
	// Delay The time delay between the current time and the latest input data time.
	Delay Duration `json:"delay,omitempty"`
	// Field The date field that is used to identify new documents in the source. In
	// general, it’s a good idea to use a field
	// that contains the ingest timestamp. If you use a different field, you might
	// need to set the delay such that it
	// accounts for data transmission delays.
	Field string `json:"field"`
}

// NewTimeSync returns a TimeSync.
func NewTimeSync() *TimeSync {
	r := &TimeSync{}

	return r
}
