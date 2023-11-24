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
	"strconv"
)

// PublishedClusterStates type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/nodes/_types/Stats.ts#L263-L276
type PublishedClusterStates struct {
	// CompatibleDiffs Number of compatible differences between published cluster states.
	CompatibleDiffs *int64 `json:"compatible_diffs,omitempty"`
	// FullStates Number of published cluster states.
	FullStates *int64 `json:"full_states,omitempty"`
	// IncompatibleDiffs Number of incompatible differences between published cluster states.
	IncompatibleDiffs *int64 `json:"incompatible_diffs,omitempty"`
}

func (s *PublishedClusterStates) UnmarshalJSON(data []byte) error {

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

		case "compatible_diffs":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.CompatibleDiffs = &value
			case float64:
				f := int64(v)
				s.CompatibleDiffs = &f
			}

		case "full_states":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.FullStates = &value
			case float64:
				f := int64(v)
				s.FullStates = &f
			}

		case "incompatible_diffs":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.IncompatibleDiffs = &value
			case float64:
				f := int64(v)
				s.IncompatibleDiffs = &f
			}

		}
	}
	return nil
}

// NewPublishedClusterStates returns a PublishedClusterStates.
func NewPublishedClusterStates() *PublishedClusterStates {
	r := &PublishedClusterStates{}

	return r
}
