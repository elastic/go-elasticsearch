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

// RrfRank type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/Rank.ts#L32-L37
type RrfRank struct {
	// RankConstant How much influence documents in individual result sets per query have over
	// the final ranked result set
	RankConstant *int64 `json:"rank_constant,omitempty"`
	// RankWindowSize Size of the individual result sets per query
	RankWindowSize *int64 `json:"rank_window_size,omitempty"`
}

func (s *RrfRank) UnmarshalJSON(data []byte) error {

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

		case "rank_constant":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "RankConstant", err)
				}
				s.RankConstant = &value
			case float64:
				f := int64(v)
				s.RankConstant = &f
			}

		case "rank_window_size":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "RankWindowSize", err)
				}
				s.RankWindowSize = &value
			case float64:
				f := int64(v)
				s.RankWindowSize = &f
			}

		}
	}
	return nil
}

// NewRrfRank returns a RrfRank.
func NewRrfRank() *RrfRank {
	r := &RrfRank{}

	return r
}
