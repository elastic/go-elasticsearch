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

// SettingsSimilarity type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexSettings.ts#L170-L178
type SettingsSimilarity struct {
	Bm25          *SettingsSimilarityBm25          `json:"bm25,omitempty"`
	Dfi           *SettingsSimilarityDfi           `json:"dfi,omitempty"`
	Dfr           *SettingsSimilarityDfr           `json:"dfr,omitempty"`
	Ib            *SettingsSimilarityIb            `json:"ib,omitempty"`
	Lmd           *SettingsSimilarityLmd           `json:"lmd,omitempty"`
	Lmj           *SettingsSimilarityLmj           `json:"lmj,omitempty"`
	ScriptedTfidf *SettingsSimilarityScriptedTfidf `json:"scripted_tfidf,omitempty"`
}

// SettingsSimilarityBuilder holds SettingsSimilarity struct and provides a builder API.
type SettingsSimilarityBuilder struct {
	v *SettingsSimilarity
}

// NewSettingsSimilarity provides a builder for the SettingsSimilarity struct.
func NewSettingsSimilarityBuilder() *SettingsSimilarityBuilder {
	r := SettingsSimilarityBuilder{
		&SettingsSimilarity{},
	}

	return &r
}

// Build finalize the chain and returns the SettingsSimilarity struct
func (rb *SettingsSimilarityBuilder) Build() SettingsSimilarity {
	return *rb.v
}

func (rb *SettingsSimilarityBuilder) Bm25(bm25 *SettingsSimilarityBm25Builder) *SettingsSimilarityBuilder {
	v := bm25.Build()
	rb.v.Bm25 = &v
	return rb
}

func (rb *SettingsSimilarityBuilder) Dfi(dfi *SettingsSimilarityDfiBuilder) *SettingsSimilarityBuilder {
	v := dfi.Build()
	rb.v.Dfi = &v
	return rb
}

func (rb *SettingsSimilarityBuilder) Dfr(dfr *SettingsSimilarityDfrBuilder) *SettingsSimilarityBuilder {
	v := dfr.Build()
	rb.v.Dfr = &v
	return rb
}

func (rb *SettingsSimilarityBuilder) Ib(ib *SettingsSimilarityIbBuilder) *SettingsSimilarityBuilder {
	v := ib.Build()
	rb.v.Ib = &v
	return rb
}

func (rb *SettingsSimilarityBuilder) Lmd(lmd *SettingsSimilarityLmdBuilder) *SettingsSimilarityBuilder {
	v := lmd.Build()
	rb.v.Lmd = &v
	return rb
}

func (rb *SettingsSimilarityBuilder) Lmj(lmj *SettingsSimilarityLmjBuilder) *SettingsSimilarityBuilder {
	v := lmj.Build()
	rb.v.Lmj = &v
	return rb
}

func (rb *SettingsSimilarityBuilder) ScriptedTfidf(scriptedtfidf *SettingsSimilarityScriptedTfidfBuilder) *SettingsSimilarityBuilder {
	v := scriptedtfidf.Build()
	rb.v.ScriptedTfidf = &v
	return rb
}
