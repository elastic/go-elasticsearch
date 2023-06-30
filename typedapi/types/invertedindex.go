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

// InvertedIndex type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/indices/field_usage_stats/IndicesFieldUsageStatsResponse.ts#L65-L73
type InvertedIndex struct {
	Offsets         uint `json:"offsets"`
	Payloads        uint `json:"payloads"`
	Positions       uint `json:"positions"`
	Postings        uint `json:"postings"`
	Proximity       uint `json:"proximity"`
	TermFrequencies uint `json:"term_frequencies"`
	Terms           uint `json:"terms"`
}

// NewInvertedIndex returns a InvertedIndex.
func NewInvertedIndex() *InvertedIndex {
	r := &InvertedIndex{}

	return r
}
