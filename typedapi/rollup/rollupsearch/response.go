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

package rollupsearch

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package rollupsearch
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/rollup/rollup_search/RollupSearchResponse.ts#L27-L36

type Response struct {
	Aggregations    map[string]types.Aggregate `json:"aggregations,omitempty"`
	Hits            types.HitsMetadata         `json:"hits"`
	Shards_         types.ShardStatistics      `json:"_shards"`
	TerminatedEarly *bool                      `json:"terminated_early,omitempty"`
	TimedOut        bool                       `json:"timed_out"`
	Took            int64                      `json:"took"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Aggregations: make(map[string]types.Aggregate, 0),
	}
	return r
}

func (s *Response) UnmarshalJSON(data []byte) error {
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

		case "aggregations":
			if s.Aggregations == nil {
				s.Aggregations = make(map[string]types.Aggregate, 0)
			}

			for dec.More() {
				tt, err := dec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return err
				}
				if value, ok := tt.(string); ok {
					if strings.Contains(value, "#") {
						elems := strings.Split(value, "#")
						if len(elems) == 2 {
							if s.Aggregations == nil {
								s.Aggregations = make(map[string]types.Aggregate, 0)
							}
							switch elems[0] {

							case "cardinality":
								o := types.NewCardinalityAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "hdr_percentiles":
								o := types.NewHdrPercentilesAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "hdr_percentile_ranks":
								o := types.NewHdrPercentileRanksAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "tdigest_percentiles":
								o := types.NewTDigestPercentilesAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "tdigest_percentile_ranks":
								o := types.NewTDigestPercentileRanksAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "percentiles_bucket":
								o := types.NewPercentilesBucketAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "median_absolute_deviation":
								o := types.NewMedianAbsoluteDeviationAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "min":
								o := types.NewMinAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "max":
								o := types.NewMaxAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "sum":
								o := types.NewSumAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "avg":
								o := types.NewAvgAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "weighted_avg":
								o := types.NewWeightedAvgAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "value_count":
								o := types.NewValueCountAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "simple_value":
								o := types.NewSimpleValueAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "derivative":
								o := types.NewDerivativeAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "bucket_metric_value":
								o := types.NewBucketMetricValueAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "stats":
								o := types.NewStatsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "stats_bucket":
								o := types.NewStatsBucketAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "extended_stats":
								o := types.NewExtendedStatsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "extended_stats_bucket":
								o := types.NewExtendedStatsBucketAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "geo_bounds":
								o := types.NewGeoBoundsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "geo_centroid":
								o := types.NewGeoCentroidAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "histogram":
								o := types.NewHistogramAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "date_histogram":
								o := types.NewDateHistogramAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "auto_date_histogram":
								o := types.NewAutoDateHistogramAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "variable_width_histogram":
								o := types.NewVariableWidthHistogramAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "sterms":
								o := types.NewStringTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "lterms":
								o := types.NewLongTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "dterms":
								o := types.NewDoubleTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "umterms":
								o := types.NewUnmappedTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "lrareterms":
								o := types.NewLongRareTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "srareterms":
								o := types.NewStringRareTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "umrareterms":
								o := types.NewUnmappedRareTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "multi_terms":
								o := types.NewMultiTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "missing":
								o := types.NewMissingAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "nested":
								o := types.NewNestedAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "reverse_nested":
								o := types.NewReverseNestedAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "global":
								o := types.NewGlobalAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "filter":
								o := types.NewFilterAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "children":
								o := types.NewChildrenAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "parent":
								o := types.NewParentAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "sampler":
								o := types.NewSamplerAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "unmapped_sampler":
								o := types.NewUnmappedSamplerAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "geohash_grid":
								o := types.NewGeoHashGridAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "geotile_grid":
								o := types.NewGeoTileGridAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "geohex_grid":
								o := types.NewGeoHexGridAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "range":
								o := types.NewRangeAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "date_range":
								o := types.NewDateRangeAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "geo_distance":
								o := types.NewGeoDistanceAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "ip_range":
								o := types.NewIpRangeAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "ip_prefix":
								o := types.NewIpPrefixAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "filters":
								o := types.NewFiltersAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "adjacency_matrix":
								o := types.NewAdjacencyMatrixAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "siglterms":
								o := types.NewSignificantLongTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "sigsterms":
								o := types.NewSignificantStringTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "umsigterms":
								o := types.NewUnmappedSignificantTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "composite":
								o := types.NewCompositeAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "frequent_item_sets":
								o := types.NewFrequentItemSetsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "scripted_metric":
								o := types.NewScriptedMetricAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "top_hits":
								o := types.NewTopHitsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "inference":
								o := types.NewInferenceAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "string_stats":
								o := types.NewStringStatsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "box_plot":
								o := types.NewBoxPlotAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "top_metrics":
								o := types.NewTopMetricsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "t_test":
								o := types.NewTTestAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "rate":
								o := types.NewRateAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "simple_long_value":
								o := types.NewCumulativeCardinalityAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "matrix_stats":
								o := types.NewMatrixStatsAggregate()
								if err := dec.Decode(&o); err != nil {
									return err
								}
								s.Aggregations[elems[1]] = o

							case "geo_line":
								o := types.NewGeoLineAggregate()
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

		case "hits":
			if err := dec.Decode(&s.Hits); err != nil {
				return err
			}

		case "_shards":
			if err := dec.Decode(&s.Shards_); err != nil {
				return err
			}

		case "terminated_early":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.TerminatedEarly = &value
			case bool:
				s.TerminatedEarly = &v
			}

		case "timed_out":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.TimedOut = value
			case bool:
				s.TimedOut = v
			}

		case "took":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Took = value
			case float64:
				f := int64(v)
				s.Took = f
			}

		}
	}
	return nil
}
