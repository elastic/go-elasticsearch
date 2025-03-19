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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _aggregations struct {
	v *types.Aggregations
}

func NewAggregations() *_aggregations {
	return &_aggregations{v: types.NewAggregations()}
}

// AdditionalAggregationsProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_aggregations) AdditionalAggregationsProperty(key string, value json.RawMessage) *_aggregations {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalAggregationsProperty = tmp
	return s
}

// A bucket aggregation returning a form of adjacency matrix.
// The request provides a collection of named filter expressions, similar to the
// `filters` aggregation.
// Each bucket in the response represents a non-empty cell in the matrix of
// intersecting filters.
func (s *_aggregations) AdjacencyMatrix(adjacencymatrix types.AdjacencyMatrixAggregationVariant) *_aggregations {

	s.v.AdjacencyMatrix = adjacencymatrix.AdjacencyMatrixAggregationCaster()

	return s
}

// Sub-aggregations for this aggregation.
// Only applies to bucket aggregations.
func (s *_aggregations) Aggregations(aggregations map[string]types.Aggregations) *_aggregations {

	s.v.Aggregations = aggregations
	return s
}

func (s *_aggregations) AddAggregation(key string, value types.AggregationsVariant) *_aggregations {

	var tmp map[string]types.Aggregations
	if s.v.Aggregations == nil {
		s.v.Aggregations = make(map[string]types.Aggregations)
	} else {
		tmp = s.v.Aggregations
	}

	tmp[key] = *value.AggregationsCaster()

	s.v.Aggregations = tmp
	return s
}

// A multi-bucket aggregation similar to the date histogram, except instead of
// providing an interval to use as the width of each bucket, a target number of
// buckets is provided.
func (s *_aggregations) AutoDateHistogram(autodatehistogram types.AutoDateHistogramAggregationVariant) *_aggregations {

	s.v.AutoDateHistogram = autodatehistogram.AutoDateHistogramAggregationCaster()

	return s
}

// A single-value metrics aggregation that computes the average of numeric
// values that are extracted from the aggregated documents.
func (s *_aggregations) Avg(avg types.AverageAggregationVariant) *_aggregations {

	s.v.Avg = avg.AverageAggregationCaster()

	return s
}

// A sibling pipeline aggregation which calculates the mean value of a specified
// metric in a sibling aggregation.
// The specified metric must be numeric and the sibling aggregation must be a
// multi-bucket aggregation.
func (s *_aggregations) AvgBucket(avgbucket types.AverageBucketAggregationVariant) *_aggregations {

	s.v.AvgBucket = avgbucket.AverageBucketAggregationCaster()

	return s
}

// A metrics aggregation that computes a box plot of numeric values extracted
// from the aggregated documents.
func (s *_aggregations) Boxplot(boxplot types.BoxplotAggregationVariant) *_aggregations {

	s.v.Boxplot = boxplot.BoxplotAggregationCaster()

	return s
}

// A sibling pipeline aggregation which runs a correlation function on the
// configured sibling multi-bucket aggregation.
func (s *_aggregations) BucketCorrelation(bucketcorrelation types.BucketCorrelationAggregationVariant) *_aggregations {

	s.v.BucketCorrelation = bucketcorrelation.BucketCorrelationAggregationCaster()

	return s
}

// A sibling pipeline aggregation which runs a two sample Kolmogorov–Smirnov
// test ("K-S test") against a provided distribution and the distribution
// implied by the documents counts in the configured sibling aggregation.
func (s *_aggregations) BucketCountKsTest(bucketcountkstest types.BucketKsAggregationVariant) *_aggregations {

	s.v.BucketCountKsTest = bucketcountkstest.BucketKsAggregationCaster()

	return s
}

// A parent pipeline aggregation which runs a script which can perform per
// bucket computations on metrics in the parent multi-bucket aggregation.
func (s *_aggregations) BucketScript(bucketscript types.BucketScriptAggregationVariant) *_aggregations {

	s.v.BucketScript = bucketscript.BucketScriptAggregationCaster()

	return s
}

// A parent pipeline aggregation which runs a script to determine whether the
// current bucket will be retained in the parent multi-bucket aggregation.
func (s *_aggregations) BucketSelector(bucketselector types.BucketSelectorAggregationVariant) *_aggregations {

	s.v.BucketSelector = bucketselector.BucketSelectorAggregationCaster()

	return s
}

