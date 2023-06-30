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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// Aggregations type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/aggregations/AggregationContainer.ts#L106-L211
type Aggregations struct {
	AdjacencyMatrix *AdjacencyMatrixAggregation `json:"adjacency_matrix,omitempty"`
	// Aggregations Sub-aggregations for this aggregation. Only applies to bucket aggregations.
	Aggregations            map[string]Aggregations             `json:"aggregations,omitempty"`
	AutoDateHistogram       *AutoDateHistogramAggregation       `json:"auto_date_histogram,omitempty"`
	Avg                     *AverageAggregation                 `json:"avg,omitempty"`
	AvgBucket               *AverageBucketAggregation           `json:"avg_bucket,omitempty"`
	Boxplot                 *BoxplotAggregation                 `json:"boxplot,omitempty"`
	BucketCorrelation       *BucketCorrelationAggregation       `json:"bucket_correlation,omitempty"`
	BucketCountKsTest       *BucketKsAggregation                `json:"bucket_count_ks_test,omitempty"`
	BucketScript            *BucketScriptAggregation            `json:"bucket_script,omitempty"`
	BucketSelector          *BucketSelectorAggregation          `json:"bucket_selector,omitempty"`
	BucketSort              *BucketSortAggregation              `json:"bucket_sort,omitempty"`
	Cardinality             *CardinalityAggregation             `json:"cardinality,omitempty"`
	CategorizeText          *CategorizeTextAggregation          `json:"categorize_text,omitempty"`
	Children                *ChildrenAggregation                `json:"children,omitempty"`
	Composite               *CompositeAggregation               `json:"composite,omitempty"`
	CumulativeCardinality   *CumulativeCardinalityAggregation   `json:"cumulative_cardinality,omitempty"`
	CumulativeSum           *CumulativeSumAggregation           `json:"cumulative_sum,omitempty"`
	DateHistogram           *DateHistogramAggregation           `json:"date_histogram,omitempty"`
	DateRange               *DateRangeAggregation               `json:"date_range,omitempty"`
	Derivative              *DerivativeAggregation              `json:"derivative,omitempty"`
	DiversifiedSampler      *DiversifiedSamplerAggregation      `json:"diversified_sampler,omitempty"`
	ExtendedStats           *ExtendedStatsAggregation           `json:"extended_stats,omitempty"`
	ExtendedStatsBucket     *ExtendedStatsBucketAggregation     `json:"extended_stats_bucket,omitempty"`
	Filter                  *Query                              `json:"filter,omitempty"`
	Filters                 *FiltersAggregation                 `json:"filters,omitempty"`
	FrequentItemSets        *FrequentItemSetsAggregation        `json:"frequent_item_sets,omitempty"`
	GeoBounds               *GeoBoundsAggregation               `json:"geo_bounds,omitempty"`
	GeoCentroid             *GeoCentroidAggregation             `json:"geo_centroid,omitempty"`
	GeoDistance             *GeoDistanceAggregation             `json:"geo_distance,omitempty"`
	GeoLine                 *GeoLineAggregation                 `json:"geo_line,omitempty"`
	GeohashGrid             *GeoHashGridAggregation             `json:"geohash_grid,omitempty"`
	GeohexGrid              *GeohexGridAggregation              `json:"geohex_grid,omitempty"`
	GeotileGrid             *GeoTileGridAggregation             `json:"geotile_grid,omitempty"`
	Global                  *GlobalAggregation                  `json:"global,omitempty"`
	Histogram               *HistogramAggregation               `json:"histogram,omitempty"`
	Inference               *InferenceAggregation               `json:"inference,omitempty"`
	IpPrefix                *IpPrefixAggregation                `json:"ip_prefix,omitempty"`
	IpRange                 *IpRangeAggregation                 `json:"ip_range,omitempty"`
	Line                    *GeoLineAggregation                 `json:"line,omitempty"`
	MatrixStats             *MatrixStatsAggregation             `json:"matrix_stats,omitempty"`
	Max                     *MaxAggregation                     `json:"max,omitempty"`
	MaxBucket               *MaxBucketAggregation               `json:"max_bucket,omitempty"`
	MedianAbsoluteDeviation *MedianAbsoluteDeviationAggregation `json:"median_absolute_deviation,omitempty"`
	Meta                    Metadata                            `json:"meta,omitempty"`
	Min                     *MinAggregation                     `json:"min,omitempty"`
	MinBucket               *MinBucketAggregation               `json:"min_bucket,omitempty"`
	Missing                 *MissingAggregation                 `json:"missing,omitempty"`
	MovingAvg               MovingAverageAggregation            `json:"moving_avg,omitempty"`
	MovingFn                *MovingFunctionAggregation          `json:"moving_fn,omitempty"`
	MovingPercentiles       *MovingPercentilesAggregation       `json:"moving_percentiles,omitempty"`
	MultiTerms              *MultiTermsAggregation              `json:"multi_terms,omitempty"`
	Nested                  *NestedAggregation                  `json:"nested,omitempty"`
	Normalize               *NormalizeAggregation               `json:"normalize,omitempty"`
	Parent                  *ParentAggregation                  `json:"parent,omitempty"`
	PercentileRanks         *PercentileRanksAggregation         `json:"percentile_ranks,omitempty"`
	Percentiles             *PercentilesAggregation             `json:"percentiles,omitempty"`
	PercentilesBucket       *PercentilesBucketAggregation       `json:"percentiles_bucket,omitempty"`
	Range                   *RangeAggregation                   `json:"range,omitempty"`
	RareTerms               *RareTermsAggregation               `json:"rare_terms,omitempty"`
	Rate                    *RateAggregation                    `json:"rate,omitempty"`
	ReverseNested           *ReverseNestedAggregation           `json:"reverse_nested,omitempty"`
	Sampler                 *SamplerAggregation                 `json:"sampler,omitempty"`
	ScriptedMetric          *ScriptedMetricAggregation          `json:"scripted_metric,omitempty"`
	SerialDiff              *SerialDifferencingAggregation      `json:"serial_diff,omitempty"`
	SignificantTerms        *SignificantTermsAggregation        `json:"significant_terms,omitempty"`
	SignificantText         *SignificantTextAggregation         `json:"significant_text,omitempty"`
	Stats                   *StatsAggregation                   `json:"stats,omitempty"`
	StatsBucket             *StatsBucketAggregation             `json:"stats_bucket,omitempty"`
	StringStats             *StringStatsAggregation             `json:"string_stats,omitempty"`
	Sum                     *SumAggregation                     `json:"sum,omitempty"`
	SumBucket               *SumBucketAggregation               `json:"sum_bucket,omitempty"`
	TTest                   *TTestAggregation                   `json:"t_test,omitempty"`
	Terms                   *TermsAggregation                   `json:"terms,omitempty"`
	TopHits                 *TopHitsAggregation                 `json:"top_hits,omitempty"`
	TopMetrics              *TopMetricsAggregation              `json:"top_metrics,omitempty"`
	ValueCount              *ValueCountAggregation              `json:"value_count,omitempty"`
	VariableWidthHistogram  *VariableWidthHistogramAggregation  `json:"variable_width_histogram,omitempty"`
	WeightedAvg             *WeightedAverageAggregation         `json:"weighted_avg,omitempty"`
}

