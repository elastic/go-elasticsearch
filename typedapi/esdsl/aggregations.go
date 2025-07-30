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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
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

func (s *_aggregations) AdjacencyMatrix(adjacencymatrix types.AdjacencyMatrixAggregationVariant) *_aggregations {

	s.v.AdjacencyMatrix = adjacencymatrix.AdjacencyMatrixAggregationCaster()

	return s
}

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

func (s *_aggregations) AutoDateHistogram(autodatehistogram types.AutoDateHistogramAggregationVariant) *_aggregations {

	s.v.AutoDateHistogram = autodatehistogram.AutoDateHistogramAggregationCaster()

	return s
}

func (s *_aggregations) Avg(avg types.AverageAggregationVariant) *_aggregations {

	s.v.Avg = avg.AverageAggregationCaster()

	return s
}

func (s *_aggregations) AvgBucket(avgbucket types.AverageBucketAggregationVariant) *_aggregations {

	s.v.AvgBucket = avgbucket.AverageBucketAggregationCaster()

	return s
}

func (s *_aggregations) Boxplot(boxplot types.BoxplotAggregationVariant) *_aggregations {

	s.v.Boxplot = boxplot.BoxplotAggregationCaster()

	return s
}

func (s *_aggregations) BucketCorrelation(bucketcorrelation types.BucketCorrelationAggregationVariant) *_aggregations {

	s.v.BucketCorrelation = bucketcorrelation.BucketCorrelationAggregationCaster()

	return s
}

func (s *_aggregations) BucketCountKsTest(bucketcountkstest types.BucketKsAggregationVariant) *_aggregations {

	s.v.BucketCountKsTest = bucketcountkstest.BucketKsAggregationCaster()

	return s
}

func (s *_aggregations) BucketScript(bucketscript types.BucketScriptAggregationVariant) *_aggregations {

	s.v.BucketScript = bucketscript.BucketScriptAggregationCaster()

	return s
}

func (s *_aggregations) BucketSelector(bucketselector types.BucketSelectorAggregationVariant) *_aggregations {

	s.v.BucketSelector = bucketselector.BucketSelectorAggregationCaster()

	return s
}

func (s *_aggregations) BucketSort(bucketsort types.BucketSortAggregationVariant) *_aggregations {

	s.v.BucketSort = bucketsort.BucketSortAggregationCaster()

	return s
}

func (s *_aggregations) Cardinality(cardinality types.CardinalityAggregationVariant) *_aggregations {

	s.v.Cardinality = cardinality.CardinalityAggregationCaster()

	return s
}

func (s *_aggregations) CategorizeText(categorizetext types.CategorizeTextAggregationVariant) *_aggregations {

	s.v.CategorizeText = categorizetext.CategorizeTextAggregationCaster()

	return s
}

func (s *_aggregations) Children(children types.ChildrenAggregationVariant) *_aggregations {

	s.v.Children = children.ChildrenAggregationCaster()

	return s
}

func (s *_aggregations) Composite(composite types.CompositeAggregationVariant) *_aggregations {

	s.v.Composite = composite.CompositeAggregationCaster()

	return s
}

func (s *_aggregations) CumulativeCardinality(cumulativecardinality types.CumulativeCardinalityAggregationVariant) *_aggregations {

	s.v.CumulativeCardinality = cumulativecardinality.CumulativeCardinalityAggregationCaster()

	return s
}

func (s *_aggregations) CumulativeSum(cumulativesum types.CumulativeSumAggregationVariant) *_aggregations {

	s.v.CumulativeSum = cumulativesum.CumulativeSumAggregationCaster()

	return s
}

func (s *_aggregations) DateHistogram(datehistogram types.DateHistogramAggregationVariant) *_aggregations {

	s.v.DateHistogram = datehistogram.DateHistogramAggregationCaster()

	return s
}

func (s *_aggregations) DateRange(daterange types.DateRangeAggregationVariant) *_aggregations {

	s.v.DateRange = daterange.DateRangeAggregationCaster()

	return s
}

func (s *_aggregations) Derivative(derivative types.DerivativeAggregationVariant) *_aggregations {

	s.v.Derivative = derivative.DerivativeAggregationCaster()

	return s
}