// A parent pipeline aggregation which sorts the buckets of its parent
// multi-bucket aggregation.
func (s *_aggregations) BucketSort(bucketsort types.BucketSortAggregationVariant) *_aggregations {

	s.v.BucketSort = bucketsort.BucketSortAggregationCaster()

	return s
}

// A single-value metrics aggregation that calculates an approximate count of
// distinct values.
func (s *_aggregations) Cardinality(cardinality types.CardinalityAggregationVariant) *_aggregations {

	s.v.Cardinality = cardinality.CardinalityAggregationCaster()

	return s
}

// A multi-bucket aggregation that groups semi-structured text into buckets.
func (s *_aggregations) CategorizeText(categorizetext types.CategorizeTextAggregationVariant) *_aggregations {

	s.v.CategorizeText = categorizetext.CategorizeTextAggregationCaster()

	return s
}

// A single bucket aggregation that selects child documents that have the
// specified type, as defined in a `join` field.
func (s *_aggregations) Children(children types.ChildrenAggregationVariant) *_aggregations {

	s.v.Children = children.ChildrenAggregationCaster()

	return s
}

// A multi-bucket aggregation that creates composite buckets from different
// sources.
// Unlike the other multi-bucket aggregations, you can use the `composite`
// aggregation to paginate *all* buckets from a multi-level aggregation
// efficiently.
func (s *_aggregations) Composite(composite types.CompositeAggregationVariant) *_aggregations {

	s.v.Composite = composite.CompositeAggregationCaster()

	return s
}

// A parent pipeline aggregation which calculates the cumulative cardinality in
// a parent `histogram` or `date_histogram` aggregation.
func (s *_aggregations) CumulativeCardinality(cumulativecardinality types.CumulativeCardinalityAggregationVariant) *_aggregations {

	s.v.CumulativeCardinality = cumulativecardinality.CumulativeCardinalityAggregationCaster()

	return s
}

// A parent pipeline aggregation which calculates the cumulative sum of a
// specified metric in a parent `histogram` or `date_histogram` aggregation.
func (s *_aggregations) CumulativeSum(cumulativesum types.CumulativeSumAggregationVariant) *_aggregations {

	s.v.CumulativeSum = cumulativesum.CumulativeSumAggregationCaster()

	return s
}

// A multi-bucket values source based aggregation that can be applied on date
// values or date range values extracted from the documents.
// It dynamically builds fixed size (interval) buckets over the values.
func (s *_aggregations) DateHistogram(datehistogram types.DateHistogramAggregationVariant) *_aggregations {

	s.v.DateHistogram = datehistogram.DateHistogramAggregationCaster()

	return s
}

// A multi-bucket value source based aggregation that enables the user to define
// a set of date ranges - each representing a bucket.
func (s *_aggregations) DateRange(daterange types.DateRangeAggregationVariant) *_aggregations {

	s.v.DateRange = daterange.DateRangeAggregationCaster()

	return s
}

// A parent pipeline aggregation which calculates the derivative of a specified
// metric in a parent `histogram` or `date_histogram` aggregation.
func (s *_aggregations) Derivative(derivative types.DerivativeAggregationVariant) *_aggregations {

	s.v.Derivative = derivative.DerivativeAggregationCaster()

	return s
}

// A filtering aggregation used to limit any sub aggregations' processing to a
// sample of the top-scoring documents.
// Similar to the `sampler` aggregation, but adds the ability to limit the
// number of matches that share a common value.
func (s *_aggregations) DiversifiedSampler(diversifiedsampler types.DiversifiedSamplerAggregationVariant) *_aggregations {

	s.v.DiversifiedSampler = diversifiedsampler.DiversifiedSamplerAggregationCaster()

	return s
}

// A multi-value metrics aggregation that computes stats over numeric values
// extracted from the aggregated documents.
func (s *_aggregations) ExtendedStats(extendedstats types.ExtendedStatsAggregationVariant) *_aggregations {

	s.v.ExtendedStats = extendedstats.ExtendedStatsAggregationCaster()

	return s
}

// A sibling pipeline aggregation which calculates a variety of stats across all
// bucket of a specified metric in a sibling aggregation.
func (s *_aggregations) ExtendedStatsBucket(extendedstatsbucket types.ExtendedStatsBucketAggregationVariant) *_aggregations {

	s.v.ExtendedStatsBucket = extendedstatsbucket.ExtendedStatsBucketAggregationCaster()

	return s
}

