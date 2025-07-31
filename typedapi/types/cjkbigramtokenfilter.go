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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/cjkbigramignoredscript"
)

// CjkBigramTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/token_filters.ts#L466-L472
type CjkBigramTokenFilter struct {
	// IgnoredScripts Array of character scripts for which to disable bigrams.
	IgnoredScripts []cjkbigramignoredscript.CjkBigramIgnoredScript `json:"ignored_scripts,omitempty"`
	// OutputUnigrams If `true`, emit tokens in both bigram and unigram form. If `false`, a CJK
	// character is output in unigram form when it has no adjacent characters.
	// Defaults to `false`.
	OutputUnigrams *bool   `json:"output_unigrams,omitempty"`
	Type           string  `json:"type,omitempty"`
	Version        *string `json:"version,omitempty"`
}

func (s *CjkBigramTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "ignored_scripts":
			if err := dec.Decode(&s.IgnoredScripts); err != nil {
				return fmt.Errorf("%s | %w", "IgnoredScripts", err)
			}

		case "output_unigrams":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "OutputUnigrams", err)
				}
				s.OutputUnigrams = &value
			case bool:
				s.OutputUnigrams = &v
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
func (s CjkBigramTokenFilter) MarshalJSON() ([]byte, error) {
	type innerCjkBigramTokenFilter CjkBigramTokenFilter
	tmp := innerCjkBigramTokenFilter{
		IgnoredScripts: s.IgnoredScripts,
		OutputUnigrams: s.OutputUnigrams,
		Type:           s.Type,
		Version:        s.Version,
	}

	tmp.Type = "cjk_bigram"

	return json.Marshal(tmp)
}

// NewCjkBigramTokenFilter returns a CjkBigramTokenFilter.
func NewCjkBigramTokenFilter() *CjkBigramTokenFilter {
	r := &CjkBigramTokenFilter{}

	return r
}
