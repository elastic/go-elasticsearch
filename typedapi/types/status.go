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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// Status type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/snapshot/_types/SnapshotStatus.ts#L26-L35
type Status struct {
	IncludeGlobalState bool                          `json:"include_global_state"`
	Indices            map[string]SnapshotIndexStats `json:"indices"`
	Repository         string                        `json:"repository"`
	ShardsStats        SnapshotShardsStats           `json:"shards_stats"`
	Snapshot           string                        `json:"snapshot"`
	State              string                        `json:"state"`
	Stats              SnapshotStats                 `json:"stats"`
	Uuid               string                        `json:"uuid"`
}

func (s *Status) UnmarshalJSON(data []byte) error {

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

		case "include_global_state":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IncludeGlobalState = value
			case bool:
				s.IncludeGlobalState = v
			}

		case "indices":
			if s.Indices == nil {
				s.Indices = make(map[string]SnapshotIndexStats, 0)
			}
			if err := dec.Decode(&s.Indices); err != nil {
				return err
			}

		case "repository":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Repository = o

		case "shards_stats":
			if err := dec.Decode(&s.ShardsStats); err != nil {
				return err
			}

		case "snapshot":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Snapshot = o

		case "state":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.State = o

		case "stats":
			if err := dec.Decode(&s.Stats); err != nil {
				return err
			}

		case "uuid":
			if err := dec.Decode(&s.Uuid); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewStatus returns a Status.
func NewStatus() *Status {
	r := &Status{
		Indices: make(map[string]SnapshotIndexStats, 0),
	}

	return r
}