func (s *_aggregations) DiversifiedSampler(diversifiedsampler types.DiversifiedSamplerAggregationVariant) *_aggregations {

	s.v.DiversifiedSampler = diversifiedsampler.DiversifiedSamplerAggregationCaster()

	return s
}

func (s *_aggregations) ExtendedStats(extendedstats types.ExtendedStatsAggregationVariant) *_aggregations {

	s.v.ExtendedStats = extendedstats.ExtendedStatsAggregationCaster()

	return s
}

func (s *_aggregations) ExtendedStatsBucket(extendedstatsbucket types.ExtendedStatsBucketAggregationVariant) *_aggregations {

	s.v.ExtendedStatsBucket = extendedstatsbucket.ExtendedStatsBucketAggregationCaster()

	return s
}

func (s *_aggregations) Filter(filter types.QueryVariant) *_aggregations {

	s.v.Filter = filter.QueryCaster()

	return s
}

func (s *_aggregations) Filters(filters types.FiltersAggregationVariant) *_aggregations {

	s.v.Filters = filters.FiltersAggregationCaster()

	return s
}

func (s *_aggregations) FrequentItemSets(frequentitemsets types.FrequentItemSetsAggregationVariant) *_aggregations {

	s.v.FrequentItemSets = frequentitemsets.FrequentItemSetsAggregationCaster()

	return s
}

func (s *_aggregations) GeoBounds(geobounds types.GeoBoundsAggregationVariant) *_aggregations {

	s.v.GeoBounds = geobounds.GeoBoundsAggregationCaster()

	return s
}

func (s *_aggregations) GeoCentroid(geocentroid types.GeoCentroidAggregationVariant) *_aggregations {

	s.v.GeoCentroid = geocentroid.GeoCentroidAggregationCaster()

	return s
}

func (s *_aggregations) GeoDistance(geodistance types.GeoDistanceAggregationVariant) *_aggregations {

	s.v.GeoDistance = geodistance.GeoDistanceAggregationCaster()

	return s
}

func (s *_aggregations) GeoLine(geoline types.GeoLineAggregationVariant) *_aggregations {

	s.v.GeoLine = geoline.GeoLineAggregationCaster()

	return s
}

func (s *_aggregations) GeohashGrid(geohashgrid types.GeoHashGridAggregationVariant) *_aggregations {

	s.v.GeohashGrid = geohashgrid.GeoHashGridAggregationCaster()

	return s
}

func (s *_aggregations) GeohexGrid(geohexgrid types.GeohexGridAggregationVariant) *_aggregations {

	s.v.GeohexGrid = geohexgrid.GeohexGridAggregationCaster()

	return s
}

func (s *_aggregations) GeotileGrid(geotilegrid types.GeoTileGridAggregationVariant) *_aggregations {

	s.v.GeotileGrid = geotilegrid.GeoTileGridAggregationCaster()

	return s
}

func (s *_aggregations) Global(global types.GlobalAggregationVariant) *_aggregations {

	s.v.Global = global.GlobalAggregationCaster()

	return s
}

func (s *_aggregations) Histogram(histogram types.HistogramAggregationVariant) *_aggregations {

	s.v.Histogram = histogram.HistogramAggregationCaster()

	return s
}

func (s *_aggregations) Inference(inference types.InferenceAggregationVariant) *_aggregations {

	s.v.Inference = inference.InferenceAggregationCaster()

	return s
}

func (s *_aggregations) IpPrefix(ipprefix types.IpPrefixAggregationVariant) *_aggregations {

	s.v.IpPrefix = ipprefix.IpPrefixAggregationCaster()

	return s
}

func (s *_aggregations) IpRange(iprange types.IpRangeAggregationVariant) *_aggregations {

	s.v.IpRange = iprange.IpRangeAggregationCaster()

	return s
}

func (s *_aggregations) Line(line types.GeoLineAggregationVariant) *_aggregations {

	s.v.Line = line.GeoLineAggregationCaster()

	return s
}

func (s *_aggregations) MatrixStats(matrixstats types.MatrixStatsAggregationVariant) *_aggregations {

	s.v.MatrixStats = matrixstats.MatrixStatsAggregationCaster()

	return s
}

