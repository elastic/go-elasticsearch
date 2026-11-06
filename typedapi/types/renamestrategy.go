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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// RenameStrategy type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/data_recovery/start_recovery/types.ts#L27-L38
type RenameStrategy struct {
	// Prefix Text to add at the start of the name of each recovered source. At least one
	// of `prefix` and `suffix` must be set to a non-empty value, and both can be
	// set together.
	Prefix *string `json:"prefix,omitempty"`
	// Suffix Text to add at the end of the name of each recovered source. At least one of
	// `prefix` and `suffix` must be set to a non-empty value, and both can be set
	// together.
	Suffix *string `json:"suffix,omitempty"`
}

func (s *RenameStrategy) UnmarshalJSON(data []byte) error {

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

		case "prefix":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Prefix", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Prefix = &o

		case "suffix":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Suffix", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Suffix = &o

		}
	}
	return nil
}

// NewRenameStrategy returns a RenameStrategy.
func NewRenameStrategy() *RenameStrategy {
	r := &RenameStrategy{}

	return r
}

type RenameStrategyVariant interface {
	RenameStrategyCaster() *RenameStrategy
}

func (s *RenameStrategy) RenameStrategyCaster() *RenameStrategy {
	return s
}
