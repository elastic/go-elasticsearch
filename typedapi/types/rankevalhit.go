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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// RankEvalHit type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/_global/rank_eval/types.ts#L141-L145
type RankEvalHit struct {
	Id_    string  `json:"_id"`
	Index_ string  `json:"_index"`
	Score_ Float64 `json:"_score"`
}

func (s *RankEvalHit) UnmarshalJSON(data []byte) error {

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

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return fmt.Errorf("%s | %w", "Id_", err)
			}

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return fmt.Errorf("%s | %w", "Index_", err)
			}

		case "_score":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Score_", err)
				}
				f := Float64(value)
				s.Score_ = f
			case float64:
				f := Float64(v)
				s.Score_ = f
			}

		}
	}
	return nil
}

// NewRankEvalHit returns a RankEvalHit.
func NewRankEvalHit() *RankEvalHit {
	r := &RankEvalHit{}

	return r
}
