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

// AggregationProfileDebug type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/profile.ts#L39-L67
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

// AggregationProfileDebugBuilder holds AggregationProfileDebug struct and provides a builder API.
type AggregationProfileDebugBuilder struct {
	v *AggregationProfileDebug
}

// NewAggregationProfileDebug provides a builder for the AggregationProfileDebug struct.
func NewAggregationProfileDebugBuilder() *AggregationProfileDebugBuilder {
	r := AggregationProfileDebugBuilder{
		&AggregationProfileDebug{},
	}

	return &r
}

// Build finalize the chain and returns the AggregationProfileDebug struct
func (rb *AggregationProfileDebugBuilder) Build() AggregationProfileDebug {
	return *rb.v
}

func (rb *AggregationProfileDebugBuilder) BuiltBuckets(builtbuckets int) *AggregationProfileDebugBuilder {
	rb.v.BuiltBuckets = &builtbuckets
	return rb
}

func (rb *AggregationProfileDebugBuilder) CharsFetched(charsfetched int) *AggregationProfileDebugBuilder {
	rb.v.CharsFetched = &charsfetched
	return rb
}

func (rb *AggregationProfileDebugBuilder) CollectAnalyzedCount(collectanalyzedcount int) *AggregationProfileDebugBuilder {
	rb.v.CollectAnalyzedCount = &collectanalyzedcount
	return rb
}

func (rb *AggregationProfileDebugBuilder) CollectAnalyzedNs(collectanalyzedns int) *AggregationProfileDebugBuilder {
	rb.v.CollectAnalyzedNs = &collectanalyzedns
	return rb
}

func (rb *AggregationProfileDebugBuilder) CollectionStrategy(collectionstrategy string) *AggregationProfileDebugBuilder {
	rb.v.CollectionStrategy = &collectionstrategy
	return rb
}

func (rb *AggregationProfileDebugBuilder) DeferredAggregators(deferred_aggregators ...string) *AggregationProfileDebugBuilder {
	rb.v.DeferredAggregators = deferred_aggregators
	return rb
}

func (rb *AggregationProfileDebugBuilder) Delegate(delegate string) *AggregationProfileDebugBuilder {
	rb.v.Delegate = &delegate
	return rb
}

func (rb *AggregationProfileDebugBuilder) DelegateDebug(delegatedebug *AggregationProfileDebugBuilder) *AggregationProfileDebugBuilder {
	v := delegatedebug.Build()
	rb.v.DelegateDebug = &v
	return rb
}

func (rb *AggregationProfileDebugBuilder) EmptyCollectorsUsed(emptycollectorsused int) *AggregationProfileDebugBuilder {
	rb.v.EmptyCollectorsUsed = &emptycollectorsused
	return rb
}

func (rb *AggregationProfileDebugBuilder) ExtractCount(extractcount int) *AggregationProfileDebugBuilder {
	rb.v.ExtractCount = &extractcount
	return rb
}

func (rb *AggregationProfileDebugBuilder) ExtractNs(extractns int) *AggregationProfileDebugBuilder {
	rb.v.ExtractNs = &extractns
	return rb
}

func (rb *AggregationProfileDebugBuilder) Filters(filters []AggregationProfileDelegateDebugFilterBuilder) *AggregationProfileDebugBuilder {
	tmp := make([]AggregationProfileDelegateDebugFilter, len(filters))
	for _, value := range filters {
		tmp = append(tmp, value.Build())
	}
	rb.v.Filters = tmp
	return rb
}

func (rb *AggregationProfileDebugBuilder) HasFilter(hasfilter bool) *AggregationProfileDebugBuilder {
	rb.v.HasFilter = &hasfilter
	return rb
}

func (rb *AggregationProfileDebugBuilder) NumericCollectorsUsed(numericcollectorsused int) *AggregationProfileDebugBuilder {
	rb.v.NumericCollectorsUsed = &numericcollectorsused
	return rb
}

func (rb *AggregationProfileDebugBuilder) OrdinalsCollectorsOverheadTooHigh(ordinalscollectorsoverheadtoohigh int) *AggregationProfileDebugBuilder {
	rb.v.OrdinalsCollectorsOverheadTooHigh = &ordinalscollectorsoverheadtoohigh
	return rb
}

func (rb *AggregationProfileDebugBuilder) OrdinalsCollectorsUsed(ordinalscollectorsused int) *AggregationProfileDebugBuilder {
	rb.v.OrdinalsCollectorsUsed = &ordinalscollectorsused
	return rb
}

func (rb *AggregationProfileDebugBuilder) ResultStrategy(resultstrategy string) *AggregationProfileDebugBuilder {
	rb.v.ResultStrategy = &resultstrategy
	return rb
}

func (rb *AggregationProfileDebugBuilder) SegmentsCollected(segmentscollected int) *AggregationProfileDebugBuilder {
	rb.v.SegmentsCollected = &segmentscollected
	return rb
}

func (rb *AggregationProfileDebugBuilder) SegmentsCounted(segmentscounted int) *AggregationProfileDebugBuilder {
	rb.v.SegmentsCounted = &segmentscounted
	return rb
}

func (rb *AggregationProfileDebugBuilder) SegmentsWithDeletedDocs(segmentswithdeleteddocs int) *AggregationProfileDebugBuilder {
	rb.v.SegmentsWithDeletedDocs = &segmentswithdeleteddocs
	return rb
}

func (rb *AggregationProfileDebugBuilder) SegmentsWithDocCountField(segmentswithdoccountfield int) *AggregationProfileDebugBuilder {
	rb.v.SegmentsWithDocCountField = &segmentswithdoccountfield
	return rb
}

func (rb *AggregationProfileDebugBuilder) SegmentsWithMultiValuedOrds(segmentswithmultivaluedords int) *AggregationProfileDebugBuilder {
	rb.v.SegmentsWithMultiValuedOrds = &segmentswithmultivaluedords
	return rb
}

func (rb *AggregationProfileDebugBuilder) SegmentsWithSingleValuedOrds(segmentswithsinglevaluedords int) *AggregationProfileDebugBuilder {
	rb.v.SegmentsWithSingleValuedOrds = &segmentswithsinglevaluedords
	return rb
}

func (rb *AggregationProfileDebugBuilder) StringHashingCollectorsUsed(stringhashingcollectorsused int) *AggregationProfileDebugBuilder {
	rb.v.StringHashingCollectorsUsed = &stringhashingcollectorsused
	return rb
}

func (rb *AggregationProfileDebugBuilder) SurvivingBuckets(survivingbuckets int) *AggregationProfileDebugBuilder {
	rb.v.SurvivingBuckets = &survivingbuckets
	return rb
}

func (rb *AggregationProfileDebugBuilder) TotalBuckets(totalbuckets int) *AggregationProfileDebugBuilder {
	rb.v.TotalBuckets = &totalbuckets
	return rb
}

func (rb *AggregationProfileDebugBuilder) ValuesFetched(valuesfetched int) *AggregationProfileDebugBuilder {
	rb.v.ValuesFetched = &valuesfetched
	return rb
}
