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

// Aggregate holds the union for the following types:
//
//	AdjacencyMatrixAggregate
//	AutoDateHistogramAggregate
//	AvgAggregate
//	BoxPlotAggregate
//	BucketMetricValueAggregate
//	CardinalityAggregate
//	ChildrenAggregate
//	CompositeAggregate
//	CumulativeCardinalityAggregate
//	DateHistogramAggregate
//	DateRangeAggregate
//	DerivativeAggregate
//	DoubleTermsAggregate
//	ExtendedStatsAggregate
//	ExtendedStatsBucketAggregate
//	FilterAggregate
//	FiltersAggregate
//	GeoBoundsAggregate
//	GeoCentroidAggregate
//	GeoDistanceAggregate
//	GeoHashGridAggregate
//	GeoLineAggregate
//	GeoTileGridAggregate
//	GlobalAggregate
//	HdrPercentileRanksAggregate
//	HdrPercentilesAggregate
//	HistogramAggregate
//	InferenceAggregate
//	IpRangeAggregate
//	LongRareTermsAggregate
//	LongTermsAggregate
//	MatrixStatsAggregate
//	MaxAggregate
//	MedianAbsoluteDeviationAggregate
//	MinAggregate
//	MissingAggregate
//	MultiTermsAggregate
//	NestedAggregate
//	ParentAggregate
//	PercentilesBucketAggregate
//	RangeAggregate
//	RateAggregate
//	ReverseNestedAggregate
//	SamplerAggregate
//	ScriptedMetricAggregate
//	SignificantLongTermsAggregate
//	SignificantStringTermsAggregate
//	SimpleValueAggregate
//	StatsAggregate
//	StatsBucketAggregate
//	StringRareTermsAggregate
//	StringStatsAggregate
//	StringTermsAggregate
//	SumAggregate
//	TDigestPercentileRanksAggregate
//	TDigestPercentilesAggregate
//	TTestAggregate
//	TopHitsAggregate
//	TopMetricsAggregate
//	UnmappedRareTermsAggregate
//	UnmappedSamplerAggregate
//	UnmappedSignificantTermsAggregate
//	UnmappedTermsAggregate
//	ValueCountAggregate
//	VariableWidthHistogramAggregate
//	WeightedAvgAggregate
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L30-L112
type Aggregate interface{}

// AggregateBuilder holds Aggregate struct and provides a builder API.
type AggregateBuilder struct {
	v Aggregate
}

// NewAggregate provides a builder for the Aggregate struct.
func NewAggregateBuilder() *AggregateBuilder {
	return &AggregateBuilder{}
}

// Build finalize the chain and returns the Aggregate struct
func (u *AggregateBuilder) Build() Aggregate {
	return u.v
}

