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

// SettingsSimilarityLmd type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexSettings.ts#L206-L209
type SettingsSimilarityLmd struct {
	Mu   int    `json:"mu"`
	Type string `json:"type,omitempty"`
}

// SettingsSimilarityLmdBuilder holds SettingsSimilarityLmd struct and provides a builder API.
type SettingsSimilarityLmdBuilder struct {
	v *SettingsSimilarityLmd
}

// NewSettingsSimilarityLmd provides a builder for the SettingsSimilarityLmd struct.
func NewSettingsSimilarityLmdBuilder() *SettingsSimilarityLmdBuilder {
	r := SettingsSimilarityLmdBuilder{
		&SettingsSimilarityLmd{},
	}

	r.v.Type = "LMDirichlet"

	return &r
}

// Build finalize the chain and returns the SettingsSimilarityLmd struct
func (rb *SettingsSimilarityLmdBuilder) Build() SettingsSimilarityLmd {
	return *rb.v
}

func (rb *SettingsSimilarityLmdBuilder) Mu(mu int) *SettingsSimilarityLmdBuilder {
	rb.v.Mu = mu
	return rb
}
