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
// https://github.com/elastic/elasticsearch-specification/tree/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33


package types

// Hit type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/_global/search/_types/hits.ts#L40-L64
type Hit struct {
	Explanation_       *Explanation               `json:"_explanation,omitempty"`
	Fields             map[string]interface{}     `json:"fields,omitempty"`
	Highlight          map[string][]string        `json:"highlight,omitempty"`
	Id_                string                     `json:"_id"`
	IgnoredFieldValues map[string][]string        `json:"ignored_field_values,omitempty"`
	Ignored_           []string                   `json:"_ignored,omitempty"`
	Index_             string                     `json:"_index"`
	InnerHits          map[string]InnerHitsResult `json:"inner_hits,omitempty"`
	MatchedQueries     []string                   `json:"matched_queries,omitempty"`
	Nested_            *NestedIdentity            `json:"_nested,omitempty"`
	Node_              *string                    `json:"_node,omitempty"`
	PrimaryTerm_       *int64                     `json:"_primary_term,omitempty"`
	Routing_           *string                    `json:"_routing,omitempty"`
	Score_             float64                    `json:"_score,omitempty"`
	SeqNo_             *int64                     `json:"_seq_no,omitempty"`
	Shard_             *string                    `json:"_shard,omitempty"`
	Sort               []FieldValue               `json:"sort,omitempty"`
	Source_            interface{}                `json:"_source,omitempty"`
	Version_           *int64                     `json:"_version,omitempty"`
}

// NewHit returns a Hit.
func NewHit() *Hit {
	r := &Hit{
		Fields:             make(map[string]interface{}, 0),
		Highlight:          make(map[string][]string, 0),
		IgnoredFieldValues: make(map[string][]string, 0),
		InnerHits:          make(map[string]InnerHitsResult, 0),
	}

	return r
}
