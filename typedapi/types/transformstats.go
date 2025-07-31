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

// TransformStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/transform/get_transform_stats/types.ts#L31-L42
type TransformStats struct {
	Checkpointing Checkpointing         `json:"checkpointing"`
	Health        *TransformStatsHealth `json:"health,omitempty"`
	Id            string                `json:"id"`
	Node          *NodeAttributes       `json:"node,omitempty"`
	Reason        *string               `json:"reason,omitempty"`
	State         string                `json:"state"`
	Stats         TransformIndexerStats `json:"stats"`
}

func (s *TransformStats) UnmarshalJSON(data []byte) error {

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

		case "checkpointing":
			if err := dec.Decode(&s.Checkpointing); err != nil {
				return fmt.Errorf("%s | %w", "Checkpointing", err)
			}

		case "health":
			if err := dec.Decode(&s.Health); err != nil {
				return fmt.Errorf("%s | %w", "Health", err)
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "node":
			if err := dec.Decode(&s.Node); err != nil {
				return fmt.Errorf("%s | %w", "Node", err)
			}

		case "reason":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Reason", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Reason = &o

		case "state":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "State", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.State = o

		case "stats":
			if err := dec.Decode(&s.Stats); err != nil {
				return fmt.Errorf("%s | %w", "Stats", err)
			}

		}
	}
	return nil
}

// NewTransformStats returns a TransformStats.
func NewTransformStats() *TransformStats {
	r := &TransformStats{}

	return r
}