// A single bucket aggregation that narrows the set of documents to those that
// match a query.
func (s *_aggregations) Filter(filter types.QueryVariant) *_aggregations {

	s.v.Filter = filter.QueryCaster()

	return s
}

// A multi-bucket aggregation where each bucket contains the documents that
// match a query.
func (s *_aggregations) Filters(filters types.FiltersAggregationVariant) *_aggregations {

	s.v.Filters = filters.FiltersAggregationCaster()

	return s
}

// A bucket aggregation which finds frequent item sets, a form of association
// rules mining that identifies items that often occur together.
func (s *_aggregations) FrequentItemSets(frequentitemsets types.FrequentItemSetsAggregationVariant) *_aggregations {

	s.v.FrequentItemSets = frequentitemsets.FrequentItemSetsAggregationCaster()

	return s
}

// A metric aggregation that computes the geographic bounding box containing all
// values for a Geopoint or Geoshape field.
func (s *_aggregations) GeoBounds(geobounds types.GeoBoundsAggregationVariant) *_aggregations {

	s.v.GeoBounds = geobounds.GeoBoundsAggregationCaster()

	return s
}

// A metric aggregation that computes the weighted centroid from all coordinate
// values for geo fields.
func (s *_aggregations) GeoCentroid(geocentroid types.GeoCentroidAggregationVariant) *_aggregations {

	s.v.GeoCentroid = geocentroid.GeoCentroidAggregationCaster()

	return s
}

// A multi-bucket aggregation that works on `geo_point` fields.
// Evaluates the distance of each document value from an origin point and
// determines the buckets it belongs to, based on ranges defined in the request.
func (s *_aggregations) GeoDistance(geodistance types.GeoDistanceAggregationVariant) *_aggregations {

	s.v.GeoDistance = geodistance.GeoDistanceAggregationCaster()

	return s
}

// Aggregates all `geo_point` values within a bucket into a `LineString` ordered
// by the chosen sort field.
func (s *_aggregations) GeoLine(geoline types.GeoLineAggregationVariant) *_aggregations {

	s.v.GeoLine = geoline.GeoLineAggregationCaster()

	return s
}

// A multi-bucket aggregation that groups `geo_point` and `geo_shape` values
// into buckets that represent a grid.
// Each cell is labeled using a geohash which is of user-definable precision.
func (s *_aggregations) GeohashGrid(geohashgrid types.GeoHashGridAggregationVariant) *_aggregations {

	s.v.GeohashGrid = geohashgrid.GeoHashGridAggregationCaster()

	return s
}

// A multi-bucket aggregation that groups `geo_point` and `geo_shape` values
// into buckets that represent a grid.
// Each cell corresponds to a H3 cell index and is labeled using the H3Index
// representation.
func (s *_aggregations) GeohexGrid(geohexgrid types.GeohexGridAggregationVariant) *_aggregations {

	s.v.GeohexGrid = geohexgrid.GeohexGridAggregationCaster()

	return s
}

// A multi-bucket aggregation that groups `geo_point` and `geo_shape` values
// into buckets that represent a grid.
// Each cell corresponds to a map tile as used by many online map sites.
func (s *_aggregations) GeotileGrid(geotilegrid types.GeoTileGridAggregationVariant) *_aggregations {

	s.v.GeotileGrid = geotilegrid.GeoTileGridAggregationCaster()

	return s
}

// Defines a single bucket of all the documents within the search execution
// context.
// This context is defined by the indices and the document types you’re
// searching on, but is not influenced by the search query itself.
func (s *_aggregations) Global(global types.GlobalAggregationVariant) *_aggregations {

	s.v.Global = global.GlobalAggregationCaster()

	return s
}

// A multi-bucket values source based aggregation that can be applied on numeric
// values or numeric range values extracted from the documents.
// It dynamically builds fixed size (interval) buckets over the values.
func (s *_aggregations) Histogram(histogram types.HistogramAggregationVariant) *_aggregations {

	s.v.Histogram = histogram.HistogramAggregationCaster()

	return s
}

// A parent pipeline aggregation which loads a pre-trained model and performs
// inference on the collated result fields from the parent bucket aggregation.
func (s *_aggregations) Inference(inference types.InferenceAggregationVariant) *_aggregations {

	s.v.Inference = inference.InferenceAggregationCaster()

	return s
}

// A bucket aggregation that groups documents based on the network or
// sub-network of an IP address.
func (s *_aggregations) IpPrefix(ipprefix types.IpPrefixAggregationVariant) *_aggregations {

	s.v.IpPrefix = ipprefix.IpPrefixAggregationCaster()

	return s
}

