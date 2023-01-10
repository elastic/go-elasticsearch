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

// Aggregations type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/aggregations/AggregationContainer.ts#L105-L209
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
	Meta                    map[string]interface{}              `json:"meta,omitempty"`
	Min                     *MinAggregation                     `json:"min,omitempty"`
	MinBucket               *MinBucketAggregation               `json:"min_bucket,omitempty"`
	Missing                 *MissingAggregation                 `json:"missing,omitempty"`
	MovingAvg               *MovingAverageAggregation           `json:"moving_avg,omitempty"`
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

// NewAggregations returns a Aggregations.
func NewAggregations() *Aggregations {
	r := &Aggregations{
		Aggregations: make(map[string]Aggregations, 0),
	}

	return r
}
