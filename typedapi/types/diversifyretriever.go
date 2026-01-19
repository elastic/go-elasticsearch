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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/diversifyretrievertypes"
)

// DiversifyRetriever type.
//
// https://github.com/elastic/elasticsearch-specification/blob/6785a6caa1fa3ca5ab3308963d79dce923a3469f/specification/_types/Retriever.ts#L216-L233
type DiversifyRetriever struct {
	// Field The document field on which to diversify results on.
	Field string `json:"field"`
	// Filter Query to filter the documents that can match.
	Filter []Query `json:"filter,omitempty"`
	// Lambda Controls the trade-off between relevance and diversity for MMR. A value of
	// 0.0 focuses solely on diversity, while a value of 1.0 focuses solely on
	// relevance. Required for MMR
	Lambda *float32 `json:"lambda,omitempty"`
	// MinScore Minimum _score for matching documents. Documents with a lower _score are not
	// included in the top documents.
	MinScore *float32 `json:"min_score,omitempty"`
	// Name_ Retriever name.
	Name_ *string `json:"_name,omitempty"`
	// QueryVector The query vector used for diversification.
	QueryVector []float32 `json:"query_vector,omitempty"`
	// QueryVectorBuilder a dense vector query vector builder to use instead of a static query_vector
	QueryVectorBuilder *QueryVectorBuilder `json:"query_vector_builder,omitempty"`
	// RankWindowSize The number of top documents from the nested retriever to consider for
	// diversification.
	RankWindowSize *int `json:"rank_window_size,omitempty"`
	// Retriever The nested retriever whose results will be diversified.
	Retriever RetrieverContainer `json:"retriever"`
	// Size The number of top documents to return after diversification.
	Size *int `json:"size,omitempty"`
	// Type The diversification strategy to apply.
	Type diversifyretrievertypes.DiversifyRetrieverTypes `json:"type"`
}

func (s *DiversifyRetriever) UnmarshalJSON(data []byte) error {

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

		case "field":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Field = o

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

		case "lambda":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Lambda", err)
				}
				f := float32(value)
				s.Lambda = &f
			case float64:
				f := float32(v)
				s.Lambda = &f
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

		case "query_vector":
			if err := dec.Decode(&s.QueryVector); err != nil {
				return fmt.Errorf("%s | %w", "QueryVector", err)
			}

		case "query_vector_builder":
			if err := dec.Decode(&s.QueryVectorBuilder); err != nil {
				return fmt.Errorf("%s | %w", "QueryVectorBuilder", err)
			}

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

		case "size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Size", err)
				}
				s.Size = &value
			case float64:
				f := int(v)
				s.Size = &f
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		}
	}
	return nil
}

// NewDiversifyRetriever returns a DiversifyRetriever.
func NewDiversifyRetriever() *DiversifyRetriever {
	r := &DiversifyRetriever{}

	return r
}

type DiversifyRetrieverVariant interface {
	DiversifyRetrieverCaster() *DiversifyRetriever
}

func (s *DiversifyRetriever) DiversifyRetrieverCaster() *DiversifyRetriever {
	return s
}
