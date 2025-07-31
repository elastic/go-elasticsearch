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

// ShardProfile type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/search/_types/profile.ts#L142-L152
type ShardProfile struct {
	Aggregations []AggregationProfile `json:"aggregations"`
	Cluster      string               `json:"cluster"`
	Dfs          *DfsProfile          `json:"dfs,omitempty"`
	Fetch        *FetchProfile        `json:"fetch,omitempty"`
	Id           string               `json:"id"`
	Index        string               `json:"index"`
	NodeId       string               `json:"node_id"`
	Searches     []SearchProfile      `json:"searches"`
	ShardId      int                  `json:"shard_id"`
}

func (s *ShardProfile) UnmarshalJSON(data []byte) error {

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

		case "aggregations":
			if err := dec.Decode(&s.Aggregations); err != nil {
				return fmt.Errorf("%s | %w", "Aggregations", err)
			}

		case "cluster":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Cluster", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Cluster = o

		case "dfs":
			if err := dec.Decode(&s.Dfs); err != nil {
				return fmt.Errorf("%s | %w", "Dfs", err)
			}

		case "fetch":
			if err := dec.Decode(&s.Fetch); err != nil {
				return fmt.Errorf("%s | %w", "Fetch", err)
			}

		case "id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Id = o

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "node_id":
			if err := dec.Decode(&s.NodeId); err != nil {
				return fmt.Errorf("%s | %w", "NodeId", err)
			}

		case "searches":
			if err := dec.Decode(&s.Searches); err != nil {
				return fmt.Errorf("%s | %w", "Searches", err)
			}

		case "shard_id":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShardId", err)
				}
				s.ShardId = value
			case float64:
				f := int(v)
				s.ShardId = f
			}

		}
	}
	return nil
}

// NewShardProfile returns a ShardProfile.
func NewShardProfile() *ShardProfile {
	r := &ShardProfile{}

	return r
}
