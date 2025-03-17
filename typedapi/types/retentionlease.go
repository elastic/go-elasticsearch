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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// RetentionLease type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ea991724f4dd4f90c496eff547d3cc2e6529f509/specification/indices/_types/IndexSettings.ts#L66-L68
type RetentionLease struct {
	Period Duration `json:"period"`
}

func (s *RetentionLease) UnmarshalJSON(data []byte) error {

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

		case "period":
			if err := dec.Decode(&s.Period); err != nil {
				return fmt.Errorf("%s | %w", "Period", err)
			}

		}
	}
	return nil
}

// NewRetentionLease returns a RetentionLease.
func NewRetentionLease() *RetentionLease {
	r := &RetentionLease{}

	return r
}

// true

type RetentionLeaseVariant interface {
	RetentionLeaseCaster() *RetentionLease
}

func (s *RetentionLease) RetentionLeaseCaster() *RetentionLease {
	return s
}
