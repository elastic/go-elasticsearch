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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

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
//	FrequentItemSetsAggregate
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
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/aggregations/Aggregate.ts#L38-L123
type Aggregate interface{}
