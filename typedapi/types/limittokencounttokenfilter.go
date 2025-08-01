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
)

// LimitTokenCountTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/analysis/token_filters.ts#L336-L342
type LimitTokenCountTokenFilter struct {
	// ConsumeAllTokens If `true`, the limit filter exhausts the token stream, even if the
	// `max_token_count` has already been reached. Defaults to `false`.
	ConsumeAllTokens *bool `json:"consume_all_tokens,omitempty"`
	// MaxTokenCount Maximum number of tokens to keep. Once this limit is reached, any remaining
	// tokens are excluded from the output. Defaults to `1`.
	MaxTokenCount Stringifiedinteger `json:"max_token_count,omitempty"`
	Type          string             `json:"type,omitempty"`
	Version       *string            `json:"version,omitempty"`
}

func (s *LimitTokenCountTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "consume_all_tokens":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ConsumeAllTokens", err)
				}
				s.ConsumeAllTokens = &value
			case bool:
				s.ConsumeAllTokens = &v
			}

		case "max_token_count":
			if err := dec.Decode(&s.MaxTokenCount); err != nil {
				return fmt.Errorf("%s | %w", "MaxTokenCount", err)
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
func (s LimitTokenCountTokenFilter) MarshalJSON() ([]byte, error) {
	type innerLimitTokenCountTokenFilter LimitTokenCountTokenFilter
	tmp := innerLimitTokenCountTokenFilter{
		ConsumeAllTokens: s.ConsumeAllTokens,
		MaxTokenCount:    s.MaxTokenCount,
		Type:             s.Type,
		Version:          s.Version,
	}

	tmp.Type = "limit"

	return json.Marshal(tmp)
}

// NewLimitTokenCountTokenFilter returns a LimitTokenCountTokenFilter.
func NewLimitTokenCountTokenFilter() *LimitTokenCountTokenFilter {
	r := &LimitTokenCountTokenFilter{}

	return r
}

type LimitTokenCountTokenFilterVariant interface {
	LimitTokenCountTokenFilterCaster() *LimitTokenCountTokenFilter
}

func (s *LimitTokenCountTokenFilter) LimitTokenCountTokenFilterCaster() *LimitTokenCountTokenFilter {
	return s
}

func (s *LimitTokenCountTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	o := TokenFilterDefinition(s)
	return &o
}
