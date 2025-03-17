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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationcollectmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationexecutionhint"
)

type _termsAggregation struct {
	v *types.TermsAggregation
}

// A multi-bucket value source based aggregation where buckets are dynamically
// built - one per unique value.
func NewTermsAggregation() *_termsAggregation {

	return &_termsAggregation{v: types.NewTermsAggregation()}

}

// Determines how child aggregations should be calculated: breadth-first or
// depth-first.
func (s *_termsAggregation) CollectMode(collectmode termsaggregationcollectmode.TermsAggregationCollectMode) *_termsAggregation {

	s.v.CollectMode = &collectmode
	return s
}

// Values to exclude.
// Accepts regular expressions and partitions.
func (s *_termsAggregation) Exclude(termsexcludes ...string) *_termsAggregation {

	s.v.Exclude = termsexcludes

	return s
}

// Determines whether the aggregation will use field values directly or global
// ordinals.
func (s *_termsAggregation) ExecutionHint(executionhint termsaggregationexecutionhint.TermsAggregationExecutionHint) *_termsAggregation {

	s.v.ExecutionHint = &executionhint
	return s
}

// The field from which to return terms.
func (s *_termsAggregation) Field(field string) *_termsAggregation {

	s.v.Field = &field

	return s
}

func (s *_termsAggregation) Format(format string) *_termsAggregation {

	s.v.Format = &format

	return s
}

// Values to include.
// Accepts regular expressions and partitions.
func (s *_termsAggregation) Include(termsinclude types.TermsIncludeVariant) *_termsAggregation {

	s.v.Include = *termsinclude.TermsIncludeCaster()

	return s
}

// Only return values that are found in more than `min_doc_count` hits.
func (s *_termsAggregation) MinDocCount(mindoccount int) *_termsAggregation {

	s.v.MinDocCount = &mindoccount

	return s
}

// The value to apply to documents that do not have a value.
// By default, documents without a value are ignored.
func (s *_termsAggregation) Missing(missing types.MissingVariant) *_termsAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_termsAggregation) MissingBucket(missingbucket bool) *_termsAggregation {

	s.v.MissingBucket = &missingbucket

	return s
}

func (s *_termsAggregation) MissingOrder(missingorder missingorder.MissingOrder) *_termsAggregation {

	s.v.MissingOrder = &missingorder
	return s
}

// Specifies the sort order of the buckets.
// Defaults to sorting by descending document count.
func (s *_termsAggregation) Order(aggregateorder types.AggregateOrderVariant) *_termsAggregation {

	s.v.Order = *aggregateorder.AggregateOrderCaster()

	return s
}

func (s *_termsAggregation) Script(script types.ScriptVariant) *_termsAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

// Regulates the certainty a shard has if the term should actually be added to
// the candidate list or not with respect to the `min_doc_count`.
// Terms will only be considered if their local shard frequency within the set
// is higher than the `shard_min_doc_count`.
func (s *_termsAggregation) ShardMinDocCount(shardmindoccount int64) *_termsAggregation {

	s.v.ShardMinDocCount = &shardmindoccount

	return s
}

// The number of candidate terms produced by each shard.
// By default, `shard_size` will be automatically estimated based on the number
// of shards and the `size` parameter.
func (s *_termsAggregation) ShardSize(shardsize int) *_termsAggregation {

	s.v.ShardSize = &shardsize

	return s
}

// Set to `true` to return the `doc_count_error_upper_bound`, which is an upper
// bound to the error on the `doc_count` returned by each shard.
func (s *_termsAggregation) ShowTermDocCountError(showtermdoccounterror bool) *_termsAggregation {

	s.v.ShowTermDocCountError = &showtermdoccounterror

	return s
}

// The number of buckets returned out of the overall terms list.
func (s *_termsAggregation) Size(size int) *_termsAggregation {

	s.v.Size = &size

	return s
}

// Coerced unmapped fields into the specified type.
func (s *_termsAggregation) ValueType(valuetype string) *_termsAggregation {

	s.v.ValueType = &valuetype

	return s
}

func (s *_termsAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Terms = s.v

	return container
}

func (s *_termsAggregation) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	container := types.NewApiKeyAggregationContainer()

	container.Terms = s.v

	return container
}

func (s *_termsAggregation) PivotGroupByContainerCaster() *types.PivotGroupByContainer {
	container := types.NewPivotGroupByContainer()

	container.Terms = s.v

	return container
}

func (s *_termsAggregation) TermsAggregationCaster() *types.TermsAggregation {
	return s.v
}