// A multi-bucket value source based aggregation that enables the user to define
// a set of IP ranges - each representing a bucket.
func (s *_aggregations) IpRange(iprange types.IpRangeAggregationVariant) *_aggregations {

	s.v.IpRange = iprange.IpRangeAggregationCaster()

	return s
}

func (s *_aggregations) Line(line types.GeoLineAggregationVariant) *_aggregations {

	s.v.Line = line.GeoLineAggregationCaster()

	return s
}

// A numeric aggregation that computes the following statistics over a set of
// document fields: `count`, `mean`, `variance`, `skewness`, `kurtosis`,
// `covariance`, and `covariance`.
func (s *_aggregations) MatrixStats(matrixstats types.MatrixStatsAggregationVariant) *_aggregations {

	s.v.MatrixStats = matrixstats.MatrixStatsAggregationCaster()

	return s
}

// A single-value metrics aggregation that returns the maximum value among the
// numeric values extracted from the aggregated documents.
func (s *_aggregations) Max(max types.MaxAggregationVariant) *_aggregations {

	s.v.Max = max.MaxAggregationCaster()

	return s
}

// A sibling pipeline aggregation which identifies the bucket(s) with the
// maximum value of a specified metric in a sibling aggregation and outputs both
// the value and the key(s) of the bucket(s).
func (s *_aggregations) MaxBucket(maxbucket types.MaxBucketAggregationVariant) *_aggregations {

	s.v.MaxBucket = maxbucket.MaxBucketAggregationCaster()

	return s
}

// A single-value aggregation that approximates the median absolute deviation of
// its search results.
func (s *_aggregations) MedianAbsoluteDeviation(medianabsolutedeviation types.MedianAbsoluteDeviationAggregationVariant) *_aggregations {

	s.v.MedianAbsoluteDeviation = medianabsolutedeviation.MedianAbsoluteDeviationAggregationCaster()

	return s
}

func (s *_aggregations) Meta(metadata types.MetadataVariant) *_aggregations {

	s.v.Meta = *metadata.MetadataCaster()

	return s
}

// A single-value metrics aggregation that returns the minimum value among
// numeric values extracted from the aggregated documents.
func (s *_aggregations) Min(min types.MinAggregationVariant) *_aggregations {

	s.v.Min = min.MinAggregationCaster()

	return s
}

// A sibling pipeline aggregation which identifies the bucket(s) with the
// minimum value of a specified metric in a sibling aggregation and outputs both
// the value and the key(s) of the bucket(s).
func (s *_aggregations) MinBucket(minbucket types.MinBucketAggregationVariant) *_aggregations {

	s.v.MinBucket = minbucket.MinBucketAggregationCaster()

	return s
}

// A field data based single bucket aggregation, that creates a bucket of all
// documents in the current document set context that are missing a field value
// (effectively, missing a field or having the configured NULL value set).
func (s *_aggregations) Missing(missing types.MissingAggregationVariant) *_aggregations {

	s.v.Missing = missing.MissingAggregationCaster()

	return s
}

func (s *_aggregations) MovingAvg(movingaverageaggregation types.MovingAverageAggregationVariant) *_aggregations {

	s.v.MovingAvg = *movingaverageaggregation.MovingAverageAggregationCaster()

	return s
}

// Given an ordered series of data, "slides" a window across the data and runs a
// custom script on each window of data.
// For convenience, a number of common functions are predefined such as `min`,
// `max`, and moving averages.
func (s *_aggregations) MovingFn(movingfn types.MovingFunctionAggregationVariant) *_aggregations {

	s.v.MovingFn = movingfn.MovingFunctionAggregationCaster()

	return s
}

// Given an ordered series of percentiles, "slides" a window across those
// percentiles and computes cumulative percentiles.
func (s *_aggregations) MovingPercentiles(movingpercentiles types.MovingPercentilesAggregationVariant) *_aggregations {

	s.v.MovingPercentiles = movingpercentiles.MovingPercentilesAggregationCaster()

	return s
}

// A multi-bucket value source based aggregation where buckets are dynamically
// built - one per unique set of values.
func (s *_aggregations) MultiTerms(multiterms types.MultiTermsAggregationVariant) *_aggregations {

	s.v.MultiTerms = multiterms.MultiTermsAggregationCaster()

	return s
}

