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

// HdrMethod type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/metric.ts#L237-L242
type HdrMethod struct {
	// NumberOfSignificantValueDigits Specifies the resolution of values for the histogram in number of significant
	// digits.
	NumberOfSignificantValueDigits *int `json:"number_of_significant_value_digits,omitempty"`
}

func (s *HdrMethod) UnmarshalJSON(data []byte) error {

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

		case "number_of_significant_value_digits":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumberOfSignificantValueDigits", err)
				}
				s.NumberOfSignificantValueDigits = &value
			case float64:
				f := int(v)
				s.NumberOfSignificantValueDigits = &f
			}

		}
	}
	return nil
}

// NewHdrMethod returns a HdrMethod.
func NewHdrMethod() *HdrMethod {
	r := &HdrMethod{}

	return r
}

type HdrMethodVariant interface {
	HdrMethodCaster() *HdrMethod
}

func (s *HdrMethod) HdrMethodCaster() *HdrMethod {
	return s
}
