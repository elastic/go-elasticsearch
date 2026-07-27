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
// https://github.com/elastic/elasticsearch-specification/tree/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// QueryLoggingConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c/specification/xpack/usage/types.ts#L92-L109
type QueryLoggingConfig struct {
	// Enabled Whether query logging is enabled.
	Enabled bool `json:"enabled"`
	// System Whether system queries are included in the query log.
	System bool `json:"system"`
	// Threshold The configured logging threshold, if any.
	Threshold Duration `json:"threshold,omitempty"`
	// User Whether user information is included in the query log.
	User bool `json:"user"`
}

func (s *QueryLoggingConfig) UnmarshalJSON(data []byte) error {

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

		case "enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Enabled", err)
				}
				s.Enabled = value
			case bool:
				s.Enabled = v
			}

		case "system":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "System", err)
				}
				s.System = value
			case bool:
				s.System = v
			}

		case "threshold":
			if err := dec.Decode(&s.Threshold); err != nil {
				return fmt.Errorf("%s | %w", "Threshold", err)
			}

		case "user":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "User", err)
				}
				s.User = value
			case bool:
				s.User = v
			}

		}
	}
	return nil
}

// NewQueryLoggingConfig returns a QueryLoggingConfig.
func NewQueryLoggingConfig() *QueryLoggingConfig {
	r := &QueryLoggingConfig{}

	return r
}
