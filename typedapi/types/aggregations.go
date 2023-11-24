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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// Aggregations type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/aggregations/AggregationContainer.ts#L106-L515
type Aggregations struct {
	// AdjacencyMatrix A bucket aggregation returning a form of adjacency matrix.
	// The request provides a collection of named filter expressions, similar to the
	// `filters` aggregation.
	// Each bucket in the response represents a non-empty cell in the matrix of
	// intersecting filters.
	AdjacencyMatrix *AdjacencyMatrixAggregation `json:"adjacency_matrix,omitempty"`
	// Aggregations Sub-aggregations for this aggregation.
	// Only applies to bucket aggregations.
	Aggregations map[string]Aggregations `json:"aggregations,omitempty"`
	// AutoDateHistogram A multi-bucket aggregation similar to the date histogram, except instead of
	// providing an interval to use as the width of each bucket, a target number of
	// buckets is provided.
	AutoDateHistogram *AutoDateHistogramAggregation `json:"auto_date_histogram,omitempty"`
	// Avg A single-value metrics aggregation that computes the average of numeric
	// values that are extracted from the aggregated documents.
	Avg *AverageAggregation `json:"avg,omitempty"`
	// AvgBucket A sibling pipeline aggregation which calculates the mean value of a specified
	// metric in a sibling aggregation.
	// The specified metric must be numeric and the sibling aggregation must be a
	// multi-bucket aggregation.
	AvgBucket *AverageBucketAggregation `json:"avg_bucket,omitempty"`
	// Boxplot A metrics aggregation that computes a box plot of numeric values extracted
	// from the aggregated documents.
	Boxplot *BoxplotAggregation `json:"boxplot,omitempty"`
	// BucketCorrelation A sibling pipeline aggregation which runs a correlation function on the
	// configured sibling multi-bucket aggregation.
	BucketCorrelation *BucketCorrelationAggregation `json:"bucket_correlation,omitempty"`
	// BucketCountKsTest A sibling pipeline aggregation which runs a two sample Kolmogorov–Smirnov
	// test ("K-S test") against a provided distribution and the distribution
	// implied by the documents counts in the configured sibling aggregation.
	BucketCountKsTest *BucketKsAggregation `json:"bucket_count_ks_test,omitempty"`
	// BucketScript A parent pipeline aggregation which runs a script which can perform per
	// bucket computations on metrics in the parent multi-bucket aggregation.
	BucketScript *BucketScriptAggregation `json:"bucket_script,omitempty"`
	// BucketSelector A parent pipeline aggregation which runs a script to determine whether the
	// current bucket will be retained in the parent multi-bucket aggregation.
	BucketSelector *BucketSelectorAggregation `json:"bucket_selector,omitempty"`
	// BucketSort A parent pipeline aggregation which sorts the buckets of its parent
	// multi-bucket aggregation.
	BucketSort *BucketSortAggregation `json:"bucket_sort,omitempty"`
	// Cardinality A single-value metrics aggregation that calculates an approximate count of
	// distinct values.
	Cardinality *CardinalityAggregation `json:"cardinality,omitempty"`
	// CategorizeText A multi-bucket aggregation that groups semi-structured text into buckets.
	CategorizeText *CategorizeTextAggregation `json:"categorize_text,omitempty"`
	// Children A single bucket aggregation that selects child documents that have the
	// specified type, as defined in a `join` field.
	Children *ChildrenAggregation `json:"children,omitempty"`
	// Composite A multi-bucket aggregation that creates composite buckets from different
	// sources.
	// Unlike the other multi-bucket aggregations, you can use the `composite`
	// aggregation to paginate *all* buckets from a multi-level aggregation
	// efficiently.
	Composite *CompositeAggregation `json:"composite,omitempty"`
	// CumulativeCardinality A parent pipeline aggregation which calculates the cumulative cardinality in
	// a parent `histogram` or `date_histogram` aggregation.
	CumulativeCardinality *CumulativeCardinalityAggregation `json:"cumulative_cardinality,omitempty"`
	// CumulativeSum A parent pipeline aggregation which calculates the cumulative sum of a
	// specified metric in a parent `histogram` or `date_histogram` aggregation.
	CumulativeSum *CumulativeSumAggregation `json:"cumulative_sum,omitempty"`
	// DateHistogram A multi-bucket values source based aggregation that can be applied on date
	// values or date range values extracted from the documents.
	// It dynamically builds fixed size (interval) buckets over the values.
	DateHistogram *DateHistogramAggregation `json:"date_histogram,omitempty"`
	// DateRange A multi-bucket value source based aggregation that enables the user to define
	// a set of date ranges - each representing a bucket.
	DateRange *DateRangeAggregation `json:"date_range,omitempty"`
	// Derivative A parent pipeline aggregation which calculates the derivative of a specified
	// metric in a parent `histogram` or `date_histogram` aggregation.
	Derivative *DerivativeAggregation `json:"derivative,omitempty"`
	// DiversifiedSampler A filtering aggregation used to limit any sub aggregations' processing to a
	// sample of the top-scoring documents.
	// Similar to the `sampler` aggregation, but adds the ability to limit the
	// number of matches that share a common value.
	DiversifiedSampler *DiversifiedSamplerAggregation `json:"diversified_sampler,omitempty"`
	// ExtendedStats A multi-value metrics aggregation that computes stats over numeric values
	// extracted from the aggregated documents.
	ExtendedStats *ExtendedStatsAggregation `json:"extended_stats,omitempty"`
	// ExtendedStatsBucket A sibling pipeline aggregation which calculates a variety of stats across all
	// bucket of a specified metric in a sibling aggregation.
	ExtendedStatsBucket *ExtendedStatsBucketAggregation `json:"extended_stats_bucket,omitempty"`
	// Filter A single bucket aggregation that narrows the set of documents to those that
	// match a query.
	Filter *Query `json:"filter,omitempty"`
	// Filters A multi-bucket aggregation where each bucket contains the documents that
	// match a query.
	Filters *FiltersAggregation `json:"filters,omitempty"`
	// FrequentItemSets A bucket aggregation which finds frequent item sets, a form of association
	// rules mining that identifies items that often occur together.
	FrequentItemSets *FrequentItemSetsAggregation `json:"frequent_item_sets,omitempty"`
	// GeoBounds A metric aggregation that computes the geographic bounding box containing all
	// values for a Geopoint or Geoshape field.
	GeoBounds *GeoBoundsAggregation `json:"geo_bounds,omitempty"`
	// GeoCentroid A metric aggregation that computes the weighted centroid from all coordinate
	// values for geo fields.
	GeoCentroid *GeoCentroidAggregation `json:"geo_centroid,omitempty"`
	// GeoDistance A multi-bucket aggregation that works on `geo_point` fields.
	// Evaluates the distance of each document value from an origin point and
	// determines the buckets it belongs to, based on ranges defined in the request.
	GeoDistance *GeoDistanceAggregation `json:"geo_distance,omitempty"`
	// GeoLine Aggregates all `geo_point` values within a bucket into a `LineString` ordered
	// by the chosen sort field.
	GeoLine *GeoLineAggregation `json:"geo_line,omitempty"`
	// GeohashGrid A multi-bucket aggregation that groups `geo_point` and `geo_shape` values
	// into buckets that represent a grid.
	// Each cell is labeled using a geohash which is of user-definable precision.
	GeohashGrid *GeoHashGridAggregation `json:"geohash_grid,omitempty"`
	// GeohexGrid A multi-bucket aggregation that groups `geo_point` and `geo_shape` values
	// into buckets that represent a grid.
	// Each cell corresponds to a H3 cell index and is labeled using the H3Index
	// representation.
	GeohexGrid *GeohexGridAggregation `json:"geohex_grid,omitempty"`
	// GeotileGrid A multi-bucket aggregation that groups `geo_point` and `geo_shape` values
	// into buckets that represent a grid.
	// Each cell corresponds to a map tile as used by many online map sites.
	GeotileGrid *GeoTileGridAggregation `json:"geotile_grid,omitempty"`
	// Global Defines a single bucket of all the documents within the search execution
	// context.
	// This context is defined by the indices and the document types you’re
	// searching on, but is not influenced by the search query itself.
	Global *GlobalAggregation `json:"global,omitempty"`
	// Histogram A multi-bucket values source based aggregation that can be applied on numeric
	// values or numeric range values extracted from the documents.
	// It dynamically builds fixed size (interval) buckets over the values.
	Histogram *HistogramAggregation `json:"histogram,omitempty"`
	// Inference A parent pipeline aggregation which loads a pre-trained model and performs
	// inference on the collated result fields from the parent bucket aggregation.
	Inference *InferenceAggregation `json:"inference,omitempty"`
	// IpPrefix A bucket aggregation that groups documents based on the network or
	// sub-network of an IP address.
	IpPrefix *IpPrefixAggregation `json:"ip_prefix,omitempty"`
	// IpRange A multi-bucket value source based aggregation that enables the user to define
	// a set of IP ranges - each representing a bucket.
	IpRange *IpRangeAggregation `json:"ip_range,omitempty"`
	Line    *GeoLineAggregation `json:"line,omitempty"`
	// MatrixStats A numeric aggregation that computes the following statistics over a set of
	// document fields: `count`, `mean`, `variance`, `skewness`, `kurtosis`,
	// `covariance`, and `covariance`.
	MatrixStats *MatrixStatsAggregation `json:"matrix_stats,omitempty"`
	// Max A single-value metrics aggregation that returns the maximum value among the
	// numeric values extracted from the aggregated documents.
	Max *MaxAggregation `json:"max,omitempty"`
	// MaxBucket A sibling pipeline aggregation which identifies the bucket(s) with the
	// maximum value of a specified metric in a sibling aggregation and outputs both
	// the value and the key(s) of the bucket(s).
	MaxBucket *MaxBucketAggregation `json:"max_bucket,omitempty"`
	// MedianAbsoluteDeviation A single-value aggregation that approximates the median absolute deviation of
	// its search results.
	MedianAbsoluteDeviation *MedianAbsoluteDeviationAggregation `json:"median_absolute_deviation,omitempty"`
	Meta                    Metadata                            `json:"meta,omitempty"`
	// Min A single-value metrics aggregation that returns the minimum value among
	// numeric values extracted from the aggregated documents.
	Min *MinAggregation `json:"min,omitempty"`
	// MinBucket A sibling pipeline aggregation which identifies the bucket(s) with the
	// minimum value of a specified metric in a sibling aggregation and outputs both
	// the value and the key(s) of the bucket(s).
	MinBucket *MinBucketAggregation `json:"min_bucket,omitempty"`
	// Missing A field data based single bucket aggregation, that creates a bucket of all
	// documents in the current document set context that are missing a field value
	// (effectively, missing a field or having the configured NULL value set).
	Missing   *MissingAggregation      `json:"missing,omitempty"`
	MovingAvg MovingAverageAggregation `json:"moving_avg,omitempty"`
	// MovingFn Given an ordered series of data, "slides" a window across the data and runs a
	// custom script on each window of data.
	// For convenience, a number of common functions are predefined such as `min`,
	// `max`, and moving averages.
	MovingFn *MovingFunctionAggregation `json:"moving_fn,omitempty"`
	// MovingPercentiles Given an ordered series of percentiles, "slides" a window across those
	// percentiles and computes cumulative percentiles.
	MovingPercentiles *MovingPercentilesAggregation `json:"moving_percentiles,omitempty"`
	// MultiTerms A multi-bucket value source based aggregation where buckets are dynamically
	// built - one per unique set of values.
	MultiTerms *MultiTermsAggregation `json:"multi_terms,omitempty"`
	// Nested A special single bucket aggregation that enables aggregating nested
	// documents.
	Nested *NestedAggregation `json:"nested,omitempty"`
	// Normalize A parent pipeline aggregation which calculates the specific
	// normalized/rescaled value for a specific bucket value.
	Normalize *NormalizeAggregation `json:"normalize,omitempty"`
	// Parent A special single bucket aggregation that selects parent documents that have
	// the specified type, as defined in a `join` field.
	Parent *ParentAggregation `json:"parent,omitempty"`
	// PercentileRanks A multi-value metrics aggregation that calculates one or more percentile
	// ranks over numeric values extracted from the aggregated documents.
	PercentileRanks *PercentileRanksAggregation `json:"percentile_ranks,omitempty"`
	// Percentiles A multi-value metrics aggregation that calculates one or more percentiles
	// over numeric values extracted from the aggregated documents.
	Percentiles *PercentilesAggregation `json:"percentiles,omitempty"`
	// PercentilesBucket A sibling pipeline aggregation which calculates percentiles across all bucket
	// of a specified metric in a sibling aggregation.
	PercentilesBucket *PercentilesBucketAggregation `json:"percentiles_bucket,omitempty"`
	// Range A multi-bucket value source based aggregation that enables the user to define
	// a set of ranges - each representing a bucket.
	Range *RangeAggregation `json:"range,omitempty"`
	// RareTerms A multi-bucket value source based aggregation which finds "rare" terms —
	// terms that are at the long-tail of the distribution and are not frequent.
	RareTerms *RareTermsAggregation `json:"rare_terms,omitempty"`
	// Rate Calculates a rate of documents or a field in each bucket.
	// Can only be used inside a `date_histogram` or `composite` aggregation.
	Rate *RateAggregation `json:"rate,omitempty"`
	// ReverseNested A special single bucket aggregation that enables aggregating on parent
	// documents from nested documents.
	// Should only be defined inside a `nested` aggregation.
	ReverseNested *ReverseNestedAggregation `json:"reverse_nested,omitempty"`
	// Sampler A filtering aggregation used to limit any sub aggregations' processing to a
	// sample of the top-scoring documents.
	Sampler *SamplerAggregation `json:"sampler,omitempty"`
	// ScriptedMetric A metric aggregation that uses scripts to provide a metric output.
	ScriptedMetric *ScriptedMetricAggregation `json:"scripted_metric,omitempty"`
	// SerialDiff An aggregation that subtracts values in a time series from themselves at
	// different time lags or periods.
	SerialDiff *SerialDifferencingAggregation `json:"serial_diff,omitempty"`
	// SignificantTerms Returns interesting or unusual occurrences of terms in a set.
	SignificantTerms *SignificantTermsAggregation `json:"significant_terms,omitempty"`
	// SignificantText Returns interesting or unusual occurrences of free-text terms in a set.
	SignificantText *SignificantTextAggregation `json:"significant_text,omitempty"`
	// Stats A multi-value metrics aggregation that computes stats over numeric values
	// extracted from the aggregated documents.
	Stats *StatsAggregation `json:"stats,omitempty"`
	// StatsBucket A sibling pipeline aggregation which calculates a variety of stats across all
	// bucket of a specified metric in a sibling aggregation.
	StatsBucket *StatsBucketAggregation `json:"stats_bucket,omitempty"`
	// StringStats A multi-value metrics aggregation that computes statistics over string values
	// extracted from the aggregated documents.
	StringStats *StringStatsAggregation `json:"string_stats,omitempty"`
	// Sum A single-value metrics aggregation that sums numeric values that are
	// extracted from the aggregated documents.
	Sum *SumAggregation `json:"sum,omitempty"`
	// SumBucket A sibling pipeline aggregation which calculates the sum of a specified metric
	// across all buckets in a sibling aggregation.
	SumBucket *SumBucketAggregation `json:"sum_bucket,omitempty"`
	// TTest A metrics aggregation that performs a statistical hypothesis test in which
	// the test statistic follows a Student’s t-distribution under the null
	// hypothesis on numeric values extracted from the aggregated documents.
	TTest *TTestAggregation `json:"t_test,omitempty"`
	// Terms A multi-bucket value source based aggregation where buckets are dynamically
	// built - one per unique value.
	Terms *TermsAggregation `json:"terms,omitempty"`
	// TopHits A metric aggregation that returns the top matching documents per bucket.
	TopHits *TopHitsAggregation `json:"top_hits,omitempty"`
	// TopMetrics A metric aggregation that selects metrics from the document with the largest
	// or smallest sort value.
	TopMetrics *TopMetricsAggregation `json:"top_metrics,omitempty"`
	// ValueCount A single-value metrics aggregation that counts the number of values that are
	// extracted from the aggregated documents.
	ValueCount *ValueCountAggregation `json:"value_count,omitempty"`
	// VariableWidthHistogram A multi-bucket aggregation similar to the histogram, except instead of
	// providing an interval to use as the width of each bucket, a target number of
	// buckets is provided.
	VariableWidthHistogram *VariableWidthHistogramAggregation `json:"variable_width_histogram,omitempty"`
	// WeightedAvg A single-value metrics aggregation that computes the weighted average of
	// numeric values that are extracted from the aggregated documents.
	WeightedAvg *WeightedAverageAggregation `json:"weighted_avg,omitempty"`
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
