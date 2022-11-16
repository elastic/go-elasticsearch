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
// https://github.com/elastic/elasticsearch-specification/tree/555082f38110f65b60d470107d211fc354a5c55a


package types

// MappingLimitSettingsDepth type.
//
// https://github.com/elastic/elasticsearch-specification/blob/555082f38110f65b60d470107d211fc354a5c55a/specification/indices/_types/IndexSettings.ts#L427-L434
type MappingLimitSettingsDepth struct {
	// Limit The maximum depth for a field, which is measured as the number of inner
	// objects. For instance, if all fields are defined
	// at the root object level, then the depth is 1. If there is one object
	// mapping, then the depth is 2, etc.
	Limit *int `json:"limit,omitempty"`
}

// NewMappingLimitSettingsDepth returns a MappingLimitSettingsDepth.
func NewMappingLimitSettingsDepth() *MappingLimitSettingsDepth {
	r := &MappingLimitSettingsDepth{}

	return r
}
