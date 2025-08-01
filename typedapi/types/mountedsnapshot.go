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

// MountedSnapshot type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/searchable_snapshots/mount/types.ts#L23-L27
type MountedSnapshot struct {
	Indices  []string        `json:"indices"`
	Shards   ShardStatistics `json:"shards"`
	Snapshot string          `json:"snapshot"`
}

func (s *MountedSnapshot) UnmarshalJSON(data []byte) error {

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

		case "indices":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Indices", err)
				}

				s.Indices = append(s.Indices, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Indices); err != nil {
					return fmt.Errorf("%s | %w", "Indices", err)
				}
			}

		case "shards":
			if err := dec.Decode(&s.Shards); err != nil {
				return fmt.Errorf("%s | %w", "Shards", err)
			}

		case "snapshot":
			if err := dec.Decode(&s.Snapshot); err != nil {
				return fmt.Errorf("%s | %w", "Snapshot", err)
			}

		}
	}
	return nil
}

// NewMountedSnapshot returns a MountedSnapshot.
func NewMountedSnapshot() *MountedSnapshot {
	r := &MountedSnapshot{}

	return r
}
