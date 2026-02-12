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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _missingAggregation struct {
	v *types.MissingAggregation
}

// A field data based single bucket aggregation, that creates a bucket of all
// documents in the current document set context that are missing a field value
// (effectively, missing a field or having the configured NULL value set).
func NewMissingAggregation() *_missingAggregation {

	return &_missingAggregation{v: types.NewMissingAggregation()}

}

func (s *_missingAggregation) Field(field string) *_missingAggregation {

	s.v.Field = &field

	return s
}

func (s *_missingAggregation) Missing(missing types.MissingVariant) *_missingAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_missingAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Missing = s.v

	return container
}

func (s *_missingAggregation) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	container := types.NewApiKeyAggregationContainer()

	container.Missing = s.v

	return container
}

func (s *_missingAggregation) MissingAggregationCaster() *types.MissingAggregation {
	return s.v
}