func (s *Aggregations) UnmarshalJSON(data []byte) error {

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

		case "adjacency_matrix":
			if err := dec.Decode(&s.AdjacencyMatrix); err != nil {
				return err
			}

		case "aggregations", "aggs":
			if s.Aggregations == nil {
				s.Aggregations = make(map[string]Aggregations, 0)
			}
			if err := dec.Decode(&s.Aggregations); err != nil {
				return err
			}

		case "auto_date_histogram":
			if err := dec.Decode(&s.AutoDateHistogram); err != nil {
				return err
			}

		case "avg":
			if err := dec.Decode(&s.Avg); err != nil {
				return err
			}

		case "avg_bucket":
			if err := dec.Decode(&s.AvgBucket); err != nil {
				return err
			}

		case "boxplot":
			if err := dec.Decode(&s.Boxplot); err != nil {
				return err
			}

		case "bucket_correlation":
			if err := dec.Decode(&s.BucketCorrelation); err != nil {
				return err
			}

		case "bucket_count_ks_test":
			if err := dec.Decode(&s.BucketCountKsTest); err != nil {
				return err
			}

		case "bucket_script":
			if err := dec.Decode(&s.BucketScript); err != nil {
				return err
			}

		case "bucket_selector":
			if err := dec.Decode(&s.BucketSelector); err != nil {
				return err
			}

		case "bucket_sort":
			if err := dec.Decode(&s.BucketSort); err != nil {
				return err
			}

		case "cardinality":
			if err := dec.Decode(&s.Cardinality); err != nil {
				return err
			}

		case "categorize_text":
			if err := dec.Decode(&s.CategorizeText); err != nil {
				return err
			}

		case "children":
			if err := dec.Decode(&s.Children); err != nil {
				return err
			}

		case "composite":
			if err := dec.Decode(&s.Composite); err != nil {
				return err
			}

		case "cumulative_cardinality":
			if err := dec.Decode(&s.CumulativeCardinality); err != nil {
				return err
			}

		case "cumulative_sum":
			if err := dec.Decode(&s.CumulativeSum); err != nil {
				return err
			}

		case "date_histogram":
			if err := dec.Decode(&s.DateHistogram); err != nil {
				return err
			}

		case "date_range":
			if err := dec.Decode(&s.DateRange); err != nil {
				return err
			}

		case "derivative":
			if err := dec.Decode(&s.Derivative); err != nil {
				return err
			}

		case "diversified_sampler":
			if err := dec.Decode(&s.DiversifiedSampler); err != nil {
				return err
			}

		case "extended_stats":
			if err := dec.Decode(&s.ExtendedStats); err != nil {
				return err
			}

		case "extended_stats_bucket":
			if err := dec.Decode(&s.ExtendedStatsBucket); err != nil {
				return err
			}

		case "filter":
			if err := dec.Decode(&s.Filter); err != nil {
				return err
			}

		case "filters":
			if err := dec.Decode(&s.Filters); err != nil {
				return err
			}

		case "frequent_item_sets":
			if err := dec.Decode(&s.FrequentItemSets); err != nil {
				return err
			}

		case "geo_bounds":
			if err := dec.Decode(&s.GeoBounds); err != nil {
				return err
			}

		case "geo_centroid":
			if err := dec.Decode(&s.GeoCentroid); err != nil {
				return err
			}

		case "geo_distance":
			if err := dec.Decode(&s.GeoDistance); err != nil {
				return err
			}

		case "geo_line":
			if err := dec.Decode(&s.GeoLine); err != nil {
				return err
			}

		case "geohash_grid":
			if err := dec.Decode(&s.GeohashGrid); err != nil {
				return err
			}

		case "geohex_grid":
			if err := dec.Decode(&s.GeohexGrid); err != nil {
				return err
			}

		case "geotile_grid":
			if err := dec.Decode(&s.GeotileGrid); err != nil {
				return err
			}

		case "global":
			if err := dec.Decode(&s.Global); err != nil {
				return err
			}

		case "histogram":
			if err := dec.Decode(&s.Histogram); err != nil {
				return err
			}

		case "inference":
			if err := dec.Decode(&s.Inference); err != nil {
				return err
			}

		case "ip_prefix":
			if err := dec.Decode(&s.IpPrefix); err != nil {
				return err
			}

		case "ip_range":
			if err := dec.Decode(&s.IpRange); err != nil {
				return err
			}

		case "line":
			if err := dec.Decode(&s.Line); err != nil {
				return err
			}

		case "matrix_stats":
			if err := dec.Decode(&s.MatrixStats); err != nil {
				return err
			}

		case "max":
			if err := dec.Decode(&s.Max); err != nil {
				return err
			}

		case "max_bucket":
			if err := dec.Decode(&s.MaxBucket); err != nil {
				return err
			}

		case "median_absolute_deviation":
			if err := dec.Decode(&s.MedianAbsoluteDeviation); err != nil {
				return err
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "min":
			if err := dec.Decode(&s.Min); err != nil {
				return err
			}

		case "min_bucket":
			if err := dec.Decode(&s.MinBucket); err != nil {
				return err
			}

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return err
			}

		case "moving_avg":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			kind := make(map[string]string, 0)
			localDec := json.NewDecoder(source)
			localDec.Decode(&kind)
			source.Seek(0, io.SeekStart)

			switch kind["model"] {

			case "linear":
				o := NewLinearMovingAverageAggregation()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.MovingAvg = *o
			case "simple":
				o := NewSimpleMovingAverageAggregation()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.MovingAvg = *o
			case "ewma":
				o := NewEwmaMovingAverageAggregation()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.MovingAvg = *o
			case "holt":
				o := NewHoltMovingAverageAggregation()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.MovingAvg = *o
			case "holt_winters":
				o := NewHoltWintersMovingAverageAggregation()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.MovingAvg = *o
			default:
				if err := localDec.Decode(&s.MovingAvg); err != nil {
					return err
				}
			}

		case "moving_fn":
			if err := dec.Decode(&s.MovingFn); err != nil {
				return err
			}

		case "moving_percentiles":
			if err := dec.Decode(&s.MovingPercentiles); err != nil {
				return err
			}

		case "multi_terms":
			if err := dec.Decode(&s.MultiTerms); err != nil {
				return err
			}

		case "nested":
			if err := dec.Decode(&s.Nested); err != nil {
				return err
			}

		case "normalize":
			if err := dec.Decode(&s.Normalize); err != nil {
				return err
			}

		case "parent":
			if err := dec.Decode(&s.Parent); err != nil {
				return err
			}

		case "percentile_ranks":
			if err := dec.Decode(&s.PercentileRanks); err != nil {
				return err
			}

		case "percentiles":
			if err := dec.Decode(&s.Percentiles); err != nil {
				return err
			}

		case "percentiles_bucket":
			if err := dec.Decode(&s.PercentilesBucket); err != nil {
				return err
			}

		case "range":
			if err := dec.Decode(&s.Range); err != nil {
				return err
			}

		case "rare_terms":
			if err := dec.Decode(&s.RareTerms); err != nil {
				return err
			}

		case "rate":
			if err := dec.Decode(&s.Rate); err != nil {
				return err
			}

		case "reverse_nested":
			if err := dec.Decode(&s.ReverseNested); err != nil {
				return err
			}

		case "sampler":
			if err := dec.Decode(&s.Sampler); err != nil {
				return err
			}

		case "scripted_metric":
			if err := dec.Decode(&s.ScriptedMetric); err != nil {
				return err
			}

		case "serial_diff":
			if err := dec.Decode(&s.SerialDiff); err != nil {
				return err
			}

		case "significant_terms":
			if err := dec.Decode(&s.SignificantTerms); err != nil {
				return err
			}

		case "significant_text":
			if err := dec.Decode(&s.SignificantText); err != nil {
				return err
			}

		case "stats":
			if err := dec.Decode(&s.Stats); err != nil {
				return err
			}

		case "stats_bucket":
			if err := dec.Decode(&s.StatsBucket); err != nil {
				return err
			}

		case "string_stats":
			if err := dec.Decode(&s.StringStats); err != nil {
				return err
			}

		case "sum":
			if err := dec.Decode(&s.Sum); err != nil {
				return err
			}

		case "sum_bucket":
			if err := dec.Decode(&s.SumBucket); err != nil {
				return err
			}

		case "t_test":
			if err := dec.Decode(&s.TTest); err != nil {
				return err
			}

		case "terms":
			if err := dec.Decode(&s.Terms); err != nil {
				return err
			}

		case "top_hits":
			if err := dec.Decode(&s.TopHits); err != nil {
				return err
			}

		case "top_metrics":
			if err := dec.Decode(&s.TopMetrics); err != nil {
				return err
			}

		case "value_count":
			if err := dec.Decode(&s.ValueCount); err != nil {
				return err
			}

		case "variable_width_histogram":
			if err := dec.Decode(&s.VariableWidthHistogram); err != nil {
				return err
			}

		case "weighted_avg":
			if err := dec.Decode(&s.WeightedAvg); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewAggregations returns a Aggregations.
func NewAggregations() *Aggregations {
	r := &Aggregations{
		Aggregations: make(map[string]Aggregations, 0),
	}

	return r
}
