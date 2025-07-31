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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

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
//	PassthroughObjectProperty
//	RankVectorProperty
//	SemanticTextProperty
//	SparseVectorProperty
//	CompletionProperty
//	ConstantKeywordProperty
//	CountedKeywordProperty
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
//	IcuCollationProperty
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/mapping/Property.ts#L120-L191
type Property any

type PropertyVariant interface {
	PropertyCaster() *Property
}
