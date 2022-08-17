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

// ShardPath type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/stats/types.ts#L128-L132
type ShardPath struct {
	DataPath         string `json:"data_path"`
	IsCustomDataPath bool   `json:"is_custom_data_path"`
	StatePath        string `json:"state_path"`
}

// ShardPathBuilder holds ShardPath struct and provides a builder API.
type ShardPathBuilder struct {
	v *ShardPath
}

// NewShardPath provides a builder for the ShardPath struct.
func NewShardPathBuilder() *ShardPathBuilder {
	r := ShardPathBuilder{
		&ShardPath{},
	}

	return &r
}

// Build finalize the chain and returns the ShardPath struct
func (rb *ShardPathBuilder) Build() ShardPath {
	return *rb.v
}

func (rb *ShardPathBuilder) DataPath(datapath string) *ShardPathBuilder {
	rb.v.DataPath = datapath
	return rb
}

func (rb *ShardPathBuilder) IsCustomDataPath(iscustomdatapath bool) *ShardPathBuilder {
	rb.v.IsCustomDataPath = iscustomdatapath
	return rb
}

func (rb *ShardPathBuilder) StatePath(statepath string) *ShardPathBuilder {
	rb.v.StatePath = statepath
	return rb
}
