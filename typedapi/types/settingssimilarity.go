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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// SettingsSimilarity type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/indices/_types/IndexSettings.ts#L170-L178
type SettingsSimilarity struct {
	Bm25          *SettingsSimilarityBm25          `json:"bm25,omitempty"`
	Dfi           *SettingsSimilarityDfi           `json:"dfi,omitempty"`
	Dfr           *SettingsSimilarityDfr           `json:"dfr,omitempty"`
	Ib            *SettingsSimilarityIb            `json:"ib,omitempty"`
	Lmd           *SettingsSimilarityLmd           `json:"lmd,omitempty"`
	Lmj           *SettingsSimilarityLmj           `json:"lmj,omitempty"`
	ScriptedTfidf *SettingsSimilarityScriptedTfidf `json:"scripted_tfidf,omitempty"`
}

// NewSettingsSimilarity returns a SettingsSimilarity.
func NewSettingsSimilarity() *SettingsSimilarity {
	r := &SettingsSimilarity{}

	return r
}
