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
// https://github.com/elastic/elasticsearch-specification/tree/0f6f3696eb685db8b944feefb6a209ad7e385b9c

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scoremode"
)

// RescoreQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/0f6f3696eb685db8b944feefb6a209ad7e385b9c/specification/_global/search/_types/rescoring.ts#L40-L62
type RescoreQuery struct {
	// Query The query to use for rescoring.
	// This query is only run on the Top-K results returned by the `query` and
	// `post_filter` phases.
	Query Query `json:"rescore_query"`
	// QueryWeight Relative importance of the original query versus the rescore query.
	QueryWeight *Float64 `json:"query_weight,omitempty"`
	// RescoreQueryWeight Relative importance of the rescore query versus the original query.
	RescoreQueryWeight *Float64 `json:"rescore_query_weight,omitempty"`
	// ScoreMode Determines how scores are combined.
	ScoreMode *scoremode.ScoreMode `json:"score_mode,omitempty"`
}

func (s *RescoreQuery) UnmarshalJSON(data []byte) error {

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

		case "rescore_query":
			if err := dec.Decode(&s.Query); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}

		case "query_weight":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "QueryWeight", err)
				}
				f := Float64(value)
				s.QueryWeight = &f
			case float64:
				f := Float64(v)
				s.QueryWeight = &f
			}

		case "rescore_query_weight":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "RescoreQueryWeight", err)
				}
				f := Float64(value)
				s.RescoreQueryWeight = &f
			case float64:
				f := Float64(v)
				s.RescoreQueryWeight = &f
			}

		case "score_mode":
			if err := dec.Decode(&s.ScoreMode); err != nil {
				return fmt.Errorf("%s | %w", "ScoreMode", err)
			}

		}
	}
	return nil
}

// NewRescoreQuery returns a RescoreQuery.
func NewRescoreQuery() *RescoreQuery {
	r := &RescoreQuery{}

	return r
}

// true

type RescoreQueryVariant interface {
	RescoreQueryCaster() *RescoreQuery
}

func (s *RescoreQuery) RescoreQueryCaster() *RescoreQuery {
	return s
}
