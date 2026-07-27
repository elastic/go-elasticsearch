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

// EsqlLoggingConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c/specification/xpack/usage/types.ts#L111-L124
type EsqlLoggingConfig struct {
	// Enabled Whether ES|QL query logging is enabled.
	Enabled bool `json:"enabled"`
	// Thresholds The configured logging thresholds, keyed by threshold name, if any.
	Thresholds map[string]Duration `json:"thresholds,omitempty"`
	// User Whether user information is included in the ES|QL query log.
	User bool `json:"user"`
}

func (s *EsqlLoggingConfig) UnmarshalJSON(data []byte) error {

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

		case "thresholds":
			if s.Thresholds == nil {
				s.Thresholds = make(map[string]Duration, 0)
			}
			if err := dec.Decode(&s.Thresholds); err != nil {
				return fmt.Errorf("%s | %w", "Thresholds", err)
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

// NewEsqlLoggingConfig returns a EsqlLoggingConfig.
func NewEsqlLoggingConfig() *EsqlLoggingConfig {
	r := &EsqlLoggingConfig{
		Thresholds: make(map[string]Duration),
	}

	return r
}
