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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/connectionscheme"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/httpinputmethod"
)

// HttpInputRequestDefinition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/watcher/_types/Input.ts#L72-L86
type HttpInputRequestDefinition struct {
	Auth              *HttpInputAuthentication           `json:"auth,omitempty"`
	Body              *string                            `json:"body,omitempty"`
	ConnectionTimeout Duration                           `json:"connection_timeout,omitempty"`
	Headers           map[string]string                  `json:"headers,omitempty"`
	Host              *string                            `json:"host,omitempty"`
	Method            *httpinputmethod.HttpInputMethod   `json:"method,omitempty"`
	Params            map[string]string                  `json:"params,omitempty"`
	Path              *string                            `json:"path,omitempty"`
	Port              *uint                              `json:"port,omitempty"`
	Proxy             *HttpInputProxy                    `json:"proxy,omitempty"`
	ReadTimeout       Duration                           `json:"read_timeout,omitempty"`
	Scheme            *connectionscheme.ConnectionScheme `json:"scheme,omitempty"`
	Url               *string                            `json:"url,omitempty"`
}

func (s *HttpInputRequestDefinition) UnmarshalJSON(data []byte) error {

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

		case "auth":
			if err := dec.Decode(&s.Auth); err != nil {
				return fmt.Errorf("%s | %w", "Auth", err)
			}

		case "body":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Body", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Body = &o

		case "connection_timeout":
			if err := dec.Decode(&s.ConnectionTimeout); err != nil {
				return fmt.Errorf("%s | %w", "ConnectionTimeout", err)
			}

		case "headers":
			if s.Headers == nil {
				s.Headers = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Headers); err != nil {
				return fmt.Errorf("%s | %w", "Headers", err)
			}

		case "host":
			if err := dec.Decode(&s.Host); err != nil {
				return fmt.Errorf("%s | %w", "Host", err)
			}

		case "method":
			if err := dec.Decode(&s.Method); err != nil {
				return fmt.Errorf("%s | %w", "Method", err)
			}

		case "params":
			if s.Params == nil {
				s.Params = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Params); err != nil {
				return fmt.Errorf("%s | %w", "Params", err)
			}

		case "path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Path", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Path = &o

		case "port":
			if err := dec.Decode(&s.Port); err != nil {
				return fmt.Errorf("%s | %w", "Port", err)
			}

		case "proxy":
			if err := dec.Decode(&s.Proxy); err != nil {
				return fmt.Errorf("%s | %w", "Proxy", err)
			}

		case "read_timeout":
			if err := dec.Decode(&s.ReadTimeout); err != nil {
				return fmt.Errorf("%s | %w", "ReadTimeout", err)
			}

		case "scheme":
			if err := dec.Decode(&s.Scheme); err != nil {
				return fmt.Errorf("%s | %w", "Scheme", err)
			}

		case "url":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Url", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Url = &o

		}
	}
	return nil
}

// NewHttpInputRequestDefinition returns a HttpInputRequestDefinition.
func NewHttpInputRequestDefinition() *HttpInputRequestDefinition {
	r := &HttpInputRequestDefinition{
		Headers: make(map[string]string),
		Params:  make(map[string]string),
	}

	return r
}
