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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// IndexSegment type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/indices/segments/types.ts#L24-L26
type IndexSegment struct {
	Shards map[string][]ShardsSegment `json:"shards"`
}

func (s *IndexSegment) UnmarshalJSON(data []byte) error {

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

		case "shards":
			if s.Shards == nil {
				s.Shards = make(map[string][]ShardsSegment, 0)
			}
			rawMsg := make(map[string]json.RawMessage, 0)
			dec.Decode(&rawMsg)
			for key, value := range rawMsg {
				switch {
				case bytes.HasPrefix(value, []byte("\"")), bytes.HasPrefix(value, []byte("{")):
					o := NewShardsSegment()
					err := json.NewDecoder(bytes.NewReader(value)).Decode(&o)
					if err != nil {
						return err
					}
					s.Shards[key] = append(s.Shards[key], *o)
				default:
					o := []ShardsSegment{}
					err := json.NewDecoder(bytes.NewReader(value)).Decode(&o)
					if err != nil {
						return err
					}
					s.Shards[key] = o
				}
			}

		}
	}
	return nil
}

// NewIndexSegment returns a IndexSegment.
func NewIndexSegment() *IndexSegment {
	r := &IndexSegment{
		Shards: make(map[string][]ShardsSegment, 0),
	}

	return r
}
