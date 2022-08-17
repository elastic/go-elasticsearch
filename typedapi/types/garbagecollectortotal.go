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

// GarbageCollectorTotal type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L361-L365
type GarbageCollectorTotal struct {
	CollectionCount        *int64  `json:"collection_count,omitempty"`
	CollectionTime         *string `json:"collection_time,omitempty"`
	CollectionTimeInMillis *int64  `json:"collection_time_in_millis,omitempty"`
}

// GarbageCollectorTotalBuilder holds GarbageCollectorTotal struct and provides a builder API.
type GarbageCollectorTotalBuilder struct {
	v *GarbageCollectorTotal
}

// NewGarbageCollectorTotal provides a builder for the GarbageCollectorTotal struct.
func NewGarbageCollectorTotalBuilder() *GarbageCollectorTotalBuilder {
	r := GarbageCollectorTotalBuilder{
		&GarbageCollectorTotal{},
	}

	return &r
}

// Build finalize the chain and returns the GarbageCollectorTotal struct
func (rb *GarbageCollectorTotalBuilder) Build() GarbageCollectorTotal {
	return *rb.v
}

func (rb *GarbageCollectorTotalBuilder) CollectionCount(collectioncount int64) *GarbageCollectorTotalBuilder {
	rb.v.CollectionCount = &collectioncount
	return rb
}

func (rb *GarbageCollectorTotalBuilder) CollectionTime(collectiontime string) *GarbageCollectorTotalBuilder {
	rb.v.CollectionTime = &collectiontime
	return rb
}

func (rb *GarbageCollectorTotalBuilder) CollectionTimeInMillis(collectiontimeinmillis int64) *GarbageCollectorTotalBuilder {
	rb.v.CollectionTimeInMillis = &collectiontimeinmillis
	return rb
}
