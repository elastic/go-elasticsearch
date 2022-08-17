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

// Latest type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/transform/_types/Transform.ts#L47-L52
type Latest struct {
	// Sort Specifies the date field that is used to identify the latest documents.
	Sort Field `json:"sort"`
	// UniqueKey Specifies an array of one or more fields that are used to group the data.
	UniqueKey []Field `json:"unique_key"`
}

// LatestBuilder holds Latest struct and provides a builder API.
type LatestBuilder struct {
	v *Latest
}

// NewLatest provides a builder for the Latest struct.
func NewLatestBuilder() *LatestBuilder {
	r := LatestBuilder{
		&Latest{},
	}

	return &r
}

// Build finalize the chain and returns the Latest struct
func (rb *LatestBuilder) Build() Latest {
	return *rb.v
}

// Sort Specifies the date field that is used to identify the latest documents.

func (rb *LatestBuilder) Sort(sort Field) *LatestBuilder {
	rb.v.Sort = sort
	return rb
}

// UniqueKey Specifies an array of one or more fields that are used to group the data.

func (rb *LatestBuilder) UniqueKey(unique_key ...Field) *LatestBuilder {
	rb.v.UniqueKey = unique_key
	return rb
}
