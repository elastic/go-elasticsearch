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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// KnnSearchQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/2514615770f18dbb4e3887cc1a279995dbfd0724/specification/_global/knn_search/_types/Knn.ts#L24-L33
type KnnSearchQuery struct {
	// Field The name of the vector field to search against
	Field string `json:"field"`
	// K The final number of nearest neighbors to return as top hits
	K int `json:"k"`
	// NumCandidates The number of nearest neighbor candidates to consider per shard
	NumCandidates int `json:"num_candidates"`
	// QueryVector The query vector
	QueryVector []float32 `json:"query_vector"`
}

func (s *KnnSearchQuery) UnmarshalJSON(data []byte) error {

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
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
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
				s.K = value
			case float64:
				f := int(v)
				s.K = f
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
				s.NumCandidates = value
			case float64:
				f := int(v)
				s.NumCandidates = f
			}

		case "query_vector":
			if err := dec.Decode(&s.QueryVector); err != nil {
				return fmt.Errorf("%s | %w", "QueryVector", err)
			}

		}
	}
	return nil
}

// NewKnnSearchQuery returns a KnnSearchQuery.
func NewKnnSearchQuery() *KnnSearchQuery {
	r := &KnnSearchQuery{}

	return r
}

type KnnSearchQueryVariant interface {
	KnnSearchQueryCaster() *KnnSearchQuery
}

func (s *KnnSearchQuery) KnnSearchQueryCaster() *KnnSearchQuery {
	return s
}
