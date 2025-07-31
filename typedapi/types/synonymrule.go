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
)

// SynonymRule type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/synonyms/_types/SynonymRule.ts#L26-L37
type SynonymRule struct {
	// Id The identifier for the synonym rule.
	// If you do not specify a synonym rule ID when you create a rule, an identifier
	// is created automatically by Elasticsearch.
	Id *string `json:"id,omitempty"`
	// Synonyms The synonyms that conform the synonym rule in Solr format.
	Synonyms string `json:"synonyms"`
}

func (s *SynonymRule) UnmarshalJSON(data []byte) error {

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

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "synonyms":
			if err := dec.Decode(&s.Synonyms); err != nil {
				return fmt.Errorf("%s | %w", "Synonyms", err)
			}

		}
	}
	return nil
}

// NewSynonymRule returns a SynonymRule.
func NewSynonymRule() *SynonymRule {
	r := &SynonymRule{}

	return r
}
