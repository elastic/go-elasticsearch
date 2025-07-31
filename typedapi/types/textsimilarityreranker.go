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

// TextSimilarityReranker type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/Retriever.ts#L146-L157
type TextSimilarityReranker struct {
	// Field The document field to be used for text similarity comparisons. This field
	// should contain the text that will be evaluated against the inference_text
	Field string `json:"field"`
	// Filter Query to filter the documents that can match.
	Filter []Query `json:"filter,omitempty"`
	// InferenceId Unique identifier of the inference endpoint created using the inference API.
	InferenceId *string `json:"inference_id,omitempty"`
	// InferenceText The text snippet used as the basis for similarity comparison
	InferenceText string `json:"inference_text"`
	// MinScore Minimum _score for matching documents. Documents with a lower _score are not
	// included in the top documents.
	MinScore *float32 `json:"min_score,omitempty"`
	// Name_ Retriever name.
	Name_ *string `json:"_name,omitempty"`
	// RankWindowSize This value determines how many documents we will consider from the nested
	// retriever.
	RankWindowSize *int `json:"rank_window_size,omitempty"`
	// Retriever The nested retriever which will produce the first-level results, that will
	// later be used for reranking.
	Retriever RetrieverContainer `json:"retriever"`
}

func (s *TextSimilarityReranker) UnmarshalJSON(data []byte) error {

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

		case "inference_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "InferenceId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.InferenceId = &o

		case "inference_text":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "InferenceText", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.InferenceText = o

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

		}
	}
	return nil
}

// NewTextSimilarityReranker returns a TextSimilarityReranker.
func NewTextSimilarityReranker() *TextSimilarityReranker {
	r := &TextSimilarityReranker{}

	return r
}
