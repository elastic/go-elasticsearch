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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package mget

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package mget
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/_global/mget/MultiGetResponse.ts#L22-L26
type Response struct {
	Docs []types.MgetResponseItem `json:"docs"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}

func (s *Response) UnmarshalJSON(data []byte) error {
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

		case "docs":
			messageArray := []json.RawMessage{}
			if err := dec.Decode(&messageArray); err != nil {
				return fmt.Errorf("%s | %w", "Docs", err)
			}
		docs:
			for _, message := range messageArray {
				keyDec := json.NewDecoder(bytes.NewReader(message))
				for {
					t, err := keyDec.Token()
					if err != nil {
						if errors.Is(err, io.EOF) {
							break
						}
						return fmt.Errorf("%s | %w", "Docs", err)
					}

					switch t {

					case "fields", "found", "_primary_term", "_routing", "_seq_no", "_source", "_version":
						o := types.NewGetResult()
						localDec := json.NewDecoder(bytes.NewReader(message))
						if err := localDec.Decode(&o); err != nil {
							return fmt.Errorf("%s | %w", "Docs", err)
						}
						s.Docs = append(s.Docs, o)
						continue docs

					case "error":
						o := types.NewMultiGetError()
						localDec := json.NewDecoder(bytes.NewReader(message))
						if err := localDec.Decode(&o); err != nil {
							return fmt.Errorf("%s | %w", "Docs", err)
						}
						s.Docs = append(s.Docs, o)
						continue docs

					}
				}
			}

		}
	}
	return nil
}
