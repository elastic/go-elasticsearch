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
	"strings"
)

// AsyncSearch type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/async_search/_types/AsyncSearch.ts#L30-L56
type AsyncSearch struct {
	// Aggregations Partial aggregations results, coming from the shards that have already
	// completed running the query.
	Aggregations map[string]Aggregate       `json:"aggregations,omitempty"`
	Clusters_    *ClusterStatistics         `json:"_clusters,omitempty"`
	Fields       map[string]json.RawMessage `json:"fields,omitempty"`
	Hits         HitsMetadata               `json:"hits"`
	MaxScore     *Float64                   `json:"max_score,omitempty"`
	// NumReducePhases Indicates how many reductions of the results have been performed.
	// If this number increases compared to the last retrieved results for a get
	// asynch search request, you can expect additional results included in the
	// search response.
	NumReducePhases *int64   `json:"num_reduce_phases,omitempty"`
	PitId           *string  `json:"pit_id,omitempty"`
	Profile         *Profile `json:"profile,omitempty"`
	ScrollId_       *string  `json:"_scroll_id,omitempty"`
	// Shards_ Indicates how many shards have run the query.
	// Note that in order for shard results to be included in the search response,
	// they need to be reduced first.
	Shards_         ShardStatistics      `json:"_shards"`
	Suggest         map[string][]Suggest `json:"suggest,omitempty"`
	TerminatedEarly *bool                `json:"terminated_early,omitempty"`
	TimedOut        bool                 `json:"timed_out"`
	Took            int64                `json:"took"`
}

