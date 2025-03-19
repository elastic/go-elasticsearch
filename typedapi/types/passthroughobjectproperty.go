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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syntheticsourcekeepenum"
)

// PassthroughObjectProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ea991724f4dd4f90c496eff547d3cc2e6529f509/specification/_types/mapping/complex.ts#L52-L57
type PassthroughObjectProperty struct {
	CopyTo      []string                       `json:"copy_to,omitempty"`
	Dynamic     *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Enabled     *bool                          `json:"enabled,omitempty"`
	Fields      map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove *int                           `json:"ignore_above,omitempty"`
	// Meta Metadata about the field.
	Meta                map[string]string                                `json:"meta,omitempty"`
	Priority            *int                                             `json:"priority,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	TimeSeriesDimension *bool                                            `json:"time_series_dimension,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *PassthroughObjectProperty) UnmarshalJSON(data []byte) error {

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

		case "dynamic":
			if err := dec.Decode(&s.Dynamic); err != nil {
				return fmt.Errorf("%s | %w", "Dynamic", err)
			}

		case "enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Enabled", err)
				}
				s.Enabled = &value
			case bool:
				s.Enabled = &v
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

		case "meta":
			if s.Meta == nil {
				s.Meta = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Meta); err != nil {
				return fmt.Errorf("%s | %w", "Meta", err)
			}

		case "priority":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Priority", err)
				}
				s.Priority = &value
			case float64:
				f := int(v)
				s.Priority = &f
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

		case "time_series_dimension":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TimeSeriesDimension", err)
				}
				s.TimeSeriesDimension = &value
			case bool:
				s.TimeSeriesDimension = &v
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
func (s PassthroughObjectProperty) MarshalJSON() ([]byte, error) {
	type innerPassthroughObjectProperty PassthroughObjectProperty
	tmp := innerPassthroughObjectProperty{
		CopyTo:              s.CopyTo,
		Dynamic:             s.Dynamic,
		Enabled:             s.Enabled,
		Fields:              s.Fields,
		IgnoreAbove:         s.IgnoreAbove,
		Meta:                s.Meta,
		Priority:            s.Priority,
		Properties:          s.Properties,
		Store:               s.Store,
		SyntheticSourceKeep: s.SyntheticSourceKeep,
		TimeSeriesDimension: s.TimeSeriesDimension,
		Type:                s.Type,
	}

	tmp.Type = "passthrough"

	return json.Marshal(tmp)
}

// NewPassthroughObjectProperty returns a PassthroughObjectProperty.
func NewPassthroughObjectProperty() *PassthroughObjectProperty {
	r := &PassthroughObjectProperty{
		Fields:     make(map[string]Property),
		Meta:       make(map[string]string),
		Properties: make(map[string]Property),
	}

	return r
}

// true

type PassthroughObjectPropertyVariant interface {
	PassthroughObjectPropertyCaster() *PassthroughObjectProperty
}

func (s *PassthroughObjectProperty) PassthroughObjectPropertyCaster() *PassthroughObjectProperty {
	return s
}
