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

package samlauthenticate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Request holds the request body struct for the package samlauthenticate
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/saml_authenticate/Request.ts#L23-L61
type Request struct {

	// Content The SAML response as it was sent by the user's browser, usually a Base64
	// encoded XML document.
	Content string `json:"content"`
	// Ids A JSON array with all the valid SAML Request Ids that the caller of the API
	// has for the current user.
	Ids []string `json:"ids"`
	// Realm The name of the realm that should authenticate the SAML response. Useful in
	// cases where many SAML realms are defined.
	Realm *string `json:"realm,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Samlauthenticate request: %w", err)
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

		case "content":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Content", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Content = o

		case "ids":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Ids", err)
				}

				s.Ids = append(s.Ids, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Ids); err != nil {
					return fmt.Errorf("%s | %w", "Ids", err)
				}
			}

		case "realm":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Realm", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Realm = &o

		}
	}
	return nil
}
