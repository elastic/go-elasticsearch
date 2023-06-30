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

// EqlFeaturesSequences type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/xpack/usage/types.ts#L130-L137
type EqlFeaturesSequences struct {
	SequenceMaxspan           uint `json:"sequence_maxspan"`
	SequenceQueriesFiveOrMore uint `json:"sequence_queries_five_or_more"`
	SequenceQueriesFour       uint `json:"sequence_queries_four"`
	SequenceQueriesThree      uint `json:"sequence_queries_three"`
	SequenceQueriesTwo        uint `json:"sequence_queries_two"`
	SequenceUntil             uint `json:"sequence_until"`
}

// NewEqlFeaturesSequences returns a EqlFeaturesSequences.
func NewEqlFeaturesSequences() *EqlFeaturesSequences {
	r := &EqlFeaturesSequences{}

	return r
}
