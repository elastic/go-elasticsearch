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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/matchtype"
)

// DynamicTemplate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/mapping/dynamic-template.ts#L23-L43
type DynamicTemplate struct {
	AdditionalDynamicTemplateProperty map[string]json.RawMessage `json:"-"`
	Mapping                           Property                   `json:"mapping,omitempty"`
	Match                             []string                   `json:"match,omitempty"`
	MatchMappingType                  []string                   `json:"match_mapping_type,omitempty"`
	MatchPattern                      *matchtype.MatchType       `json:"match_pattern,omitempty"`
	PathMatch                         []string                   `json:"path_match,omitempty"`
	PathUnmatch                       []string                   `json:"path_unmatch,omitempty"`
	Runtime                           *RuntimeField              `json:"runtime,omitempty"`
	Unmatch                           []string                   `json:"unmatch,omitempty"`
	UnmatchMappingType                []string                   `json:"unmatch_mapping_type,omitempty"`
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
					return fmt.Errorf("%s | %w", "binary", err)
				}
				s.Mapping = *o
			case "boolean":
				o := NewBooleanProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "boolean", err)
				}
				s.Mapping = *o
			case "{dynamic_type}":
				o := NewDynamicProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "{dynamic_type}", err)
				}
				s.Mapping = *o
			case "join":
				o := NewJoinProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "join", err)
				}
				s.Mapping = *o
			case "keyword":
				o := NewKeywordProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "keyword", err)
				}
				s.Mapping = *o
			case "match_only_text":
				o := NewMatchOnlyTextProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "match_only_text", err)
				}
				s.Mapping = *o
			case "percolator":
				o := NewPercolatorProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "percolator", err)
				}
				s.Mapping = *o
			case "rank_feature":
				o := NewRankFeatureProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "rank_feature", err)
				}
				s.Mapping = *o
			case "rank_features":
				o := NewRankFeaturesProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "rank_features", err)
				}
				s.Mapping = *o
			case "search_as_you_type":
				o := NewSearchAsYouTypeProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "search_as_you_type", err)
				}
				s.Mapping = *o
			case "text":
				o := NewTextProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "text", err)
				}
				s.Mapping = *o
			case "version":
				o := NewVersionProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "version", err)
				}
				s.Mapping = *o
			case "wildcard":
				o := NewWildcardProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "wildcard", err)
				}
				s.Mapping = *o
			case "date_nanos":
				o := NewDateNanosProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "date_nanos", err)
				}
				s.Mapping = *o
			case "date":
				o := NewDateProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "date", err)
				}
				s.Mapping = *o
			case "aggregate_metric_double":
				o := NewAggregateMetricDoubleProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "aggregate_metric_double", err)
				}
				s.Mapping = *o
			case "dense_vector":
				o := NewDenseVectorProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "dense_vector", err)
				}
				s.Mapping = *o
			case "flattened":
				o := NewFlattenedProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "flattened", err)
				}
				s.Mapping = *o
			case "nested":
				o := NewNestedProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "nested", err)
				}
				s.Mapping = *o
			case "object":
				o := NewObjectProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "object", err)
				}
				s.Mapping = *o
			case "passthrough":
				o := NewPassthroughObjectProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "passthrough", err)
				}
				s.Mapping = *o
			case "rank_vectors":
				o := NewRankVectorProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "rank_vectors", err)
				}
				s.Mapping = *o
			case "semantic_text":
				o := NewSemanticTextProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "semantic_text", err)
				}
				s.Mapping = *o
			case "sparse_vector":
				o := NewSparseVectorProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "sparse_vector", err)
				}
				s.Mapping = *o
			case "completion":
				o := NewCompletionProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "completion", err)
				}
				s.Mapping = *o
			case "constant_keyword":
				o := NewConstantKeywordProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "constant_keyword", err)
				}
				s.Mapping = *o
			case "counted_keyword":
				o := NewCountedKeywordProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "counted_keyword", err)
				}
				s.Mapping = *o
			case "alias":
				o := NewFieldAliasProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "alias", err)
				}
				s.Mapping = *o
			case "histogram":
				o := NewHistogramProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "histogram", err)
				}
				s.Mapping = *o
			case "ip":
				o := NewIpProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "ip", err)
				}
				s.Mapping = *o
			case "murmur3":
				o := NewMurmur3HashProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "murmur3", err)
				}
				s.Mapping = *o
			case "token_count":
				o := NewTokenCountProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "token_count", err)
				}
				s.Mapping = *o
			case "geo_point":
				o := NewGeoPointProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "geo_point", err)
				}
				s.Mapping = *o
			case "geo_shape":
				o := NewGeoShapeProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "geo_shape", err)
				}
				s.Mapping = *o
			case "point":
				o := NewPointProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "point", err)
				}
				s.Mapping = *o
			case "shape":
				o := NewShapeProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "shape", err)
				}
				s.Mapping = *o
			case "byte":
				o := NewByteNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "byte", err)
				}
				s.Mapping = *o
			case "double":
				o := NewDoubleNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "double", err)
				}
				s.Mapping = *o
			case "float":
				o := NewFloatNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "float", err)
				}
				s.Mapping = *o
			case "half_float":
				o := NewHalfFloatNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "half_float", err)
				}
				s.Mapping = *o
			case "integer":
				o := NewIntegerNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "integer", err)
				}
				s.Mapping = *o
			case "long":
				o := NewLongNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "long", err)
				}
				s.Mapping = *o
			case "scaled_float":
				o := NewScaledFloatNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "scaled_float", err)
				}
				s.Mapping = *o
			case "short":
				o := NewShortNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "short", err)
				}
				s.Mapping = *o
			case "unsigned_long":
				o := NewUnsignedLongNumberProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "unsigned_long", err)
				}
				s.Mapping = *o
			case "date_range":
				o := NewDateRangeProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "date_range", err)
				}
				s.Mapping = *o
			case "double_range":
				o := NewDoubleRangeProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "double_range", err)
				}
				s.Mapping = *o
			case "float_range":
				o := NewFloatRangeProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "float_range", err)
				}
				s.Mapping = *o
			case "integer_range":
				o := NewIntegerRangeProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "integer_range", err)
				}
				s.Mapping = *o
			case "ip_range":
				o := NewIpRangeProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "ip_range", err)
				}
				s.Mapping = *o
			case "long_range":
				o := NewLongRangeProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "long_range", err)
				}
				s.Mapping = *o
			case "icu_collation_keyword":
				o := NewIcuCollationProperty()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "icu_collation_keyword", err)
				}
				s.Mapping = *o
			default:
				if err := localDec.Decode(&s.Mapping); err != nil {
					return fmt.Errorf("Mapping | %w", err)
				}
			}

		case "match":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Match", err)
				}

				s.Match = append(s.Match, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Match); err != nil {
					return fmt.Errorf("%s | %w", "Match", err)
				}
			}

		case "match_mapping_type":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "MatchMappingType", err)
				}

				s.MatchMappingType = append(s.MatchMappingType, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.MatchMappingType); err != nil {
					return fmt.Errorf("%s | %w", "MatchMappingType", err)
				}
			}

		case "match_pattern":
			if err := dec.Decode(&s.MatchPattern); err != nil {
				return fmt.Errorf("%s | %w", "MatchPattern", err)
			}

		case "path_match":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "PathMatch", err)
				}

				s.PathMatch = append(s.PathMatch, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.PathMatch); err != nil {
					return fmt.Errorf("%s | %w", "PathMatch", err)
				}
			}

		case "path_unmatch":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "PathUnmatch", err)
				}

				s.PathUnmatch = append(s.PathUnmatch, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.PathUnmatch); err != nil {
					return fmt.Errorf("%s | %w", "PathUnmatch", err)
				}
			}

		case "runtime":
			if err := dec.Decode(&s.Runtime); err != nil {
				return fmt.Errorf("%s | %w", "Runtime", err)
			}

		case "unmatch":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Unmatch", err)
				}

				s.Unmatch = append(s.Unmatch, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Unmatch); err != nil {
					return fmt.Errorf("%s | %w", "Unmatch", err)
				}
			}

		case "unmatch_mapping_type":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "UnmatchMappingType", err)
				}

				s.UnmatchMappingType = append(s.UnmatchMappingType, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.UnmatchMappingType); err != nil {
					return fmt.Errorf("%s | %w", "UnmatchMappingType", err)
				}
			}

		default:

			if key, ok := t.(string); ok {
				if s.AdditionalDynamicTemplateProperty == nil {
					s.AdditionalDynamicTemplateProperty = make(map[string]json.RawMessage, 0)
				}
				raw := new(json.RawMessage)
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "AdditionalDynamicTemplateProperty", err)
				}
				s.AdditionalDynamicTemplateProperty[key] = *raw
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s DynamicTemplate) MarshalJSON() ([]byte, error) {
	type opt DynamicTemplate
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.AdditionalDynamicTemplateProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalDynamicTemplateProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewDynamicTemplate returns a DynamicTemplate.
func NewDynamicTemplate() *DynamicTemplate {
	r := &DynamicTemplate{
		AdditionalDynamicTemplateProperty: make(map[string]json.RawMessage),
	}

	return r
}
