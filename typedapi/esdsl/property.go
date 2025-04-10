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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// This is provide all the types that are part of the union.
type _property struct {
	v types.Property
}

func NewProperty() *_property {
	return &_property{v: nil}
}

// UnknownProperty is used to set the unknown value of the union.
// Highlited as @non_exhaustive in the specification.
func (u *_property) UnknownProperty(unknown json.RawMessage) *_property {
	u.v = unknown
	return u
}

func (u *_property) BinaryProperty(binaryproperty types.BinaryPropertyVariant) *_property {

	u.v = &binaryproperty

	return u
}

// Interface implementation for BinaryProperty in Property union
func (u *_binaryProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) BooleanProperty(booleanproperty types.BooleanPropertyVariant) *_property {

	u.v = &booleanproperty

	return u
}

// Interface implementation for BooleanProperty in Property union
func (u *_booleanProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) DynamicProperty(dynamicproperty types.DynamicPropertyVariant) *_property {

	u.v = &dynamicproperty

	return u
}

// Interface implementation for DynamicProperty in Property union
func (u *_dynamicProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) JoinProperty(joinproperty types.JoinPropertyVariant) *_property {

	u.v = &joinproperty

	return u
}

// Interface implementation for JoinProperty in Property union
func (u *_joinProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) KeywordProperty(keywordproperty types.KeywordPropertyVariant) *_property {

	u.v = &keywordproperty

	return u
}

// Interface implementation for KeywordProperty in Property union
func (u *_keywordProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) MatchOnlyTextProperty(matchonlytextproperty types.MatchOnlyTextPropertyVariant) *_property {

	u.v = &matchonlytextproperty

	return u
}

// Interface implementation for MatchOnlyTextProperty in Property union
func (u *_matchOnlyTextProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) PercolatorProperty(percolatorproperty types.PercolatorPropertyVariant) *_property {

	u.v = &percolatorproperty

	return u
}

// Interface implementation for PercolatorProperty in Property union
func (u *_percolatorProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) RankFeatureProperty(rankfeatureproperty types.RankFeaturePropertyVariant) *_property {

	u.v = &rankfeatureproperty

	return u
}

// Interface implementation for RankFeatureProperty in Property union
func (u *_rankFeatureProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) RankFeaturesProperty(rankfeaturesproperty types.RankFeaturesPropertyVariant) *_property {

	u.v = &rankfeaturesproperty

	return u
}

// Interface implementation for RankFeaturesProperty in Property union
func (u *_rankFeaturesProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) SearchAsYouTypeProperty(searchasyoutypeproperty types.SearchAsYouTypePropertyVariant) *_property {

	u.v = &searchasyoutypeproperty

	return u
}

// Interface implementation for SearchAsYouTypeProperty in Property union
func (u *_searchAsYouTypeProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) TextProperty(textproperty types.TextPropertyVariant) *_property {

	u.v = &textproperty

	return u
}

// Interface implementation for TextProperty in Property union
func (u *_textProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) VersionProperty(versionproperty types.VersionPropertyVariant) *_property {

	u.v = &versionproperty

	return u
}

// Interface implementation for VersionProperty in Property union
func (u *_versionProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) WildcardProperty(wildcardproperty types.WildcardPropertyVariant) *_property {

	u.v = &wildcardproperty

	return u
}

// Interface implementation for WildcardProperty in Property union
func (u *_wildcardProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) DateNanosProperty(datenanosproperty types.DateNanosPropertyVariant) *_property {

	u.v = &datenanosproperty

	return u
}

// Interface implementation for DateNanosProperty in Property union
func (u *_dateNanosProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) DateProperty(dateproperty types.DatePropertyVariant) *_property {

	u.v = &dateproperty

	return u
}

// Interface implementation for DateProperty in Property union
func (u *_dateProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) AggregateMetricDoubleProperty(aggregatemetricdoubleproperty types.AggregateMetricDoublePropertyVariant) *_property {

	u.v = &aggregatemetricdoubleproperty

	return u
}

// Interface implementation for AggregateMetricDoubleProperty in Property union
func (u *_aggregateMetricDoubleProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) DenseVectorProperty(densevectorproperty types.DenseVectorPropertyVariant) *_property {

	u.v = &densevectorproperty

	return u
}

