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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termvectoroption"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeseriesmetrictype"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// DynamicProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/mapping/core.ts#L275-L306
type DynamicProperty struct {
	Analyzer            *string                        `json:"analyzer,omitempty"`
	Boost               *Float64                       `json:"boost,omitempty"`
	Coerce              *bool                          `json:"coerce,omitempty"`
	CopyTo              []string                       `json:"copy_to,omitempty"`
	DocValues           *bool                          `json:"doc_values,omitempty"`
	Dynamic             *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	EagerGlobalOrdinals *bool                          `json:"eager_global_ordinals,omitempty"`
	Enabled             *bool                          `json:"enabled,omitempty"`
	Fields              map[string]Property            `json:"fields,omitempty"`
	Format              *string                        `json:"format,omitempty"`
	IgnoreAbove         *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed     *bool                          `json:"ignore_malformed,omitempty"`
	Index               *bool                          `json:"index,omitempty"`
	IndexOptions        *indexoptions.IndexOptions     `json:"index_options,omitempty"`
	IndexPhrases        *bool                          `json:"index_phrases,omitempty"`
	IndexPrefixes       *TextIndexPrefixes             `json:"index_prefixes,omitempty"`
	Locale              *string                        `json:"locale,omitempty"`
	// Meta Metadata about the field.
	Meta                 map[string]string                          `json:"meta,omitempty"`
	Norms                *bool                                      `json:"norms,omitempty"`
	NullValue            FieldValue                                 `json:"null_value,omitempty"`
	OnScriptError        *onscripterror.OnScriptError               `json:"on_script_error,omitempty"`
	PositionIncrementGap *int                                       `json:"position_increment_gap,omitempty"`
	PrecisionStep        *int                                       `json:"precision_step,omitempty"`
	Properties           map[string]Property                        `json:"properties,omitempty"`
	Script               Script                                     `json:"script,omitempty"`
	SearchAnalyzer       *string                                    `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer  *string                                    `json:"search_quote_analyzer,omitempty"`
	Similarity           *string                                    `json:"similarity,omitempty"`
	Store                *bool                                      `json:"store,omitempty"`
	TermVector           *termvectoroption.TermVectorOption         `json:"term_vector,omitempty"`
	TimeSeriesMetric     *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
	Type                 string                                     `json:"type,omitempty"`
}

func (s *DynamicProperty) UnmarshalJSON(data []byte) error {
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
			if err := dec.Decode(&s.Analyzer); err != nil {
				return err
			}

		case "boost":
			if err := dec.Decode(&s.Boost); err != nil {
				return err
			}

		case "coerce":
			if err := dec.Decode(&s.Coerce); err != nil {
				return err
			}

		case "copy_to":
			if err := dec.Decode(&s.CopyTo); err != nil {
				return err
			}

		case "doc_values":
			if err := dec.Decode(&s.DocValues); err != nil {
				return err
			}

		case "dynamic":
			if err := dec.Decode(&s.Dynamic); err != nil {
				return err
			}

		case "eager_global_ordinals":
			if err := dec.Decode(&s.EagerGlobalOrdinals); err != nil {
				return err
			}

		case "enabled":
			if err := dec.Decode(&s.Enabled); err != nil {
				return err
			}

		case "fields":
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]interface{})
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)

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
					if err := dec.Decode(&s.Fields); err != nil {
						return err
					}
				}
			}

		case "format":
			if err := dec.Decode(&s.Format); err != nil {
				return err
			}

		case "ignore_above":
			if err := dec.Decode(&s.IgnoreAbove); err != nil {
				return err
			}

		case "ignore_malformed":
			if err := dec.Decode(&s.IgnoreMalformed); err != nil {
				return err
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "index_options":
			if err := dec.Decode(&s.IndexOptions); err != nil {
				return err
			}

		case "index_phrases":
			if err := dec.Decode(&s.IndexPhrases); err != nil {
				return err
			}

		case "index_prefixes":
			if err := dec.Decode(&s.IndexPrefixes); err != nil {
				return err
			}

		case "locale":
			if err := dec.Decode(&s.Locale); err != nil {
				return err
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "norms":
			if err := dec.Decode(&s.Norms); err != nil {
				return err
			}

		case "null_value":
			if err := dec.Decode(&s.NullValue); err != nil {
				return err
			}

		case "on_script_error":
			if err := dec.Decode(&s.OnScriptError); err != nil {
				return err
			}

		case "position_increment_gap":
			if err := dec.Decode(&s.PositionIncrementGap); err != nil {
				return err
			}

		case "precision_step":
			if err := dec.Decode(&s.PrecisionStep); err != nil {
				return err
			}

		case "properties":
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]interface{})
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)

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
					if err := dec.Decode(&s.Properties); err != nil {
						return err
					}
				}
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return err
			}

		case "search_analyzer":
			if err := dec.Decode(&s.SearchAnalyzer); err != nil {
				return err
			}

		case "search_quote_analyzer":
			if err := dec.Decode(&s.SearchQuoteAnalyzer); err != nil {
				return err
			}

		case "similarity":
			if err := dec.Decode(&s.Similarity); err != nil {
				return err
			}

		case "store":
			if err := dec.Decode(&s.Store); err != nil {
				return err
			}

		case "term_vector":
			if err := dec.Decode(&s.TermVector); err != nil {
				return err
			}

		case "time_series_metric":
			if err := dec.Decode(&s.TimeSeriesMetric); err != nil {
				return err
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDynamicProperty returns a DynamicProperty.
func NewDynamicProperty() *DynamicProperty {
	r := &DynamicProperty{
		Fields:     make(map[string]Property, 0),
		Meta:       make(map[string]string, 0),
		Properties: make(map[string]Property, 0),
	}

	r.Type = "{dynamic_property}"

	return r
}
