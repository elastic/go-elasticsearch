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
// https://github.com/elastic/elasticsearch-specification/tree/37285cbd3fd155f913b50d880b40ec45f9df64b3

package types

// Coordinator snapshot of the original search request, serialized under
// `profile.request` when profiling is enabled. Introduced in Elasticsearch 9.5;
// omitted when the cluster contains mixed-version nodes that do not serialize
// this metadata.
//
// https://github.com/elastic/elasticsearch-specification/blob/37285cbd3fd155f913b50d880b40ec45f9df64b3/specification/_global/search/_types/profile.ts#L102-L115
type SearchRequestCoordinatorMetadata struct {
	// Indices Target index expressions from the request (before index resolution).
	Indices []string `json:"indices,omitempty"`
	// Source Original query source from the search request (`SearchSourceBuilder` as
	// JSON).
	Source *SearchRequestBody `json:"source,omitempty"`
}

// NewSearchRequestCoordinatorMetadata returns a SearchRequestCoordinatorMetadata.
func NewSearchRequestCoordinatorMetadata() *SearchRequestCoordinatorMetadata {
	r := &SearchRequestCoordinatorMetadata{}

	return r
}