func (s *AsyncSearch) UnmarshalJSON(data []byte) error {

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
				s.Aggregations = make(map[string]Aggregate, 0)
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
								s.Aggregations = make(map[string]Aggregate, 0)
							}
							switch elems[0] {

							case "cardinality":
								o := NewCardinalityAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "hdr_percentiles":
								o := NewHdrPercentilesAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "hdr_percentile_ranks":
								o := NewHdrPercentileRanksAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "tdigest_percentiles":
								o := NewTDigestPercentilesAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "tdigest_percentile_ranks":
								o := NewTDigestPercentileRanksAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "percentiles_bucket":
								o := NewPercentilesBucketAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "median_absolute_deviation":
								o := NewMedianAbsoluteDeviationAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "min":
								o := NewMinAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "max":
								o := NewMaxAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "sum":
								o := NewSumAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "avg":
								o := NewAvgAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "weighted_avg":
								o := NewWeightedAvgAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "value_count":
								o := NewValueCountAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "simple_value":
								o := NewSimpleValueAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "derivative":
								o := NewDerivativeAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "bucket_metric_value":
								o := NewBucketMetricValueAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "stats":
								o := NewStatsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "stats_bucket":
								o := NewStatsBucketAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "extended_stats":
								o := NewExtendedStatsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "extended_stats_bucket":
								o := NewExtendedStatsBucketAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "geo_bounds":
								o := NewGeoBoundsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "geo_centroid":
								o := NewGeoCentroidAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "histogram":
								o := NewHistogramAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "date_histogram":
								o := NewDateHistogramAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "auto_date_histogram":
								o := NewAutoDateHistogramAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "variable_width_histogram":
								o := NewVariableWidthHistogramAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "sterms":
								o := NewStringTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "lterms":
								o := NewLongTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "dterms":
								o := NewDoubleTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "umterms":
								o := NewUnmappedTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "lrareterms":
								o := NewLongRareTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "srareterms":
								o := NewStringRareTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "umrareterms":
								o := NewUnmappedRareTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "multi_terms":
								o := NewMultiTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "missing":
								o := NewMissingAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "nested":
								o := NewNestedAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "reverse_nested":
								o := NewReverseNestedAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "global":
								o := NewGlobalAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "filter":
								o := NewFilterAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "children":
								o := NewChildrenAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "parent":
								o := NewParentAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "sampler":
								o := NewSamplerAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "unmapped_sampler":
								o := NewUnmappedSamplerAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "geohash_grid":
								o := NewGeoHashGridAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "geotile_grid":
								o := NewGeoTileGridAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "geohex_grid":
								o := NewGeoHexGridAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "range":
								o := NewRangeAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "date_range":
								o := NewDateRangeAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "geo_distance":
								o := NewGeoDistanceAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "ip_range":
								o := NewIpRangeAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "ip_prefix":
								o := NewIpPrefixAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "filters":
								o := NewFiltersAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "adjacency_matrix":
								o := NewAdjacencyMatrixAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "siglterms":
								o := NewSignificantLongTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "sigsterms":
								o := NewSignificantStringTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "umsigterms":
								o := NewUnmappedSignificantTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "composite":
								o := NewCompositeAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "frequent_item_sets":
								o := NewFrequentItemSetsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "time_series":
								o := NewTimeSeriesAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "scripted_metric":
								o := NewScriptedMetricAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "top_hits":
								o := NewTopHitsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "inference":
								o := NewInferenceAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "string_stats":
								o := NewStringStatsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "boxplot":
								o := NewBoxPlotAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "top_metrics":
								o := NewTopMetricsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "t_test":
								o := NewTTestAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "rate":
								o := NewRateAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "simple_long_value":
								o := NewCumulativeCardinalityAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "matrix_stats":
								o := NewMatrixStatsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "geo_line":
								o := NewGeoLineAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							default:
								o := make(map[string]any, 0)
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o
							}
						} else {
							return errors.New("cannot decode JSON for field Aggregations")
						}
					} else {
						o := make(map[string]any, 0)
						if err := dec.Decode(&o); err != nil {
							return fmt.Errorf("%s | %w", "Aggregations", err)
						}
						s.Aggregations[value] = o
					}
				}
			}

		case "_clusters":
			if err := dec.Decode(&s.Clusters_); err != nil {
				return fmt.Errorf("%s | %w", "Clusters_", err)
			}

		case "fields":
			if s.Fields == nil {
				s.Fields = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Fields); err != nil {
				return fmt.Errorf("%s | %w", "Fields", err)
			}

		case "hits":
			if err := dec.Decode(&s.Hits); err != nil {
				return fmt.Errorf("%s | %w", "Hits", err)
			}

		case "max_score":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxScore", err)
				}
				f := Float64(value)
				s.MaxScore = &f
			case float64:
				f := Float64(v)
				s.MaxScore = &f
			}

		case "num_reduce_phases":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumReducePhases", err)
				}
				s.NumReducePhases = &value
			case float64:
				f := int64(v)
				s.NumReducePhases = &f
			}

		case "pit_id":
			if err := dec.Decode(&s.PitId); err != nil {
				return fmt.Errorf("%s | %w", "PitId", err)
			}

		case "profile":
			if err := dec.Decode(&s.Profile); err != nil {
				return fmt.Errorf("%s | %w", "Profile", err)
			}

		case "_scroll_id":
			if err := dec.Decode(&s.ScrollId_); err != nil {
				return fmt.Errorf("%s | %w", "ScrollId_", err)
			}

		case "_shards":
			if err := dec.Decode(&s.Shards_); err != nil {
				return fmt.Errorf("%s | %w", "Shards_", err)
			}

		case "suggest":
			if s.Suggest == nil {
				s.Suggest = make(map[string][]Suggest, 0)
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
							if s.Suggest == nil {
								s.Suggest = make(map[string][]Suggest, 0)
							}
							switch elems[0] {

							case "completion":
								o := NewCompletionSuggest()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Suggest", err)
								}
								s.Suggest[elems[1]] = append(s.Suggest[elems[1]], o)

							case "phrase":
								o := NewPhraseSuggest()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Suggest", err)
								}
								s.Suggest[elems[1]] = append(s.Suggest[elems[1]], o)

							case "term":
								o := NewTermSuggest()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Suggest", err)
								}
								s.Suggest[elems[1]] = append(s.Suggest[elems[1]], o)

							default:
								o := make(map[string]any, 0)
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Suggest", err)
								}
								s.Suggest[elems[1]] = append(s.Suggest[elems[1]], o)
							}
						} else {
							return errors.New("cannot decode JSON for field Suggest")
						}
					} else {
						o := make(map[string]any, 0)
						if err := dec.Decode(&o); err != nil {
							return fmt.Errorf("%s | %w", "Suggest", err)
						}
						s.Suggest[value] = append(s.Suggest[value], o)
					}
				}
			}

		case "terminated_early":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TerminatedEarly", err)
				}
				s.TerminatedEarly = &value
			case bool:
				s.TerminatedEarly = &v
			}

		case "timed_out":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TimedOut", err)
				}
				s.TimedOut = value
			case bool:
				s.TimedOut = v
			}

		case "took":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Took", err)
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

// NewAsyncSearch returns a AsyncSearch.
func NewAsyncSearch() *AsyncSearch {
	r := &AsyncSearch{
		Aggregations: make(map[string]Aggregate),
		Fields:       make(map[string]json.RawMessage),
		Suggest:      make(map[string][]Suggest),
	}

	return r
}
