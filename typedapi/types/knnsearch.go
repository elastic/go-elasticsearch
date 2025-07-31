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

// KnnSearch type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/Knn.ts#L35-L62
type KnnSearch struct {
	// Boost Boost value to apply to kNN scores
	Boost *float32 `json:"boost,omitempty"`
	// Field The name of the vector field to search against
	Field string `json:"field"`
	// Filter Filters for the kNN search query
	Filter []Query `json:"filter,omitempty"`
	// InnerHits If defined, each search hit will contain inner hits.
	InnerHits *InnerHits `json:"inner_hits,omitempty"`
	// K The final number of nearest neighbors to return as top hits
	K *int `json:"k,omitempty"`
	// NumCandidates The number of nearest neighbor candidates to consider per shard
	NumCandidates *int `json:"num_candidates,omitempty"`
	// QueryVector The query vector
	QueryVector []float32 `json:"query_vector,omitempty"`
	// QueryVectorBuilder The query vector builder. You must provide a query_vector_builder or
	// query_vector, but not both.
	QueryVectorBuilder *QueryVectorBuilder `json:"query_vector_builder,omitempty"`
	// RescoreVector Apply oversampling and rescoring to quantized vectors *
	RescoreVector *RescoreVector `json:"rescore_vector,omitempty"`
	// Similarity The minimum similarity for a vector to be considered a match
	Similarity *float32 `json:"similarity,omitempty"`
}

func (s *KnnSearch) UnmarshalJSON(data []byte) error {

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

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
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

		case "inner_hits":
			if err := dec.Decode(&s.InnerHits); err != nil {
				return fmt.Errorf("%s | %w", "InnerHits", err)
			}

		case "k":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "K", err)
				}
				s.K = &value
			case float64:
				f := int(v)
				s.K = &f
			}

		case "num_candidates":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumCandidates", err)
				}
				s.NumCandidates = &value
			case float64:
				f := int(v)
				s.NumCandidates = &f
			}

		case "query_vector":
			if err := dec.Decode(&s.QueryVector); err != nil {
				return fmt.Errorf("%s | %w", "QueryVector", err)
			}

		case "query_vector_builder":
			if err := dec.Decode(&s.QueryVectorBuilder); err != nil {
				return fmt.Errorf("%s | %w", "QueryVectorBuilder", err)
			}

		case "rescore_vector":
			if err := dec.Decode(&s.RescoreVector); err != nil {
				return fmt.Errorf("%s | %w", "RescoreVector", err)
			}

		case "similarity":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Similarity", err)
				}
				f := float32(value)
				s.Similarity = &f
			case float64:
				f := float32(v)
				s.Similarity = &f
			}

		}
	}
	return nil
}

// NewKnnSearch returns a KnnSearch.
func NewKnnSearch() *KnnSearch {
	r := &KnnSearch{}

	return r
}
