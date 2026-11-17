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
// https://github.com/elastic/elasticsearch-specification/tree/1cdb0310e2ed76dea2491f2551fca3f9143b6dc4

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// TextContent type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1cdb0310e2ed76dea2491f2551fca3f9143b6dc4/specification/inference/_types/CommonTypes.ts#L167-L173
type TextContent struct {
	// Text The text content.
	Text string `json:"text"`
	Type string `json:"type,omitempty"`
}

func (s *TextContent) UnmarshalJSON(data []byte) error {

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

// MarshalJSON override marshalling to include literal value
func (s TextContent) MarshalJSON() ([]byte, error) {
	type innerTextContent TextContent
	tmp := innerTextContent{
		Text: s.Text,
		Type: s.Type,
	}

	tmp.Type = "text"

	return json.Marshal(tmp)
}

// NewTextContent returns a TextContent.
func NewTextContent() *TextContent {
	r := &TextContent{}

	return r
}

type TextContentVariant interface {
	TextContentCaster() *TextContent
}

func (s *TextContent) TextContentCaster() *TextContent {
	return s
}

func (s *TextContent) ContentObjectCaster() *ContentObject {
	if s == nil {
		return nil
	}
	o := ContentObject(s)
	return &o
}
