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
	"io"
	"strconv"
)

// AnalyticsStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/xpack/usage/types.ts#L61-L71
type AnalyticsStatistics struct {
	BoxplotUsage               int64  `json:"boxplot_usage"`
	CumulativeCardinalityUsage int64  `json:"cumulative_cardinality_usage"`
	MovingPercentilesUsage     int64  `json:"moving_percentiles_usage"`
	MultiTermsUsage            *int64 `json:"multi_terms_usage,omitempty"`
	NormalizeUsage             int64  `json:"normalize_usage"`
	RateUsage                  int64  `json:"rate_usage"`
	StringStatsUsage           int64  `json:"string_stats_usage"`
	TTestUsage                 int64  `json:"t_test_usage"`
	TopMetricsUsage            int64  `json:"top_metrics_usage"`
}

func (s *AnalyticsStatistics) UnmarshalJSON(data []byte) error {

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

		case "boxplot_usage":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.BoxplotUsage = value
			case float64:
				f := int64(v)
				s.BoxplotUsage = f
			}

		case "cumulative_cardinality_usage":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.CumulativeCardinalityUsage = value
			case float64:
				f := int64(v)
				s.CumulativeCardinalityUsage = f
			}

		case "moving_percentiles_usage":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MovingPercentilesUsage = value
			case float64:
				f := int64(v)
				s.MovingPercentilesUsage = f
			}

		case "multi_terms_usage":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MultiTermsUsage = &value
			case float64:
				f := int64(v)
				s.MultiTermsUsage = &f
			}

		case "normalize_usage":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.NormalizeUsage = value
			case float64:
				f := int64(v)
				s.NormalizeUsage = f
			}

		case "rate_usage":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.RateUsage = value
			case float64:
				f := int64(v)
				s.RateUsage = f
			}

		case "string_stats_usage":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.StringStatsUsage = value
			case float64:
				f := int64(v)
				s.StringStatsUsage = f
			}

		case "t_test_usage":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TTestUsage = value
			case float64:
				f := int64(v)
				s.TTestUsage = f
			}

		case "top_metrics_usage":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TopMetricsUsage = value
			case float64:
				f := int64(v)
				s.TopMetricsUsage = f
			}

		}
	}
	return nil
}

// NewAnalyticsStatistics returns a AnalyticsStatistics.
func NewAnalyticsStatistics() *AnalyticsStatistics {
	r := &AnalyticsStatistics{}

	return r
}
