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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/subobjects"
)

// TypeMapping type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/mapping/TypeMapping.ts#L34-L57
type TypeMapping struct {
	AllField             *AllField                      `json:"all_field,omitempty"`
	DataStreamTimestamp_ *DataStreamTimestamp           `json:"_data_stream_timestamp,omitempty"`
	DateDetection        *bool                          `json:"date_detection,omitempty"`
	Dynamic              *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	DynamicDateFormats   []string                       `json:"dynamic_date_formats,omitempty"`
	DynamicTemplates     []map[string]DynamicTemplate   `json:"dynamic_templates,omitempty"`
	Enabled              *bool                          `json:"enabled,omitempty"`
	FieldNames_          *FieldNamesField               `json:"_field_names,omitempty"`
	IndexField           *IndexField                    `json:"index_field,omitempty"`
	Meta_                Metadata                       `json:"_meta,omitempty"`
	NumericDetection     *bool                          `json:"numeric_detection,omitempty"`
	Properties           map[string]Property            `json:"properties,omitempty"`
	Routing_             *RoutingField                  `json:"_routing,omitempty"`
	Runtime              map[string]RuntimeField        `json:"runtime,omitempty"`
	Size_                *SizeField                     `json:"_size,omitempty"`
	Source_              *SourceField                   `json:"_source,omitempty"`
	Subobjects           *subobjects.Subobjects         `json:"subobjects,omitempty"`
}

func (s *TypeMapping) UnmarshalJSON(data []byte) error {

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

		case "all_field":
			if err := dec.Decode(&s.AllField); err != nil {
				return fmt.Errorf("%s | %w", "AllField", err)
			}

		case "_data_stream_timestamp":
			if err := dec.Decode(&s.DataStreamTimestamp_); err != nil {
				return fmt.Errorf("%s | %w", "DataStreamTimestamp_", err)
			}

		case "date_detection":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DateDetection", err)
				}
				s.DateDetection = &value
			case bool:
				s.DateDetection = &v
			}

		case "dynamic":
			if err := dec.Decode(&s.Dynamic); err != nil {
				return fmt.Errorf("%s | %w", "Dynamic", err)
			}

		case "dynamic_date_formats":
			if err := dec.Decode(&s.DynamicDateFormats); err != nil {
				return fmt.Errorf("%s | %w", "DynamicDateFormats", err)
			}

		case "dynamic_templates":
			if err := dec.Decode(&s.DynamicTemplates); err != nil {
				return fmt.Errorf("%s | %w", "DynamicTemplates", err)
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

		case "_field_names":
			if err := dec.Decode(&s.FieldNames_); err != nil {
				return fmt.Errorf("%s | %w", "FieldNames_", err)
			}

		case "index_field":
			if err := dec.Decode(&s.IndexField); err != nil {
				return fmt.Errorf("%s | %w", "IndexField", err)
			}

		case "_meta":
			if err := dec.Decode(&s.Meta_); err != nil {
				return fmt.Errorf("%s | %w", "Meta_", err)
			}

		case "numeric_detection":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumericDetection", err)
				}
				s.NumericDetection = &value
			case bool:
				s.NumericDetection = &v
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

		case "_routing":
			if err := dec.Decode(&s.Routing_); err != nil {
				return fmt.Errorf("%s | %w", "Routing_", err)
			}

		case "runtime":
			if s.Runtime == nil {
				s.Runtime = make(map[string]RuntimeField, 0)
			}
			if err := dec.Decode(&s.Runtime); err != nil {
				return fmt.Errorf("%s | %w", "Runtime", err)
			}

		case "_size":
			if err := dec.Decode(&s.Size_); err != nil {
				return fmt.Errorf("%s | %w", "Size_", err)
			}

		case "_source":
			if err := dec.Decode(&s.Source_); err != nil {
				return fmt.Errorf("%s | %w", "Source_", err)
			}

		case "subobjects":
			if err := dec.Decode(&s.Subobjects); err != nil {
				return fmt.Errorf("%s | %w", "Subobjects", err)
			}

		}
	}
	return nil
}

// NewTypeMapping returns a TypeMapping.
func NewTypeMapping() *TypeMapping {
	r := &TypeMapping{
		Properties: make(map[string]Property),
		Runtime:    make(map[string]RuntimeField),
	}

	return r
}
