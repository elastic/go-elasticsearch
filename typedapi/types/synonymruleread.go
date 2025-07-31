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

// SynonymRuleRead type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/synonyms/_types/SynonymRule.ts#L40-L50
type SynonymRuleRead struct {
	// Id Synonym Rule identifier
	Id string `json:"id"`
	// Synonyms Synonyms, in Solr format, that conform the synonym rule.
	Synonyms string `json:"synonyms"`
}

func (s *SynonymRuleRead) UnmarshalJSON(data []byte) error {

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

// NewSynonymRuleRead returns a SynonymRuleRead.
func NewSynonymRuleRead() *SynonymRuleRead {
	r := &SynonymRuleRead{}

	return r
}
