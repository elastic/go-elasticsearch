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

// HoltLinearModelSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/pipeline.ts#L297-L300
type HoltLinearModelSettings struct {
	Alpha *float32 `json:"alpha,omitempty"`
	Beta  *float32 `json:"beta,omitempty"`
}

func (s *HoltLinearModelSettings) UnmarshalJSON(data []byte) error {

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

		case "alpha":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Alpha", err)
				}
				f := float32(value)
				s.Alpha = &f
			case float64:
				f := float32(v)
				s.Alpha = &f
			}

		case "beta":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Beta", err)
				}
				f := float32(value)
				s.Beta = &f
			case float64:
				f := float32(v)
				s.Beta = &f
			}

		}
	}
	return nil
}

// NewHoltLinearModelSettings returns a HoltLinearModelSettings.
func NewHoltLinearModelSettings() *HoltLinearModelSettings {
	r := &HoltLinearModelSettings{}

	return r
}

type HoltLinearModelSettingsVariant interface {
	HoltLinearModelSettingsCaster() *HoltLinearModelSettings
}

func (s *HoltLinearModelSettings) HoltLinearModelSettingsCaster() *HoltLinearModelSettings {
	return s
}
