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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/embeddingcontentformat"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/embeddingcontenttype"
)

// EmbeddingContentObjectContents type.
//
// https://github.com/elastic/elasticsearch-specification/blob/2514615770f18dbb4e3887cc1a279995dbfd0724/specification/inference/_types/CommonTypes.ts#L385-L401
type EmbeddingContentObjectContents struct {
	// Format The format of the input. For the `text` type this defaults to `text`. For the
	// `image` type, this defaults to `base64`.
	Format *embeddingcontentformat.EmbeddingContentFormat `json:"format,omitempty"`
	// Type The type of input to embed.
	Type embeddingcontenttype.EmbeddingContentType `json:"type"`
	// Value The value of the input to embed.
	Value string `json:"value"`
}

func (s *EmbeddingContentObjectContents) UnmarshalJSON(data []byte) error {

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

		case "format":
			if err := dec.Decode(&s.Format); err != nil {
				return fmt.Errorf("%s | %w", "Format", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Value", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Value = o

		}
	}
	return nil
}

// NewEmbeddingContentObjectContents returns a EmbeddingContentObjectContents.
func NewEmbeddingContentObjectContents() *EmbeddingContentObjectContents {
	r := &EmbeddingContentObjectContents{}

	return r
}

type EmbeddingContentObjectContentsVariant interface {
	EmbeddingContentObjectContentsCaster() *EmbeddingContentObjectContents
}

func (s *EmbeddingContentObjectContents) EmbeddingContentObjectContentsCaster() *EmbeddingContentObjectContents {
	return s
}
