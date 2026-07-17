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
// https://github.com/elastic/elasticsearch-specification/tree/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// EsqlApproximationSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c/specification/esql/_types/types.ts#L221-L233
type EsqlApproximationSettings struct {
	// ConfidenceLevel The confidence level of the computed confidence intervals. A null value
	// disables computing confidence intervals.
	ConfidenceLevel *Float64 `json:"confidence_level,omitempty"`
	// Rows The number of sampled rows used for approximating the query. It must be at
	// least 10,000. A null value uses the system default.
	Rows *int `json:"rows,omitempty"`
}

func (s *EsqlApproximationSettings) UnmarshalJSON(data []byte) error {

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

		case "confidence_level":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ConfidenceLevel", err)
				}
				f := Float64(value)
				s.ConfidenceLevel = &f
			case float64:
				f := Float64(v)
				s.ConfidenceLevel = &f
			}

		case "rows":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Rows", err)
				}
				s.Rows = &value
			case float64:
				f := int(v)
				s.Rows = &f
			}

		}
	}
	return nil
}

// NewEsqlApproximationSettings returns a EsqlApproximationSettings.
func NewEsqlApproximationSettings() *EsqlApproximationSettings {
	r := &EsqlApproximationSettings{}

	return r
}

type EsqlApproximationSettingsVariant interface {
	EsqlApproximationSettingsCaster() *EsqlApproximationSettings
}

func (s *EsqlApproximationSettings) EsqlApproximationSettingsCaster() *EsqlApproximationSettings {
	return s
}

func (s *EsqlApproximationSettings) EsqlApproximationCaster() *EsqlApproximation {
	if s == nil {
		return nil
	}
	o := EsqlApproximation(s)
	return &o
}
