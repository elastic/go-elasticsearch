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

// NodeInfoDiscover type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/nodes/info/types.ts#L183-L192
type NodeInfoDiscover struct {
	NodeInfoDiscover map[string]json.RawMessage `json:"-"`
	SeedHosts        []string                   `json:"seed_hosts,omitempty"`
	SeedProviders    []string                   `json:"seed_providers,omitempty"`
	Type             *string                    `json:"type,omitempty"`
}

func (s *NodeInfoDiscover) UnmarshalJSON(data []byte) error {

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

		case "seed_hosts":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "SeedHosts", err)
				}

				s.SeedHosts = append(s.SeedHosts, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.SeedHosts); err != nil {
					return fmt.Errorf("%s | %w", "SeedHosts", err)
				}
			}

		case "seed_providers":
			if err := dec.Decode(&s.SeedProviders); err != nil {
				return fmt.Errorf("%s | %w", "SeedProviders", err)
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
			s.Type = &o

		default:

			if key, ok := t.(string); ok {
				if s.NodeInfoDiscover == nil {
					s.NodeInfoDiscover = make(map[string]json.RawMessage, 0)
				}
				raw := new(json.RawMessage)
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "NodeInfoDiscover", err)
				}
				s.NodeInfoDiscover[key] = *raw
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s NodeInfoDiscover) MarshalJSON() ([]byte, error) {
	type opt NodeInfoDiscover
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.NodeInfoDiscover {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "NodeInfoDiscover")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewNodeInfoDiscover returns a NodeInfoDiscover.
func NewNodeInfoDiscover() *NodeInfoDiscover {
	r := &NodeInfoDiscover{
		NodeInfoDiscover: make(map[string]json.RawMessage),
	}

	return r
}
