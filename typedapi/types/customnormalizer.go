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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"encoding/json"
)

// CustomNormalizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/analysis/normalizers.ts#L30-L34
type CustomNormalizer struct {
	CharFilter []string `json:"char_filter,omitempty"`
	Filter     []string `json:"filter,omitempty"`
	Type       string   `json:"type,omitempty"`
}

// MarshalJSON override marshalling to include literal value
func (s CustomNormalizer) MarshalJSON() ([]byte, error) {
	type innerCustomNormalizer CustomNormalizer
	tmp := innerCustomNormalizer{
		CharFilter: s.CharFilter,
		Filter:     s.Filter,
		Type:       s.Type,
	}

	tmp.Type = "custom"

	return json.Marshal(tmp)
}

// NewCustomNormalizer returns a CustomNormalizer.
func NewCustomNormalizer() *CustomNormalizer {
	r := &CustomNormalizer{}

	return r
}

type CustomNormalizerVariant interface {
	CustomNormalizerCaster() *CustomNormalizer
}

func (s *CustomNormalizer) CustomNormalizerCaster() *CustomNormalizer {
	return s
}

func (s *CustomNormalizer) NormalizerCaster() *Normalizer {
	o := Normalizer(s)
	return &o
}
