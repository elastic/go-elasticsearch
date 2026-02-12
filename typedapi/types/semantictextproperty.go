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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// SemanticTextProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/_types/mapping/core.ts#L239-L274
type SemanticTextProperty struct {
	// ChunkingSettings Settings for chunking text into smaller passages. If specified, these will
	// override the
	// chunking settings sent in the inference endpoint associated with
	// inference_id. If chunking settings are updated,
	// they will not be applied to existing documents until they are reindexed.
	ChunkingSettings *ChunkingSettings `json:"chunking_settings,omitempty"`
	// Fields Multi-fields allow the same string value to be indexed in multiple ways for
	// different purposes, such as one
	// field for search and a multi-field for sorting and aggregations, or the same
	// string value analyzed by different analyzers.
	Fields map[string]Property `json:"fields,omitempty"`
	// IndexOptions Settings for index_options that override any defaults used by semantic_text,
	// for example
	// specific quantization settings.
	IndexOptions *SemanticTextIndexOptions `json:"index_options,omitempty"`
	// InferenceId Inference endpoint that will be used to generate embeddings for the field.
	// This parameter cannot be updated. Use the Create inference API to create the
	// endpoint.
	// If `search_inference_id` is specified, the inference endpoint will only be
	// used at index time.
	InferenceId *string           `json:"inference_id,omitempty"`
	Meta        map[string]string `json:"meta,omitempty"`
	// SearchInferenceId Inference endpoint that will be used to generate embeddings at query time.
	// You can update this parameter by using the Update mapping API. Use the Create
	// inference API to create the endpoint.
	// If not specified, the inference endpoint defined by inference_id will be used
	// at both index and query time.
	SearchInferenceId *string `json:"search_inference_id,omitempty"`
	Type              string  `json:"type,omitempty"`
}

func (s *SemanticTextProperty) UnmarshalJSON(data []byte) error {

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

		case "chunking_settings":
			if err := dec.Decode(&s.ChunkingSettings); err != nil {
				return fmt.Errorf("%s | %w", "ChunkingSettings", err)
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
				case "exponential_histogram":
					oo := NewExponentialHistogramProperty()
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

		case "index_options":
			if err := dec.Decode(&s.IndexOptions); err != nil {
				return fmt.Errorf("%s | %w", "IndexOptions", err)
			}

		case "inference_id":
			if err := dec.Decode(&s.InferenceId); err != nil {
				return fmt.Errorf("%s | %w", "InferenceId", err)
			}

		case "meta":
			if s.Meta == nil {
				s.Meta = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Meta); err != nil {
				return fmt.Errorf("%s | %w", "Meta", err)
			}

		case "search_inference_id":
			if err := dec.Decode(&s.SearchInferenceId); err != nil {
				return fmt.Errorf("%s | %w", "SearchInferenceId", err)
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
func (s SemanticTextProperty) MarshalJSON() ([]byte, error) {
	type innerSemanticTextProperty SemanticTextProperty
	tmp := innerSemanticTextProperty{
		ChunkingSettings:  s.ChunkingSettings,
		Fields:            s.Fields,
		IndexOptions:      s.IndexOptions,
		InferenceId:       s.InferenceId,
		Meta:              s.Meta,
		SearchInferenceId: s.SearchInferenceId,
		Type:              s.Type,
	}

	tmp.Type = "semantic_text"

	return json.Marshal(tmp)
}

// NewSemanticTextProperty returns a SemanticTextProperty.
func NewSemanticTextProperty() *SemanticTextProperty {
	r := &SemanticTextProperty{
		Fields: make(map[string]Property),
		Meta:   make(map[string]string),
	}

	return r
}

type SemanticTextPropertyVariant interface {
	SemanticTextPropertyCaster() *SemanticTextProperty
}

func (s *SemanticTextProperty) SemanticTextPropertyCaster() *SemanticTextProperty {
	return s
}

func (s *SemanticTextProperty) PropertyCaster() *Property {
	if s == nil {
		return nil
	}
	o := Property(s)
	return &o
}
