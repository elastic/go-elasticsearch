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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// MgetOperation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_global/mget/types.ts#L32-L55
type MgetOperation struct {
	// Id_ The unique document ID.
	Id_ string `json:"_id"`
	// Index_ The index that contains the document.
	Index_ *string `json:"_index,omitempty"`
	// Routing The key for the primary shard the document resides on. Required if routing is
	// used during indexing.
	Routing *string `json:"routing,omitempty"`
	// Source_ If `false`, excludes all _source fields.
	Source_ SourceConfig `json:"_source,omitempty"`
	// StoredFields The stored fields you want to retrieve.
	StoredFields []string                 `json:"stored_fields,omitempty"`
	Version      *int64                   `json:"version,omitempty"`
	VersionType  *versiontype.VersionType `json:"version_type,omitempty"`
}

// NewMgetOperation returns a MgetOperation.
func NewMgetOperation() *MgetOperation {
	r := &MgetOperation{}

	return r
}
