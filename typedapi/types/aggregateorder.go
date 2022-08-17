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
// https://github.com/elastic/elasticsearch-specification/tree/e0ea3dc890d394d682096cc862b3bd879d9422e9


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

// AggregateOrder holds the union for the following types:
//
//	map[Field]sortorder.SortOrder
//	[]map[Field]sortorder.SortOrder
//
// https://github.com/elastic/elasticsearch-specification/blob/e0ea3dc890d394d682096cc862b3bd879d9422e9/specification/_types/aggregations/bucket.ts#L399-L401
type AggregateOrder interface{}

// AggregateOrderBuilder holds AggregateOrder struct and provides a builder API.
type AggregateOrderBuilder struct {
	v AggregateOrder
}

// NewAggregateOrder provides a builder for the AggregateOrder struct.
func NewAggregateOrderBuilder() *AggregateOrderBuilder {
	return &AggregateOrderBuilder{}
}

// Build finalize the chain and returns the AggregateOrder struct
func (u *AggregateOrderBuilder) Build() AggregateOrder {
	return u.v
}

func (u *AggregateOrderBuilder) Map(value map[Field]sortorder.SortOrder) *AggregateOrderBuilder {
	u.v = value
	return u
}

func (u *AggregateOrderBuilder) SortOrders(value ...map[Field]sortorder.SortOrder) *AggregateOrderBuilder {
	u.v = value
	return u
}
