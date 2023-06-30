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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// IndexDetails type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/snapshot/_types/SnapshotIndexDetails.ts#L23-L28
type IndexDetails struct {
	MaxSegmentsPerShard int64    `json:"max_segments_per_shard"`
	ShardCount          int      `json:"shard_count"`
	Size                ByteSize `json:"size,omitempty"`
	SizeInBytes         int64    `json:"size_in_bytes"`
}

func (s *IndexDetails) UnmarshalJSON(data []byte) error {

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

		case "max_segments_per_shard":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MaxSegmentsPerShard = value
			case float64:
				f := int64(v)
				s.MaxSegmentsPerShard = f
			}

		case "shard_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ShardCount = value
			case float64:
				f := int(v)
				s.ShardCount = f
			}

		case "size":
			if err := dec.Decode(&s.Size); err != nil {
				return err
			}

		case "size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.SizeInBytes = value
			case float64:
				f := int64(v)
				s.SizeInBytes = f
			}

		}
	}
	return nil
}

// NewIndexDetails returns a IndexDetails.
func NewIndexDetails() *IndexDetails {
	r := &IndexDetails{}

	return r
}
