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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/synonymformat"
)

// SynonymGraphTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/token_filters.ts#L163-L165
type SynonymGraphTokenFilter struct {
	// Expand Expands definitions for equivalent synonym rules. Defaults to `true`.
	Expand *bool `json:"expand,omitempty"`
	// Format Sets the synonym rules format.
	Format *synonymformat.SynonymFormat `json:"format,omitempty"`
	// Lenient If `true` ignores errors while parsing the synonym rules. It is important to
	// note that only those synonym rules which cannot get parsed are ignored.
	// Defaults to the value of the `updateable` setting.
	Lenient *bool `json:"lenient,omitempty"`
	// Synonyms Used to define inline synonyms.
	Synonyms []string `json:"synonyms,omitempty"`
	// SynonymsPath Used to provide a synonym file. This path must be absolute or relative to the
	// `config` location.
	SynonymsPath *string `json:"synonyms_path,omitempty"`
	// SynonymsSet Provide a synonym set created via Synonyms Management APIs.
	SynonymsSet *string `json:"synonyms_set,omitempty"`
	// Tokenizer Controls the tokenizers that will be used to tokenize the synonym, this
	// parameter is for backwards compatibility for indices that created before 6.0.
	Tokenizer *string `json:"tokenizer,omitempty"`
	Type      string  `json:"type,omitempty"`
	// Updateable If `true` allows reloading search analyzers to pick up changes to synonym
	// files. Only to be used for search analyzers. Defaults to `false`.
	Updateable *bool   `json:"updateable,omitempty"`
	Version    *string `json:"version,omitempty"`
}

func (s *SynonymGraphTokenFilter) UnmarshalJSON(data []byte) error {

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
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Expand", err)
				}
				s.Expand = &value
			case bool:
				s.Expand = &v
			}

		case "format":
			if err := dec.Decode(&s.Format); err != nil {
				return fmt.Errorf("%s | %w", "Format", err)
			}

		case "lenient":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Lenient", err)
				}
				s.Lenient = &value
			case bool:
				s.Lenient = &v
			}

		case "synonyms":
			if err := dec.Decode(&s.Synonyms); err != nil {
				return fmt.Errorf("%s | %w", "Synonyms", err)
			}

		case "synonyms_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "SynonymsPath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SynonymsPath = &o

		case "synonyms_set":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "SynonymsSet", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SynonymsSet = &o

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
			s.Tokenizer = &o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "updateable":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Updateable", err)
				}
				s.Updateable = &value
			case bool:
				s.Updateable = &v
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
func (s SynonymGraphTokenFilter) MarshalJSON() ([]byte, error) {
	type innerSynonymGraphTokenFilter SynonymGraphTokenFilter
	tmp := innerSynonymGraphTokenFilter{
		Expand:       s.Expand,
		Format:       s.Format,
		Lenient:      s.Lenient,
		Synonyms:     s.Synonyms,
		SynonymsPath: s.SynonymsPath,
		SynonymsSet:  s.SynonymsSet,
		Tokenizer:    s.Tokenizer,
		Type:         s.Type,
		Updateable:   s.Updateable,
		Version:      s.Version,
	}

	tmp.Type = "synonym_graph"

	return json.Marshal(tmp)
}

// NewSynonymGraphTokenFilter returns a SynonymGraphTokenFilter.
func NewSynonymGraphTokenFilter() *SynonymGraphTokenFilter {
	r := &SynonymGraphTokenFilter{}

	return r
}
