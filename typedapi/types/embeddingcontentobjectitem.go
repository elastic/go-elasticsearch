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
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/embeddingcontentformat"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/embeddingcontenttype"
)

// An object containing the input data for a single item for the model to embed.
//
// https://github.com/elastic/elasticsearch-specification/blob/37285cbd3fd155f913b50d880b40ec45f9df64b3/specification/inference/_types/CommonTypes.ts#L647-L664
type EmbeddingContentObjectItem struct {
	// Format The format of the input. For the `text` type this must be `text`. For all
	// other types, this must be `base64`. If not specified, this will default to
	// `text` for the `text` type and `base64` for all other types.
	Format *embeddingcontentformat.EmbeddingContentFormat `json:"format,omitempty"`
	// Type The type of input to embed. Not all models support all input types.
	Type embeddingcontenttype.EmbeddingContentType `json:"type"`
	// Value The value of the input to embed. For images, this must be a base64-encoded
	// data URI, i.e. "data:content/type;base64,..."
	Value string `json:"value"`
}

func (s *EmbeddingContentObjectItem) UnmarshalJSON(data []byte) error {

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

// NewEmbeddingContentObjectItem returns a EmbeddingContentObjectItem.
func NewEmbeddingContentObjectItem() *EmbeddingContentObjectItem {
	r := &EmbeddingContentObjectItem{}

	return r
}

type EmbeddingContentObjectItemVariant interface {
	EmbeddingContentObjectItemCaster() *EmbeddingContentObjectItem
}

func (s *EmbeddingContentObjectItem) EmbeddingContentObjectItemCaster() *EmbeddingContentObjectItem {
	return s
}

func (s *EmbeddingContentObjectItem) EmbeddingContentObjectGroupCaster() *EmbeddingContentObjectGroup {
	if s == nil {
		return nil
	}
	o := EmbeddingContentObjectGroup{*s}
	return &o
}
