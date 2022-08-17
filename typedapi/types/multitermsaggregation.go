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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationcollectmode"
)

// MultiTermsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L262-L271
type MultiTermsAggregation struct {
	CollectMode           *termsaggregationcollectmode.TermsAggregationCollectMode `json:"collect_mode,omitempty"`
	Meta                  *Metadata                                                `json:"meta,omitempty"`
	MinDocCount           *int64                                                   `json:"min_doc_count,omitempty"`
	Name                  *string                                                  `json:"name,omitempty"`
	Order                 *AggregateOrder                                          `json:"order,omitempty"`
	ShardMinDocCount      *int64                                                   `json:"shard_min_doc_count,omitempty"`
	ShardSize             *int                                                     `json:"shard_size,omitempty"`
	ShowTermDocCountError *bool                                                    `json:"show_term_doc_count_error,omitempty"`
	Size                  *int                                                     `json:"size,omitempty"`
	Terms                 []MultiTermLookup                                        `json:"terms"`
}

// MultiTermsAggregationBuilder holds MultiTermsAggregation struct and provides a builder API.
type MultiTermsAggregationBuilder struct {
	v *MultiTermsAggregation
}

// NewMultiTermsAggregation provides a builder for the MultiTermsAggregation struct.
func NewMultiTermsAggregationBuilder() *MultiTermsAggregationBuilder {
	r := MultiTermsAggregationBuilder{
		&MultiTermsAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the MultiTermsAggregation struct
func (rb *MultiTermsAggregationBuilder) Build() MultiTermsAggregation {
	return *rb.v
}

func (rb *MultiTermsAggregationBuilder) CollectMode(collectmode termsaggregationcollectmode.TermsAggregationCollectMode) *MultiTermsAggregationBuilder {
	rb.v.CollectMode = &collectmode
	return rb
}

func (rb *MultiTermsAggregationBuilder) Meta(meta *MetadataBuilder) *MultiTermsAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *MultiTermsAggregationBuilder) MinDocCount(mindoccount int64) *MultiTermsAggregationBuilder {
	rb.v.MinDocCount = &mindoccount
	return rb
}

func (rb *MultiTermsAggregationBuilder) Name(name string) *MultiTermsAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *MultiTermsAggregationBuilder) Order(order *AggregateOrderBuilder) *MultiTermsAggregationBuilder {
	v := order.Build()
	rb.v.Order = &v
	return rb
}

func (rb *MultiTermsAggregationBuilder) ShardMinDocCount(shardmindoccount int64) *MultiTermsAggregationBuilder {
	rb.v.ShardMinDocCount = &shardmindoccount
	return rb
}

func (rb *MultiTermsAggregationBuilder) ShardSize(shardsize int) *MultiTermsAggregationBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}

func (rb *MultiTermsAggregationBuilder) ShowTermDocCountError(showtermdoccounterror bool) *MultiTermsAggregationBuilder {
	rb.v.ShowTermDocCountError = &showtermdoccounterror
	return rb
}

func (rb *MultiTermsAggregationBuilder) Size(size int) *MultiTermsAggregationBuilder {
	rb.v.Size = &size
	return rb
}

func (rb *MultiTermsAggregationBuilder) Terms(terms []MultiTermLookupBuilder) *MultiTermsAggregationBuilder {
	tmp := make([]MultiTermLookup, len(terms))
	for _, value := range terms {
		tmp = append(tmp, value.Build())
	}
	rb.v.Terms = tmp
	return rb
}
