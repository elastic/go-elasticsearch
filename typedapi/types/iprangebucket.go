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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// IpRangeBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/aggregations/Aggregate.ts#L560-L564
type IpRangeBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	From         *string              `json:"from,omitempty"`
	Key          *string              `json:"key,omitempty"`
	To           *string              `json:"to,omitempty"`
}

func (s *IpRangeBucket) UnmarshalJSON(data []byte) error {

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

		case "doc_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DocCount = value
			case float64:
				f := int64(v)
				s.DocCount = f
			}

		case "from":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.From = &o

		case "key":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Key = &o

		case "to":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.To = &o

		default:

			if value, ok := t.(string); ok {
				if strings.Contains(value, "#") {
					elems := strings.Split(value, "#")
					if len(elems) == 2 {
						if s.Aggregations == nil {
							s.Aggregations = make(map[string]Aggregate, 0)
						}
						switch elems[0] {

						case "cardinality":
							o := NewCardinalityAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "hdr_percentiles":
							o := NewHdrPercentilesAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "hdr_percentile_ranks":
							o := NewHdrPercentileRanksAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "tdigest_percentiles":
							o := NewTDigestPercentilesAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "tdigest_percentile_ranks":
							o := NewTDigestPercentileRanksAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "percentiles_bucket":
							o := NewPercentilesBucketAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "median_absolute_deviation":
							o := NewMedianAbsoluteDeviationAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "min":
							o := NewMinAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "max":
							o := NewMaxAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "sum":
							o := NewSumAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "avg":
							o := NewAvgAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "weighted_avg":
							o := NewWeightedAvgAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "value_count":
							o := NewValueCountAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "simple_value":
							o := NewSimpleValueAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "derivative":
							o := NewDerivativeAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "bucket_metric_value":
							o := NewBucketMetricValueAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "stats":
							o := NewStatsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "stats_bucket":
							o := NewStatsBucketAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "extended_stats":
							o := NewExtendedStatsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "extended_stats_bucket":
							o := NewExtendedStatsBucketAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "geo_bounds":
							o := NewGeoBoundsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "geo_centroid":
							o := NewGeoCentroidAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "histogram":
							o := NewHistogramAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "date_histogram":
							o := NewDateHistogramAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "auto_date_histogram":
							o := NewAutoDateHistogramAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "variable_width_histogram":
							o := NewVariableWidthHistogramAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "sterms":
							o := NewStringTermsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "lterms":
							o := NewLongTermsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "dterms":
							o := NewDoubleTermsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "umterms":
							o := NewUnmappedTermsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "lrareterms":
							o := NewLongRareTermsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "srareterms":
							o := NewStringRareTermsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "umrareterms":
							o := NewUnmappedRareTermsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "multi_terms":
							o := NewMultiTermsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "missing":
							o := NewMissingAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "nested":
							o := NewNestedAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "reverse_nested":
							o := NewReverseNestedAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "global":
							o := NewGlobalAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "filter":
							o := NewFilterAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "children":
							o := NewChildrenAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "parent":
							o := NewParentAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "sampler":
							o := NewSamplerAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "unmapped_sampler":
							o := NewUnmappedSamplerAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "geohash_grid":
							o := NewGeoHashGridAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "geotile_grid":
							o := NewGeoTileGridAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "geohex_grid":
							o := NewGeoHexGridAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "range":
							o := NewRangeAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "date_range":
							o := NewDateRangeAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "geo_distance":
							o := NewGeoDistanceAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "ip_range":
							o := NewIpRangeAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "ip_prefix":
							o := NewIpPrefixAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "filters":
							o := NewFiltersAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "adjacency_matrix":
							o := NewAdjacencyMatrixAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "siglterms":
							o := NewSignificantLongTermsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "sigsterms":
							o := NewSignificantStringTermsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "umsigterms":
							o := NewUnmappedSignificantTermsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "composite":
							o := NewCompositeAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "frequent_item_sets":
							o := NewFrequentItemSetsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "scripted_metric":
							o := NewScriptedMetricAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "top_hits":
							o := NewTopHitsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "inference":
							o := NewInferenceAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "string_stats":
							o := NewStringStatsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "box_plot":
							o := NewBoxPlotAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "top_metrics":
							o := NewTopMetricsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "t_test":
							o := NewTTestAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "rate":
							o := NewRateAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "simple_long_value":
							o := NewCumulativeCardinalityAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "matrix_stats":
							o := NewMatrixStatsAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						case "geo_line":
							o := NewGeoLineAggregate()
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o

						default:
							o := make(map[string]interface{}, 0)
							if err := dec.Decode(&o); err != nil {
								return err
							}
							s.Aggregations[elems[1]] = o
						}
					} else {
						return errors.New("cannot decode JSON for field Aggregations")
					}
				} else {
					o := make(map[string]interface{}, 0)
					if err := dec.Decode(&o); err != nil {
						return err
					}
					s.Aggregations[value] = o
				}
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s IpRangeBucket) MarshalJSON() ([]byte, error) {
	type opt IpRangeBucket
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.Aggregations {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "Aggregations")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewIpRangeBucket returns a IpRangeBucket.
func NewIpRangeBucket() *IpRangeBucket {
	r := &IpRangeBucket{
		Aggregations: make(map[string]Aggregate, 0),
	}

	return r
}
