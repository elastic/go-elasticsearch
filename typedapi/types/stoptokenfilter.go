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

// StopTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/analysis/token_filters.ts#L97-L103
type StopTokenFilter struct {
	IgnoreCase     *bool    `json:"ignore_case,omitempty"`
	RemoveTrailing *bool    `json:"remove_trailing,omitempty"`
	Stopwords      []string `json:"stopwords,omitempty"`
	StopwordsPath  *string  `json:"stopwords_path,omitempty"`
	Type           string   `json:"type,omitempty"`
	Version        *string  `json:"version,omitempty"`
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
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IgnoreCase = &value
			case bool:
				s.IgnoreCase = &v
			}

		case "remove_trailing":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.RemoveTrailing = &value
			case bool:
				s.RemoveTrailing = &v
			}

		case "stopwords":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Stopwords = append(s.Stopwords, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Stopwords); err != nil {
					return err
				}
			}

		case "stopwords_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.StopwordsPath = &o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
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
