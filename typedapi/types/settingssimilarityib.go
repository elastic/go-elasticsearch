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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/ibdistribution"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/iblambda"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/normalization"
)

// SettingsSimilarityIb type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/indices/_types/IndexSettings.ts#L199-L204
type SettingsSimilarityIb struct {
	Distribution  ibdistribution.IBDistribution `json:"distribution"`
	Lambda        iblambda.IBLambda             `json:"lambda"`
	Normalization normalization.Normalization   `json:"normalization"`
	Type          string                        `json:"type,omitempty"`
}

// MarshalJSON override marshalling to include literal value
func (s SettingsSimilarityIb) MarshalJSON() ([]byte, error) {
	type innerSettingsSimilarityIb SettingsSimilarityIb
	tmp := innerSettingsSimilarityIb{
		Distribution:  s.Distribution,
		Lambda:        s.Lambda,
		Normalization: s.Normalization,
		Type:          s.Type,
	}

	tmp.Type = "IB"

	return json.Marshal(tmp)
}

// NewSettingsSimilarityIb returns a SettingsSimilarityIb.
func NewSettingsSimilarityIb() *SettingsSimilarityIb {
	r := &SettingsSimilarityIb{}

	return r
}
