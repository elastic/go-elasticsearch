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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
)

// IndicesOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/common.ts#L297-L324
type IndicesOptions struct {
	// AllowNoIndices If false, the request returns an error if any wildcard expression, index
	// alias, or `_all` value targets only
	// missing or closed indices. This behavior applies even if the request targets
	// other open indices. For example,
	// a request targeting `foo*,bar*` returns an error if an index starts with
	// `foo` but no index starts with `bar`.
	AllowNoIndices *bool `json:"allow_no_indices,omitempty"`
	// ExpandWildcards Type of index that wildcard patterns can match. If the request can target
	// data streams, this argument
	// determines whether wildcard expressions match hidden data streams. Supports
	// comma-separated values,
	// such as `open,hidden`.
	ExpandWildcards []expandwildcard.ExpandWildcard `json:"expand_wildcards,omitempty"`
	// IgnoreThrottled If true, concrete, expanded or aliased indices are ignored when frozen.
	IgnoreThrottled *bool `json:"ignore_throttled,omitempty"`
	// IgnoreUnavailable If true, missing or closed indices are not included in the response.
	IgnoreUnavailable *bool `json:"ignore_unavailable,omitempty"`
}

// NewIndicesOptions returns a IndicesOptions.
func NewIndicesOptions() *IndicesOptions {
	r := &IndicesOptions{}

	return r
}
