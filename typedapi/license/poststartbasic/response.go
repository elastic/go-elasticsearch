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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package poststartbasic

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/licensetype"
)

// Response holds the response body struct for the package poststartbasic
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/license/post_start_basic/StartBasicLicenseResponse.ts#L23-L31
type Response struct {
	Acknowledge     map[string][]string      `json:"acknowledge,omitempty"`
	Acknowledged    bool                     `json:"acknowledged"`
	BasicWasStarted bool                     `json:"basic_was_started"`
	ErrorMessage    *string                  `json:"error_message,omitempty"`
	Type            *licensetype.LicenseType `json:"type,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Acknowledge: make(map[string][]string, 0),
	}
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

		case "acknowledge":
			if s.Acknowledge == nil {
				s.Acknowledge = make(map[string][]string, 0)
			}
			rawMsg := make(map[string]json.RawMessage, 0)
			dec.Decode(&rawMsg)
			for key, value := range rawMsg {
				switch {
				case bytes.HasPrefix(value, []byte("\"")), bytes.HasPrefix(value, []byte("{")):
					o := new(string)
					err := json.NewDecoder(bytes.NewReader(value)).Decode(&o)
					if err != nil {
						return fmt.Errorf("%s | %w", "Acknowledge", err)
					}
					s.Acknowledge[key] = append(s.Acknowledge[key], *o)
				default:
					o := []string{}
					err := json.NewDecoder(bytes.NewReader(value)).Decode(&o)
					if err != nil {
						return fmt.Errorf("%s | %w", "Acknowledge", err)
					}
					s.Acknowledge[key] = o
				}
			}

		case "acknowledged":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Acknowledged", err)
				}
				s.Acknowledged = value
			case bool:
				s.Acknowledged = v
			}

		case "basic_was_started":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "BasicWasStarted", err)
				}
				s.BasicWasStarted = value
			case bool:
				s.BasicWasStarted = v
			}

		case "error_message":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ErrorMessage", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ErrorMessage = &o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		}
	}
	return nil
}
