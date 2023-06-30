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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/synonymformat"
)

// SynonymTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/analysis/token_filters.ts#L121-L130
type SynonymTokenFilter struct {
	Expand       *bool                        `json:"expand,omitempty"`
	Format       *synonymformat.SynonymFormat `json:"format,omitempty"`
	Lenient      *bool                        `json:"lenient,omitempty"`
	Synonyms     []string                     `json:"synonyms,omitempty"`
	SynonymsPath *string                      `json:"synonyms_path,omitempty"`
	Tokenizer    *string                      `json:"tokenizer,omitempty"`
	Type         string                       `json:"type,omitempty"`
	Updateable   *bool                        `json:"updateable,omitempty"`
	Version      *string                      `json:"version,omitempty"`
}

func (s *SynonymTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "expand":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Expand = &value
			case bool:
				s.Expand = &v
			}

		case "format":
			if err := dec.Decode(&s.Format); err != nil {
				return err
			}

		case "lenient":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Lenient = &value
			case bool:
				s.Lenient = &v
			}

		case "synonyms":
			if err := dec.Decode(&s.Synonyms); err != nil {
				return err
			}

		case "synonyms_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SynonymsPath = &o

		case "tokenizer":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Tokenizer = &o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "updateable":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Updateable = &value
			case bool:
				s.Updateable = &v
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewSynonymTokenFilter returns a SynonymTokenFilter.
func NewSynonymTokenFilter() *SynonymTokenFilter {
	r := &SynonymTokenFilter{}

	r.Type = "synonym"

	return r
}
