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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// AggregationProfileDebug type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/search/_types/profile.ts#L42-L76
type AggregationProfileDebug struct {
	BruteForceUsed                    *int                                    `json:"brute_force_used,omitempty"`
	BuiltBuckets                      *int                                    `json:"built_buckets,omitempty"`
	CharsFetched                      *int                                    `json:"chars_fetched,omitempty"`
	CollectAnalyzedCount              *int                                    `json:"collect_analyzed_count,omitempty"`
	CollectAnalyzedNs                 *int                                    `json:"collect_analyzed_ns,omitempty"`
	CollectionStrategy                *string                                 `json:"collection_strategy,omitempty"`
	DeferredAggregators               []string                                `json:"deferred_aggregators,omitempty"`
	Delegate                          *string                                 `json:"delegate,omitempty"`
	DelegateDebug                     *AggregationProfileDebug                `json:"delegate_debug,omitempty"`
	DynamicPruningAttempted           *int                                    `json:"dynamic_pruning_attempted,omitempty"`
	DynamicPruningUsed                *int                                    `json:"dynamic_pruning_used,omitempty"`
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
	SkippedDueToNoData                *int                                    `json:"skipped_due_to_no_data,omitempty"`
	StringHashingCollectorsUsed       *int                                    `json:"string_hashing_collectors_used,omitempty"`
	SurvivingBuckets                  *int                                    `json:"surviving_buckets,omitempty"`
	TotalBuckets                      *int                                    `json:"total_buckets,omitempty"`
	ValuesFetched                     *int                                    `json:"values_fetched,omitempty"`
}

