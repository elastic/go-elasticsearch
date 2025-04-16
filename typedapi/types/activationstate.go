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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ActivationState type.
//
// https://github.com/elastic/elasticsearch-specification/blob/52c473efb1fb5320a5bac12572d0b285882862fb/specification/watcher/_types/Activation.ts#L24-L27
type ActivationState struct {
	Active    bool     `json:"active"`
	Timestamp DateTime `json:"timestamp"`
}

func (s *ActivationState) UnmarshalJSON(data []byte) error {

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

		case "active":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Active", err)
				}
				s.Active = value
			case bool:
				s.Active = v
			}

		case "timestamp":
			if err := dec.Decode(&s.Timestamp); err != nil {
				return fmt.Errorf("%s | %w", "Timestamp", err)
			}

		}
	}
	return nil
}

// NewActivationState returns a ActivationState.
func NewActivationState() *ActivationState {
	r := &ActivationState{}

	return r
}

// true

type ActivationStateVariant interface {
	ActivationStateCaster() *ActivationState
}

func (s *ActivationState) ActivationStateCaster() *ActivationState {
	return s
}
