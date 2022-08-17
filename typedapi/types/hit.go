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

// Hit type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/hits.ts#L40-L60
type Hit struct {
	Explanation_       *Explanation               `json:"_explanation,omitempty"`
	Fields             map[string]interface{}     `json:"fields,omitempty"`
	Highlight          map[string][]string        `json:"highlight,omitempty"`
	Id_                Id                         `json:"_id"`
	IgnoredFieldValues map[string][]string        `json:"ignored_field_values,omitempty"`
	Ignored_           []string                   `json:"_ignored,omitempty"`
	Index_             IndexName                  `json:"_index"`
	InnerHits          map[string]InnerHitsResult `json:"inner_hits,omitempty"`
	MatchedQueries     []string                   `json:"matched_queries,omitempty"`
	Nested_            *NestedIdentity            `json:"_nested,omitempty"`
	Node_              *string                    `json:"_node,omitempty"`
	PrimaryTerm_       *int64                     `json:"_primary_term,omitempty"`
	Routing_           *string                    `json:"_routing,omitempty"`
	Score_             float64                    `json:"_score,omitempty"`
	SeqNo_             *SequenceNumber            `json:"_seq_no,omitempty"`
	Shard_             *string                    `json:"_shard,omitempty"`
	Sort               *SortResults               `json:"sort,omitempty"`
	Source_            interface{}                `json:"_source,omitempty"`
	Version_           *VersionNumber             `json:"_version,omitempty"`
}

// HitBuilder holds Hit struct and provides a builder API.
type HitBuilder struct {
	v *Hit
}

// NewHit provides a builder for the Hit struct.
func NewHitBuilder() *HitBuilder {
	r := HitBuilder{
		&Hit{
			Fields:             make(map[string]interface{}, 0),
			Highlight:          make(map[string][]string, 0),
			IgnoredFieldValues: make(map[string][]string, 0),
			InnerHits:          make(map[string]InnerHitsResult, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Hit struct
func (rb *HitBuilder) Build() Hit {
	return *rb.v
}

func (rb *HitBuilder) Explanation_(explanation_ *ExplanationBuilder) *HitBuilder {
	v := explanation_.Build()
	rb.v.Explanation_ = &v
	return rb
}

func (rb *HitBuilder) Fields(value map[string]interface{}) *HitBuilder {
	rb.v.Fields = value
	return rb
}

func (rb *HitBuilder) Highlight(value map[string][]string) *HitBuilder {
	rb.v.Highlight = value
	return rb
}

func (rb *HitBuilder) Id_(id_ Id) *HitBuilder {
	rb.v.Id_ = id_
	return rb
}

func (rb *HitBuilder) IgnoredFieldValues(value map[string][]string) *HitBuilder {
	rb.v.IgnoredFieldValues = value
	return rb
}

func (rb *HitBuilder) Ignored_(_ignored ...string) *HitBuilder {
	rb.v.Ignored_ = _ignored
	return rb
}

func (rb *HitBuilder) Index_(index_ IndexName) *HitBuilder {
	rb.v.Index_ = index_
	return rb
}

func (rb *HitBuilder) InnerHits(values map[string]*InnerHitsResultBuilder) *HitBuilder {
	tmp := make(map[string]InnerHitsResult, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.InnerHits = tmp
	return rb
}

func (rb *HitBuilder) MatchedQueries(matched_queries ...string) *HitBuilder {
	rb.v.MatchedQueries = matched_queries
	return rb
}

func (rb *HitBuilder) Nested_(nested_ *NestedIdentityBuilder) *HitBuilder {
	v := nested_.Build()
	rb.v.Nested_ = &v
	return rb
}

func (rb *HitBuilder) Node_(node_ string) *HitBuilder {
	rb.v.Node_ = &node_
	return rb
}

func (rb *HitBuilder) PrimaryTerm_(primaryterm_ int64) *HitBuilder {
	rb.v.PrimaryTerm_ = &primaryterm_
	return rb
}

func (rb *HitBuilder) Routing_(routing_ string) *HitBuilder {
	rb.v.Routing_ = &routing_
	return rb
}

func (rb *HitBuilder) Score_(score_ float64) *HitBuilder {
	rb.v.Score_ = score_
	return rb
}

func (rb *HitBuilder) SeqNo_(seqno_ SequenceNumber) *HitBuilder {
	rb.v.SeqNo_ = &seqno_
	return rb
}

func (rb *HitBuilder) Shard_(shard_ string) *HitBuilder {
	rb.v.Shard_ = &shard_
	return rb
}

func (rb *HitBuilder) Sort(sort *SortResultsBuilder) *HitBuilder {
	v := sort.Build()
	rb.v.Sort = &v
	return rb
}

func (rb *HitBuilder) Source_(source_ interface{}) *HitBuilder {
	rb.v.Source_ = source_
	return rb
}

func (rb *HitBuilder) Version_(version_ VersionNumber) *HitBuilder {
	rb.v.Version_ = &version_
	return rb
}
