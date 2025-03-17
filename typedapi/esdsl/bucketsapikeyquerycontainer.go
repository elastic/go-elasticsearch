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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

// This is provide all the types that are part of the union.
type _bucketsApiKeyQueryContainer struct {
	v types.BucketsApiKeyQueryContainer
}

func NewBucketsApiKeyQueryContainer() *_bucketsApiKeyQueryContainer {
	return &_bucketsApiKeyQueryContainer{v: nil}
}

func (u *_bucketsApiKeyQueryContainer) Map(value map[string]types.ApiKeyQueryContainerVariant) *_bucketsApiKeyQueryContainer { // union map

	u.v = make(map[string]types.ApiKeyQueryContainerVariant)
	for k, v := range value {
		u.v.(map[string]types.ApiKeyQueryContainer)[k] = *v.ApiKeyQueryContainerCaster()
	}

	return u
}

func (u *_bucketsApiKeyQueryContainer) ApiKeyQueryContainers(apikeyquerycontainers ...types.ApiKeyQueryContainerVariant) *_bucketsApiKeyQueryContainer {

	u.v = make([]types.ApiKeyQueryContainer, len(apikeyquerycontainers))
	for i, v := range apikeyquerycontainers {
		u.v.([]types.ApiKeyQueryContainer)[i] = *v.ApiKeyQueryContainerCaster()
	}

	return u
}

func (u *_bucketsApiKeyQueryContainer) BucketsApiKeyQueryContainerCaster() *types.BucketsApiKeyQueryContainer {
	return &u.v
}
