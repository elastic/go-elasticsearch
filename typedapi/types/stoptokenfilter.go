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

// StopTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/token_filters.ts#L125-L136
type StopTokenFilter struct {
	// IgnoreCase If `true`, stop word matching is case insensitive. For example, if `true`, a
	// stop word of the matches and removes `The`, `THE`, or `the`. Defaults to
	// `false`.
	IgnoreCase *bool `json:"ignore_case,omitempty"`
	// RemoveTrailing If `true`, the last token of a stream is removed if it’s a stop word.
	// Defaults to `true`.
	RemoveTrailing *bool `json:"remove_trailing,omitempty"`
	// Stopwords Language value, such as `_arabic_` or `_thai_`. Defaults to `_english_`.
	Stopwords StopWords `json:"stopwords,omitempty"`
	// StopwordsPath Path to a file that contains a list of stop words to remove.
	// This path must be absolute or relative to the `config` location, and the file
	// must be UTF-8 encoded. Each stop word in the file must be separated by a line
	// break.
	StopwordsPath *string `json:"stopwords_path,omitempty"`
	Type          string  `json:"type,omitempty"`
	Version       *string `json:"version,omitempty"`
}

func (s *StopTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "ignore_case":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreCase", err)
				}
				s.IgnoreCase = &value
			case bool:
				s.IgnoreCase = &v
			}

		case "remove_trailing":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "RemoveTrailing", err)
				}
				s.RemoveTrailing = &value
			case bool:
				s.RemoveTrailing = &v
			}

		case "stopwords":
			if err := dec.Decode(&s.Stopwords); err != nil {
				return fmt.Errorf("%s | %w", "Stopwords", err)
			}

		case "stopwords_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "StopwordsPath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.StopwordsPath = &o

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
func (s StopTokenFilter) MarshalJSON() ([]byte, error) {
	type innerStopTokenFilter StopTokenFilter
	tmp := innerStopTokenFilter{
		IgnoreCase:     s.IgnoreCase,
		RemoveTrailing: s.RemoveTrailing,
		Stopwords:      s.Stopwords,
		StopwordsPath:  s.StopwordsPath,
		Type:           s.Type,
		Version:        s.Version,
	}

	tmp.Type = "stop"

	return json.Marshal(tmp)
}

// NewStopTokenFilter returns a StopTokenFilter.
func NewStopTokenFilter() *StopTokenFilter {
	r := &StopTokenFilter{}

	return r
}
