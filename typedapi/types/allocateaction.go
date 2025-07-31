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

// AllocateAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ilm/_types/Phase.ts#L133-L139
type AllocateAction struct {
	Exclude            map[string]string `json:"exclude,omitempty"`
	Include            map[string]string `json:"include,omitempty"`
	NumberOfReplicas   *int              `json:"number_of_replicas,omitempty"`
	Require            map[string]string `json:"require,omitempty"`
	TotalShardsPerNode *int              `json:"total_shards_per_node,omitempty"`
}

func (s *AllocateAction) UnmarshalJSON(data []byte) error {

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

		case "exclude":
			if s.Exclude == nil {
				s.Exclude = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Exclude); err != nil {
				return fmt.Errorf("%s | %w", "Exclude", err)
			}

		case "include":
			if s.Include == nil {
				s.Include = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Include); err != nil {
				return fmt.Errorf("%s | %w", "Include", err)
			}

		case "number_of_replicas":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumberOfReplicas", err)
				}
				s.NumberOfReplicas = &value
			case float64:
				f := int(v)
				s.NumberOfReplicas = &f
			}

		case "require":
			if s.Require == nil {
				s.Require = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Require); err != nil {
				return fmt.Errorf("%s | %w", "Require", err)
			}

		case "total_shards_per_node":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalShardsPerNode", err)
				}
				s.TotalShardsPerNode = &value
			case float64:
				f := int(v)
				s.TotalShardsPerNode = &f
			}

		}
	}
	return nil
}

// NewAllocateAction returns a AllocateAction.
func NewAllocateAction() *AllocateAction {
	r := &AllocateAction{
		Exclude: make(map[string]string),
		Include: make(map[string]string),
		Require: make(map[string]string),
	}

	return r
}
