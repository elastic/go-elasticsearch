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

// Aggregate holds the union for the following types:
//
//	CardinalityAggregate
//	HdrPercentilesAggregate
//	HdrPercentileRanksAggregate
//	TDigestPercentilesAggregate
//	TDigestPercentileRanksAggregate
//	PercentilesBucketAggregate
//	MedianAbsoluteDeviationAggregate
//	MinAggregate
//	MaxAggregate
//	SumAggregate
//	AvgAggregate
//	WeightedAvgAggregate
//	ValueCountAggregate
//	SimpleValueAggregate
//	DerivativeAggregate
//	BucketMetricValueAggregate
//	StatsAggregate
//	StatsBucketAggregate
//	ExtendedStatsAggregate
//	ExtendedStatsBucketAggregate
//	GeoBoundsAggregate
//	GeoCentroidAggregate
//	HistogramAggregate
//	DateHistogramAggregate
//	AutoDateHistogramAggregate
//	VariableWidthHistogramAggregate
//	StringTermsAggregate
//	LongTermsAggregate
//	DoubleTermsAggregate
//	UnmappedTermsAggregate
//	LongRareTermsAggregate
//	StringRareTermsAggregate
//	UnmappedRareTermsAggregate
//	MultiTermsAggregate
//	MissingAggregate
//	NestedAggregate
//	ReverseNestedAggregate
//	GlobalAggregate
//	FilterAggregate
//	ChildrenAggregate
//	ParentAggregate
//	SamplerAggregate
//	UnmappedSamplerAggregate
//	GeoHashGridAggregate
//	GeoTileGridAggregate
//	GeoHexGridAggregate
//	RangeAggregate
//	DateRangeAggregate
//	GeoDistanceAggregate
//	IpRangeAggregate
//	IpPrefixAggregate
//	FiltersAggregate
//	AdjacencyMatrixAggregate
//	SignificantLongTermsAggregate
//	SignificantStringTermsAggregate
//	UnmappedSignificantTermsAggregate
//	CompositeAggregate
//	ScriptedMetricAggregate
//	TopHitsAggregate
//	InferenceAggregate
//	StringStatsAggregate
//	BoxPlotAggregate
//	TopMetricsAggregate
//	TTestAggregate
//	RateAggregate
//	CumulativeCardinalityAggregate
//	MatrixStatsAggregate
//	GeoLineAggregate
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/aggregations/Aggregate.ts#L38-L122
type Aggregate interface {
	isAggregate()
}

func (i CardinalityAggregate) isAggregate() {}

func (i HdrPercentilesAggregate) isAggregate() {}

func (i HdrPercentileRanksAggregate) isAggregate() {}

func (i TDigestPercentilesAggregate) isAggregate() {}

func (i TDigestPercentileRanksAggregate) isAggregate() {}

func (i PercentilesBucketAggregate) isAggregate() {}

func (i MedianAbsoluteDeviationAggregate) isAggregate() {}

func (i MinAggregate) isAggregate() {}

func (i MaxAggregate) isAggregate() {}

func (i SumAggregate) isAggregate() {}

func (i AvgAggregate) isAggregate() {}

func (i WeightedAvgAggregate) isAggregate() {}

func (i ValueCountAggregate) isAggregate() {}

func (i SimpleValueAggregate) isAggregate() {}

func (i DerivativeAggregate) isAggregate() {}

func (i BucketMetricValueAggregate) isAggregate() {}

func (i StatsAggregate) isAggregate() {}

func (i StatsBucketAggregate) isAggregate() {}

func (i ExtendedStatsAggregate) isAggregate() {}

func (i ExtendedStatsBucketAggregate) isAggregate() {}

func (i GeoBoundsAggregate) isAggregate() {}

func (i GeoCentroidAggregate) isAggregate() {}

func (i HistogramAggregate) isAggregate() {}

func (i DateHistogramAggregate) isAggregate() {}

func (i AutoDateHistogramAggregate) isAggregate() {}

func (i VariableWidthHistogramAggregate) isAggregate() {}

func (i StringTermsAggregate) isAggregate() {}

func (i LongTermsAggregate) isAggregate() {}

func (i DoubleTermsAggregate) isAggregate() {}

func (i UnmappedTermsAggregate) isAggregate() {}

func (i LongRareTermsAggregate) isAggregate() {}

func (i StringRareTermsAggregate) isAggregate() {}

func (i UnmappedRareTermsAggregate) isAggregate() {}

func (i MultiTermsAggregate) isAggregate() {}

func (i MissingAggregate) isAggregate() {}

func (i NestedAggregate) isAggregate() {}

func (i ReverseNestedAggregate) isAggregate() {}

func (i GlobalAggregate) isAggregate() {}

func (i FilterAggregate) isAggregate() {}

func (i ChildrenAggregate) isAggregate() {}

func (i ParentAggregate) isAggregate() {}

func (i SamplerAggregate) isAggregate() {}

func (i UnmappedSamplerAggregate) isAggregate() {}

func (i GeoHashGridAggregate) isAggregate() {}

func (i GeoTileGridAggregate) isAggregate() {}

func (i GeoHexGridAggregate) isAggregate() {}

func (i RangeAggregate) isAggregate() {}

func (i DateRangeAggregate) isAggregate() {}

func (i GeoDistanceAggregate) isAggregate() {}

func (i IpRangeAggregate) isAggregate() {}

func (i IpPrefixAggregate) isAggregate() {}

func (i FiltersAggregate) isAggregate() {}

func (i AdjacencyMatrixAggregate) isAggregate() {}

func (i SignificantLongTermsAggregate) isAggregate() {}

func (i SignificantStringTermsAggregate) isAggregate() {}

func (i UnmappedSignificantTermsAggregate) isAggregate() {}

func (i CompositeAggregate) isAggregate() {}

func (i ScriptedMetricAggregate) isAggregate() {}

func (i TopHitsAggregate) isAggregate() {}

func (i InferenceAggregate) isAggregate() {}

func (i StringStatsAggregate) isAggregate() {}

func (i BoxPlotAggregate) isAggregate() {}

func (i TopMetricsAggregate) isAggregate() {}

func (i TTestAggregate) isAggregate() {}

func (i RateAggregate) isAggregate() {}

func (i CumulativeCardinalityAggregate) isAggregate() {}

func (i MatrixStatsAggregate) isAggregate() {}

func (i GeoLineAggregate) isAggregate() {}
