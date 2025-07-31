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
)

// NodeInfoTransport type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/nodes/info/types.ts#L362-L366
type NodeInfoTransport struct {
	BoundAddress   []string          `json:"bound_address"`
	Profiles       map[string]string `json:"profiles"`
	PublishAddress string            `json:"publish_address"`
}

func (s *NodeInfoTransport) UnmarshalJSON(data []byte) error {

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

		case "bound_address":
			if err := dec.Decode(&s.BoundAddress); err != nil {
				return fmt.Errorf("%s | %w", "BoundAddress", err)
			}

		case "profiles":
			if s.Profiles == nil {
				s.Profiles = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Profiles); err != nil {
				return fmt.Errorf("%s | %w", "Profiles", err)
			}

		case "publish_address":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "PublishAddress", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PublishAddress = o

		}
	}
	return nil
}

// NewNodeInfoTransport returns a NodeInfoTransport.
func NewNodeInfoTransport() *NodeInfoTransport {
	r := &NodeInfoTransport{
		Profiles: make(map[string]string),
	}

	return r
}
