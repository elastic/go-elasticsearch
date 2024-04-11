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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// CoordsGeoBounds type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/_types/Geo.ts#L154-L159
type CoordsGeoBounds struct {
	Bottom Float64 `json:"bottom"`
	Left   Float64 `json:"left"`
	Right  Float64 `json:"right"`
	Top    Float64 `json:"top"`
}

func (s *CoordsGeoBounds) UnmarshalJSON(data []byte) error {

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

		case "bottom":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Bottom", err)
				}
				f := Float64(value)
				s.Bottom = f
			case float64:
				f := Float64(v)
				s.Bottom = f
			}

		case "left":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Left", err)
				}
				f := Float64(value)
				s.Left = f
			case float64:
				f := Float64(v)
				s.Left = f
			}

		case "right":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Right", err)
				}
				f := Float64(value)
				s.Right = f
			case float64:
				f := Float64(v)
				s.Right = f
			}

		case "top":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Top", err)
				}
				f := Float64(value)
				s.Top = f
			case float64:
				f := Float64(v)
				s.Top = f
			}

		}
	}
	return nil
}

// NewCoordsGeoBounds returns a CoordsGeoBounds.
func NewCoordsGeoBounds() *CoordsGeoBounds {
	r := &CoordsGeoBounds{}

	return r
}
