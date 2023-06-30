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

// AllocationRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/cat/allocation/types.ts#L24-L69
type AllocationRecord struct {
	// DiskAvail disk available
	DiskAvail ByteSize `json:"disk.avail,omitempty"`
	// DiskIndices disk used by ES indices
	DiskIndices ByteSize `json:"disk.indices,omitempty"`
	// DiskPercent percent disk used
	DiskPercent Percentage `json:"disk.percent,omitempty"`
	// DiskTotal total capacity of all volumes
	DiskTotal ByteSize `json:"disk.total,omitempty"`
	// DiskUsed disk used (total, not just ES)
	DiskUsed ByteSize `json:"disk.used,omitempty"`
	// Host host of node
	Host string `json:"host,omitempty"`
	// Ip ip of node
	Ip string `json:"ip,omitempty"`
	// Node name of node
	Node *string `json:"node,omitempty"`
	// Shards number of shards on node
	Shards *string `json:"shards,omitempty"`
}

func (s *AllocationRecord) UnmarshalJSON(data []byte) error {

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

		case "disk.avail", "da", "diskAvail":
			if err := dec.Decode(&s.DiskAvail); err != nil {
				return err
			}

		case "disk.indices", "di", "diskIndices":
			if err := dec.Decode(&s.DiskIndices); err != nil {
				return err
			}

		case "disk.percent", "dp", "diskPercent":
			if err := dec.Decode(&s.DiskPercent); err != nil {
				return err
			}

		case "disk.total", "dt", "diskTotal":
			if err := dec.Decode(&s.DiskTotal); err != nil {
				return err
			}

		case "disk.used", "du", "diskUsed":
			if err := dec.Decode(&s.DiskUsed); err != nil {
				return err
			}

		case "host", "h":
			if err := dec.Decode(&s.Host); err != nil {
				return err
			}

		case "ip":
			if err := dec.Decode(&s.Ip); err != nil {
				return err
			}

		case "node", "n":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Node = &o

		case "shards", "s":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Shards = &o

		}
	}
	return nil
}

// NewAllocationRecord returns a AllocationRecord.
func NewAllocationRecord() *AllocationRecord {
	r := &AllocationRecord{}

	return r
}
