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
// https://github.com/elastic/elasticsearch-specification/tree/e16324dcde9297dd1149c1ef3d6d58afe272e646

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/edgengramside"
)

// EdgeNGramTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e16324dcde9297dd1149c1ef3d6d58afe272e646/specification/_types/analysis/token_filters.ts#L79-L85
type EdgeNGramTokenFilter struct {
	MaxGram          *int                         `json:"max_gram,omitempty"`
	MinGram          *int                         `json:"min_gram,omitempty"`
	PreserveOriginal Stringifiedboolean           `json:"preserve_original,omitempty"`
	Side             *edgengramside.EdgeNGramSide `json:"side,omitempty"`
	Type             string                       `json:"type,omitempty"`
	Version          *string                      `json:"version,omitempty"`
}

func (s *EdgeNGramTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "max_gram":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxGram = &value
			case float64:
				f := int(v)
				s.MaxGram = &f
			}

		case "min_gram":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MinGram = &value
			case float64:
				f := int(v)
				s.MinGram = &f
			}

		case "preserve_original":
			if err := dec.Decode(&s.PreserveOriginal); err != nil {
				return err
			}

		case "side":
			if err := dec.Decode(&s.Side); err != nil {
				return err
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s EdgeNGramTokenFilter) MarshalJSON() ([]byte, error) {
	type innerEdgeNGramTokenFilter EdgeNGramTokenFilter
	tmp := innerEdgeNGramTokenFilter{
		MaxGram:          s.MaxGram,
		MinGram:          s.MinGram,
		PreserveOriginal: s.PreserveOriginal,
		Side:             s.Side,
		Type:             s.Type,
		Version:          s.Version,
	}

	tmp.Type = "edge_ngram"

	return json.Marshal(tmp)
}

// NewEdgeNGramTokenFilter returns a EdgeNGramTokenFilter.
func NewEdgeNGramTokenFilter() *EdgeNGramTokenFilter {
	r := &EdgeNGramTokenFilter{}

	return r
}
