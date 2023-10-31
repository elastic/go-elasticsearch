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

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// CompletionProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/mapping/specialized.ts#L27-L35
type CompletionProperty struct {
	Analyzer       *string                        `json:"analyzer,omitempty"`
	Contexts       []SuggestContext               `json:"contexts,omitempty"`
	CopyTo         []string                       `json:"copy_to,omitempty"`
	DocValues      *bool                          `json:"doc_values,omitempty"`
	Dynamic        *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields         map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove    *int                           `json:"ignore_above,omitempty"`
	MaxInputLength *int                           `json:"max_input_length,omitempty"`
	// Meta Metadata about the field.
	Meta                       map[string]string   `json:"meta,omitempty"`
	PreservePositionIncrements *bool               `json:"preserve_position_increments,omitempty"`
	PreserveSeparators         *bool               `json:"preserve_separators,omitempty"`
	Properties                 map[string]Property `json:"properties,omitempty"`
	SearchAnalyzer             *string             `json:"search_analyzer,omitempty"`
	Similarity                 *string             `json:"similarity,omitempty"`
	Store                      *bool               `json:"store,omitempty"`
	Type                       string              `json:"type,omitempty"`
}

func (s *CompletionProperty) UnmarshalJSON(data []byte) error {

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

		case "analyzer":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Analyzer = &o

		case "contexts":
			if err := dec.Decode(&s.Contexts); err != nil {
				return err
			}

		case "copy_to":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.CopyTo = append(s.CopyTo, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.CopyTo); err != nil {
					return err
				}
			}

		case "doc_values":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.DocValues = &value
			case bool:
				s.DocValues = &v
			}

		case "dynamic":
			if err := dec.Decode(&s.Dynamic); err != nil {
				return err
			}

		case "fields":
			if s.Fields == nil {
				s.Fields = make(map[string]Property, 0)
			}
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]interface{})
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
						return err
					}
					s.Fields[key] = oo
				case "boolean":
					oo := NewBooleanProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "{dynamic_property}":
					oo := NewDynamicProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "join":
					oo := NewJoinProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "keyword":
					oo := NewKeywordProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "match_only_text":
					oo := NewMatchOnlyTextProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "percolator":
					oo := NewPercolatorProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "rank_feature":
					oo := NewRankFeatureProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "rank_features":
					oo := NewRankFeaturesProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "search_as_you_type":
					oo := NewSearchAsYouTypeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "text":
					oo := NewTextProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "version":
					oo := NewVersionProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "wildcard":
					oo := NewWildcardProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "date_nanos":
					oo := NewDateNanosProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "date":
					oo := NewDateProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "aggregate_metric_double":
					oo := NewAggregateMetricDoubleProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "dense_vector":
					oo := NewDenseVectorProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "flattened":
					oo := NewFlattenedProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "nested":
					oo := NewNestedProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "object":
					oo := NewObjectProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "completion":
					oo := NewCompletionProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "constant_keyword":
					oo := NewConstantKeywordProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "alias":
					oo := NewFieldAliasProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "histogram":
					oo := NewHistogramProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "ip":
					oo := NewIpProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "murmur3":
					oo := NewMurmur3HashProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "token_count":
					oo := NewTokenCountProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "geo_point":
					oo := NewGeoPointProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "geo_shape":
					oo := NewGeoShapeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "point":
					oo := NewPointProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "shape":
					oo := NewShapeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "byte":
					oo := NewByteNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "double":
					oo := NewDoubleNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "float":
					oo := NewFloatNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "half_float":
					oo := NewHalfFloatNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "integer":
					oo := NewIntegerNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "long":
					oo := NewLongNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "scaled_float":
					oo := NewScaledFloatNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "short":
					oo := NewShortNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "unsigned_long":
					oo := NewUnsignedLongNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "date_range":
					oo := NewDateRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "double_range":
					oo := NewDoubleRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "float_range":
					oo := NewFloatRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "integer_range":
					oo := NewIntegerRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "ip_range":
					oo := NewIpRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				case "long_range":
					oo := NewLongRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				default:
					oo := new(Property)
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Fields[key] = oo
				}
			}

		case "ignore_above":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.IgnoreAbove = &value
			case float64:
				f := int(v)
				s.IgnoreAbove = &f
			}

		case "max_input_length":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxInputLength = &value
			case float64:
				f := int(v)
				s.MaxInputLength = &f
			}

		case "meta":
			if s.Meta == nil {
				s.Meta = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "preserve_position_increments":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.PreservePositionIncrements = &value
			case bool:
				s.PreservePositionIncrements = &v
			}

		case "preserve_separators":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.PreserveSeparators = &value
			case bool:
				s.PreserveSeparators = &v
			}

		case "properties":
			if s.Properties == nil {
				s.Properties = make(map[string]Property, 0)
			}
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]interface{})
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
						return err
					}
					s.Properties[key] = oo
				case "boolean":
					oo := NewBooleanProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "{dynamic_property}":
					oo := NewDynamicProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "join":
					oo := NewJoinProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "keyword":
					oo := NewKeywordProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "match_only_text":
					oo := NewMatchOnlyTextProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "percolator":
					oo := NewPercolatorProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "rank_feature":
					oo := NewRankFeatureProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "rank_features":
					oo := NewRankFeaturesProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "search_as_you_type":
					oo := NewSearchAsYouTypeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "text":
					oo := NewTextProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "version":
					oo := NewVersionProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "wildcard":
					oo := NewWildcardProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "date_nanos":
					oo := NewDateNanosProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "date":
					oo := NewDateProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "aggregate_metric_double":
					oo := NewAggregateMetricDoubleProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "dense_vector":
					oo := NewDenseVectorProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "flattened":
					oo := NewFlattenedProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "nested":
					oo := NewNestedProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "object":
					oo := NewObjectProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "completion":
					oo := NewCompletionProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "constant_keyword":
					oo := NewConstantKeywordProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "alias":
					oo := NewFieldAliasProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "histogram":
					oo := NewHistogramProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "ip":
					oo := NewIpProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "murmur3":
					oo := NewMurmur3HashProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "token_count":
					oo := NewTokenCountProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "geo_point":
					oo := NewGeoPointProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "geo_shape":
					oo := NewGeoShapeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "point":
					oo := NewPointProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "shape":
					oo := NewShapeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "byte":
					oo := NewByteNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "double":
					oo := NewDoubleNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "float":
					oo := NewFloatNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "half_float":
					oo := NewHalfFloatNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "integer":
					oo := NewIntegerNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "long":
					oo := NewLongNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "scaled_float":
					oo := NewScaledFloatNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "short":
					oo := NewShortNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "unsigned_long":
					oo := NewUnsignedLongNumberProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "date_range":
					oo := NewDateRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "double_range":
					oo := NewDoubleRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "float_range":
					oo := NewFloatRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "integer_range":
					oo := NewIntegerRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "ip_range":
					oo := NewIpRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				case "long_range":
					oo := NewLongRangeProperty()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				default:
					oo := new(Property)
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Properties[key] = oo
				}
			}

		case "search_analyzer":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchAnalyzer = &o

		case "similarity":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Similarity = &o

		case "store":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Store = &value
			case bool:
				s.Store = &v
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s CompletionProperty) MarshalJSON() ([]byte, error) {
	type innerCompletionProperty CompletionProperty
	tmp := innerCompletionProperty{
		Analyzer:                   s.Analyzer,
		Contexts:                   s.Contexts,
		CopyTo:                     s.CopyTo,
		DocValues:                  s.DocValues,
		Dynamic:                    s.Dynamic,
		Fields:                     s.Fields,
		IgnoreAbove:                s.IgnoreAbove,
		MaxInputLength:             s.MaxInputLength,
		Meta:                       s.Meta,
		PreservePositionIncrements: s.PreservePositionIncrements,
		PreserveSeparators:         s.PreserveSeparators,
		Properties:                 s.Properties,
		SearchAnalyzer:             s.SearchAnalyzer,
		Similarity:                 s.Similarity,
		Store:                      s.Store,
		Type:                       s.Type,
	}

	tmp.Type = "completion"

	return json.Marshal(tmp)
}

// NewCompletionProperty returns a CompletionProperty.
func NewCompletionProperty() *CompletionProperty {
	r := &CompletionProperty{
		Fields:     make(map[string]Property, 0),
		Meta:       make(map[string]string, 0),
		Properties: make(map[string]Property, 0),
	}

	return r
}
