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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/matchtype"
)

// DynamicTemplate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/mapping/dynamic-template.ts#L22-L30
type DynamicTemplate struct {
	Mapping          Property             `json:"mapping,omitempty"`
	Match            *string              `json:"match,omitempty"`
	MatchMappingType *string              `json:"match_mapping_type,omitempty"`
	MatchPattern     *matchtype.MatchType `json:"match_pattern,omitempty"`
	PathMatch        *string              `json:"path_match,omitempty"`
	PathUnmatch      *string              `json:"path_unmatch,omitempty"`
	Unmatch          *string              `json:"unmatch,omitempty"`
}

func (s *DynamicTemplate) UnmarshalJSON(data []byte) error {

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

		case "mapping":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			kind := make(map[string]string, 0)
			localDec := json.NewDecoder(source)
			localDec.Decode(&kind)
			source.Seek(0, io.SeekStart)
			if _, ok := kind["type"]; !ok {
				kind["type"] = "object"
			}
			switch kind["type"] {

			case "binary":
				o := NewBinaryProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "boolean":
				o := NewBooleanProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "{dynamic_property}":
				o := NewDynamicProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "join":
				o := NewJoinProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "keyword":
				o := NewKeywordProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "match_only_text":
				o := NewMatchOnlyTextProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "percolator":
				o := NewPercolatorProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "rank_feature":
				o := NewRankFeatureProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "rank_features":
				o := NewRankFeaturesProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "search_as_you_type":
				o := NewSearchAsYouTypeProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "text":
				o := NewTextProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "version":
				o := NewVersionProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "wildcard":
				o := NewWildcardProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "date_nanos":
				o := NewDateNanosProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "date":
				o := NewDateProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "aggregate_metric_double":
				o := NewAggregateMetricDoubleProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "dense_vector":
				o := NewDenseVectorProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "flattened":
				o := NewFlattenedProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "nested":
				o := NewNestedProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "object":
				o := NewObjectProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "completion":
				o := NewCompletionProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "constant_keyword":
				o := NewConstantKeywordProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "alias":
				o := NewFieldAliasProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "histogram":
				o := NewHistogramProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "ip":
				o := NewIpProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "murmur3":
				o := NewMurmur3HashProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "token_count":
				o := NewTokenCountProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "geo_point":
				o := NewGeoPointProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "geo_shape":
				o := NewGeoShapeProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "point":
				o := NewPointProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "shape":
				o := NewShapeProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "byte":
				o := NewByteNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "double":
				o := NewDoubleNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "float":
				o := NewFloatNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "half_float":
				o := NewHalfFloatNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "integer":
				o := NewIntegerNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "long":
				o := NewLongNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "scaled_float":
				o := NewScaledFloatNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "short":
				o := NewShortNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "unsigned_long":
				o := NewUnsignedLongNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "date_range":
				o := NewDateRangeProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "double_range":
				o := NewDoubleRangeProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "float_range":
				o := NewFloatRangeProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "integer_range":
				o := NewIntegerRangeProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "ip_range":
				o := NewIpRangeProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			case "long_range":
				o := NewLongRangeProperty()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Mapping = *o
			default:
				if err := localDec.Decode(&s.Mapping); err != nil {
					return err
				}
			}

		case "match":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Match = &o

		case "match_mapping_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MatchMappingType = &o

		case "match_pattern":
			if err := dec.Decode(&s.MatchPattern); err != nil {
				return err
			}

		case "path_match":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PathMatch = &o

		case "path_unmatch":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PathUnmatch = &o

		case "unmatch":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Unmatch = &o

		}
	}
	return nil
}

// NewDynamicTemplate returns a DynamicTemplate.
func NewDynamicTemplate() *DynamicTemplate {
	r := &DynamicTemplate{}

	return r
}
