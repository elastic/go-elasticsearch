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

// ShrinkAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ilm/_types/Phase.ts#L117-L121
type ShrinkAction struct {
	AllowWriteAfterShrink *bool    `json:"allow_write_after_shrink,omitempty"`
	MaxPrimaryShardSize   ByteSize `json:"max_primary_shard_size,omitempty"`
	NumberOfShards        *int     `json:"number_of_shards,omitempty"`
}

func (s *ShrinkAction) UnmarshalJSON(data []byte) error {

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

		case "allow_write_after_shrink":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllowWriteAfterShrink", err)
				}
				s.AllowWriteAfterShrink = &value
			case bool:
				s.AllowWriteAfterShrink = &v
			}

		case "max_primary_shard_size":
			if err := dec.Decode(&s.MaxPrimaryShardSize); err != nil {
				return fmt.Errorf("%s | %w", "MaxPrimaryShardSize", err)
			}

		case "number_of_shards":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumberOfShards", err)
				}
				s.NumberOfShards = &value
			case float64:
				f := int(v)
				s.NumberOfShards = &f
			}

		}
	}
	return nil
}

// NewShrinkAction returns a ShrinkAction.
func NewShrinkAction() *ShrinkAction {
	r := &ShrinkAction{}

	return r
}
