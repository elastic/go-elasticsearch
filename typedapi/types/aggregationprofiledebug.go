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

// AggregationProfileDebug type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_global/search/_types/profile.ts#L39-L68
type AggregationProfileDebug struct {
	BuiltBuckets                      *int                                    `json:"built_buckets,omitempty"`
	CharsFetched                      *int                                    `json:"chars_fetched,omitempty"`
	CollectAnalyzedCount              *int                                    `json:"collect_analyzed_count,omitempty"`
	CollectAnalyzedNs                 *int                                    `json:"collect_analyzed_ns,omitempty"`
	CollectionStrategy                *string                                 `json:"collection_strategy,omitempty"`
	DeferredAggregators               []string                                `json:"deferred_aggregators,omitempty"`
	Delegate                          *string                                 `json:"delegate,omitempty"`
	DelegateDebug                     *AggregationProfileDebug                `json:"delegate_debug,omitempty"`
	EmptyCollectorsUsed               *int                                    `json:"empty_collectors_used,omitempty"`
	ExtractCount                      *int                                    `json:"extract_count,omitempty"`
	ExtractNs                         *int                                    `json:"extract_ns,omitempty"`
	Filters                           []AggregationProfileDelegateDebugFilter `json:"filters,omitempty"`
	HasFilter                         *bool                                   `json:"has_filter,omitempty"`
	MapReducer                        *string                                 `json:"map_reducer,omitempty"`
	NumericCollectorsUsed             *int                                    `json:"numeric_collectors_used,omitempty"`
	OrdinalsCollectorsOverheadTooHigh *int                                    `json:"ordinals_collectors_overhead_too_high,omitempty"`
	OrdinalsCollectorsUsed            *int                                    `json:"ordinals_collectors_used,omitempty"`
	ResultStrategy                    *string                                 `json:"result_strategy,omitempty"`
	SegmentsCollected                 *int                                    `json:"segments_collected,omitempty"`
	SegmentsCounted                   *int                                    `json:"segments_counted,omitempty"`
	SegmentsWithDeletedDocs           *int                                    `json:"segments_with_deleted_docs,omitempty"`
	SegmentsWithDocCountField         *int                                    `json:"segments_with_doc_count_field,omitempty"`
	SegmentsWithMultiValuedOrds       *int                                    `json:"segments_with_multi_valued_ords,omitempty"`
	SegmentsWithSingleValuedOrds      *int                                    `json:"segments_with_single_valued_ords,omitempty"`
	StringHashingCollectorsUsed       *int                                    `json:"string_hashing_collectors_used,omitempty"`
	SurvivingBuckets                  *int                                    `json:"surviving_buckets,omitempty"`
	TotalBuckets                      *int                                    `json:"total_buckets,omitempty"`
	ValuesFetched                     *int                                    `json:"values_fetched,omitempty"`
}

// NewAggregationProfileDebug returns a AggregationProfileDebug.
func NewAggregationProfileDebug() *AggregationProfileDebug {
	r := &AggregationProfileDebug{}

	return r
}
