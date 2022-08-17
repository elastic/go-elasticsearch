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

// Property holds the union for the following types:
//
//	AggregateMetricDoubleProperty
//	BinaryProperty
//	BooleanProperty
//	ByteNumberProperty
//	CompletionProperty
//	ConstantKeywordProperty
//	DateNanosProperty
//	DateProperty
//	DateRangeProperty
//	DenseVectorProperty
//	DoubleNumberProperty
//	DoubleRangeProperty
//	DynamicProperty
//	FieldAliasProperty
//	FlattenedProperty
//	FloatNumberProperty
//	FloatRangeProperty
//	GeoPointProperty
//	GeoShapeProperty
//	HalfFloatNumberProperty
//	HistogramProperty
//	IntegerNumberProperty
//	IntegerRangeProperty
//	IpProperty
//	IpRangeProperty
//	JoinProperty
//	KeywordProperty
//	LongNumberProperty
//	LongRangeProperty
//	MatchOnlyTextProperty
//	Murmur3HashProperty
//	NestedProperty
//	ObjectProperty
//	PercolatorProperty
//	PointProperty
//	RankFeatureProperty
//	RankFeaturesProperty
//	ScaledFloatNumberProperty
//	SearchAsYouTypeProperty
//	ShapeProperty
//	ShortNumberProperty
//	TextProperty
//	TokenCountProperty
//	UnsignedLongNumberProperty
//	VersionProperty
//	WildcardProperty
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/Property.ts#L91-L154
type Property interface{}

// PropertyBuilder holds Property struct and provides a builder API.
type PropertyBuilder struct {
	v Property
}

// NewProperty provides a builder for the Property struct.
func NewPropertyBuilder() *PropertyBuilder {
	return &PropertyBuilder{}
}

// Build finalize the chain and returns the Property struct
func (u *PropertyBuilder) Build() Property {
	return u.v
}

