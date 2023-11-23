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
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// ExtendedBoundsFieldDateMath type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/aggregations/bucket.ts#L489-L498
type ExtendedBoundsFieldDateMath struct {
	// Max Maximum value for the bound.
	Max FieldDateMath `json:"max"`
	// Min Minimum value for the bound.
	Min FieldDateMath `json:"min"`
}

func (s *ExtendedBoundsFieldDateMath) UnmarshalJSON(data []byte) error {

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

		case "max":
			if err := dec.Decode(&s.Max); err != nil {
				return err
			}

		case "min":
			if err := dec.Decode(&s.Min); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewExtendedBoundsFieldDateMath returns a ExtendedBoundsFieldDateMath.
func NewExtendedBoundsFieldDateMath() *ExtendedBoundsFieldDateMath {
	r := &ExtendedBoundsFieldDateMath{}

	return r
}
