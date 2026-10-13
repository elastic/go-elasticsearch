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
// https://github.com/elastic/elasticsearch-specification/tree/56c1eabdd35f941d1fbb3ad7ad8a9676664223f6

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Configuration object for doc values when sub-parameters are needed.
//
// https://github.com/elastic/elasticsearch-specification/blob/56c1eabdd35f941d1fbb3ad7ad8a9676664223f6/specification/_types/mapping/core.ts#L53-L78
type DocValuesConfig struct {
	// MultiValue If `false`, the field is treated as single-valued, enabling optimized
	// storage. Only has an effect when columnar index mode is active.
	MultiValue *bool `json:"multi_value,omitempty"`
	// Nullability If `false`, every document must provide a non-null value for the field: a
	// document that omits the field, sets it to `null`, or supplies only null
	// values (an empty array or an all-null array) is rejected at index time. A
	// field that defines `null_value` is always exempt, since the configured
	// default removes the absence of a value. Only has an effect when columnar
	// index mode is active.
	Nullability *bool `json:"nullability,omitempty"`
}

func (s *DocValuesConfig) UnmarshalJSON(data []byte) error {

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

		case "multi_value":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MultiValue", err)
				}
				s.MultiValue = &value
			case bool:
				s.MultiValue = &v
			}

		case "nullability":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Nullability", err)
				}
				s.Nullability = &value
			case bool:
				s.Nullability = &v
			}

		}
	}
	return nil
}

// NewDocValuesConfig returns a DocValuesConfig.
func NewDocValuesConfig() *DocValuesConfig {
	r := &DocValuesConfig{}

	return r
}

type DocValuesConfigVariant interface {
	DocValuesConfigCaster() *DocValuesConfig
}

func (s *DocValuesConfig) DocValuesConfigCaster() *DocValuesConfig {
	return s
}

func (s *DocValuesConfig) DocValuesCaster() *DocValues {
	if s == nil {
		return nil
	}
	o := DocValues(s)
	return &o
}
