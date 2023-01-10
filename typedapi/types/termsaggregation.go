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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationcollectmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationexecutionhint"
)

// TermsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/aggregations/bucket.ts#L380-L397
type TermsAggregation struct {
	CollectMode           *termsaggregationcollectmode.TermsAggregationCollectMode     `json:"collect_mode,omitempty"`
	Exclude               []string                                                     `json:"exclude,omitempty"`
	ExecutionHint         *termsaggregationexecutionhint.TermsAggregationExecutionHint `json:"execution_hint,omitempty"`
	Field                 *string                                                      `json:"field,omitempty"`
	Format                *string                                                      `json:"format,omitempty"`
	Include               *TermsInclude                                                `json:"include,omitempty"`
	Meta                  map[string]interface{}                                       `json:"meta,omitempty"`
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

// NewTermsAggregation returns a TermsAggregation.
func NewTermsAggregation() *TermsAggregation {
	r := &TermsAggregation{}

	return r
}
