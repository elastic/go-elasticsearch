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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// NodeAttributes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/Node.ts#L41-L52
type NodeAttributes struct {
	// Attributes Lists node attributes.
	Attributes map[string]string `json:"attributes"`
	// EphemeralId The ephemeral ID of the node.
	EphemeralId string `json:"ephemeral_id"`
	// Id The unique identifier of the node.
	Id *string `json:"id,omitempty"`
	// Name The unique identifier of the node.
	Name string `json:"name"`
	// TransportAddress The host and port where transport HTTP connections are accepted.
	TransportAddress string `json:"transport_address"`
}

func (s *NodeAttributes) UnmarshalJSON(data []byte) error {

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

		case "attributes":
			if s.Attributes == nil {
				s.Attributes = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Attributes); err != nil {
				return fmt.Errorf("%s | %w", "Attributes", err)
			}

		case "ephemeral_id":
			if err := dec.Decode(&s.EphemeralId); err != nil {
				return fmt.Errorf("%s | %w", "EphemeralId", err)
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "transport_address":
			if err := dec.Decode(&s.TransportAddress); err != nil {
				return fmt.Errorf("%s | %w", "TransportAddress", err)
			}

		}
	}
	return nil
}

// NewNodeAttributes returns a NodeAttributes.
func NewNodeAttributes() *NodeAttributes {
	r := &NodeAttributes{
		Attributes: make(map[string]string),
	}

	return r
}
