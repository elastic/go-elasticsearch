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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/densevectorsimilarity"
)

// DenseVectorQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_types/query_dsl/DenseVectorQuery.ts#L25-L58
type DenseVectorQuery struct {
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query. Boost values are relative to the default value of 1.0. A boost
	// value between 0 and 1.0 decreases the relevance score. A value greater than
	// 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	// Field The name of the `dense_vector` field to search against. The field can be
	// indexed, or not indexed (`index: false`).
	Field string `json:"field"`
	// Quantized If `false`, the documents are scored with the full-precision vectors,
	// regardless of the index options of the field. If `true`, the documents are
	// scored with the quantized vectors of a quantized field, which is faster but
	// less accurate. For a field that is not quantized, `true` has no effect.
	Quantized  *bool   `json:"quantized,omitempty"`
	QueryName_ *string `json:"_name,omitempty"`
	// QueryVector The query vector. It must have the same number of dimensions as the field.
	// You must provide either `query_vector` or `query_vector_builder`, but not
	// both.
	QueryVector QueryVector `json:"query_vector,omitempty"`
	// QueryVectorBuilder The configuration to build the query vector, for example by running a text
	// embedding model. You must provide either `query_vector` or
	// `query_vector_builder`, but not both.
	QueryVectorBuilder *QueryVectorBuilder `json:"query_vector_builder,omitempty"`
	// SimilarityFunction The similarity used for scoring, which overrides the similarity of the field
	// for this query. It defaults to the similarity of the field, or to `cosine`
	// for a field that is not indexed. For `bit` fields, only `l2_norm` is
	// supported. It can't be combined with `quantized: true`.
	SimilarityFunction *densevectorsimilarity.DenseVectorSimilarity `json:"similarity_function,omitempty"`
}

func (s *DenseVectorQuery) UnmarshalJSON(data []byte) error {

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

		case "quantized":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Quantized", err)
				}
				s.Quantized = &value
			case bool:
				s.Quantized = &v
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

		case "query_vector":
			if err := dec.Decode(&s.QueryVector); err != nil {
				return fmt.Errorf("%s | %w", "QueryVector", err)
			}

		case "query_vector_builder":
			if err := dec.Decode(&s.QueryVectorBuilder); err != nil {
				return fmt.Errorf("%s | %w", "QueryVectorBuilder", err)
			}

		case "similarity_function":
			if err := dec.Decode(&s.SimilarityFunction); err != nil {
				return fmt.Errorf("%s | %w", "SimilarityFunction", err)
			}

		}
	}
	return nil
}

// NewDenseVectorQuery returns a DenseVectorQuery.
func NewDenseVectorQuery() *DenseVectorQuery {
	r := &DenseVectorQuery{}

	return r
}

type DenseVectorQueryVariant interface {
	DenseVectorQueryCaster() *DenseVectorQuery
}

func (s *DenseVectorQuery) DenseVectorQueryCaster() *DenseVectorQuery {
	return s
}