// Interface implementation for DenseVectorProperty in Property union
func (u *_denseVectorProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) FlattenedProperty(flattenedproperty types.FlattenedPropertyVariant) *_property {

	u.v = &flattenedproperty

	return u
}

// Interface implementation for FlattenedProperty in Property union
func (u *_flattenedProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) NestedProperty(nestedproperty types.NestedPropertyVariant) *_property {

	u.v = &nestedproperty

	return u
}

// Interface implementation for NestedProperty in Property union
func (u *_nestedProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) ObjectProperty(objectproperty types.ObjectPropertyVariant) *_property {

	u.v = &objectproperty

	return u
}

// Interface implementation for ObjectProperty in Property union
func (u *_objectProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) PassthroughObjectProperty(passthroughobjectproperty types.PassthroughObjectPropertyVariant) *_property {

	u.v = &passthroughobjectproperty

	return u
}

// Interface implementation for PassthroughObjectProperty in Property union
func (u *_passthroughObjectProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) SemanticTextProperty(semantictextproperty types.SemanticTextPropertyVariant) *_property {

	u.v = &semantictextproperty

	return u
}

// Interface implementation for SemanticTextProperty in Property union
func (u *_semanticTextProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) SparseVectorProperty(sparsevectorproperty types.SparseVectorPropertyVariant) *_property {

	u.v = &sparsevectorproperty

	return u
}

// Interface implementation for SparseVectorProperty in Property union
func (u *_sparseVectorProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) CompletionProperty(completionproperty types.CompletionPropertyVariant) *_property {

	u.v = &completionproperty

	return u
}

// Interface implementation for CompletionProperty in Property union
func (u *_completionProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) ConstantKeywordProperty(constantkeywordproperty types.ConstantKeywordPropertyVariant) *_property {

	u.v = &constantkeywordproperty

	return u
}

// Interface implementation for ConstantKeywordProperty in Property union
func (u *_constantKeywordProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) CountedKeywordProperty(countedkeywordproperty types.CountedKeywordPropertyVariant) *_property {

	u.v = &countedkeywordproperty

	return u
}

// Interface implementation for CountedKeywordProperty in Property union
func (u *_countedKeywordProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) FieldAliasProperty(fieldaliasproperty types.FieldAliasPropertyVariant) *_property {

	u.v = &fieldaliasproperty

	return u
}

// Interface implementation for FieldAliasProperty in Property union
func (u *_fieldAliasProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) HistogramProperty(histogramproperty types.HistogramPropertyVariant) *_property {

	u.v = &histogramproperty

	return u
}

// Interface implementation for HistogramProperty in Property union
func (u *_histogramProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) IpProperty(ipproperty types.IpPropertyVariant) *_property {

	u.v = &ipproperty

	return u
}

// Interface implementation for IpProperty in Property union
func (u *_ipProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) Murmur3HashProperty(murmur3hashproperty types.Murmur3HashPropertyVariant) *_property {

	u.v = &murmur3hashproperty

	return u
}

// Interface implementation for Murmur3HashProperty in Property union
func (u *_murmur3HashProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) TokenCountProperty(tokencountproperty types.TokenCountPropertyVariant) *_property {

	u.v = &tokencountproperty

	return u
}

// Interface implementation for TokenCountProperty in Property union
func (u *_tokenCountProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) GeoPointProperty(geopointproperty types.GeoPointPropertyVariant) *_property {

	u.v = &geopointproperty

	return u
}

// Interface implementation for GeoPointProperty in Property union
func (u *_geoPointProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) GeoShapeProperty(geoshapeproperty types.GeoShapePropertyVariant) *_property {

	u.v = &geoshapeproperty

	return u
}

// Interface implementation for GeoShapeProperty in Property union
func (u *_geoShapeProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) PointProperty(pointproperty types.PointPropertyVariant) *_property {

	u.v = &pointproperty

	return u
}

// Interface implementation for PointProperty in Property union
func (u *_pointProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) ShapeProperty(shapeproperty types.ShapePropertyVariant) *_property {

	u.v = &shapeproperty

	return u
}

