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

// RerouteParameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/reroute/types.ts#L98-L105
type RerouteParameters struct {
	AllowPrimary bool    `json:"allow_primary"`
	FromNode     *string `json:"from_node,omitempty"`
	Index        string  `json:"index"`
	Node         string  `json:"node"`
	Shard        int     `json:"shard"`
	ToNode       *string `json:"to_node,omitempty"`
}

func (s *RerouteParameters) UnmarshalJSON(data []byte) error {

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

		case "allow_primary":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllowPrimary", err)
				}
				s.AllowPrimary = value
			case bool:
				s.AllowPrimary = v
			}

		case "from_node":
			if err := dec.Decode(&s.FromNode); err != nil {
				return fmt.Errorf("%s | %w", "FromNode", err)
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "node":
			if err := dec.Decode(&s.Node); err != nil {
				return fmt.Errorf("%s | %w", "Node", err)
			}

		case "shard":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Shard", err)
				}
				s.Shard = value
			case float64:
				f := int(v)
				s.Shard = f
			}

		case "to_node":
			if err := dec.Decode(&s.ToNode); err != nil {
				return fmt.Errorf("%s | %w", "ToNode", err)
			}

		}
	}
	return nil
}

// NewRerouteParameters returns a RerouteParameters.
func NewRerouteParameters() *RerouteParameters {
	r := &RerouteParameters{}

	return r
}
