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

package updateconfiguration

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package updateconfiguration
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/connector/update_configuration/ConnectorUpdateConfigurationRequest.ts#L25-L55
type Request struct {
	Configuration types.ConnectorConfiguration `json:"configuration,omitempty"`
	Values        map[string]json.RawMessage   `json:"values,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Values: make(map[string]json.RawMessage, 0),
	}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Updateconfiguration request: %w", err)
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

		case "configuration":
			if err := dec.Decode(&s.Configuration); err != nil {
				return fmt.Errorf("%s | %w", "Configuration", err)
			}

		case "values":
			if s.Values == nil {
				s.Values = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Values); err != nil {
				return fmt.Errorf("%s | %w", "Values", err)
			}

		}
	}
	return nil
}
