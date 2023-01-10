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

// Property holds the union for the following types:
//
//	BinaryProperty
//	BooleanProperty
//	DynamicProperty
//	JoinProperty
//	KeywordProperty
//	MatchOnlyTextProperty
//	PercolatorProperty
//	RankFeatureProperty
//	RankFeaturesProperty
//	SearchAsYouTypeProperty
//	TextProperty
//	VersionProperty
//	WildcardProperty
//	DateNanosProperty
//	DateProperty
//	AggregateMetricDoubleProperty
//	DenseVectorProperty
//	FlattenedProperty
//	NestedProperty
//	ObjectProperty
//	CompletionProperty
//	ConstantKeywordProperty
//	FieldAliasProperty
//	HistogramProperty
//	IpProperty
//	Murmur3HashProperty
//	TokenCountProperty
//	GeoPointProperty
//	GeoShapeProperty
//	PointProperty
//	ShapeProperty
//	ByteNumberProperty
//	DoubleNumberProperty
//	FloatNumberProperty
//	HalfFloatNumberProperty
//	IntegerNumberProperty
//	LongNumberProperty
//	ScaledFloatNumberProperty
//	ShortNumberProperty
//	UnsignedLongNumberProperty
//	DateRangeProperty
//	DoubleRangeProperty
//	FloatRangeProperty
//	IntegerRangeProperty
//	IpRangeProperty
//	LongRangeProperty
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/mapping/Property.ts#L93-L156
type Property interface {
	isProperty()
}

func (i BinaryProperty) isProperty() {}

func (i BooleanProperty) isProperty() {}

func (i DynamicProperty) isProperty() {}

func (i JoinProperty) isProperty() {}

func (i KeywordProperty) isProperty() {}

func (i MatchOnlyTextProperty) isProperty() {}

func (i PercolatorProperty) isProperty() {}

func (i RankFeatureProperty) isProperty() {}

func (i RankFeaturesProperty) isProperty() {}

func (i SearchAsYouTypeProperty) isProperty() {}

func (i TextProperty) isProperty() {}

func (i VersionProperty) isProperty() {}

func (i WildcardProperty) isProperty() {}

func (i DateNanosProperty) isProperty() {}

func (i DateProperty) isProperty() {}

func (i AggregateMetricDoubleProperty) isProperty() {}

func (i DenseVectorProperty) isProperty() {}

func (i FlattenedProperty) isProperty() {}

func (i NestedProperty) isProperty() {}

func (i ObjectProperty) isProperty() {}

func (i CompletionProperty) isProperty() {}

func (i ConstantKeywordProperty) isProperty() {}

func (i FieldAliasProperty) isProperty() {}

func (i HistogramProperty) isProperty() {}

func (i IpProperty) isProperty() {}

func (i Murmur3HashProperty) isProperty() {}

func (i TokenCountProperty) isProperty() {}

func (i GeoPointProperty) isProperty() {}

func (i GeoShapeProperty) isProperty() {}

func (i PointProperty) isProperty() {}

func (i ShapeProperty) isProperty() {}

func (i ByteNumberProperty) isProperty() {}

func (i DoubleNumberProperty) isProperty() {}

func (i FloatNumberProperty) isProperty() {}

func (i HalfFloatNumberProperty) isProperty() {}

func (i IntegerNumberProperty) isProperty() {}

func (i LongNumberProperty) isProperty() {}

func (i ScaledFloatNumberProperty) isProperty() {}

func (i ShortNumberProperty) isProperty() {}

func (i UnsignedLongNumberProperty) isProperty() {}

func (i DateRangeProperty) isProperty() {}

func (i DoubleRangeProperty) isProperty() {}

func (i FloatRangeProperty) isProperty() {}

func (i IntegerRangeProperty) isProperty() {}

func (i IpRangeProperty) isProperty() {}

func (i LongRangeProperty) isProperty() {}
