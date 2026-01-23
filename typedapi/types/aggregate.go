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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

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
//	ChangePointAggregate
//	StatsAggregate
//	StatsBucketAggregate
//	ExtendedStatsAggregate
//	ExtendedStatsBucketAggregate
//	CartesianBoundsAggregate
//	CartesianCentroidAggregate
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
//	FrequentItemSetsAggregate
//	TimeSeriesAggregate
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
// https://github.com/elastic/elasticsearch-specification/blob/b1811e10a0722431d79d1c234dd412ff47d8656f/specification/_types/aggregations/Aggregate.ts#L40-L130
type Aggregate any