func (u *PropertyBuilder) AggregateMetricDoubleProperty(aggregatemetricdoubleproperty *AggregateMetricDoublePropertyBuilder) *PropertyBuilder {
	v := aggregatemetricdoubleproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) BinaryProperty(binaryproperty *BinaryPropertyBuilder) *PropertyBuilder {
	v := binaryproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) BooleanProperty(booleanproperty *BooleanPropertyBuilder) *PropertyBuilder {
	v := booleanproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) ByteNumberProperty(bytenumberproperty *ByteNumberPropertyBuilder) *PropertyBuilder {
	v := bytenumberproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) CompletionProperty(completionproperty *CompletionPropertyBuilder) *PropertyBuilder {
	v := completionproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) ConstantKeywordProperty(constantkeywordproperty *ConstantKeywordPropertyBuilder) *PropertyBuilder {
	v := constantkeywordproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) DateNanosProperty(datenanosproperty *DateNanosPropertyBuilder) *PropertyBuilder {
	v := datenanosproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) DateProperty(dateproperty *DatePropertyBuilder) *PropertyBuilder {
	v := dateproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) DateRangeProperty(daterangeproperty *DateRangePropertyBuilder) *PropertyBuilder {
	v := daterangeproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) DenseVectorProperty(densevectorproperty *DenseVectorPropertyBuilder) *PropertyBuilder {
	v := densevectorproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) DoubleNumberProperty(doublenumberproperty *DoubleNumberPropertyBuilder) *PropertyBuilder {
	v := doublenumberproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) DoubleRangeProperty(doublerangeproperty *DoubleRangePropertyBuilder) *PropertyBuilder {
	v := doublerangeproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) DynamicProperty(dynamicproperty *DynamicPropertyBuilder) *PropertyBuilder {
	v := dynamicproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) FieldAliasProperty(fieldaliasproperty *FieldAliasPropertyBuilder) *PropertyBuilder {
	v := fieldaliasproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) FlattenedProperty(flattenedproperty *FlattenedPropertyBuilder) *PropertyBuilder {
	v := flattenedproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) FloatNumberProperty(floatnumberproperty *FloatNumberPropertyBuilder) *PropertyBuilder {
	v := floatnumberproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) FloatRangeProperty(floatrangeproperty *FloatRangePropertyBuilder) *PropertyBuilder {
	v := floatrangeproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) GeoPointProperty(geopointproperty *GeoPointPropertyBuilder) *PropertyBuilder {
	v := geopointproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) GeoShapeProperty(geoshapeproperty *GeoShapePropertyBuilder) *PropertyBuilder {
	v := geoshapeproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) HalfFloatNumberProperty(halffloatnumberproperty *HalfFloatNumberPropertyBuilder) *PropertyBuilder {
	v := halffloatnumberproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) HistogramProperty(histogramproperty *HistogramPropertyBuilder) *PropertyBuilder {
	v := histogramproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) IntegerNumberProperty(integernumberproperty *IntegerNumberPropertyBuilder) *PropertyBuilder {
	v := integernumberproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) IntegerRangeProperty(integerrangeproperty *IntegerRangePropertyBuilder) *PropertyBuilder {
	v := integerrangeproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) IpProperty(ipproperty *IpPropertyBuilder) *PropertyBuilder {
	v := ipproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) IpRangeProperty(iprangeproperty *IpRangePropertyBuilder) *PropertyBuilder {
	v := iprangeproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) JoinProperty(joinproperty *JoinPropertyBuilder) *PropertyBuilder {
	v := joinproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) KeywordProperty(keywordproperty *KeywordPropertyBuilder) *PropertyBuilder {
	v := keywordproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) LongNumberProperty(longnumberproperty *LongNumberPropertyBuilder) *PropertyBuilder {
	v := longnumberproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) LongRangeProperty(longrangeproperty *LongRangePropertyBuilder) *PropertyBuilder {
	v := longrangeproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) MatchOnlyTextProperty(matchonlytextproperty *MatchOnlyTextPropertyBuilder) *PropertyBuilder {
	v := matchonlytextproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) Murmur3HashProperty(murmur3hashproperty *Murmur3HashPropertyBuilder) *PropertyBuilder {
	v := murmur3hashproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) NestedProperty(nestedproperty *NestedPropertyBuilder) *PropertyBuilder {
	v := nestedproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) ObjectProperty(objectproperty *ObjectPropertyBuilder) *PropertyBuilder {
	v := objectproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) PercolatorProperty(percolatorproperty *PercolatorPropertyBuilder) *PropertyBuilder {
	v := percolatorproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) PointProperty(pointproperty *PointPropertyBuilder) *PropertyBuilder {
	v := pointproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) RankFeatureProperty(rankfeatureproperty *RankFeaturePropertyBuilder) *PropertyBuilder {
	v := rankfeatureproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) RankFeaturesProperty(rankfeaturesproperty *RankFeaturesPropertyBuilder) *PropertyBuilder {
	v := rankfeaturesproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) ScaledFloatNumberProperty(scaledfloatnumberproperty *ScaledFloatNumberPropertyBuilder) *PropertyBuilder {
	v := scaledfloatnumberproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) SearchAsYouTypeProperty(searchasyoutypeproperty *SearchAsYouTypePropertyBuilder) *PropertyBuilder {
	v := searchasyoutypeproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) ShapeProperty(shapeproperty *ShapePropertyBuilder) *PropertyBuilder {
	v := shapeproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) ShortNumberProperty(shortnumberproperty *ShortNumberPropertyBuilder) *PropertyBuilder {
	v := shortnumberproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) TextProperty(textproperty *TextPropertyBuilder) *PropertyBuilder {
	v := textproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) TokenCountProperty(tokencountproperty *TokenCountPropertyBuilder) *PropertyBuilder {
	v := tokencountproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) UnsignedLongNumberProperty(unsignedlongnumberproperty *UnsignedLongNumberPropertyBuilder) *PropertyBuilder {
	v := unsignedlongnumberproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) VersionProperty(versionproperty *VersionPropertyBuilder) *PropertyBuilder {
	v := versionproperty.Build()
	u.v = &v
	return u
}

func (u *PropertyBuilder) WildcardProperty(wildcardproperty *WildcardPropertyBuilder) *PropertyBuilder {
	v := wildcardproperty.Build()
	u.v = &v
	return u
}