func (s *_aggregations) Max(max types.MaxAggregationVariant) *_aggregations {

	s.v.Max = max.MaxAggregationCaster()

	return s
}

func (s *_aggregations) MaxBucket(maxbucket types.MaxBucketAggregationVariant) *_aggregations {

	s.v.MaxBucket = maxbucket.MaxBucketAggregationCaster()

	return s
}

func (s *_aggregations) MedianAbsoluteDeviation(medianabsolutedeviation types.MedianAbsoluteDeviationAggregationVariant) *_aggregations {

	s.v.MedianAbsoluteDeviation = medianabsolutedeviation.MedianAbsoluteDeviationAggregationCaster()

	return s
}

func (s *_aggregations) Meta(metadata types.MetadataVariant) *_aggregations {

	s.v.Meta = *metadata.MetadataCaster()

	return s
}

func (s *_aggregations) Min(min types.MinAggregationVariant) *_aggregations {

	s.v.Min = min.MinAggregationCaster()

	return s
}

func (s *_aggregations) MinBucket(minbucket types.MinBucketAggregationVariant) *_aggregations {

	s.v.MinBucket = minbucket.MinBucketAggregationCaster()

	return s
}

func (s *_aggregations) Missing(missing types.MissingAggregationVariant) *_aggregations {

	s.v.Missing = missing.MissingAggregationCaster()

	return s
}

func (s *_aggregations) MovingAvg(movingaverageaggregation types.MovingAverageAggregationVariant) *_aggregations {

	s.v.MovingAvg = *movingaverageaggregation.MovingAverageAggregationCaster()

	return s
}

func (s *_aggregations) MovingFn(movingfn types.MovingFunctionAggregationVariant) *_aggregations {

	s.v.MovingFn = movingfn.MovingFunctionAggregationCaster()

	return s
}

func (s *_aggregations) MovingPercentiles(movingpercentiles types.MovingPercentilesAggregationVariant) *_aggregations {

	s.v.MovingPercentiles = movingpercentiles.MovingPercentilesAggregationCaster()

	return s
}

func (s *_aggregations) MultiTerms(multiterms types.MultiTermsAggregationVariant) *_aggregations {

	s.v.MultiTerms = multiterms.MultiTermsAggregationCaster()

	return s
}

func (s *_aggregations) Nested(nested types.NestedAggregationVariant) *_aggregations {

	s.v.Nested = nested.NestedAggregationCaster()

	return s
}

func (s *_aggregations) Normalize(normalize types.NormalizeAggregationVariant) *_aggregations {

	s.v.Normalize = normalize.NormalizeAggregationCaster()

	return s
}

func (s *_aggregations) Parent(parent types.ParentAggregationVariant) *_aggregations {

	s.v.Parent = parent.ParentAggregationCaster()

	return s
}

func (s *_aggregations) PercentileRanks(percentileranks types.PercentileRanksAggregationVariant) *_aggregations {

	s.v.PercentileRanks = percentileranks.PercentileRanksAggregationCaster()

	return s
}

func (s *_aggregations) Percentiles(percentiles types.PercentilesAggregationVariant) *_aggregations {

	s.v.Percentiles = percentiles.PercentilesAggregationCaster()

	return s
}

func (s *_aggregations) PercentilesBucket(percentilesbucket types.PercentilesBucketAggregationVariant) *_aggregations {

	s.v.PercentilesBucket = percentilesbucket.PercentilesBucketAggregationCaster()

	return s
}

func (s *_aggregations) RandomSampler(randomsampler types.RandomSamplerAggregationVariant) *_aggregations {

	s.v.RandomSampler = randomsampler.RandomSamplerAggregationCaster()

	return s
}

func (s *_aggregations) Range(range_ types.RangeAggregationVariant) *_aggregations {

	s.v.Range = range_.RangeAggregationCaster()

	return s
}

func (s *_aggregations) RareTerms(rareterms types.RareTermsAggregationVariant) *_aggregations {

	s.v.RareTerms = rareterms.RareTermsAggregationCaster()

	return s
}