// Interface implementation for ShapeProperty in Property union
func (u *_shapeProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) ByteNumberProperty(bytenumberproperty types.ByteNumberPropertyVariant) *_property {

	u.v = &bytenumberproperty

	return u
}

// Interface implementation for ByteNumberProperty in Property union
func (u *_byteNumberProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) DoubleNumberProperty(doublenumberproperty types.DoubleNumberPropertyVariant) *_property {

	u.v = &doublenumberproperty

	return u
}

// Interface implementation for DoubleNumberProperty in Property union
func (u *_doubleNumberProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) FloatNumberProperty(floatnumberproperty types.FloatNumberPropertyVariant) *_property {

	u.v = &floatnumberproperty

	return u
}

// Interface implementation for FloatNumberProperty in Property union
func (u *_floatNumberProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) HalfFloatNumberProperty(halffloatnumberproperty types.HalfFloatNumberPropertyVariant) *_property {

	u.v = &halffloatnumberproperty

	return u
}

// Interface implementation for HalfFloatNumberProperty in Property union
func (u *_halfFloatNumberProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) IntegerNumberProperty(integernumberproperty types.IntegerNumberPropertyVariant) *_property {

	u.v = &integernumberproperty

	return u
}

// Interface implementation for IntegerNumberProperty in Property union
func (u *_integerNumberProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) LongNumberProperty(longnumberproperty types.LongNumberPropertyVariant) *_property {

	u.v = &longnumberproperty

	return u
}

// Interface implementation for LongNumberProperty in Property union
func (u *_longNumberProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) ScaledFloatNumberProperty(scaledfloatnumberproperty types.ScaledFloatNumberPropertyVariant) *_property {

	u.v = &scaledfloatnumberproperty

	return u
}

// Interface implementation for ScaledFloatNumberProperty in Property union
func (u *_scaledFloatNumberProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) ShortNumberProperty(shortnumberproperty types.ShortNumberPropertyVariant) *_property {

	u.v = &shortnumberproperty

	return u
}

// Interface implementation for ShortNumberProperty in Property union
func (u *_shortNumberProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) UnsignedLongNumberProperty(unsignedlongnumberproperty types.UnsignedLongNumberPropertyVariant) *_property {

	u.v = &unsignedlongnumberproperty

	return u
}

// Interface implementation for UnsignedLongNumberProperty in Property union
func (u *_unsignedLongNumberProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) DateRangeProperty(daterangeproperty types.DateRangePropertyVariant) *_property {

	u.v = &daterangeproperty

	return u
}

// Interface implementation for DateRangeProperty in Property union
func (u *_dateRangeProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) DoubleRangeProperty(doublerangeproperty types.DoubleRangePropertyVariant) *_property {

	u.v = &doublerangeproperty

	return u
}

// Interface implementation for DoubleRangeProperty in Property union
func (u *_doubleRangeProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) FloatRangeProperty(floatrangeproperty types.FloatRangePropertyVariant) *_property {

	u.v = &floatrangeproperty

	return u
}

// Interface implementation for FloatRangeProperty in Property union
func (u *_floatRangeProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) IntegerRangeProperty(integerrangeproperty types.IntegerRangePropertyVariant) *_property {

	u.v = &integerrangeproperty

	return u
}

// Interface implementation for IntegerRangeProperty in Property union
func (u *_integerRangeProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) IpRangeProperty(iprangeproperty types.IpRangePropertyVariant) *_property {

	u.v = &iprangeproperty

	return u
}

// Interface implementation for IpRangeProperty in Property union
func (u *_ipRangeProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) LongRangeProperty(longrangeproperty types.LongRangePropertyVariant) *_property {

	u.v = &longrangeproperty

	return u
}

// Interface implementation for LongRangeProperty in Property union
func (u *_longRangeProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) IcuCollationProperty(icucollationproperty types.IcuCollationPropertyVariant) *_property {

	u.v = &icucollationproperty

	return u
}

// Interface implementation for IcuCollationProperty in Property union
func (u *_icuCollationProperty) PropertyCaster() *types.Property {
	t := types.Property(u.v)
	return &t
}

func (u *_property) PropertyCaster() *types.Property {
	return &u.v
}
