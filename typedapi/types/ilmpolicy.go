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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// IlmPolicy type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ilm/_types/Policy.ts#L23-L29
type IlmPolicy struct {
	// Meta_ Arbitrary metadata that is not automatically generated or used by
	// Elasticsearch.
	Meta_  Metadata `json:"_meta,omitempty"`
	Phases Phases   `json:"phases"`
}

func (s *IlmPolicy) UnmarshalJSON(data []byte) error {

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

		case "_meta":
			if err := dec.Decode(&s.Meta_); err != nil {
				return fmt.Errorf("%s | %w", "Meta_", err)
			}

		case "phases":
			if err := dec.Decode(&s.Phases); err != nil {
				return fmt.Errorf("%s | %w", "Phases", err)
			}

		}
	}
	return nil
}

// NewIlmPolicy returns a IlmPolicy.
func NewIlmPolicy() *IlmPolicy {
	r := &IlmPolicy{}

	return r
}

type IlmPolicyVariant interface {
	IlmPolicyCaster() *IlmPolicy
}

func (s *IlmPolicy) IlmPolicyCaster() *IlmPolicy {
	return s
}
