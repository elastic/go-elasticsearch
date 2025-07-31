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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/densevectorelementtype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/densevectorsimilarity"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syntheticsourcekeepenum"
)

// DenseVectorProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/mapping/DenseVectorProperty.ts#L23-L62
type DenseVectorProperty struct {
	// Dims Number of vector dimensions. Can't exceed `4096`. If `dims` is not specified,
	// it will be set to the length of
	// the first vector added to the field.
	Dims    *int                           `json:"dims,omitempty"`
	Dynamic *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	// ElementType The data type used to encode vectors. The supported data types are `float`
	// (default), `byte`, and `bit`.
	ElementType *densevectorelementtype.DenseVectorElementType `json:"element_type,omitempty"`
	Fields      map[string]Property                            `json:"fields,omitempty"`
	IgnoreAbove *int                                           `json:"ignore_above,omitempty"`
	// Index If `true`, you can search this field using the kNN search API.
	Index *bool `json:"index,omitempty"`
	// IndexOptions An optional section that configures the kNN indexing algorithm. The HNSW
	// algorithm has two internal parameters
	// that influence how the data structure is built. These can be adjusted to
	// improve the accuracy of results, at the
	// expense of slower indexing speed.
	//
	// This parameter can only be specified when `index` is `true`.
	IndexOptions *DenseVectorIndexOptions `json:"index_options,omitempty"`
	// Meta Metadata about the field.
	Meta       map[string]string   `json:"meta,omitempty"`
	Properties map[string]Property `json:"properties,omitempty"`
	// Similarity The vector similarity metric to use in kNN search.
	//
	// Documents are ranked by their vector field's similarity to the query vector.
	// The `_score` of each document will
	// be derived from the similarity, in a way that ensures scores are positive and
	// that a larger score corresponds
	// to a higher ranking.
	//
	// Defaults to `l2_norm` when `element_type` is `bit` otherwise defaults to
	// `cosine`.
	//
	// `bit` vectors only support `l2_norm` as their similarity metric.
	//
	// This parameter can only be specified when `index` is `true`.
	Similarity          *densevectorsimilarity.DenseVectorSimilarity     `json:"similarity,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *DenseVectorProperty) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "dims":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Dims", err)
				}
				s.Dims = &value
			case float64:
				f := int(v)
				s.Dims = &f
			}

		case "dynamic":
			if err := dec.Decode(&s.Dynamic); err != nil {
				return fmt.Errorf("%s | %w", "Dynamic", err)
			}

		case "element_type":
			if err := dec.Decode(&s.ElementType); err != nil {
				return fmt.Errorf("%s | %w", "ElementType", err)
			}

		case "fields":
			if s.Fields == nil {
				s.Fields = make(map[string]Property, 0)
			}
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]any)
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)
				if _, ok := kind["type"]; !ok {
					kind["type"] = "object"
				}
				switch kind["type"] {
				case "binary":
					oo := NewBinaryProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "boolean":
					oo := NewBooleanProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "{dynamic_type}":
					oo := NewDynamicProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "join":
					oo := NewJoinProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "keyword":
					oo := NewKeywordProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "match_only_text":
					oo := NewMatchOnlyTextProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "percolator":
					oo := NewPercolatorProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "rank_feature":
					oo := NewRankFeatureProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "rank_features":
					oo := NewRankFeaturesProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "search_as_you_type":
					oo := NewSearchAsYouTypeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "text":
					oo := NewTextProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "version":
					oo := NewVersionProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "wildcard":
					oo := NewWildcardProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "date_nanos":
					oo := NewDateNanosProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "date":
					oo := NewDateProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "aggregate_metric_double":
					oo := NewAggregateMetricDoubleProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "dense_vector":
					oo := NewDenseVectorProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "flattened":
					oo := NewFlattenedProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "nested":
					oo := NewNestedProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "object":
					oo := NewObjectProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "passthrough":
					oo := NewPassthroughObjectProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "rank_vectors":
					oo := NewRankVectorProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "semantic_text":
					oo := NewSemanticTextProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "sparse_vector":
					oo := NewSparseVectorProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "completion":
					oo := NewCompletionProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "constant_keyword":
					oo := NewConstantKeywordProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "counted_keyword":
					oo := NewCountedKeywordProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "alias":
					oo := NewFieldAliasProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "histogram":
					oo := NewHistogramProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "ip":
					oo := NewIpProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "murmur3":
					oo := NewMurmur3HashProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "token_count":
					oo := NewTokenCountProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "geo_point":
					oo := NewGeoPointProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "geo_shape":
					oo := NewGeoShapeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "point":
					oo := NewPointProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "shape":
					oo := NewShapeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "byte":
					oo := NewByteNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "double":
					oo := NewDoubleNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "float":
					oo := NewFloatNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "half_float":
					oo := NewHalfFloatNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "integer":
					oo := NewIntegerNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "long":
					oo := NewLongNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "scaled_float":
					oo := NewScaledFloatNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "short":
					oo := NewShortNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "unsigned_long":
					oo := NewUnsignedLongNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "date_range":
					oo := NewDateRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "double_range":
					oo := NewDoubleRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "float_range":
					oo := NewFloatRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "integer_range":
					oo := NewIntegerRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "ip_range":
					oo := NewIpRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "long_range":
					oo := NewLongRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				case "icu_collation_keyword":
					oo := NewIcuCollationProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Fields | %w", err)
					}
					s.Fields[key] = oo
				default:
					oo := new(Property)
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("new(Property) | %w", err)
					}
					s.Fields[key] = oo
				}
			}

		case "ignore_above":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreAbove", err)
				}
				s.IgnoreAbove = &value
			case float64:
				f := int(v)
				s.IgnoreAbove = &f
			}

		case "index":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Index", err)
				}
				s.Index = &value
			case bool:
				s.Index = &v
			}

		case "index_options":
			if err := dec.Decode(&s.IndexOptions); err != nil {
				return fmt.Errorf("%s | %w", "IndexOptions", err)
			}

		case "meta":
			if s.Meta == nil {
				s.Meta = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Meta); err != nil {
				return fmt.Errorf("%s | %w", "Meta", err)
			}

		case "properties":
			if s.Properties == nil {
				s.Properties = make(map[string]Property, 0)
			}
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]any)
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)
				if _, ok := kind["type"]; !ok {
					kind["type"] = "object"
				}
				switch kind["type"] {
				case "binary":
					oo := NewBinaryProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "boolean":
					oo := NewBooleanProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "{dynamic_type}":
					oo := NewDynamicProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "join":
					oo := NewJoinProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "keyword":
					oo := NewKeywordProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "match_only_text":
					oo := NewMatchOnlyTextProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "percolator":
					oo := NewPercolatorProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "rank_feature":
					oo := NewRankFeatureProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "rank_features":
					oo := NewRankFeaturesProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "search_as_you_type":
					oo := NewSearchAsYouTypeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "text":
					oo := NewTextProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "version":
					oo := NewVersionProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "wildcard":
					oo := NewWildcardProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "date_nanos":
					oo := NewDateNanosProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "date":
					oo := NewDateProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "aggregate_metric_double":
					oo := NewAggregateMetricDoubleProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "dense_vector":
					oo := NewDenseVectorProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "flattened":
					oo := NewFlattenedProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "nested":
					oo := NewNestedProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "object":
					oo := NewObjectProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "passthrough":
					oo := NewPassthroughObjectProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "rank_vectors":
					oo := NewRankVectorProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "semantic_text":
					oo := NewSemanticTextProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "sparse_vector":
					oo := NewSparseVectorProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "completion":
					oo := NewCompletionProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "constant_keyword":
					oo := NewConstantKeywordProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "counted_keyword":
					oo := NewCountedKeywordProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "alias":
					oo := NewFieldAliasProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "histogram":
					oo := NewHistogramProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "ip":
					oo := NewIpProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "murmur3":
					oo := NewMurmur3HashProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "token_count":
					oo := NewTokenCountProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "geo_point":
					oo := NewGeoPointProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "geo_shape":
					oo := NewGeoShapeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "point":
					oo := NewPointProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "shape":
					oo := NewShapeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "byte":
					oo := NewByteNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "double":
					oo := NewDoubleNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "float":
					oo := NewFloatNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "half_float":
					oo := NewHalfFloatNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "integer":
					oo := NewIntegerNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "long":
					oo := NewLongNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "scaled_float":
					oo := NewScaledFloatNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "short":
					oo := NewShortNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "unsigned_long":
					oo := NewUnsignedLongNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "date_range":
					oo := NewDateRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "double_range":
					oo := NewDoubleRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "float_range":
					oo := NewFloatRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "integer_range":
					oo := NewIntegerRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "ip_range":
					oo := NewIpRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "long_range":
					oo := NewLongRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				case "icu_collation_keyword":
					oo := NewIcuCollationProperty()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Properties | %w", err)
					}
					s.Properties[key] = oo
				default:
					oo := new(Property)
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("new(Property) | %w", err)
					}
					s.Properties[key] = oo
				}
			}

		case "similarity":
			if err := dec.Decode(&s.Similarity); err != nil {
				return fmt.Errorf("%s | %w", "Similarity", err)
			}

		case "synthetic_source_keep":
			if err := dec.Decode(&s.SyntheticSourceKeep); err != nil {
				return fmt.Errorf("%s | %w", "SyntheticSourceKeep", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s DenseVectorProperty) MarshalJSON() ([]byte, error) {
	type innerDenseVectorProperty DenseVectorProperty
	tmp := innerDenseVectorProperty{
		Dims:                s.Dims,
		Dynamic:             s.Dynamic,
		ElementType:         s.ElementType,
		Fields:              s.Fields,
		IgnoreAbove:         s.IgnoreAbove,
		Index:               s.Index,
		IndexOptions:        s.IndexOptions,
		Meta:                s.Meta,
		Properties:          s.Properties,
		Similarity:          s.Similarity,
		SyntheticSourceKeep: s.SyntheticSourceKeep,
		Type:                s.Type,
	}

	tmp.Type = "dense_vector"

	return json.Marshal(tmp)
}

// NewDenseVectorProperty returns a DenseVectorProperty.
func NewDenseVectorProperty() *DenseVectorProperty {
	r := &DenseVectorProperty{
		Fields:     make(map[string]Property),
		Meta:       make(map[string]string),
		Properties: make(map[string]Property),
	}

	return r
}
