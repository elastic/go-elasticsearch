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
// https://github.com/elastic/elasticsearch-specification/tree/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// SettingsSimilarityBm25 type.
//
// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/indices/_types/IndexSettings.ts#L179-L184
type SettingsSimilarityBm25 struct {
	B                Float64 `json:"b"`
	DiscountOverlaps bool    `json:"discount_overlaps"`
	K1               Float64 `json:"k1"`
	Type             string  `json:"type,omitempty"`
}

func (s *SettingsSimilarityBm25) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "b":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.B = f
			case float64:
				f := Float64(v)
				s.B = f
			}

		case "discount_overlaps":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.DiscountOverlaps = value
			case bool:
				s.DiscountOverlaps = v
			}

		case "k1":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.K1 = f
			case float64:
				f := Float64(v)
				s.K1 = f
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s SettingsSimilarityBm25) MarshalJSON() ([]byte, error) {
	type innerSettingsSimilarityBm25 SettingsSimilarityBm25
	tmp := innerSettingsSimilarityBm25{
		B:                s.B,
		DiscountOverlaps: s.DiscountOverlaps,
		K1:               s.K1,
		Type:             s.Type,
	}

	tmp.Type = "BM25"

	return json.Marshal(tmp)
}

// NewSettingsSimilarityBm25 returns a SettingsSimilarityBm25.
func NewSettingsSimilarityBm25() *SettingsSimilarityBm25 {
	r := &SettingsSimilarityBm25{}

	return r
}
