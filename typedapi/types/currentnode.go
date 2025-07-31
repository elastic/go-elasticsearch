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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/noderole"
)

// CurrentNode type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/allocation_explain/types.ts#L79-L90
type CurrentNode struct {
	Attributes       map[string]string   `json:"attributes"`
	Id               string              `json:"id"`
	Name             string              `json:"name"`
	Roles            []noderole.NodeRole `json:"roles"`
	TransportAddress string              `json:"transport_address"`
	WeightRanking    int                 `json:"weight_ranking"`
}

func (s *CurrentNode) UnmarshalJSON(data []byte) error {

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

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "roles":
			if err := dec.Decode(&s.Roles); err != nil {
				return fmt.Errorf("%s | %w", "Roles", err)
			}

		case "transport_address":
			if err := dec.Decode(&s.TransportAddress); err != nil {
				return fmt.Errorf("%s | %w", "TransportAddress", err)
			}

		case "weight_ranking":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "WeightRanking", err)
				}
				s.WeightRanking = value
			case float64:
				f := int(v)
				s.WeightRanking = f
			}

		}
	}
	return nil
}

// NewCurrentNode returns a CurrentNode.
func NewCurrentNode() *CurrentNode {
	r := &CurrentNode{
		Attributes: make(map[string]string),
	}

	return r
}
