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
// https://github.com/elastic/elasticsearch-specification/tree/9a0362eb2579c6604966a8fb307caee92de04270

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// NodeInfoXpackSecurityAuthcToken type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9a0362eb2579c6604966a8fb307caee92de04270/specification/nodes/info/types.ts#L264-L266
type NodeInfoXpackSecurityAuthcToken struct {
	Enabled string `json:"enabled"`
}

func (s *NodeInfoXpackSecurityAuthcToken) UnmarshalJSON(data []byte) error {

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

		}
	}
	return nil
}

// NewNodeInfoXpackSecurityAuthcToken returns a NodeInfoXpackSecurityAuthcToken.
func NewNodeInfoXpackSecurityAuthcToken() *NodeInfoXpackSecurityAuthcToken {
	r := &NodeInfoXpackSecurityAuthcToken{}

	return r
}
