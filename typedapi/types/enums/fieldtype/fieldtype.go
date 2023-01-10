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


// Package fieldtype
package fieldtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/mapping/Property.ts#L158-L201
type FieldType struct {
	Name string
}

var (
	None = FieldType{"none"}

	Geopoint = FieldType{"geo_point"}

	Geoshape = FieldType{"geo_shape"}

	Ip = FieldType{"ip"}

	Binary = FieldType{"binary"}

	Keyword = FieldType{"keyword"}

	Text = FieldType{"text"}

	Searchasyoutype = FieldType{"search_as_you_type"}

	Date = FieldType{"date"}

	Datenanos = FieldType{"date_nanos"}

	Boolean = FieldType{"boolean"}

	Completion = FieldType{"completion"}

	Nested = FieldType{"nested"}

	Object = FieldType{"object"}

	Murmur3 = FieldType{"murmur3"}

	Tokencount = FieldType{"token_count"}

	Percolator = FieldType{"percolator"}

	Integer = FieldType{"integer"}

	Long = FieldType{"long"}

	Short = FieldType{"short"}

	Byte = FieldType{"byte"}

	Float = FieldType{"float"}

	Halffloat = FieldType{"half_float"}

	Scaledfloat = FieldType{"scaled_float"}

	Double = FieldType{"double"}

	Integerrange = FieldType{"integer_range"}

	Floatrange = FieldType{"float_range"}

	Longrange = FieldType{"long_range"}

	Doublerange = FieldType{"double_range"}

	Daterange = FieldType{"date_range"}

	Iprange = FieldType{"ip_range"}

	Alias = FieldType{"alias"}

	Join = FieldType{"join"}

	Rankfeature = FieldType{"rank_feature"}

	Rankfeatures = FieldType{"rank_features"}

	Flattened = FieldType{"flattened"}

	Shape = FieldType{"shape"}

	Histogram = FieldType{"histogram"}

	Constantkeyword = FieldType{"constant_keyword"}

	Aggregatemetricdouble = FieldType{"aggregate_metric_double"}

	Densevector = FieldType{"dense_vector"}

	Matchonlytext = FieldType{"match_only_text"}
)

func (f FieldType) MarshalText() (text []byte, err error) {
	return []byte(f.String()), nil
}

func (f *FieldType) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "none":
		*f = None
	case "geo_point":
		*f = Geopoint
	case "geo_shape":
		*f = Geoshape
	case "ip":
		*f = Ip
	case "binary":
		*f = Binary
	case "keyword":
		*f = Keyword
	case "text":
		*f = Text
	case "search_as_you_type":
		*f = Searchasyoutype
	case "date":
		*f = Date
	case "date_nanos":
		*f = Datenanos
	case "boolean":
		*f = Boolean
	case "completion":
		*f = Completion
	case "nested":
		*f = Nested
	case "object":
		*f = Object
	case "murmur3":
		*f = Murmur3
	case "token_count":
		*f = Tokencount
	case "percolator":
		*f = Percolator
	case "integer":
		*f = Integer
	case "long":
		*f = Long
	case "short":
		*f = Short
	case "byte":
		*f = Byte
	case "float":
		*f = Float
	case "half_float":
		*f = Halffloat
	case "scaled_float":
		*f = Scaledfloat
	case "double":
		*f = Double
	case "integer_range":
		*f = Integerrange
	case "float_range":
		*f = Floatrange
	case "long_range":
		*f = Longrange
	case "double_range":
		*f = Doublerange
	case "date_range":
		*f = Daterange
	case "ip_range":
		*f = Iprange
	case "alias":
		*f = Alias
	case "join":
		*f = Join
	case "rank_feature":
		*f = Rankfeature
	case "rank_features":
		*f = Rankfeatures
	case "flattened":
		*f = Flattened
	case "shape":
		*f = Shape
	case "histogram":
		*f = Histogram
	case "constant_keyword":
		*f = Constantkeyword
	case "aggregate_metric_double":
		*f = Aggregatemetricdouble
	case "dense_vector":
		*f = Densevector
	case "match_only_text":
		*f = Matchonlytext
	default:
		*f = FieldType{string(text)}
	}

	return nil
}

func (f FieldType) String() string {
	return f.Name
}
