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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationcollectmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationexecutionhint"
)

// TermsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L376-L393
type TermsAggregation struct {
	CollectMode           *termsaggregationcollectmode.TermsAggregationCollectMode     `json:"collect_mode,omitempty"`
	Exclude               *TermsExclude                                                `json:"exclude,omitempty"`
	ExecutionHint         *termsaggregationexecutionhint.TermsAggregationExecutionHint `json:"execution_hint,omitempty"`
	Field                 *Field                                                       `json:"field,omitempty"`
	Format                *string                                                      `json:"format,omitempty"`
	Include               *TermsInclude                                                `json:"include,omitempty"`
	Meta                  *Metadata                                                    `json:"meta,omitempty"`
	MinDocCount           *int                                                         `json:"min_doc_count,omitempty"`
	Missing               *Missing                                                     `json:"missing,omitempty"`
	MissingBucket         *bool                                                        `json:"missing_bucket,omitempty"`
	MissingOrder          *missingorder.MissingOrder                                   `json:"missing_order,omitempty"`
	Name                  *string                                                      `json:"name,omitempty"`
	Order                 *AggregateOrder                                              `json:"order,omitempty"`
	Script                *Script                                                      `json:"script,omitempty"`
	ShardSize             *int                                                         `json:"shard_size,omitempty"`
	ShowTermDocCountError *bool                                                        `json:"show_term_doc_count_error,omitempty"`
	Size                  *int                                                         `json:"size,omitempty"`
	ValueType             *string                                                      `json:"value_type,omitempty"`
}

// TermsAggregationBuilder holds TermsAggregation struct and provides a builder API.
type TermsAggregationBuilder struct {
	v *TermsAggregation
}

// NewTermsAggregation provides a builder for the TermsAggregation struct.
func NewTermsAggregationBuilder() *TermsAggregationBuilder {
	r := TermsAggregationBuilder{
		&TermsAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the TermsAggregation struct
func (rb *TermsAggregationBuilder) Build() TermsAggregation {
	return *rb.v
}

func (rb *TermsAggregationBuilder) CollectMode(collectmode termsaggregationcollectmode.TermsAggregationCollectMode) *TermsAggregationBuilder {
	rb.v.CollectMode = &collectmode
	return rb
}

func (rb *TermsAggregationBuilder) Exclude(exclude *TermsExcludeBuilder) *TermsAggregationBuilder {
	v := exclude.Build()
	rb.v.Exclude = &v
	return rb
}

func (rb *TermsAggregationBuilder) ExecutionHint(executionhint termsaggregationexecutionhint.TermsAggregationExecutionHint) *TermsAggregationBuilder {
	rb.v.ExecutionHint = &executionhint
	return rb
}

func (rb *TermsAggregationBuilder) Field(field Field) *TermsAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *TermsAggregationBuilder) Format(format string) *TermsAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *TermsAggregationBuilder) Include(include *TermsIncludeBuilder) *TermsAggregationBuilder {
	v := include.Build()
	rb.v.Include = &v
	return rb
}

func (rb *TermsAggregationBuilder) Meta(meta *MetadataBuilder) *TermsAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *TermsAggregationBuilder) MinDocCount(mindoccount int) *TermsAggregationBuilder {
	rb.v.MinDocCount = &mindoccount
	return rb
}

func (rb *TermsAggregationBuilder) Missing(missing *MissingBuilder) *TermsAggregationBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *TermsAggregationBuilder) MissingBucket(missingbucket bool) *TermsAggregationBuilder {
	rb.v.MissingBucket = &missingbucket
	return rb
}

func (rb *TermsAggregationBuilder) MissingOrder(missingorder missingorder.MissingOrder) *TermsAggregationBuilder {
	rb.v.MissingOrder = &missingorder
	return rb
}

func (rb *TermsAggregationBuilder) Name(name string) *TermsAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *TermsAggregationBuilder) Order(order *AggregateOrderBuilder) *TermsAggregationBuilder {
	v := order.Build()
	rb.v.Order = &v
	return rb
}

func (rb *TermsAggregationBuilder) Script(script *ScriptBuilder) *TermsAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *TermsAggregationBuilder) ShardSize(shardsize int) *TermsAggregationBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}

func (rb *TermsAggregationBuilder) ShowTermDocCountError(showtermdoccounterror bool) *TermsAggregationBuilder {
	rb.v.ShowTermDocCountError = &showtermdoccounterror
	return rb
}

func (rb *TermsAggregationBuilder) Size(size int) *TermsAggregationBuilder {
	rb.v.Size = &size
	return rb
}

func (rb *TermsAggregationBuilder) ValueType(valuetype string) *TermsAggregationBuilder {
	rb.v.ValueType = &valuetype
	return rb
}
