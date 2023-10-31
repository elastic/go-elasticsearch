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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// CommandAllocatePrimaryAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cluster/reroute/types.ts#L78-L84
type CommandAllocatePrimaryAction struct {
	// AcceptDataLoss If a node which has a copy of the data rejoins the cluster later on, that
	// data will be deleted. To ensure that these implications are well-understood,
	// this command requires the flag accept_data_loss to be explicitly set to true
	AcceptDataLoss bool   `json:"accept_data_loss"`
	Index          string `json:"index"`
	Node           string `json:"node"`
	Shard          int    `json:"shard"`
}

func (s *CommandAllocatePrimaryAction) UnmarshalJSON(data []byte) error {

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

		case "accept_data_loss":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AcceptDataLoss = value
			case bool:
				s.AcceptDataLoss = v
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "node":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Node = o

		case "shard":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Shard = value
			case float64:
				f := int(v)
				s.Shard = f
			}

		}
	}
	return nil
}

// NewCommandAllocatePrimaryAction returns a CommandAllocatePrimaryAction.
func NewCommandAllocatePrimaryAction() *CommandAllocatePrimaryAction {
	r := &CommandAllocatePrimaryAction{}

	return r
}
