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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

// This is provide all the types that are part of the union.
type _aggregateOrder struct {
	v types.AggregateOrder
}

func NewAggregateOrder() *_aggregateOrder {
	return &_aggregateOrder{v: nil}
}

func (u *_aggregateOrder) Map(value map[string]sortorder.SortOrder) *_aggregateOrder { // union map

	u.v = make(map[string]sortorder.SortOrder)
	for k, v := range value {
		u.v.(map[string]sortorder.SortOrder)[k] = v
	}

	return u
}

func (u *_aggregateOrder) SortOrders(sortorders ...map[string]sortorder.SortOrder) *_aggregateOrder {

	u.v = make([]map[string]sortorder.SortOrder, len(sortorders))
	u.v = sortorders

	return u
}

func (u *_aggregateOrder) AggregateOrderCaster() *types.AggregateOrder {
	return &u.v
}
