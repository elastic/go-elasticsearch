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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Status type.
//
// https://github.com/elastic/elasticsearch-specification/blob/cbfcc73d01310bed2a480ec35aaef98138b598e5/specification/snapshot/_types/SnapshotStatus.ts#L26-L60
type Status struct {
	// IncludeGlobalState Indicates whether the current cluster state is included in the snapshot.
	IncludeGlobalState bool                          `json:"include_global_state"`
	Indices            map[string]SnapshotIndexStats `json:"indices"`
	// Repository The name of the repository that includes the snapshot.
	Repository string `json:"repository"`
	// ShardsStats Statistics for the shards in the snapshot.
	ShardsStats SnapshotShardsStats `json:"shards_stats"`
	// Snapshot The name of the snapshot.
	Snapshot string `json:"snapshot"`
	// State The current snapshot state:
	//
	// * `FAILED`: The snapshot finished with an error and failed to store any data.
	// * `STARTED`: The snapshot is currently running.
	// * `SUCCESS`: The snapshot completed.
	State string `json:"state"`
	// Stats Details about the number (`file_count`) and size (`size_in_bytes`) of files
	// included in the snapshot.
	Stats SnapshotStats `json:"stats"`
	// Uuid The universally unique identifier (UUID) for the snapshot.
	Uuid string `json:"uuid"`
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
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IncludeGlobalState", err)
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
				return fmt.Errorf("%s | %w", "Indices", err)
			}

		case "repository":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Repository", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Repository = o

		case "shards_stats":
			if err := dec.Decode(&s.ShardsStats); err != nil {
				return fmt.Errorf("%s | %w", "ShardsStats", err)
			}

		case "snapshot":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Snapshot", err)
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

		case "uuid":
			if err := dec.Decode(&s.Uuid); err != nil {
				return fmt.Errorf("%s | %w", "Uuid", err)
			}

		}
	}
	return nil
}

// NewStatus returns a Status.
func NewStatus() *Status {
	r := &Status{
		Indices: make(map[string]SnapshotIndexStats),
	}

	return r
}

// false
