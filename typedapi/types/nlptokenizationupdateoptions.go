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
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenizationtruncate"
)

// NlpTokenizationUpdateOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/_types/inference.ts#L375-L380
type NlpTokenizationUpdateOptions struct {
	// Span Span options to apply
	Span *int `json:"span,omitempty"`
	// Truncate Truncate options to apply
	Truncate *tokenizationtruncate.TokenizationTruncate `json:"truncate,omitempty"`
}

func (s *NlpTokenizationUpdateOptions) UnmarshalJSON(data []byte) error {

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

		case "span":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Span", err)
				}
				s.Span = &value
			case float64:
				f := int(v)
				s.Span = &f
			}

		case "truncate":
			if err := dec.Decode(&s.Truncate); err != nil {
				return fmt.Errorf("%s | %w", "Truncate", err)
			}

		}
	}
	return nil
}

// NewNlpTokenizationUpdateOptions returns a NlpTokenizationUpdateOptions.
func NewNlpTokenizationUpdateOptions() *NlpTokenizationUpdateOptions {
	r := &NlpTokenizationUpdateOptions{}

	return r
}

type NlpTokenizationUpdateOptionsVariant interface {
	NlpTokenizationUpdateOptionsCaster() *NlpTokenizationUpdateOptions
}

func (s *NlpTokenizationUpdateOptions) NlpTokenizationUpdateOptionsCaster() *NlpTokenizationUpdateOptions {
	return s
}
