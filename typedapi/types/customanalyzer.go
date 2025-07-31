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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// CustomAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/analyzers.ts#L28-L35
type CustomAnalyzer struct {
	CharFilter           []string `json:"char_filter,omitempty"`
	Filter               []string `json:"filter,omitempty"`
	PositionIncrementGap *int     `json:"position_increment_gap,omitempty"`
	PositionOffsetGap    *int     `json:"position_offset_gap,omitempty"`
	Tokenizer            string   `json:"tokenizer"`
	Type                 string   `json:"type,omitempty"`
}

func (s *CustomAnalyzer) UnmarshalJSON(data []byte) error {

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

		case "char_filter":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "CharFilter", err)
				}

				s.CharFilter = append(s.CharFilter, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.CharFilter); err != nil {
					return fmt.Errorf("%s | %w", "CharFilter", err)
				}
			}

		case "filter":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Filter", err)
				}

				s.Filter = append(s.Filter, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Filter); err != nil {
					return fmt.Errorf("%s | %w", "Filter", err)
				}
			}

		case "position_increment_gap":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "PositionIncrementGap", err)
				}
				s.PositionIncrementGap = &value
			case float64:
				f := int(v)
				s.PositionIncrementGap = &f
			}

		case "position_offset_gap":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "PositionOffsetGap", err)
				}
				s.PositionOffsetGap = &value
			case float64:
				f := int(v)
				s.PositionOffsetGap = &f
			}

		case "tokenizer":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Tokenizer", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Tokenizer = o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s CustomAnalyzer) MarshalJSON() ([]byte, error) {
	type innerCustomAnalyzer CustomAnalyzer
	tmp := innerCustomAnalyzer{
		CharFilter:           s.CharFilter,
		Filter:               s.Filter,
		PositionIncrementGap: s.PositionIncrementGap,
		PositionOffsetGap:    s.PositionOffsetGap,
		Tokenizer:            s.Tokenizer,
		Type:                 s.Type,
	}

	tmp.Type = "custom"

	return json.Marshal(tmp)
}

// NewCustomAnalyzer returns a CustomAnalyzer.
func NewCustomAnalyzer() *CustomAnalyzer {
	r := &CustomAnalyzer{}

	return r
}
