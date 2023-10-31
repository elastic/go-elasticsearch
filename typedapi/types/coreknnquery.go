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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// CoreKnnQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_global/knn_search/_types/Knn.ts#L24-L33
type CoreKnnQuery struct {
	// Field The name of the vector field to search against
	Field string `json:"field"`
	// K The final number of nearest neighbors to return as top hits
	K int64 `json:"k"`
	// NumCandidates The number of nearest neighbor candidates to consider per shard
	NumCandidates int64 `json:"num_candidates"`
	// QueryVector The query vector
	QueryVector []float32 `json:"query_vector"`
}

func (s *CoreKnnQuery) UnmarshalJSON(data []byte) error {

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
				return err
			}

		case "k":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.K = value
			case float64:
				f := int64(v)
				s.K = f
			}

		case "num_candidates":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.NumCandidates = value
			case float64:
				f := int64(v)
				s.NumCandidates = f
			}

		case "query_vector":
			if err := dec.Decode(&s.QueryVector); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewCoreKnnQuery returns a CoreKnnQuery.
func NewCoreKnnQuery() *CoreKnnQuery {
	r := &CoreKnnQuery{}

	return r
}
