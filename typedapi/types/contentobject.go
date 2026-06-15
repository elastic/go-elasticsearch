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
// https://github.com/elastic/elasticsearch-specification/tree/eb2e22fb2ac404e676d19bcc7bb089647f029026

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/contenttype"
)

// An object style representation of a single portion of a conversation.
//
// https://github.com/elastic/elasticsearch-specification/blob/eb2e22fb2ac404e676d19bcc7bb089647f029026/specification/inference/_types/CommonTypes.ts#L159-L180
type ContentObject struct {
	// File The file content. Only applicable for the `file` type
	File FileContent `json:"file"`
	// ImageUrl The image content. Only applicable for the `image_url` type
	ImageUrl ImageUrl `json:"image_url"`
	// Text The text content. Only applicable for the `text` type
	Text string `json:"text"`
	// Type The type of content. Must be one of `text`, `image_url` or `file`. Not all
	// services/models support content types other than "text"
	Type contenttype.ContentType `json:"type"`
}

func (s *ContentObject) UnmarshalJSON(data []byte) error {

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

		case "file":
			if err := dec.Decode(&s.File); err != nil {
				return fmt.Errorf("%s | %w", "File", err)
			}

		case "image_url":
			if err := dec.Decode(&s.ImageUrl); err != nil {
				return fmt.Errorf("%s | %w", "ImageUrl", err)
			}

		case "text":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Text", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Text = o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		}
	}
	return nil
}

// NewContentObject returns a ContentObject.
func NewContentObject() *ContentObject {
	r := &ContentObject{}

	return r
}

type ContentObjectVariant interface {
	ContentObjectCaster() *ContentObject
}

func (s *ContentObject) ContentObjectCaster() *ContentObject {
	return s
}
