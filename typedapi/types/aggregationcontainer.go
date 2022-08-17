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

// AggregationContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/AggregationContainer.ts#L104-L207
type AggregationContainer struct {
	AdjacencyMatrix *AdjacencyMatrixAggregation `json:"adjacency_matrix,omitempty"`
	// Aggregations Sub-aggregations for this aggregation. Only applies to bucket aggregations.
	Aggregations            map[string]AggregationContainer     `json:"aggregations,omitempty"`
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
	Filter                  *QueryContainer                     `json:"filter,omitempty"`
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
	IpRange                 *IpRangeAggregation                 `json:"ip_range,omitempty"`
	Line                    *GeoLineAggregation                 `json:"line,omitempty"`
	MatrixStats             *MatrixStatsAggregation             `json:"matrix_stats,omitempty"`
	Max                     *MaxAggregation                     `json:"max,omitempty"`
	MaxBucket               *MaxBucketAggregation               `json:"max_bucket,omitempty"`
	MedianAbsoluteDeviation *MedianAbsoluteDeviationAggregation `json:"median_absolute_deviation,omitempty"`
	Meta                    *Metadata                           `json:"meta,omitempty"`
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

// AggregationContainerBuilder holds AggregationContainer struct and provides a builder API.
type AggregationContainerBuilder struct {
	v *AggregationContainer
}

// NewAggregationContainer provides a builder for the AggregationContainer struct.
func NewAggregationContainerBuilder() *AggregationContainerBuilder {
	r := AggregationContainerBuilder{
		&AggregationContainer{
			Aggregations: make(map[string]AggregationContainer, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the AggregationContainer struct
func (rb *AggregationContainerBuilder) Build() AggregationContainer {
	return *rb.v
}

func (rb *AggregationContainerBuilder) AdjacencyMatrix(adjacencymatrix *AdjacencyMatrixAggregationBuilder) *AggregationContainerBuilder {
	v := adjacencymatrix.Build()
	rb.v.AdjacencyMatrix = &v
	return rb
}

// Aggregations Sub-aggregations for this aggregation. Only applies to bucket aggregations.

func (rb *AggregationContainerBuilder) Aggregations(values map[string]*AggregationContainerBuilder) *AggregationContainerBuilder {
	tmp := make(map[string]AggregationContainer, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *AggregationContainerBuilder) AutoDateHistogram(autodatehistogram *AutoDateHistogramAggregationBuilder) *AggregationContainerBuilder {
	v := autodatehistogram.Build()
	rb.v.AutoDateHistogram = &v
	return rb
}

func (rb *AggregationContainerBuilder) Avg(avg *AverageAggregationBuilder) *AggregationContainerBuilder {
	v := avg.Build()
	rb.v.Avg = &v
	return rb
}

func (rb *AggregationContainerBuilder) AvgBucket(avgbucket *AverageBucketAggregationBuilder) *AggregationContainerBuilder {
	v := avgbucket.Build()
	rb.v.AvgBucket = &v
	return rb
}

func (rb *AggregationContainerBuilder) Boxplot(boxplot *BoxplotAggregationBuilder) *AggregationContainerBuilder {
	v := boxplot.Build()
	rb.v.Boxplot = &v
	return rb
}

func (rb *AggregationContainerBuilder) BucketCorrelation(bucketcorrelation *BucketCorrelationAggregationBuilder) *AggregationContainerBuilder {
	v := bucketcorrelation.Build()
	rb.v.BucketCorrelation = &v
	return rb
}

func (rb *AggregationContainerBuilder) BucketCountKsTest(bucketcountkstest *BucketKsAggregationBuilder) *AggregationContainerBuilder {
	v := bucketcountkstest.Build()
	rb.v.BucketCountKsTest = &v
	return rb
}

func (rb *AggregationContainerBuilder) BucketScript(bucketscript *BucketScriptAggregationBuilder) *AggregationContainerBuilder {
	v := bucketscript.Build()
	rb.v.BucketScript = &v
	return rb
}

func (rb *AggregationContainerBuilder) BucketSelector(bucketselector *BucketSelectorAggregationBuilder) *AggregationContainerBuilder {
	v := bucketselector.Build()
	rb.v.BucketSelector = &v
	return rb
}

func (rb *AggregationContainerBuilder) BucketSort(bucketsort *BucketSortAggregationBuilder) *AggregationContainerBuilder {
	v := bucketsort.Build()
	rb.v.BucketSort = &v
	return rb
}

func (rb *AggregationContainerBuilder) Cardinality(cardinality *CardinalityAggregationBuilder) *AggregationContainerBuilder {
	v := cardinality.Build()
	rb.v.Cardinality = &v
	return rb
}

func (rb *AggregationContainerBuilder) CategorizeText(categorizetext *CategorizeTextAggregationBuilder) *AggregationContainerBuilder {
	v := categorizetext.Build()
	rb.v.CategorizeText = &v
	return rb
}

func (rb *AggregationContainerBuilder) Children(children *ChildrenAggregationBuilder) *AggregationContainerBuilder {
	v := children.Build()
	rb.v.Children = &v
	return rb
}

func (rb *AggregationContainerBuilder) Composite(composite *CompositeAggregationBuilder) *AggregationContainerBuilder {
	v := composite.Build()
	rb.v.Composite = &v
	return rb
}

func (rb *AggregationContainerBuilder) CumulativeCardinality(cumulativecardinality *CumulativeCardinalityAggregationBuilder) *AggregationContainerBuilder {
	v := cumulativecardinality.Build()
	rb.v.CumulativeCardinality = &v
	return rb
}

func (rb *AggregationContainerBuilder) CumulativeSum(cumulativesum *CumulativeSumAggregationBuilder) *AggregationContainerBuilder {
	v := cumulativesum.Build()
	rb.v.CumulativeSum = &v
	return rb
}

func (rb *AggregationContainerBuilder) DateHistogram(datehistogram *DateHistogramAggregationBuilder) *AggregationContainerBuilder {
	v := datehistogram.Build()
	rb.v.DateHistogram = &v
	return rb
}

func (rb *AggregationContainerBuilder) DateRange(daterange *DateRangeAggregationBuilder) *AggregationContainerBuilder {
	v := daterange.Build()
	rb.v.DateRange = &v
	return rb
}

func (rb *AggregationContainerBuilder) Derivative(derivative *DerivativeAggregationBuilder) *AggregationContainerBuilder {
	v := derivative.Build()
	rb.v.Derivative = &v
	return rb
}

func (rb *AggregationContainerBuilder) DiversifiedSampler(diversifiedsampler *DiversifiedSamplerAggregationBuilder) *AggregationContainerBuilder {
	v := diversifiedsampler.Build()
	rb.v.DiversifiedSampler = &v
	return rb
}

func (rb *AggregationContainerBuilder) ExtendedStats(extendedstats *ExtendedStatsAggregationBuilder) *AggregationContainerBuilder {
	v := extendedstats.Build()
	rb.v.ExtendedStats = &v
	return rb
}

func (rb *AggregationContainerBuilder) ExtendedStatsBucket(extendedstatsbucket *ExtendedStatsBucketAggregationBuilder) *AggregationContainerBuilder {
	v := extendedstatsbucket.Build()
	rb.v.ExtendedStatsBucket = &v
	return rb
}

func (rb *AggregationContainerBuilder) Filter(filter *QueryContainerBuilder) *AggregationContainerBuilder {
	v := filter.Build()
	rb.v.Filter = &v
	return rb
}

func (rb *AggregationContainerBuilder) Filters(filters *FiltersAggregationBuilder) *AggregationContainerBuilder {
	v := filters.Build()
	rb.v.Filters = &v
	return rb
}

func (rb *AggregationContainerBuilder) GeoBounds(geobounds *GeoBoundsAggregationBuilder) *AggregationContainerBuilder {
	v := geobounds.Build()
	rb.v.GeoBounds = &v
	return rb
}

func (rb *AggregationContainerBuilder) GeoCentroid(geocentroid *GeoCentroidAggregationBuilder) *AggregationContainerBuilder {
	v := geocentroid.Build()
	rb.v.GeoCentroid = &v
	return rb
}

func (rb *AggregationContainerBuilder) GeoDistance(geodistance *GeoDistanceAggregationBuilder) *AggregationContainerBuilder {
	v := geodistance.Build()
	rb.v.GeoDistance = &v
	return rb
}

func (rb *AggregationContainerBuilder) GeoLine(geoline *GeoLineAggregationBuilder) *AggregationContainerBuilder {
	v := geoline.Build()
	rb.v.GeoLine = &v
	return rb
}

func (rb *AggregationContainerBuilder) GeohashGrid(geohashgrid *GeoHashGridAggregationBuilder) *AggregationContainerBuilder {
	v := geohashgrid.Build()
	rb.v.GeohashGrid = &v
	return rb
}

func (rb *AggregationContainerBuilder) GeohexGrid(geohexgrid *GeohexGridAggregationBuilder) *AggregationContainerBuilder {
	v := geohexgrid.Build()
	rb.v.GeohexGrid = &v
	return rb
}

func (rb *AggregationContainerBuilder) GeotileGrid(geotilegrid *GeoTileGridAggregationBuilder) *AggregationContainerBuilder {
	v := geotilegrid.Build()
	rb.v.GeotileGrid = &v
	return rb
}

func (rb *AggregationContainerBuilder) Global(global *GlobalAggregationBuilder) *AggregationContainerBuilder {
	v := global.Build()
	rb.v.Global = &v
	return rb
}

func (rb *AggregationContainerBuilder) Histogram(histogram *HistogramAggregationBuilder) *AggregationContainerBuilder {
	v := histogram.Build()
	rb.v.Histogram = &v
	return rb
}

func (rb *AggregationContainerBuilder) Inference(inference *InferenceAggregationBuilder) *AggregationContainerBuilder {
	v := inference.Build()
	rb.v.Inference = &v
	return rb
}

func (rb *AggregationContainerBuilder) IpRange(iprange *IpRangeAggregationBuilder) *AggregationContainerBuilder {
	v := iprange.Build()
	rb.v.IpRange = &v
	return rb
}

func (rb *AggregationContainerBuilder) Line(line *GeoLineAggregationBuilder) *AggregationContainerBuilder {
	v := line.Build()
	rb.v.Line = &v
	return rb
}

func (rb *AggregationContainerBuilder) MatrixStats(matrixstats *MatrixStatsAggregationBuilder) *AggregationContainerBuilder {
	v := matrixstats.Build()
	rb.v.MatrixStats = &v
	return rb
}

func (rb *AggregationContainerBuilder) Max(max *MaxAggregationBuilder) *AggregationContainerBuilder {
	v := max.Build()
	rb.v.Max = &v
	return rb
}

func (rb *AggregationContainerBuilder) MaxBucket(maxbucket *MaxBucketAggregationBuilder) *AggregationContainerBuilder {
	v := maxbucket.Build()
	rb.v.MaxBucket = &v
	return rb
}

func (rb *AggregationContainerBuilder) MedianAbsoluteDeviation(medianabsolutedeviation *MedianAbsoluteDeviationAggregationBuilder) *AggregationContainerBuilder {
	v := medianabsolutedeviation.Build()
	rb.v.MedianAbsoluteDeviation = &v
	return rb
}

func (rb *AggregationContainerBuilder) Meta(meta *MetadataBuilder) *AggregationContainerBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *AggregationContainerBuilder) Min(min *MinAggregationBuilder) *AggregationContainerBuilder {
	v := min.Build()
	rb.v.Min = &v
	return rb
}

func (rb *AggregationContainerBuilder) MinBucket(minbucket *MinBucketAggregationBuilder) *AggregationContainerBuilder {
	v := minbucket.Build()
	rb.v.MinBucket = &v
	return rb
}

func (rb *AggregationContainerBuilder) Missing(missing *MissingAggregationBuilder) *AggregationContainerBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *AggregationContainerBuilder) MovingAvg(movingavg *MovingAverageAggregationBuilder) *AggregationContainerBuilder {
	v := movingavg.Build()
	rb.v.MovingAvg = &v
	return rb
}

func (rb *AggregationContainerBuilder) MovingFn(movingfn *MovingFunctionAggregationBuilder) *AggregationContainerBuilder {
	v := movingfn.Build()
	rb.v.MovingFn = &v
	return rb
}

func (rb *AggregationContainerBuilder) MovingPercentiles(movingpercentiles *MovingPercentilesAggregationBuilder) *AggregationContainerBuilder {
	v := movingpercentiles.Build()
	rb.v.MovingPercentiles = &v
	return rb
}

func (rb *AggregationContainerBuilder) MultiTerms(multiterms *MultiTermsAggregationBuilder) *AggregationContainerBuilder {
	v := multiterms.Build()
	rb.v.MultiTerms = &v
	return rb
}

func (rb *AggregationContainerBuilder) Nested(nested *NestedAggregationBuilder) *AggregationContainerBuilder {
	v := nested.Build()
	rb.v.Nested = &v
	return rb
}

func (rb *AggregationContainerBuilder) Normalize(normalize *NormalizeAggregationBuilder) *AggregationContainerBuilder {
	v := normalize.Build()
	rb.v.Normalize = &v
	return rb
}

func (rb *AggregationContainerBuilder) Parent(parent *ParentAggregationBuilder) *AggregationContainerBuilder {
	v := parent.Build()
	rb.v.Parent = &v
	return rb
}

func (rb *AggregationContainerBuilder) PercentileRanks(percentileranks *PercentileRanksAggregationBuilder) *AggregationContainerBuilder {
	v := percentileranks.Build()
	rb.v.PercentileRanks = &v
	return rb
}

func (rb *AggregationContainerBuilder) Percentiles(percentiles *PercentilesAggregationBuilder) *AggregationContainerBuilder {
	v := percentiles.Build()
	rb.v.Percentiles = &v
	return rb
}

func (rb *AggregationContainerBuilder) PercentilesBucket(percentilesbucket *PercentilesBucketAggregationBuilder) *AggregationContainerBuilder {
	v := percentilesbucket.Build()
	rb.v.PercentilesBucket = &v
	return rb
}

func (rb *AggregationContainerBuilder) Range_(range_ *RangeAggregationBuilder) *AggregationContainerBuilder {
	v := range_.Build()
	rb.v.Range = &v
	return rb
}

func (rb *AggregationContainerBuilder) RareTerms(rareterms *RareTermsAggregationBuilder) *AggregationContainerBuilder {
	v := rareterms.Build()
	rb.v.RareTerms = &v
	return rb
}

func (rb *AggregationContainerBuilder) Rate(rate *RateAggregationBuilder) *AggregationContainerBuilder {
	v := rate.Build()
	rb.v.Rate = &v
	return rb
}

func (rb *AggregationContainerBuilder) ReverseNested(reversenested *ReverseNestedAggregationBuilder) *AggregationContainerBuilder {
	v := reversenested.Build()
	rb.v.ReverseNested = &v
	return rb
}

func (rb *AggregationContainerBuilder) Sampler(sampler *SamplerAggregationBuilder) *AggregationContainerBuilder {
	v := sampler.Build()
	rb.v.Sampler = &v
	return rb
}

func (rb *AggregationContainerBuilder) ScriptedMetric(scriptedmetric *ScriptedMetricAggregationBuilder) *AggregationContainerBuilder {
	v := scriptedmetric.Build()
	rb.v.ScriptedMetric = &v
	return rb
}

func (rb *AggregationContainerBuilder) SerialDiff(serialdiff *SerialDifferencingAggregationBuilder) *AggregationContainerBuilder {
	v := serialdiff.Build()
	rb.v.SerialDiff = &v
	return rb
}

func (rb *AggregationContainerBuilder) SignificantTerms(significantterms *SignificantTermsAggregationBuilder) *AggregationContainerBuilder {
	v := significantterms.Build()
	rb.v.SignificantTerms = &v
	return rb
}

func (rb *AggregationContainerBuilder) SignificantText(significanttext *SignificantTextAggregationBuilder) *AggregationContainerBuilder {
	v := significanttext.Build()
	rb.v.SignificantText = &v
	return rb
}

func (rb *AggregationContainerBuilder) Stats(stats *StatsAggregationBuilder) *AggregationContainerBuilder {
	v := stats.Build()
	rb.v.Stats = &v
	return rb
}

func (rb *AggregationContainerBuilder) StatsBucket(statsbucket *StatsBucketAggregationBuilder) *AggregationContainerBuilder {
	v := statsbucket.Build()
	rb.v.StatsBucket = &v
	return rb
}

func (rb *AggregationContainerBuilder) StringStats(stringstats *StringStatsAggregationBuilder) *AggregationContainerBuilder {
	v := stringstats.Build()
	rb.v.StringStats = &v
	return rb
}

func (rb *AggregationContainerBuilder) Sum(sum *SumAggregationBuilder) *AggregationContainerBuilder {
	v := sum.Build()
	rb.v.Sum = &v
	return rb
}

func (rb *AggregationContainerBuilder) SumBucket(sumbucket *SumBucketAggregationBuilder) *AggregationContainerBuilder {
	v := sumbucket.Build()
	rb.v.SumBucket = &v
	return rb
}

func (rb *AggregationContainerBuilder) TTest(ttest *TTestAggregationBuilder) *AggregationContainerBuilder {
	v := ttest.Build()
	rb.v.TTest = &v
	return rb
}

func (rb *AggregationContainerBuilder) Terms(terms *TermsAggregationBuilder) *AggregationContainerBuilder {
	v := terms.Build()
	rb.v.Terms = &v
	return rb
}

func (rb *AggregationContainerBuilder) TopHits(tophits *TopHitsAggregationBuilder) *AggregationContainerBuilder {
	v := tophits.Build()
	rb.v.TopHits = &v
	return rb
}

func (rb *AggregationContainerBuilder) TopMetrics(topmetrics *TopMetricsAggregationBuilder) *AggregationContainerBuilder {
	v := topmetrics.Build()
	rb.v.TopMetrics = &v
	return rb
}

func (rb *AggregationContainerBuilder) ValueCount(valuecount *ValueCountAggregationBuilder) *AggregationContainerBuilder {
	v := valuecount.Build()
	rb.v.ValueCount = &v
	return rb
}

func (rb *AggregationContainerBuilder) VariableWidthHistogram(variablewidthhistogram *VariableWidthHistogramAggregationBuilder) *AggregationContainerBuilder {
	v := variablewidthhistogram.Build()
	rb.v.VariableWidthHistogram = &v
	return rb
}

func (rb *AggregationContainerBuilder) WeightedAvg(weightedavg *WeightedAverageAggregationBuilder) *AggregationContainerBuilder {
	v := weightedavg.Build()
	rb.v.WeightedAvg = &v
	return rb
}
