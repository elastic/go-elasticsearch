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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

// TermsAggregationOrder holds the union for the following types:
//     map[Field]sortorder.SortOrder
//     []map[Field]sortorder.SortOrder
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/aggregations/bucket.ts#L392-L394
type TermsAggregationOrder interface{}

// TermsAggregationOrderBuilder holds TermsAggregationOrder struct and provides a builder API.
type TermsAggregationOrderBuilder struct {
	v TermsAggregationOrder
}

// NewTermsAggregationOrder provides a builder for the TermsAggregationOrder struct.
func NewTermsAggregationOrderBuilder() *TermsAggregationOrderBuilder {
	return &TermsAggregationOrderBuilder{}
}

// Build finalize the chain and returns the TermsAggregationOrder struct
func (u *TermsAggregationOrderBuilder) Build() TermsAggregationOrder {
	return u.v
}

func (u *TermsAggregationOrderBuilder) Map(value map[Field]sortorder.SortOrder) *TermsAggregationOrderBuilder {
	u.v = value
	return u
}

func (u *TermsAggregationOrderBuilder) SortOrders(value ...map[Field]sortorder.SortOrder) *TermsAggregationOrderBuilder {
	u.v = value
	return u
}