func (s *_aggregations) Rate(rate types.RateAggregationVariant) *_aggregations {

	s.v.Rate = rate.RateAggregationCaster()

	return s
}

func (s *_aggregations) ReverseNested(reversenested types.ReverseNestedAggregationVariant) *_aggregations {

	s.v.ReverseNested = reversenested.ReverseNestedAggregationCaster()

	return s
}

func (s *_aggregations) Sampler(sampler types.SamplerAggregationVariant) *_aggregations {

	s.v.Sampler = sampler.SamplerAggregationCaster()

	return s
}

func (s *_aggregations) ScriptedMetric(scriptedmetric types.ScriptedMetricAggregationVariant) *_aggregations {

	s.v.ScriptedMetric = scriptedmetric.ScriptedMetricAggregationCaster()

	return s
}

func (s *_aggregations) SerialDiff(serialdiff types.SerialDifferencingAggregationVariant) *_aggregations {

	s.v.SerialDiff = serialdiff.SerialDifferencingAggregationCaster()

	return s
}

func (s *_aggregations) SignificantTerms(significantterms types.SignificantTermsAggregationVariant) *_aggregations {

	s.v.SignificantTerms = significantterms.SignificantTermsAggregationCaster()

	return s
}

func (s *_aggregations) SignificantText(significanttext types.SignificantTextAggregationVariant) *_aggregations {

	s.v.SignificantText = significanttext.SignificantTextAggregationCaster()

	return s
}

func (s *_aggregations) Stats(stats types.StatsAggregationVariant) *_aggregations {

	s.v.Stats = stats.StatsAggregationCaster()

	return s
}

func (s *_aggregations) StatsBucket(statsbucket types.StatsBucketAggregationVariant) *_aggregations {

	s.v.StatsBucket = statsbucket.StatsBucketAggregationCaster()

	return s
}

func (s *_aggregations) StringStats(stringstats types.StringStatsAggregationVariant) *_aggregations {

	s.v.StringStats = stringstats.StringStatsAggregationCaster()

	return s
}

func (s *_aggregations) Sum(sum types.SumAggregationVariant) *_aggregations {

	s.v.Sum = sum.SumAggregationCaster()

	return s
}

func (s *_aggregations) SumBucket(sumbucket types.SumBucketAggregationVariant) *_aggregations {

	s.v.SumBucket = sumbucket.SumBucketAggregationCaster()

	return s
}

func (s *_aggregations) TTest(ttest types.TTestAggregationVariant) *_aggregations {

	s.v.TTest = ttest.TTestAggregationCaster()

	return s
}

func (s *_aggregations) Terms(terms types.TermsAggregationVariant) *_aggregations {

	s.v.Terms = terms.TermsAggregationCaster()

	return s
}

func (s *_aggregations) TimeSeries(timeseries types.TimeSeriesAggregationVariant) *_aggregations {

	s.v.TimeSeries = timeseries.TimeSeriesAggregationCaster()

	return s
}

func (s *_aggregations) TopHits(tophits types.TopHitsAggregationVariant) *_aggregations {

	s.v.TopHits = tophits.TopHitsAggregationCaster()

	return s
}

func (s *_aggregations) TopMetrics(topmetrics types.TopMetricsAggregationVariant) *_aggregations {

	s.v.TopMetrics = topmetrics.TopMetricsAggregationCaster()

	return s
}

func (s *_aggregations) ValueCount(valuecount types.ValueCountAggregationVariant) *_aggregations {

	s.v.ValueCount = valuecount.ValueCountAggregationCaster()

	return s
}

func (s *_aggregations) VariableWidthHistogram(variablewidthhistogram types.VariableWidthHistogramAggregationVariant) *_aggregations {

	s.v.VariableWidthHistogram = variablewidthhistogram.VariableWidthHistogramAggregationCaster()

	return s
}

func (s *_aggregations) WeightedAvg(weightedavg types.WeightedAverageAggregationVariant) *_aggregations {

	s.v.WeightedAvg = weightedavg.WeightedAverageAggregationCaster()

	return s
}

func (s *_aggregations) AggregationsCaster() *types.Aggregations {
	return s.v
}
