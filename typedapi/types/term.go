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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// Term type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_global/termvectors/types.ts#L34-L40
type Term struct {
	DocFreq  *int               `json:"doc_freq,omitempty"`
	Score    *Float64           `json:"score,omitempty"`
	TermFreq int                `json:"term_freq"`
	Tokens   []TermVectorsToken `json:"tokens,omitempty"`
	Ttf      *int               `json:"ttf,omitempty"`
}

func (s *Term) UnmarshalJSON(data []byte) error {

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

		case "doc_freq":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.DocFreq = &value
			case float64:
				f := int(v)
				s.DocFreq = &f
			}

		case "score":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Score = &f
			case float64:
				f := Float64(v)
				s.Score = &f
			}

		case "term_freq":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.TermFreq = value
			case float64:
				f := int(v)
				s.TermFreq = f
			}

		case "tokens":
			if err := dec.Decode(&s.Tokens); err != nil {
				return err
			}

		case "ttf":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Ttf = &value
			case float64:
				f := int(v)
				s.Ttf = &f
			}

		}
	}
	return nil
}

// NewTerm returns a Term.
func NewTerm() *Term {
	r := &Term{}

	return r
}
