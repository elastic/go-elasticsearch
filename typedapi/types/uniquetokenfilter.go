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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// UniqueTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/c75a0abec670d027d13eb8d6f23374f86621c76b/specification/_types/analysis/token_filters.ts#L336-L339
type UniqueTokenFilter struct {
	OnlyOnSamePosition *bool   `json:"only_on_same_position,omitempty"`
	Type               string  `json:"type,omitempty"`
	Version            *string `json:"version,omitempty"`
}

func (s *UniqueTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "only_on_same_position":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "OnlyOnSamePosition", err)
				}
				s.OnlyOnSamePosition = &value
			case bool:
				s.OnlyOnSamePosition = &v
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s UniqueTokenFilter) MarshalJSON() ([]byte, error) {
	type innerUniqueTokenFilter UniqueTokenFilter
	tmp := innerUniqueTokenFilter{
		OnlyOnSamePosition: s.OnlyOnSamePosition,
		Type:               s.Type,
		Version:            s.Version,
	}

	tmp.Type = "unique"

	return json.Marshal(tmp)
}

// NewUniqueTokenFilter returns a UniqueTokenFilter.
func NewUniqueTokenFilter() *UniqueTokenFilter {
	r := &UniqueTokenFilter{}

	return r
}

// true

type UniqueTokenFilterVariant interface {
	UniqueTokenFilterCaster() *UniqueTokenFilter
}

func (s *UniqueTokenFilter) UniqueTokenFilterCaster() *UniqueTokenFilter {
	return s
}
