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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// AnalyticsStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/xpack/usage/types.ts#L57-L67
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

// NewAnalyticsStatistics returns a AnalyticsStatistics.
func NewAnalyticsStatistics() *AnalyticsStatistics {
	r := &AnalyticsStatistics{}

	return r
}
