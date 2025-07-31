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

// HitsMetadata type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/search/_types/hits.ts#L68-L74
type HitsMetadata struct {
	Hits     []Hit    `json:"hits"`
	MaxScore *Float64 `json:"max_score,omitempty"`
	// Total Total hit count information, present only if `track_total_hits` wasn't
	// `false` in the search request.
	Total *TotalHits `json:"total,omitempty"`
}

func (s *HitsMetadata) UnmarshalJSON(data []byte) error {

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

		case "hits":
			if err := dec.Decode(&s.Hits); err != nil {
				return fmt.Errorf("%s | %w", "Hits", err)
			}

		case "max_score":
			if err := dec.Decode(&s.MaxScore); err != nil {
				return fmt.Errorf("%s | %w", "MaxScore", err)
			}

		case "total":
			if err := dec.Decode(&s.Total); err != nil {
				return fmt.Errorf("%s | %w", "Total", err)
			}

		}
	}
	return nil
}

// NewHitsMetadata returns a HitsMetadata.
func NewHitsMetadata() *HitsMetadata {
	r := &HitsMetadata{}

	return r
}
