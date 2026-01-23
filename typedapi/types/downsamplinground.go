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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// DownsamplingRound type.
//
// https://github.com/elastic/elasticsearch-specification/blob/b1811e10a0722431d79d1c234dd412ff47d8656f/specification/indices/_types/DownsamplingRound.ts#L22-L31
type DownsamplingRound struct {
	// After The duration since rollover when this downsampling round should execute
	After Duration `json:"after"`
	// FixedInterval The downsample interval.
	FixedInterval string `json:"fixed_interval"`
}

func (s *DownsamplingRound) UnmarshalJSON(data []byte) error {

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

		case "after":
			if err := dec.Decode(&s.After); err != nil {
				return fmt.Errorf("%s | %w", "After", err)
			}

		case "fixed_interval":
			if err := dec.Decode(&s.FixedInterval); err != nil {
				return fmt.Errorf("%s | %w", "FixedInterval", err)
			}

		}
	}
	return nil
}

// NewDownsamplingRound returns a DownsamplingRound.
func NewDownsamplingRound() *DownsamplingRound {
	r := &DownsamplingRound{}

	return r
}

type DownsamplingRoundVariant interface {
	DownsamplingRoundCaster() *DownsamplingRound
}

func (s *DownsamplingRound) DownsamplingRoundCaster() *DownsamplingRound {
	return s
}
