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

// ResponseBody type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/SearchResponse.ts#L38-L54
type ResponseBody struct {
	Aggregations    map[AggregateName]Aggregate  `json:"aggregations,omitempty"`
	Clusters_       *ClusterStatistics           `json:"_clusters,omitempty"`
	Fields          map[string]interface{}       `json:"fields,omitempty"`
	Hits            HitsMetadata                 `json:"hits"`
	MaxScore        *float64                     `json:"max_score,omitempty"`
	NumReducePhases *int64                       `json:"num_reduce_phases,omitempty"`
	PitId           *Id                          `json:"pit_id,omitempty"`
	Profile         *Profile                     `json:"profile,omitempty"`
	ScrollId_       *ScrollId                    `json:"_scroll_id,omitempty"`
	Shards_         ShardStatistics              `json:"_shards"`
	Suggest         map[SuggestionName][]Suggest `json:"suggest,omitempty"`
	TerminatedEarly *bool                        `json:"terminated_early,omitempty"`
	TimedOut        bool                         `json:"timed_out"`
	Took            int64                        `json:"took"`
}

// ResponseBodyBuilder holds ResponseBody struct and provides a builder API.
type ResponseBodyBuilder struct {
	v *ResponseBody
}

// NewResponseBody provides a builder for the ResponseBody struct.
func NewResponseBodyBuilder() *ResponseBodyBuilder {
	r := ResponseBodyBuilder{
		&ResponseBody{
			Aggregations: make(map[AggregateName]Aggregate, 0),
			Fields:       make(map[string]interface{}, 0),
			Suggest:      make(map[SuggestionName][]Suggest, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ResponseBody struct
func (rb *ResponseBodyBuilder) Build() ResponseBody {
	return *rb.v
}

func (rb *ResponseBodyBuilder) Aggregations(values map[AggregateName]*AggregateBuilder) *ResponseBodyBuilder {
	tmp := make(map[AggregateName]Aggregate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *ResponseBodyBuilder) Clusters_(clusters_ *ClusterStatisticsBuilder) *ResponseBodyBuilder {
	v := clusters_.Build()
	rb.v.Clusters_ = &v
	return rb
}

func (rb *ResponseBodyBuilder) Fields(value map[string]interface{}) *ResponseBodyBuilder {
	rb.v.Fields = value
	return rb
}

func (rb *ResponseBodyBuilder) Hits(hits *HitsMetadataBuilder) *ResponseBodyBuilder {
	v := hits.Build()
	rb.v.Hits = v
	return rb
}

func (rb *ResponseBodyBuilder) MaxScore(maxscore float64) *ResponseBodyBuilder {
	rb.v.MaxScore = &maxscore
	return rb
}

func (rb *ResponseBodyBuilder) NumReducePhases(numreducephases int64) *ResponseBodyBuilder {
	rb.v.NumReducePhases = &numreducephases
	return rb
}

func (rb *ResponseBodyBuilder) PitId(pitid Id) *ResponseBodyBuilder {
	rb.v.PitId = &pitid
	return rb
}

func (rb *ResponseBodyBuilder) Profile(profile *ProfileBuilder) *ResponseBodyBuilder {
	v := profile.Build()
	rb.v.Profile = &v
	return rb
}

func (rb *ResponseBodyBuilder) ScrollId_(scrollid_ ScrollId) *ResponseBodyBuilder {
	rb.v.ScrollId_ = &scrollid_
	return rb
}

func (rb *ResponseBodyBuilder) Shards_(shards_ *ShardStatisticsBuilder) *ResponseBodyBuilder {
	v := shards_.Build()
	rb.v.Shards_ = v
	return rb
}

func (rb *ResponseBodyBuilder) Suggest(value map[SuggestionName][]Suggest) *ResponseBodyBuilder {
	rb.v.Suggest = value
	return rb
}

func (rb *ResponseBodyBuilder) TerminatedEarly(terminatedearly bool) *ResponseBodyBuilder {
	rb.v.TerminatedEarly = &terminatedearly
	return rb
}

func (rb *ResponseBodyBuilder) TimedOut(timedout bool) *ResponseBodyBuilder {
	rb.v.TimedOut = timedout
	return rb
}

func (rb *ResponseBodyBuilder) Took(took int64) *ResponseBodyBuilder {
	rb.v.Took = took
	return rb
}
