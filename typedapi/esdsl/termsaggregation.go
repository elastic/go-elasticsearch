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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termsaggregationcollectmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termsaggregationexecutionhint"
)

type _termsAggregation struct {
	v *types.TermsAggregation
}

// A multi-bucket value source based aggregation where buckets are dynamically
// built - one per unique value.
func NewTermsAggregation() *_termsAggregation {

	return &_termsAggregation{v: types.NewTermsAggregation()}

}

func (s *_termsAggregation) CollectMode(collectmode termsaggregationcollectmode.TermsAggregationCollectMode) *_termsAggregation {

	s.v.CollectMode = &collectmode
	return s
}

func (s *_termsAggregation) Exclude(termsexcludes ...string) *_termsAggregation {

	s.v.Exclude = termsexcludes

	return s
}

func (s *_termsAggregation) ExecutionHint(executionhint termsaggregationexecutionhint.TermsAggregationExecutionHint) *_termsAggregation {

	s.v.ExecutionHint = &executionhint
	return s
}

func (s *_termsAggregation) Field(field string) *_termsAggregation {

	s.v.Field = &field

	return s
}

func (s *_termsAggregation) Format(format string) *_termsAggregation {

	s.v.Format = &format

	return s
}

func (s *_termsAggregation) Include(termsinclude types.TermsIncludeVariant) *_termsAggregation {

	s.v.Include = *termsinclude.TermsIncludeCaster()

	return s
}

func (s *_termsAggregation) MinDocCount(mindoccount int) *_termsAggregation {

	s.v.MinDocCount = &mindoccount

	return s
}

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

func (s *_termsAggregation) Order(aggregateorder types.AggregateOrderVariant) *_termsAggregation {

	s.v.Order = *aggregateorder.AggregateOrderCaster()

	return s
}

func (s *_termsAggregation) Script(script types.ScriptVariant) *_termsAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_termsAggregation) ShardMinDocCount(shardmindoccount int64) *_termsAggregation {

	s.v.ShardMinDocCount = &shardmindoccount

	return s
}

func (s *_termsAggregation) ShardSize(shardsize int) *_termsAggregation {

	s.v.ShardSize = &shardsize

	return s
}

func (s *_termsAggregation) ShowTermDocCountError(showtermdoccounterror bool) *_termsAggregation {

	s.v.ShowTermDocCountError = &showtermdoccounterror

	return s
}

func (s *_termsAggregation) Size(size int) *_termsAggregation {

	s.v.Size = &size

	return s
}

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
