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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/shardstoreallocation"
)

// ShardStore type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/shard_stores/types.ts#L29-L38
type ShardStore struct {
	Allocation       shardstoreallocation.ShardStoreAllocation `json:"allocation"`
	AllocationId     Id                                        `json:"allocation_id"`
	Attributes       map[string]interface{}                    `json:"attributes"`
	Id               Id                                        `json:"id"`
	LegacyVersion    VersionNumber                             `json:"legacy_version"`
	Name             Name                                      `json:"name"`
	StoreException   ShardStoreException                       `json:"store_exception"`
	TransportAddress TransportAddress                          `json:"transport_address"`
}

// ShardStoreBuilder holds ShardStore struct and provides a builder API.
type ShardStoreBuilder struct {
	v *ShardStore
}

// NewShardStore provides a builder for the ShardStore struct.
func NewShardStoreBuilder() *ShardStoreBuilder {
	r := ShardStoreBuilder{
		&ShardStore{
			Attributes: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ShardStore struct
func (rb *ShardStoreBuilder) Build() ShardStore {
	return *rb.v
}

func (rb *ShardStoreBuilder) Allocation(allocation shardstoreallocation.ShardStoreAllocation) *ShardStoreBuilder {
	rb.v.Allocation = allocation
	return rb
}

func (rb *ShardStoreBuilder) AllocationId(allocationid Id) *ShardStoreBuilder {
	rb.v.AllocationId = allocationid
	return rb
}

func (rb *ShardStoreBuilder) Attributes(value map[string]interface{}) *ShardStoreBuilder {
	rb.v.Attributes = value
	return rb
}

func (rb *ShardStoreBuilder) Id(id Id) *ShardStoreBuilder {
	rb.v.Id = id
	return rb
}

func (rb *ShardStoreBuilder) LegacyVersion(legacyversion VersionNumber) *ShardStoreBuilder {
	rb.v.LegacyVersion = legacyversion
	return rb
}

func (rb *ShardStoreBuilder) Name(name Name) *ShardStoreBuilder {
	rb.v.Name = name
	return rb
}

func (rb *ShardStoreBuilder) StoreException(storeexception *ShardStoreExceptionBuilder) *ShardStoreBuilder {
	v := storeexception.Build()
	rb.v.StoreException = v
	return rb
}

func (rb *ShardStoreBuilder) TransportAddress(transportaddress TransportAddress) *ShardStoreBuilder {
	rb.v.TransportAddress = transportaddress
	return rb
}
