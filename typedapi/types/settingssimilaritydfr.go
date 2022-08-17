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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dfraftereffect"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dfrbasicmodel"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/normalization"
)

// SettingsSimilarityDfr type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexSettings.ts#L192-L197
type SettingsSimilarityDfr struct {
	AfterEffect   dfraftereffect.DFRAfterEffect `json:"after_effect"`
	BasicModel    dfrbasicmodel.DFRBasicModel   `json:"basic_model"`
	Normalization normalization.Normalization   `json:"normalization"`
	Type          string                        `json:"type,omitempty"`
}

// SettingsSimilarityDfrBuilder holds SettingsSimilarityDfr struct and provides a builder API.
type SettingsSimilarityDfrBuilder struct {
	v *SettingsSimilarityDfr
}

// NewSettingsSimilarityDfr provides a builder for the SettingsSimilarityDfr struct.
func NewSettingsSimilarityDfrBuilder() *SettingsSimilarityDfrBuilder {
	r := SettingsSimilarityDfrBuilder{
		&SettingsSimilarityDfr{},
	}

	r.v.Type = "DFR"

	return &r
}

// Build finalize the chain and returns the SettingsSimilarityDfr struct
func (rb *SettingsSimilarityDfrBuilder) Build() SettingsSimilarityDfr {
	return *rb.v
}

func (rb *SettingsSimilarityDfrBuilder) AfterEffect(aftereffect dfraftereffect.DFRAfterEffect) *SettingsSimilarityDfrBuilder {
	rb.v.AfterEffect = aftereffect
	return rb
}

func (rb *SettingsSimilarityDfrBuilder) BasicModel(basicmodel dfrbasicmodel.DFRBasicModel) *SettingsSimilarityDfrBuilder {
	rb.v.BasicModel = basicmodel
	return rb
}

func (rb *SettingsSimilarityDfrBuilder) Normalization(normalization normalization.Normalization) *SettingsSimilarityDfrBuilder {
	rb.v.Normalization = normalization
	return rb
}
