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
// https://github.com/elastic/elasticsearch-specification/tree/37285cbd3fd155f913b50d880b40ec45f9df64b3

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// A wrapper object which contains the fields required to specify multimodal
// inputs
//
// https://github.com/elastic/elasticsearch-specification/blob/37285cbd3fd155f913b50d880b40ec45f9df64b3/specification/inference/_types/CommonTypes.ts#L630-L638
type EmbeddingContentObject struct {
	// Content An object or an array of objects containing the input data for the model to
	// embed
	Content []EmbeddingContentObjectItem `json:"content"`
}

func (s *EmbeddingContentObject) UnmarshalJSON(data []byte) error {

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

		case "content":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewEmbeddingContentObjectItem()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Content", err)
				}

				s.Content = append(s.Content, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Content); err != nil {
					return fmt.Errorf("%s | %w", "Content", err)
				}
			}

		}
	}
	return nil
}

// NewEmbeddingContentObject returns a EmbeddingContentObject.
func NewEmbeddingContentObject() *EmbeddingContentObject {
	r := &EmbeddingContentObject{}

	return r
}

type EmbeddingContentObjectVariant interface {
	EmbeddingContentObjectCaster() *EmbeddingContentObject
}

func (s *EmbeddingContentObject) EmbeddingContentObjectCaster() *EmbeddingContentObject {
	return s
}

func (s *EmbeddingContentObject) EmbeddingContentInputCaster() *EmbeddingContentInput {
	if s == nil {
		return nil
	}
	o := EmbeddingContentInput{*s}
	return &o
}
