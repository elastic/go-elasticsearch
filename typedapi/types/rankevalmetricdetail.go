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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// RankEvalMetricDetail type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/rank_eval/types.ts#L125-L134
type RankEvalMetricDetail struct {
	// Hits The hits section shows a grouping of the search results with their supplied
	// ratings
	Hits []RankEvalHitItem `json:"hits"`
	// MetricDetails The metric_details give additional information about the calculated quality
	// metric (e.g. how many of the retrieved documents were relevant). The content
	// varies for each metric but allows for better interpretation of the results
	MetricDetails map[string]map[string]interface{} `json:"metric_details"`
	// MetricScore The metric_score in the details section shows the contribution of this query
	// to the global quality metric score
	MetricScore float64 `json:"metric_score"`
	// UnratedDocs The unrated_docs section contains an _index and _id entry for each document
	// in the search result for this query that didn’t have a ratings value. This
	// can be used to ask the user to supply ratings for these documents
	UnratedDocs []UnratedDocument `json:"unrated_docs"`
}

// RankEvalMetricDetailBuilder holds RankEvalMetricDetail struct and provides a builder API.
type RankEvalMetricDetailBuilder struct {
	v *RankEvalMetricDetail
}

// NewRankEvalMetricDetail provides a builder for the RankEvalMetricDetail struct.
func NewRankEvalMetricDetailBuilder() *RankEvalMetricDetailBuilder {
	r := RankEvalMetricDetailBuilder{
		&RankEvalMetricDetail{
			MetricDetails: make(map[string]map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the RankEvalMetricDetail struct
func (rb *RankEvalMetricDetailBuilder) Build() RankEvalMetricDetail {
	return *rb.v
}

// Hits The hits section shows a grouping of the search results with their supplied
// ratings

func (rb *RankEvalMetricDetailBuilder) Hits(hits []RankEvalHitItemBuilder) *RankEvalMetricDetailBuilder {
	tmp := make([]RankEvalHitItem, len(hits))
	for _, value := range hits {
		tmp = append(tmp, value.Build())
	}
	rb.v.Hits = tmp
	return rb
}

// MetricDetails The metric_details give additional information about the calculated quality
// metric (e.g. how many of the retrieved documents were relevant). The content
// varies for each metric but allows for better interpretation of the results

func (rb *RankEvalMetricDetailBuilder) MetricDetails(value map[string]map[string]interface{}) *RankEvalMetricDetailBuilder {
	rb.v.MetricDetails = value
	return rb
}

// MetricScore The metric_score in the details section shows the contribution of this query
// to the global quality metric score

func (rb *RankEvalMetricDetailBuilder) MetricScore(metricscore float64) *RankEvalMetricDetailBuilder {
	rb.v.MetricScore = metricscore
	return rb
}

// UnratedDocs The unrated_docs section contains an _index and _id entry for each document
// in the search result for this query that didn’t have a ratings value. This
// can be used to ask the user to supply ratings for these documents

func (rb *RankEvalMetricDetailBuilder) UnratedDocs(unrated_docs []UnratedDocumentBuilder) *RankEvalMetricDetailBuilder {
	tmp := make([]UnratedDocument, len(unrated_docs))
	for _, value := range unrated_docs {
		tmp = append(tmp, value.Build())
	}
	rb.v.UnratedDocs = tmp
	return rb
}
