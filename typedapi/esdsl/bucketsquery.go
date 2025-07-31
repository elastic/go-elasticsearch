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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _bucketsQuery struct {
	v types.BucketsQuery
}

func NewBucketsQuery() *_bucketsQuery {
	return &_bucketsQuery{v: nil}
}

func (u *_bucketsQuery) Map(value map[string]types.QueryVariant) *_bucketsQuery { // union map

	u.v = make(map[string]types.QueryVariant)
	for k, v := range value {
		u.v.(map[string]types.Query)[k] = *v.QueryCaster()
	}

	return u
}

func (u *_bucketsQuery) QueryContainers(querycontainers ...types.QueryVariant) *_bucketsQuery {

	u.v = make([]types.Query, len(querycontainers))
	for i, v := range querycontainers {
		u.v.([]types.Query)[i] = *v.QueryCaster()
	}

	return u
}

func (u *_bucketsQuery) BucketsQueryCaster() *types.BucketsQuery {
	return &u.v
}
