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

// Limits type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/info/types.ts#L34-L40
type Limits struct {
	EffectiveMaxModelMemoryLimit ByteSize `json:"effective_max_model_memory_limit,omitempty"`
	MaxModelMemoryLimit          ByteSize `json:"max_model_memory_limit,omitempty"`
	MaxSingleMlNodeProcessors    *int     `json:"max_single_ml_node_processors,omitempty"`
	TotalMlMemory                ByteSize `json:"total_ml_memory"`
	TotalMlProcessors            *int     `json:"total_ml_processors,omitempty"`
}

func (s *Limits) UnmarshalJSON(data []byte) error {

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

		case "effective_max_model_memory_limit":
			if err := dec.Decode(&s.EffectiveMaxModelMemoryLimit); err != nil {
				return fmt.Errorf("%s | %w", "EffectiveMaxModelMemoryLimit", err)
			}

		case "max_model_memory_limit":
			if err := dec.Decode(&s.MaxModelMemoryLimit); err != nil {
				return fmt.Errorf("%s | %w", "MaxModelMemoryLimit", err)
			}

		case "max_single_ml_node_processors":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxSingleMlNodeProcessors", err)
				}
				s.MaxSingleMlNodeProcessors = &value
			case float64:
				f := int(v)
				s.MaxSingleMlNodeProcessors = &f
			}

		case "total_ml_memory":
			if err := dec.Decode(&s.TotalMlMemory); err != nil {
				return fmt.Errorf("%s | %w", "TotalMlMemory", err)
			}

		case "total_ml_processors":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalMlProcessors", err)
				}
				s.TotalMlProcessors = &value
			case float64:
				f := int(v)
				s.TotalMlProcessors = &f
			}

		}
	}
	return nil
}

// NewLimits returns a Limits.
func NewLimits() *Limits {
	r := &Limits{}

	return r
}
