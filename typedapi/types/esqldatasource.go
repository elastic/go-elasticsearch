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
// https://github.com/elastic/elasticsearch-specification/tree/7560c979602e6941815872bdaec801200bc7ec4e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Represents a data source definition stored in cluster state. A data source
// holds connection settings (credentials, endpoints, auth) for an external data
// provider.
//
// https://github.com/elastic/elasticsearch-specification/blob/7560c979602e6941815872bdaec801200bc7ec4e/specification/esql/_types/types.ts#L132-L145
type ESQLDataSource struct {
	// Description A free-text description.
	Description *string `json:"description,omitempty"`
	// Name The data source name.
	Name string `json:"name"`
	// Settings Type-specific connection and authentication settings.
	Settings map[string]json.RawMessage `json:"settings"`
	// Type The data source type. Currently, `s3` is supported.
	Type string `json:"type"`
}

func (s *ESQLDataSource) UnmarshalJSON(data []byte) error {

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

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "settings":
			if s.Settings == nil {
				s.Settings = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Settings); err != nil {
				return fmt.Errorf("%s | %w", "Settings", err)
			}

		case "type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Type = o

		}
	}
	return nil
}

// NewESQLDataSource returns a ESQLDataSource.
func NewESQLDataSource() *ESQLDataSource {
	r := &ESQLDataSource{
		Settings: make(map[string]json.RawMessage),
	}

	return r
}