func (u *AggregateBuilder) AdjacencyMatrixAggregate(adjacencymatrixaggregate *AdjacencyMatrixAggregateBuilder) *AggregateBuilder {
	v := adjacencymatrixaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) AutoDateHistogramAggregate(autodatehistogramaggregate *AutoDateHistogramAggregateBuilder) *AggregateBuilder {
	v := autodatehistogramaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) AvgAggregate(avgaggregate *AvgAggregateBuilder) *AggregateBuilder {
	v := avgaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) BoxPlotAggregate(boxplotaggregate *BoxPlotAggregateBuilder) *AggregateBuilder {
	v := boxplotaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) BucketMetricValueAggregate(bucketmetricvalueaggregate *BucketMetricValueAggregateBuilder) *AggregateBuilder {
	v := bucketmetricvalueaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) CardinalityAggregate(cardinalityaggregate *CardinalityAggregateBuilder) *AggregateBuilder {
	v := cardinalityaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) ChildrenAggregate(childrenaggregate *ChildrenAggregateBuilder) *AggregateBuilder {
	v := childrenaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) CompositeAggregate(compositeaggregate *CompositeAggregateBuilder) *AggregateBuilder {
	v := compositeaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) CumulativeCardinalityAggregate(cumulativecardinalityaggregate *CumulativeCardinalityAggregateBuilder) *AggregateBuilder {
	v := cumulativecardinalityaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) DateHistogramAggregate(datehistogramaggregate *DateHistogramAggregateBuilder) *AggregateBuilder {
	v := datehistogramaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) DateRangeAggregate(daterangeaggregate *DateRangeAggregateBuilder) *AggregateBuilder {
	v := daterangeaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) DerivativeAggregate(derivativeaggregate *DerivativeAggregateBuilder) *AggregateBuilder {
	v := derivativeaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) DoubleTermsAggregate(doubletermsaggregate *DoubleTermsAggregateBuilder) *AggregateBuilder {
	v := doubletermsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) ExtendedStatsAggregate(extendedstatsaggregate *ExtendedStatsAggregateBuilder) *AggregateBuilder {
	v := extendedstatsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) ExtendedStatsBucketAggregate(extendedstatsbucketaggregate *ExtendedStatsBucketAggregateBuilder) *AggregateBuilder {
	v := extendedstatsbucketaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) FilterAggregate(filteraggregate *FilterAggregateBuilder) *AggregateBuilder {
	v := filteraggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) FiltersAggregate(filtersaggregate *FiltersAggregateBuilder) *AggregateBuilder {
	v := filtersaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) GeoBoundsAggregate(geoboundsaggregate *GeoBoundsAggregateBuilder) *AggregateBuilder {
	v := geoboundsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) GeoCentroidAggregate(geocentroidaggregate *GeoCentroidAggregateBuilder) *AggregateBuilder {
	v := geocentroidaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) GeoDistanceAggregate(geodistanceaggregate *GeoDistanceAggregateBuilder) *AggregateBuilder {
	v := geodistanceaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) GeoHashGridAggregate(geohashgridaggregate *GeoHashGridAggregateBuilder) *AggregateBuilder {
	v := geohashgridaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) GeoLineAggregate(geolineaggregate *GeoLineAggregateBuilder) *AggregateBuilder {
	v := geolineaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) GeoTileGridAggregate(geotilegridaggregate *GeoTileGridAggregateBuilder) *AggregateBuilder {
	v := geotilegridaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) GlobalAggregate(globalaggregate *GlobalAggregateBuilder) *AggregateBuilder {
	v := globalaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) HdrPercentileRanksAggregate(hdrpercentileranksaggregate *HdrPercentileRanksAggregateBuilder) *AggregateBuilder {
	v := hdrpercentileranksaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) HdrPercentilesAggregate(hdrpercentilesaggregate *HdrPercentilesAggregateBuilder) *AggregateBuilder {
	v := hdrpercentilesaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) HistogramAggregate(histogramaggregate *HistogramAggregateBuilder) *AggregateBuilder {
	v := histogramaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) InferenceAggregate(inferenceaggregate *InferenceAggregateBuilder) *AggregateBuilder {
	v := inferenceaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) IpRangeAggregate(iprangeaggregate *IpRangeAggregateBuilder) *AggregateBuilder {
	v := iprangeaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) LongRareTermsAggregate(longraretermsaggregate *LongRareTermsAggregateBuilder) *AggregateBuilder {
	v := longraretermsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) LongTermsAggregate(longtermsaggregate *LongTermsAggregateBuilder) *AggregateBuilder {
	v := longtermsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) MatrixStatsAggregate(matrixstatsaggregate *MatrixStatsAggregateBuilder) *AggregateBuilder {
	v := matrixstatsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) MaxAggregate(maxaggregate *MaxAggregateBuilder) *AggregateBuilder {
	v := maxaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) MedianAbsoluteDeviationAggregate(medianabsolutedeviationaggregate *MedianAbsoluteDeviationAggregateBuilder) *AggregateBuilder {
	v := medianabsolutedeviationaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) MinAggregate(minaggregate *MinAggregateBuilder) *AggregateBuilder {
	v := minaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) MissingAggregate(missingaggregate *MissingAggregateBuilder) *AggregateBuilder {
	v := missingaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) MultiTermsAggregate(multitermsaggregate *MultiTermsAggregateBuilder) *AggregateBuilder {
	v := multitermsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) NestedAggregate(nestedaggregate *NestedAggregateBuilder) *AggregateBuilder {
	v := nestedaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) ParentAggregate(parentaggregate *ParentAggregateBuilder) *AggregateBuilder {
	v := parentaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) PercentilesBucketAggregate(percentilesbucketaggregate *PercentilesBucketAggregateBuilder) *AggregateBuilder {
	v := percentilesbucketaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) RangeAggregate(rangeaggregate *RangeAggregateBuilder) *AggregateBuilder {
	v := rangeaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) RateAggregate(rateaggregate *RateAggregateBuilder) *AggregateBuilder {
	v := rateaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) ReverseNestedAggregate(reversenestedaggregate *ReverseNestedAggregateBuilder) *AggregateBuilder {
	v := reversenestedaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) SamplerAggregate(sampleraggregate *SamplerAggregateBuilder) *AggregateBuilder {
	v := sampleraggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) ScriptedMetricAggregate(scriptedmetricaggregate *ScriptedMetricAggregateBuilder) *AggregateBuilder {
	v := scriptedmetricaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) SignificantLongTermsAggregate(significantlongtermsaggregate *SignificantLongTermsAggregateBuilder) *AggregateBuilder {
	v := significantlongtermsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) SignificantStringTermsAggregate(significantstringtermsaggregate *SignificantStringTermsAggregateBuilder) *AggregateBuilder {
	v := significantstringtermsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) SimpleValueAggregate(simplevalueaggregate *SimpleValueAggregateBuilder) *AggregateBuilder {
	v := simplevalueaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) StatsAggregate(statsaggregate *StatsAggregateBuilder) *AggregateBuilder {
	v := statsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) StatsBucketAggregate(statsbucketaggregate *StatsBucketAggregateBuilder) *AggregateBuilder {
	v := statsbucketaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) StringRareTermsAggregate(stringraretermsaggregate *StringRareTermsAggregateBuilder) *AggregateBuilder {
	v := stringraretermsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) StringStatsAggregate(stringstatsaggregate *StringStatsAggregateBuilder) *AggregateBuilder {
	v := stringstatsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) StringTermsAggregate(stringtermsaggregate *StringTermsAggregateBuilder) *AggregateBuilder {
	v := stringtermsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) SumAggregate(sumaggregate *SumAggregateBuilder) *AggregateBuilder {
	v := sumaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) TDigestPercentileRanksAggregate(tdigestpercentileranksaggregate *TDigestPercentileRanksAggregateBuilder) *AggregateBuilder {
	v := tdigestpercentileranksaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) TDigestPercentilesAggregate(tdigestpercentilesaggregate *TDigestPercentilesAggregateBuilder) *AggregateBuilder {
	v := tdigestpercentilesaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) TTestAggregate(ttestaggregate *TTestAggregateBuilder) *AggregateBuilder {
	v := ttestaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) TopHitsAggregate(tophitsaggregate *TopHitsAggregateBuilder) *AggregateBuilder {
	v := tophitsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) TopMetricsAggregate(topmetricsaggregate *TopMetricsAggregateBuilder) *AggregateBuilder {
	v := topmetricsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) UnmappedRareTermsAggregate(unmappedraretermsaggregate *UnmappedRareTermsAggregateBuilder) *AggregateBuilder {
	v := unmappedraretermsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) UnmappedSamplerAggregate(unmappedsampleraggregate *UnmappedSamplerAggregateBuilder) *AggregateBuilder {
	v := unmappedsampleraggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) UnmappedSignificantTermsAggregate(unmappedsignificanttermsaggregate *UnmappedSignificantTermsAggregateBuilder) *AggregateBuilder {
	v := unmappedsignificanttermsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) UnmappedTermsAggregate(unmappedtermsaggregate *UnmappedTermsAggregateBuilder) *AggregateBuilder {
	v := unmappedtermsaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) ValueCountAggregate(valuecountaggregate *ValueCountAggregateBuilder) *AggregateBuilder {
	v := valuecountaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) VariableWidthHistogramAggregate(variablewidthhistogramaggregate *VariableWidthHistogramAggregateBuilder) *AggregateBuilder {
	v := variablewidthhistogramaggregate.Build()
	u.v = &v
	return u
}

func (u *AggregateBuilder) WeightedAvgAggregate(weightedavgaggregate *WeightedAvgAggregateBuilder) *AggregateBuilder {
	v := weightedavgaggregate.Build()
	u.v = &v
	return u
}
