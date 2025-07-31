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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scorenormalizer"
)

// LinearRetriever type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/Retriever.ts#L68-L75
type LinearRetriever struct {
	Fields []string `json:"fields,omitempty"`
	// Filter Query to filter the documents that can match.
	Filter []Query `json:"filter,omitempty"`
	// MinScore Minimum _score for matching documents. Documents with a lower _score are not
	// included in the top documents.
	MinScore *float32 `json:"min_score,omitempty"`
	// Name_ Retriever name.
	Name_          *string                          `json:"_name,omitempty"`
	Normalizer     *scorenormalizer.ScoreNormalizer `json:"normalizer,omitempty"`
	Query          *string                          `json:"query,omitempty"`
	RankWindowSize *int                             `json:"rank_window_size,omitempty"`
	// Retrievers Inner retrievers.
	Retrievers []InnerRetriever `json:"retrievers,omitempty"`
}

func (s *LinearRetriever) UnmarshalJSON(data []byte) error {

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

		case "fields":
			if err := dec.Decode(&s.Fields); err != nil {
				return fmt.Errorf("%s | %w", "Fields", err)
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

		case "normalizer":
			if err := dec.Decode(&s.Normalizer); err != nil {
				return fmt.Errorf("%s | %w", "Normalizer", err)
			}

		case "query":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Query = &o

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

		case "retrievers":
			if err := dec.Decode(&s.Retrievers); err != nil {
				return fmt.Errorf("%s | %w", "Retrievers", err)
			}

		}
	}
	return nil
}

// NewLinearRetriever returns a LinearRetriever.
func NewLinearRetriever() *LinearRetriever {
	r := &LinearRetriever{}

	return r
}

type LinearRetrieverVariant interface {
	LinearRetrieverCaster() *LinearRetriever
}

func (s *LinearRetriever) LinearRetrieverCaster() *LinearRetriever {
	return s
}
