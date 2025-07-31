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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// SettingsSimilarityLmj type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/_types/IndexSettings.ts#L229-L232
type SettingsSimilarityLmj struct {
	Lambda *Float64 `json:"lambda,omitempty"`
	Type   string   `json:"type,omitempty"`
}

func (s *SettingsSimilarityLmj) UnmarshalJSON(data []byte) error {

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

		case "lambda":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Lambda", err)
				}
				f := Float64(value)
				s.Lambda = &f
			case float64:
				f := Float64(v)
				s.Lambda = &f
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
func (s SettingsSimilarityLmj) MarshalJSON() ([]byte, error) {
	type innerSettingsSimilarityLmj SettingsSimilarityLmj
	tmp := innerSettingsSimilarityLmj{
		Lambda: s.Lambda,
		Type:   s.Type,
	}

	tmp.Type = "LMJelinekMercer"

	return json.Marshal(tmp)
}

// NewSettingsSimilarityLmj returns a SettingsSimilarityLmj.
func NewSettingsSimilarityLmj() *SettingsSimilarityLmj {
	r := &SettingsSimilarityLmj{}

	return r
}

type SettingsSimilarityLmjVariant interface {
	SettingsSimilarityLmjCaster() *SettingsSimilarityLmj
}

func (s *SettingsSimilarityLmj) SettingsSimilarityLmjCaster() *SettingsSimilarityLmj {
	return s
}

func (s *SettingsSimilarityLmj) SettingsSimilarityCaster() *SettingsSimilarity {
	o := SettingsSimilarity(s)
	return &o
}