// A special single bucket aggregation that enables aggregating nested
// documents.
func (s *_aggregations) Nested(nested types.NestedAggregationVariant) *_aggregations {

	s.v.Nested = nested.NestedAggregationCaster()

	return s
}

// A parent pipeline aggregation which calculates the specific
// normalized/rescaled value for a specific bucket value.
func (s *_aggregations) Normalize(normalize types.NormalizeAggregationVariant) *_aggregations {

	s.v.Normalize = normalize.NormalizeAggregationCaster()

	return s
}

// A special single bucket aggregation that selects parent documents that have
// the specified type, as defined in a `join` field.
func (s *_aggregations) Parent(parent types.ParentAggregationVariant) *_aggregations {

	s.v.Parent = parent.ParentAggregationCaster()

	return s
}

// A multi-value metrics aggregation that calculates one or more percentile
// ranks over numeric values extracted from the aggregated documents.
func (s *_aggregations) PercentileRanks(percentileranks types.PercentileRanksAggregationVariant) *_aggregations {

	s.v.PercentileRanks = percentileranks.PercentileRanksAggregationCaster()

	return s
}

// A multi-value metrics aggregation that calculates one or more percentiles
// over numeric values extracted from the aggregated documents.
func (s *_aggregations) Percentiles(percentiles types.PercentilesAggregationVariant) *_aggregations {

	s.v.Percentiles = percentiles.PercentilesAggregationCaster()

	return s
}

// A sibling pipeline aggregation which calculates percentiles across all bucket
// of a specified metric in a sibling aggregation.
func (s *_aggregations) PercentilesBucket(percentilesbucket types.PercentilesBucketAggregationVariant) *_aggregations {

	s.v.PercentilesBucket = percentilesbucket.PercentilesBucketAggregationCaster()

	return s
}

// A single bucket aggregation that randomly includes documents in the
// aggregated results.
// Sampling provides significant speed improvement at the cost of accuracy.
func (s *_aggregations) RandomSampler(randomsampler types.RandomSamplerAggregationVariant) *_aggregations {

	s.v.RandomSampler = randomsampler.RandomSamplerAggregationCaster()

	return s
}

// A multi-bucket value source based aggregation that enables the user to define
// a set of ranges - each representing a bucket.
func (s *_aggregations) Range(range_ types.RangeAggregationVariant) *_aggregations {

	s.v.Range = range_.RangeAggregationCaster()

	return s
}

// A multi-bucket value source based aggregation which finds "rare" terms —
// terms that are at the long-tail of the distribution and are not frequent.
func (s *_aggregations) RareTerms(rareterms types.RareTermsAggregationVariant) *_aggregations {

	s.v.RareTerms = rareterms.RareTermsAggregationCaster()

	return s
}

// Calculates a rate of documents or a field in each bucket.
// Can only be used inside a `date_histogram` or `composite` aggregation.
func (s *_aggregations) Rate(rate types.RateAggregationVariant) *_aggregations {

	s.v.Rate = rate.RateAggregationCaster()

	return s
}

// A special single bucket aggregation that enables aggregating on parent
// documents from nested documents.
// Should only be defined inside a `nested` aggregation.
func (s *_aggregations) ReverseNested(reversenested types.ReverseNestedAggregationVariant) *_aggregations {

	s.v.ReverseNested = reversenested.ReverseNestedAggregationCaster()

	return s
}

// A filtering aggregation used to limit any sub aggregations' processing to a
// sample of the top-scoring documents.
func (s *_aggregations) Sampler(sampler types.SamplerAggregationVariant) *_aggregations {

	s.v.Sampler = sampler.SamplerAggregationCaster()

	return s
}

// A metric aggregation that uses scripts to provide a metric output.
func (s *_aggregations) ScriptedMetric(scriptedmetric types.ScriptedMetricAggregationVariant) *_aggregations {

	s.v.ScriptedMetric = scriptedmetric.ScriptedMetricAggregationCaster()

	return s
}

// An aggregation that subtracts values in a time series from themselves at
// different time lags or periods.
func (s *_aggregations) SerialDiff(serialdiff types.SerialDifferencingAggregationVariant) *_aggregations {

	s.v.SerialDiff = serialdiff.SerialDifferencingAggregationCaster()

	return s
}

// Returns interesting or unusual occurrences of terms in a set.
func (s *_aggregations) SignificantTerms(significantterms types.SignificantTermsAggregationVariant) *_aggregations {

	s.v.SignificantTerms = significantterms.SignificantTermsAggregationCaster()

	return s
}

