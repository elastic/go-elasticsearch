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

// PinnedRetriever type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/Retriever.ts#L77-L83
type PinnedRetriever struct {
	Docs []SpecifiedDocument `json:"docs,omitempty"`
	// Filter Query to filter the documents that can match.
	Filter []Query  `json:"filter,omitempty"`
	Ids    []string `json:"ids,omitempty"`
	// MinScore Minimum _score for matching documents. Documents with a lower _score are not
	// included in the top documents.
	MinScore *float32 `json:"min_score,omitempty"`
	// Name_ Retriever name.
	Name_          *string `json:"_name,omitempty"`
	RankWindowSize *int    `json:"rank_window_size,omitempty"`
	// Retriever Inner retriever.
	Retriever RetrieverContainer `json:"retriever"`
}

func (s *PinnedRetriever) UnmarshalJSON(data []byte) error {

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

		case "docs":
			if err := dec.Decode(&s.Docs); err != nil {
				return fmt.Errorf("%s | %w", "Docs", err)
			}

		case "filter":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewQuery()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Filter", err)
				}

				s.Filter = append(s.Filter, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Filter); err != nil {
					return fmt.Errorf("%s | %w", "Filter", err)
				}
			}

		case "ids":
			if err := dec.Decode(&s.Ids); err != nil {
				return fmt.Errorf("%s | %w", "Ids", err)
			}

		case "min_score":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinScore", err)
				}
				f := float32(value)
				s.MinScore = &f
			case float64:
				f := float32(v)
				s.MinScore = &f
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Name_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name_ = &o

		case "rank_window_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "RankWindowSize", err)
				}
				s.RankWindowSize = &value
			case float64:
				f := int(v)
				s.RankWindowSize = &f
			}

		case "retriever":
			if err := dec.Decode(&s.Retriever); err != nil {
				return fmt.Errorf("%s | %w", "Retriever", err)
			}

		}
	}
	return nil
}

// NewPinnedRetriever returns a PinnedRetriever.
func NewPinnedRetriever() *PinnedRetriever {
	r := &PinnedRetriever{}

	return r
}
