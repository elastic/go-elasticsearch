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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// KnnEvalQueryVector type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_global/knn_eval/types.ts#L59-L68
type KnnEvalQueryVector struct {
	// Id The identifier of the query, which identifies it in the `failures` of the
	// response.
	Id string `json:"id"`
	// QueryVector The query vector.
	QueryVector QueryVector `json:"query_vector"`
}

func (s *KnnEvalQueryVector) UnmarshalJSON(data []byte) error {

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

		case "id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Id = o

		case "query_vector":
			if err := dec.Decode(&s.QueryVector); err != nil {
				return fmt.Errorf("%s | %w", "QueryVector", err)
			}

		}
	}
	return nil
}

// NewKnnEvalQueryVector returns a KnnEvalQueryVector.
func NewKnnEvalQueryVector() *KnnEvalQueryVector {
	r := &KnnEvalQueryVector{}

	return r
}

type KnnEvalQueryVectorVariant interface {
	KnnEvalQueryVectorCaster() *KnnEvalQueryVector
}

func (s *KnnEvalQueryVector) KnnEvalQueryVectorCaster() *KnnEvalQueryVector {
	return s
}
