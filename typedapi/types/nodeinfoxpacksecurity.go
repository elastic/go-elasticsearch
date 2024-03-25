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
// https://github.com/elastic/elasticsearch-specification/tree/5fb8f1ce9c4605abcaa44aa0f17dbfc60497a757

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// NodeInfoXpackSecurity type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fb8f1ce9c4605abcaa44aa0f17dbfc60497a757/specification/nodes/info/types.ts#L242-L247
type NodeInfoXpackSecurity struct {
	Authc     *NodeInfoXpackSecurityAuthc `json:"authc,omitempty"`
	Enabled   string                      `json:"enabled"`
	Http      NodeInfoXpackSecuritySsl    `json:"http"`
	Transport *NodeInfoXpackSecuritySsl   `json:"transport,omitempty"`
}

func (s *NodeInfoXpackSecurity) UnmarshalJSON(data []byte) error {

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

		case "authc":
			if err := dec.Decode(&s.Authc); err != nil {
				return fmt.Errorf("%s | %w", "Authc", err)
			}

		case "enabled":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Enabled", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Enabled = o

		case "http":
			if err := dec.Decode(&s.Http); err != nil {
				return fmt.Errorf("%s | %w", "Http", err)
			}

		case "transport":
			if err := dec.Decode(&s.Transport); err != nil {
				return fmt.Errorf("%s | %w", "Transport", err)
			}

		}
	}
	return nil
}

// NewNodeInfoXpackSecurity returns a NodeInfoXpackSecurity.
func NewNodeInfoXpackSecurity() *NodeInfoXpackSecurity {
	r := &NodeInfoXpackSecurity{}

	return r
}
