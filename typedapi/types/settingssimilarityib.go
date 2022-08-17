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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/ibdistribution"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/iblambda"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/normalization"
)

// SettingsSimilarityIb type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexSettings.ts#L199-L204
type SettingsSimilarityIb struct {
	Distribution  ibdistribution.IBDistribution `json:"distribution"`
	Lambda        iblambda.IBLambda             `json:"lambda"`
	Normalization normalization.Normalization   `json:"normalization"`
	Type          string                        `json:"type,omitempty"`
}

// SettingsSimilarityIbBuilder holds SettingsSimilarityIb struct and provides a builder API.
type SettingsSimilarityIbBuilder struct {
	v *SettingsSimilarityIb
}

// NewSettingsSimilarityIb provides a builder for the SettingsSimilarityIb struct.
func NewSettingsSimilarityIbBuilder() *SettingsSimilarityIbBuilder {
	r := SettingsSimilarityIbBuilder{
		&SettingsSimilarityIb{},
	}

	r.v.Type = "IB"

	return &r
}

// Build finalize the chain and returns the SettingsSimilarityIb struct
func (rb *SettingsSimilarityIbBuilder) Build() SettingsSimilarityIb {
	return *rb.v
}

func (rb *SettingsSimilarityIbBuilder) Distribution(distribution ibdistribution.IBDistribution) *SettingsSimilarityIbBuilder {
	rb.v.Distribution = distribution
	return rb
}

func (rb *SettingsSimilarityIbBuilder) Lambda(lambda iblambda.IBLambda) *SettingsSimilarityIbBuilder {
	rb.v.Lambda = lambda
	return rb
}

func (rb *SettingsSimilarityIbBuilder) Normalization(normalization normalization.Normalization) *SettingsSimilarityIbBuilder {
	rb.v.Normalization = normalization
	return rb
}
