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

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

// PointProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/mapping/geo.ts#L74-L79
type PointProperty struct {
	CopyTo          []string                       `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	IgnoreZValue    *bool                          `json:"ignore_z_value,omitempty"`
	// Meta Metadata about the field.
	Meta                map[string]string                                `json:"meta,omitempty"`
	NullValue           *string                                          `json:"null_value,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *PointProperty) UnmarshalJSON(data []byte) error {

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

		case "copy_to":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "CopyTo", err)
				}

				s.CopyTo = append(s.CopyTo, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.CopyTo); err != nil {
					return fmt.Errorf("%s | %w", "CopyTo", err)
				}
			}

		case "doc_values":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DocValues", err)
				}
				s.DocValues = &value
			case bool:
				s.DocValues = &v
			}

		case "dynamic":
			if err := dec.Decode(&s.Dynamic); err != nil {
				return fmt.Errorf("%s | %w", "Dynamic", err)
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

		case "ignore_malformed":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreMalformed", err)
				}
				s.IgnoreMalformed = &value
			case bool:
				s.IgnoreMalformed = &v
			}

		case "ignore_z_value":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreZValue", err)
				}
				s.IgnoreZValue = &value
			case bool:
				s.IgnoreZValue = &v
			}

		case "meta":
			if s.Meta == nil {
				s.Meta = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Meta); err != nil {
				return fmt.Errorf("%s | %w", "Meta", err)
			}

		case "null_value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "NullValue", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.NullValue = &o

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

		case "store":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Store", err)
				}
				s.Store = &value
			case bool:
				s.Store = &v
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
func (s PointProperty) MarshalJSON() ([]byte, error) {
	type innerPointProperty PointProperty
	tmp := innerPointProperty{
		CopyTo:              s.CopyTo,
		DocValues:           s.DocValues,
		Dynamic:             s.Dynamic,
		Fields:              s.Fields,
		IgnoreAbove:         s.IgnoreAbove,
		IgnoreMalformed:     s.IgnoreMalformed,
		IgnoreZValue:        s.IgnoreZValue,
		Meta:                s.Meta,
		NullValue:           s.NullValue,
		Properties:          s.Properties,
		Store:               s.Store,
		SyntheticSourceKeep: s.SyntheticSourceKeep,
		Type:                s.Type,
	}

	tmp.Type = "point"

	return json.Marshal(tmp)
}

// NewPointProperty returns a PointProperty.
func NewPointProperty() *PointProperty {
	r := &PointProperty{
		Fields:     make(map[string]Property),
		Meta:       make(map[string]string),
		Properties: make(map[string]Property),
	}

	return r
}

type PointPropertyVariant interface {
	PointPropertyCaster() *PointProperty
}

func (s *PointProperty) PointPropertyCaster() *PointProperty {
	return s
}

func (s *PointProperty) PropertyCaster() *Property {
	o := Property(s)
	return &o
}
