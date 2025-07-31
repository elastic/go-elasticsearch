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

// PhraseSuggestCollateQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/search/_types/suggester.ts#L349-L358
type PhraseSuggestCollateQuery struct {
	// Id The search template ID.
	Id *string `json:"id,omitempty"`
	// Source The query source.
	Source ScriptSource `json:"source,omitempty"`
}

func (s *PhraseSuggestCollateQuery) UnmarshalJSON(data []byte) error {

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

		case "source":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Source", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		source_field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Source", err)
				}

				switch t {

				case "aggregations", "collapse", "docvalue_fields", "explain", "ext", "fields", "from", "highlight", "indices_boost", "knn", "min_score", "pit", "post_filter", "profile", "query", "rank", "rescore", "retriever", "runtime_mappings", "script_fields", "search_after", "seq_no_primary_term", "size", "slice", "sort", "_source", "stats", "stored_fields", "suggest", "terminate_after", "timeout", "track_scores", "track_total_hits", "version":
					o := NewSearchRequestBody()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Source", err)
					}
					s.Source = o
					break source_field

				}
			}
			if s.Source == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.Source); err != nil {
					return fmt.Errorf("%s | %w", "Source", err)
				}
			}

		}
	}
	return nil
}

// NewPhraseSuggestCollateQuery returns a PhraseSuggestCollateQuery.
func NewPhraseSuggestCollateQuery() *PhraseSuggestCollateQuery {
	r := &PhraseSuggestCollateQuery{}

	return r
}

type PhraseSuggestCollateQueryVariant interface {
	PhraseSuggestCollateQueryCaster() *PhraseSuggestCollateQuery
}

func (s *PhraseSuggestCollateQuery) PhraseSuggestCollateQueryCaster() *PhraseSuggestCollateQuery {
	return s
}
