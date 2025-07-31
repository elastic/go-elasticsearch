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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// SettingsSimilarityBm25 type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/indices/_types/IndexSettings.ts#L196-L201
type SettingsSimilarityBm25 struct {
	B                *Float64 `json:"b,omitempty"`
	DiscountOverlaps *bool    `json:"discount_overlaps,omitempty"`
	K1               *Float64 `json:"k1,omitempty"`
	Type             string   `json:"type,omitempty"`
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
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "B", err)
				}
				f := Float64(value)
				s.B = &f
			case float64:
				f := Float64(v)
				s.B = &f
			}

		case "discount_overlaps":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DiscountOverlaps", err)
				}
				s.DiscountOverlaps = &value
			case bool:
				s.DiscountOverlaps = &v
			}

		case "k1":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "K1", err)
				}
				f := Float64(value)
				s.K1 = &f
			case float64:
				f := Float64(v)
				s.K1 = &f
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
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
