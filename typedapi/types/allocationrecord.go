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

// AllocationRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cat/allocation/types.ts#L24-L75
type AllocationRecord struct {
	// DiskAvail Free disk space available to Elasticsearch.
	// Elasticsearch retrieves this metric from the node’s operating system.
	// Disk-based shard allocation uses this metric to assign shards to nodes based
	// on available disk space.
	DiskAvail ByteSize `json:"disk.avail,omitempty"`
	// DiskIndices Disk space used by the node’s shards. Does not include disk space for the
	// translog or unassigned shards.
	// IMPORTANT: This metric double-counts disk space for hard-linked files, such
	// as those created when shrinking, splitting, or cloning an index.
	DiskIndices ByteSize `json:"disk.indices,omitempty"`
	// DiskPercent Total percentage of disk space in use. Calculated as `disk.used /
	// disk.total`.
	DiskPercent Percentage `json:"disk.percent,omitempty"`
	// DiskTotal Total disk space for the node, including in-use and available space.
	DiskTotal ByteSize `json:"disk.total,omitempty"`
	// DiskUsed Total disk space in use.
	// Elasticsearch retrieves this metric from the node’s operating system (OS).
	// The metric includes disk space for: Elasticsearch, including the translog and
	// unassigned shards; the node’s operating system; any other applications or
	// files on the node.
	// Unlike `disk.indices`, this metric does not double-count disk space for
	// hard-linked files.
	DiskUsed ByteSize `json:"disk.used,omitempty"`
	// Host Network host for the node. Set using the `network.host` setting.
	Host string `json:"host,omitempty"`
	// Ip IP address and port for the node.
	Ip string `json:"ip,omitempty"`
	// Node Name for the node. Set using the `node.name` setting.
	Node *string `json:"node,omitempty"`
	// Shards Number of primary and replica shards assigned to the node.
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
