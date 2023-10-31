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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// SettingsSimilarityLmd type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/indices/_types/IndexSettings.ts#L206-L209
type SettingsSimilarityLmd struct {
	Mu   int    `json:"mu"`
	Type string `json:"type,omitempty"`
}

func (s *SettingsSimilarityLmd) UnmarshalJSON(data []byte) error {

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

		case "mu":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Mu = value
			case float64:
				f := int(v)
				s.Mu = f
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
func (s SettingsSimilarityLmd) MarshalJSON() ([]byte, error) {
	type innerSettingsSimilarityLmd SettingsSimilarityLmd
	tmp := innerSettingsSimilarityLmd{
		Mu:   s.Mu,
		Type: s.Type,
	}

	tmp.Type = "LMDirichlet"

	return json.Marshal(tmp)
}

// NewSettingsSimilarityLmd returns a SettingsSimilarityLmd.
func NewSettingsSimilarityLmd() *SettingsSimilarityLmd {
	r := &SettingsSimilarityLmd{}

	return r
}