// Returns interesting or unusual occurrences of free-text terms in a set.
func (s *_aggregations) SignificantText(significanttext types.SignificantTextAggregationVariant) *_aggregations {

	s.v.SignificantText = significanttext.SignificantTextAggregationCaster()

	return s
}

// A multi-value metrics aggregation that computes stats over numeric values
// extracted from the aggregated documents.
func (s *_aggregations) Stats(stats types.StatsAggregationVariant) *_aggregations {

	s.v.Stats = stats.StatsAggregationCaster()

	return s
}

// A sibling pipeline aggregation which calculates a variety of stats across all
// bucket of a specified metric in a sibling aggregation.
func (s *_aggregations) StatsBucket(statsbucket types.StatsBucketAggregationVariant) *_aggregations {

	s.v.StatsBucket = statsbucket.StatsBucketAggregationCaster()

	return s
}

// A multi-value metrics aggregation that computes statistics over string values
// extracted from the aggregated documents.
func (s *_aggregations) StringStats(stringstats types.StringStatsAggregationVariant) *_aggregations {

	s.v.StringStats = stringstats.StringStatsAggregationCaster()

	return s
}

// A single-value metrics aggregation that sums numeric values that are
// extracted from the aggregated documents.
func (s *_aggregations) Sum(sum types.SumAggregationVariant) *_aggregations {

	s.v.Sum = sum.SumAggregationCaster()

	return s
}

// A sibling pipeline aggregation which calculates the sum of a specified metric
// across all buckets in a sibling aggregation.
func (s *_aggregations) SumBucket(sumbucket types.SumBucketAggregationVariant) *_aggregations {

	s.v.SumBucket = sumbucket.SumBucketAggregationCaster()

	return s
}

// A metrics aggregation that performs a statistical hypothesis test in which
// the test statistic follows a Student’s t-distribution under the null
// hypothesis on numeric values extracted from the aggregated documents.
func (s *_aggregations) TTest(ttest types.TTestAggregationVariant) *_aggregations {

	s.v.TTest = ttest.TTestAggregationCaster()

	return s
}

// A multi-bucket value source based aggregation where buckets are dynamically
// built - one per unique value.
func (s *_aggregations) Terms(terms types.TermsAggregationVariant) *_aggregations {

	s.v.Terms = terms.TermsAggregationCaster()

	return s
}

// The time series aggregation queries data created using a time series index.
// This is typically data such as metrics or other data streams with a time
// component, and requires creating an index using the time series mode.
func (s *_aggregations) TimeSeries(timeseries types.TimeSeriesAggregationVariant) *_aggregations {

	s.v.TimeSeries = timeseries.TimeSeriesAggregationCaster()

	return s
}

// A metric aggregation that returns the top matching documents per bucket.
func (s *_aggregations) TopHits(tophits types.TopHitsAggregationVariant) *_aggregations {

	s.v.TopHits = tophits.TopHitsAggregationCaster()

	return s
}

// A metric aggregation that selects metrics from the document with the largest
// or smallest sort value.
func (s *_aggregations) TopMetrics(topmetrics types.TopMetricsAggregationVariant) *_aggregations {

	s.v.TopMetrics = topmetrics.TopMetricsAggregationCaster()

	return s
}

// A single-value metrics aggregation that counts the number of values that are
// extracted from the aggregated documents.
func (s *_aggregations) ValueCount(valuecount types.ValueCountAggregationVariant) *_aggregations {

	s.v.ValueCount = valuecount.ValueCountAggregationCaster()

	return s
}

// A multi-bucket aggregation similar to the histogram, except instead of
// providing an interval to use as the width of each bucket, a target number of
// buckets is provided.
func (s *_aggregations) VariableWidthHistogram(variablewidthhistogram types.VariableWidthHistogramAggregationVariant) *_aggregations {

	s.v.VariableWidthHistogram = variablewidthhistogram.VariableWidthHistogramAggregationCaster()

	return s
}

// A single-value metrics aggregation that computes the weighted average of
// numeric values that are extracted from the aggregated documents.
func (s *_aggregations) WeightedAvg(weightedavg types.WeightedAverageAggregationVariant) *_aggregations {

	s.v.WeightedAvg = weightedavg.WeightedAverageAggregationCaster()

	return s
}

func (s *_aggregations) AggregationsCaster() *types.Aggregations {
	return s.v
}
