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

package knneval

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Request holds the request body struct for the package knneval
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_global/knn_eval/KnnEvalRequest.ts#L25-L117
type Request struct {
	// AllowExactBaseline If `true`, the `baseline` can be `exact`. An exact baseline compares every
	// vector at full precision, and can severely slow down other searches.
	AllowExactBaseline *bool `json:"allow_exact_baseline,omitempty"`
	// Baseline The settings of the search that the `knn_settings` entries are compared with.
	// If not specified, it uses a `visit_percentage` of 20 and a
	// `rescore_vector.oversample` of 100.
	Baseline *types.KnnEvalSettings `json:"baseline,omitempty"`
	// Field The `dense_vector` field to evaluate.
	Field string `json:"field"`
	// K The number of nearest neighbors to compute the recall for, from 1 to 1000.
	K int `json:"k"`
	// KnnSettings The kNN search settings to evaluate, from 1 to 32 unique entries.
	KnnSettings []types.KnnEvalSettings `json:"knn_settings"`
	// QuerySource The source of the query vectors.
	QuerySource types.KnnEvalQuerySource `json:"query_source"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Knneval request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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

		case "allow_exact_baseline":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllowExactBaseline", err)
				}
				s.AllowExactBaseline = &value
			case bool:
				s.AllowExactBaseline = &v
			}

		case "baseline":
			if err := dec.Decode(&s.Baseline); err != nil {
				return fmt.Errorf("%s | %w", "Baseline", err)
			}

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

		case "knn_settings":
			if err := dec.Decode(&s.KnnSettings); err != nil {
				return fmt.Errorf("%s | %w", "KnnSettings", err)
			}

		case "query_source":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			kind := make(map[string]string, 0)
			localDec := json.NewDecoder(source)
			localDec.Decode(&kind)
			source.Seek(0, io.SeekStart)

			switch kind["from"] {

			case "docs":
				o := types.NewKnnEvalDocsQuerySource()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "docs", err)
				}
				s.QuerySource = *o
			case "vectors":
				o := types.NewKnnEvalVectorsQuerySource()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "vectors", err)
				}
				s.QuerySource = *o
			default:
				if err := localDec.Decode(&s.QuerySource); err != nil {
					return fmt.Errorf("QuerySource | %w", err)
				}
			}

		}
	}
	return nil
}
