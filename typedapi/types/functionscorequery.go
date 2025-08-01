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
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/functionboostmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/functionscoremode"
)

// FunctionScoreQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/query_dsl/compound.ts#L105-L135
type FunctionScoreQuery struct {
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	// BoostMode Defines how he newly computed score is combined with the score of the query
	BoostMode *functionboostmode.FunctionBoostMode `json:"boost_mode,omitempty"`
	// Functions One or more functions that compute a new score for each document returned by
	// the query.
	Functions []FunctionScore `json:"functions,omitempty"`
	// MaxBoost Restricts the new score to not exceed the provided limit.
	MaxBoost *Float64 `json:"max_boost,omitempty"`
	// MinScore Excludes documents that do not meet the provided score threshold.
	MinScore *Float64 `json:"min_score,omitempty"`
	// Query A query that determines the documents for which a new score is computed.
	Query      *Query  `json:"query,omitempty"`
	QueryName_ *string `json:"_name,omitempty"`
	// ScoreMode Specifies how the computed scores are combined
	ScoreMode *functionscoremode.FunctionScoreMode `json:"score_mode,omitempty"`
}

func (s *FunctionScoreQuery) UnmarshalJSON(data []byte) error {

	if !bytes.HasPrefix(data, []byte(`{`)) {
		err := json.NewDecoder(bytes.NewReader(data)).Decode(&s.Functions)
		return err
	}

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

		case "boost":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Boost", err)
				}
				f := float32(value)
				s.Boost = &f
			case float64:
				f := float32(v)
				s.Boost = &f
			}

		case "boost_mode":
			if err := dec.Decode(&s.BoostMode); err != nil {
				return fmt.Errorf("%s | %w", "BoostMode", err)
			}

		case "functions":
			if err := dec.Decode(&s.Functions); err != nil {
				return fmt.Errorf("%s | %w", "Functions", err)
			}

		case "max_boost":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxBoost", err)
				}
				f := Float64(value)
				s.MaxBoost = &f
			case float64:
				f := Float64(v)
				s.MaxBoost = &f
			}

		case "min_score":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinScore", err)
				}
				f := Float64(value)
				s.MinScore = &f
			case float64:
				f := Float64(v)
				s.MinScore = &f
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "QueryName_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueryName_ = &o

		case "score_mode":
			if err := dec.Decode(&s.ScoreMode); err != nil {
				return fmt.Errorf("%s | %w", "ScoreMode", err)
			}

		}
	}
	return nil
}

// NewFunctionScoreQuery returns a FunctionScoreQuery.
func NewFunctionScoreQuery() *FunctionScoreQuery {
	r := &FunctionScoreQuery{}

	return r
}

type FunctionScoreQueryVariant interface {
	FunctionScoreQueryCaster() *FunctionScoreQuery
}

func (s *FunctionScoreQuery) FunctionScoreQueryCaster() *FunctionScoreQuery {
	return s
}
