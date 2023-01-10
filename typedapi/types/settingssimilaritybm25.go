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

// SettingsSimilarityBm25 type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/indices/_types/IndexSettings.ts#L180-L185
type SettingsSimilarityBm25 struct {
	B                float64 `json:"b"`
	DiscountOverlaps bool    `json:"discount_overlaps"`
	K1               float64 `json:"k1"`
	Type             string  `json:"type,omitempty"`
}

// NewSettingsSimilarityBm25 returns a SettingsSimilarityBm25.
func NewSettingsSimilarityBm25() *SettingsSimilarityBm25 {
	r := &SettingsSimilarityBm25{}

	r.Type = "BM25"

	return r
}
