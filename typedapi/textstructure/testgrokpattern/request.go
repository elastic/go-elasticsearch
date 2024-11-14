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
// https://github.com/elastic/elasticsearch-specification/tree/4fcf747dfafc951e1dcf3077327e3dcee9107db3

package testgrokpattern

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Request holds the request body struct for the package testgrokpattern
//
// https://github.com/elastic/elasticsearch-specification/blob/4fcf747dfafc951e1dcf3077327e3dcee9107db3/specification/text_structure/test_grok_pattern/TestGrokPatternRequest.ts#L23-L44
type Request struct {

	// GrokPattern Grok pattern to run on the text.
	GrokPattern string `json:"grok_pattern"`
	// Text Lines of text to run the Grok pattern on.
	Text []string `json:"text"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Testgrokpattern request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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

		case "grok_pattern":
			if err := dec.Decode(&s.GrokPattern); err != nil {
				return fmt.Errorf("%s | %w", "GrokPattern", err)
			}

		case "text":
			if err := dec.Decode(&s.Text); err != nil {
				return fmt.Errorf("%s | %w", "Text", err)
			}

		}
	}
	return nil
}
