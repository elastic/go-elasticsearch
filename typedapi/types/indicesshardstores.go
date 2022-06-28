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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// IndicesShardStores type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/indices/shard_stores/types.ts#L25-L27
type IndicesShardStores struct {
	Shards map[string]ShardStoreWrapper `json:"shards"`
}

// IndicesShardStoresBuilder holds IndicesShardStores struct and provides a builder API.
type IndicesShardStoresBuilder struct {
	v *IndicesShardStores
}

// NewIndicesShardStores provides a builder for the IndicesShardStores struct.
func NewIndicesShardStoresBuilder() *IndicesShardStoresBuilder {
	r := IndicesShardStoresBuilder{
		&IndicesShardStores{
			Shards: make(map[string]ShardStoreWrapper, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the IndicesShardStores struct
func (rb *IndicesShardStoresBuilder) Build() IndicesShardStores {
	return *rb.v
}

func (rb *IndicesShardStoresBuilder) Shards(values map[string]*ShardStoreWrapperBuilder) *IndicesShardStoresBuilder {
	tmp := make(map[string]ShardStoreWrapper, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Shards = tmp
	return rb
}
