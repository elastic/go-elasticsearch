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

// ApiKey type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/security/_types/ApiKey.ts#L27-L41
type ApiKey struct {
	Creation        *int64                      `json:"creation,omitempty"`
	Expiration      *int64                      `json:"expiration,omitempty"`
	Id              string                      `json:"id"`
	Invalidated     *bool                       `json:"invalidated,omitempty"`
	LimitedBy       []map[string]RoleDescriptor `json:"limited_by,omitempty"`
	Metadata        map[string]interface{}      `json:"metadata,omitempty"`
	Name            string                      `json:"name"`
	Realm           *string                     `json:"realm,omitempty"`
	RoleDescriptors map[string]RoleDescriptor   `json:"role_descriptors,omitempty"`
	Sort_           []FieldValue                `json:"_sort,omitempty"`
	Username        *string                     `json:"username,omitempty"`
}

// NewApiKey returns a ApiKey.
func NewApiKey() *ApiKey {
	r := &ApiKey{
		RoleDescriptors: make(map[string]RoleDescriptor, 0),
	}

	return r
}
