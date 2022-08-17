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

// BucketsQueryContainer holds the union for the following types:
//
//	map[string]QueryContainer
//	[]QueryContainer
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L303-L312
type BucketsQueryContainer interface{}

// BucketsQueryContainerBuilder holds BucketsQueryContainer struct and provides a builder API.
type BucketsQueryContainerBuilder struct {
	v BucketsQueryContainer
}

// NewBucketsQueryContainer provides a builder for the BucketsQueryContainer struct.
func NewBucketsQueryContainerBuilder() *BucketsQueryContainerBuilder {
	return &BucketsQueryContainerBuilder{}
}

// Build finalize the chain and returns the BucketsQueryContainer struct
func (u *BucketsQueryContainerBuilder) Build() BucketsQueryContainer {
	return u.v
}

func (u *BucketsQueryContainerBuilder) Map(values map[string]*QueryContainerBuilder) *BucketsQueryContainerBuilder {
	tmp := make(map[string]QueryContainer, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	u.v = tmp
	return u
}

func (u *BucketsQueryContainerBuilder) QueryContainers(querycontainers []QueryContainerBuilder) *BucketsQueryContainerBuilder {
	tmp := make([]QueryContainer, len(querycontainers))
	for _, value := range querycontainers {
		tmp = append(tmp, value.Build())
	}
	u.v = tmp
	return u
}
