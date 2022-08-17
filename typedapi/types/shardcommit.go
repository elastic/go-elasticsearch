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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// ShardCommit type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/stats/types.ts#L100-L105
type ShardCommit struct {
	Generation int               `json:"generation"`
	Id         Id                `json:"id"`
	NumDocs    int64             `json:"num_docs"`
	UserData   map[string]string `json:"user_data"`
}

// ShardCommitBuilder holds ShardCommit struct and provides a builder API.
type ShardCommitBuilder struct {
	v *ShardCommit
}

// NewShardCommit provides a builder for the ShardCommit struct.
func NewShardCommitBuilder() *ShardCommitBuilder {
	r := ShardCommitBuilder{
		&ShardCommit{
			UserData: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ShardCommit struct
func (rb *ShardCommitBuilder) Build() ShardCommit {
	return *rb.v
}

func (rb *ShardCommitBuilder) Generation(generation int) *ShardCommitBuilder {
	rb.v.Generation = generation
	return rb
}

func (rb *ShardCommitBuilder) Id(id Id) *ShardCommitBuilder {
	rb.v.Id = id
	return rb
}

func (rb *ShardCommitBuilder) NumDocs(numdocs int64) *ShardCommitBuilder {
	rb.v.NumDocs = numdocs
	return rb
}

func (rb *ShardCommitBuilder) UserData(value map[string]string) *ShardCommitBuilder {
	rb.v.UserData = value
	return rb
}