func (s *AggregationProfileDebug) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "brute_force_used":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "BruteForceUsed", err)
				}
				s.BruteForceUsed = &value
			case float64:
				f := int(v)
				s.BruteForceUsed = &f
			}

		case "built_buckets":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "BuiltBuckets", err)
				}
				s.BuiltBuckets = &value
			case float64:
				f := int(v)
				s.BuiltBuckets = &f
			}

		case "chars_fetched":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "CharsFetched", err)
				}
				s.CharsFetched = &value
			case float64:
				f := int(v)
				s.CharsFetched = &f
			}

		case "collect_analyzed_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "CollectAnalyzedCount", err)
				}
				s.CollectAnalyzedCount = &value
			case float64:
				f := int(v)
				s.CollectAnalyzedCount = &f
			}

		case "collect_analyzed_ns":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "CollectAnalyzedNs", err)
				}
				s.CollectAnalyzedNs = &value
			case float64:
				f := int(v)
				s.CollectAnalyzedNs = &f
			}

		case "collection_strategy":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "CollectionStrategy", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CollectionStrategy = &o

		case "deferred_aggregators":
			if err := dec.Decode(&s.DeferredAggregators); err != nil {
				return fmt.Errorf("%s | %w", "DeferredAggregators", err)
			}

		case "delegate":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Delegate", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Delegate = &o

		case "delegate_debug":
			if err := dec.Decode(&s.DelegateDebug); err != nil {
				return fmt.Errorf("%s | %w", "DelegateDebug", err)
			}

		case "dynamic_pruning_attempted":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DynamicPruningAttempted", err)
				}
				s.DynamicPruningAttempted = &value
			case float64:
				f := int(v)
				s.DynamicPruningAttempted = &f
			}

		case "dynamic_pruning_used":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DynamicPruningUsed", err)
				}
				s.DynamicPruningUsed = &value
			case float64:
				f := int(v)
				s.DynamicPruningUsed = &f
			}

		case "empty_collectors_used":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "EmptyCollectorsUsed", err)
				}
				s.EmptyCollectorsUsed = &value
			case float64:
				f := int(v)
				s.EmptyCollectorsUsed = &f
			}

		case "extract_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ExtractCount", err)
				}
				s.ExtractCount = &value
			case float64:
				f := int(v)
				s.ExtractCount = &f
			}

		case "extract_ns":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ExtractNs", err)
				}
				s.ExtractNs = &value
			case float64:
				f := int(v)
				s.ExtractNs = &f
			}

		case "filters":
			if err := dec.Decode(&s.Filters); err != nil {
				return fmt.Errorf("%s | %w", "Filters", err)
			}

		case "has_filter":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "HasFilter", err)
				}
				s.HasFilter = &value
			case bool:
				s.HasFilter = &v
			}

		case "map_reducer":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MapReducer", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MapReducer = &o

		case "numeric_collectors_used":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumericCollectorsUsed", err)
				}
				s.NumericCollectorsUsed = &value
			case float64:
				f := int(v)
				s.NumericCollectorsUsed = &f
			}

		case "ordinals_collectors_overhead_too_high":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "OrdinalsCollectorsOverheadTooHigh", err)
				}
				s.OrdinalsCollectorsOverheadTooHigh = &value
			case float64:
				f := int(v)
				s.OrdinalsCollectorsOverheadTooHigh = &f
			}

		case "ordinals_collectors_used":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "OrdinalsCollectorsUsed", err)
				}
				s.OrdinalsCollectorsUsed = &value
			case float64:
				f := int(v)
				s.OrdinalsCollectorsUsed = &f
			}

		case "result_strategy":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ResultStrategy", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ResultStrategy = &o

		case "segments_collected":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SegmentsCollected", err)
				}
				s.SegmentsCollected = &value
			case float64:
				f := int(v)
				s.SegmentsCollected = &f
			}

		case "segments_counted":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SegmentsCounted", err)
				}
				s.SegmentsCounted = &value
			case float64:
				f := int(v)
				s.SegmentsCounted = &f
			}

		case "segments_with_deleted_docs":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SegmentsWithDeletedDocs", err)
				}
				s.SegmentsWithDeletedDocs = &value
			case float64:
				f := int(v)
				s.SegmentsWithDeletedDocs = &f
			}

		case "segments_with_doc_count_field":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SegmentsWithDocCountField", err)
				}
				s.SegmentsWithDocCountField = &value
			case float64:
				f := int(v)
				s.SegmentsWithDocCountField = &f
			}

		case "segments_with_multi_valued_ords":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SegmentsWithMultiValuedOrds", err)
				}
				s.SegmentsWithMultiValuedOrds = &value
			case float64:
				f := int(v)
				s.SegmentsWithMultiValuedOrds = &f
			}

		case "segments_with_single_valued_ords":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SegmentsWithSingleValuedOrds", err)
				}
				s.SegmentsWithSingleValuedOrds = &value
			case float64:
				f := int(v)
				s.SegmentsWithSingleValuedOrds = &f
			}

		case "skipped_due_to_no_data":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SkippedDueToNoData", err)
				}
				s.SkippedDueToNoData = &value
			case float64:
				f := int(v)
				s.SkippedDueToNoData = &f
			}

		case "string_hashing_collectors_used":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "StringHashingCollectorsUsed", err)
				}
				s.StringHashingCollectorsUsed = &value
			case float64:
				f := int(v)
				s.StringHashingCollectorsUsed = &f
			}

		case "surviving_buckets":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SurvivingBuckets", err)
				}
				s.SurvivingBuckets = &value
			case float64:
				f := int(v)
				s.SurvivingBuckets = &f
			}

		case "total_buckets":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalBuckets", err)
				}
				s.TotalBuckets = &value
			case float64:
				f := int(v)
				s.TotalBuckets = &f
			}

		case "values_fetched":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ValuesFetched", err)
				}
				s.ValuesFetched = &value
			case float64:
				f := int(v)
				s.ValuesFetched = &f
			}

		}
	}
	return nil
}

// NewAggregationProfileDebug returns a AggregationProfileDebug.
func NewAggregationProfileDebug() *AggregationProfileDebug {
	r := &AggregationProfileDebug{}

	return r
}
