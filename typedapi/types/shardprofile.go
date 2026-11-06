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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

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
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_global/search/_types/profile.ts#L239-L255
type ShardProfile struct {
	Aggregations []AggregationProfile `json:"aggregations"`
	Cluster      string               `json:"cluster"`
	Dfs          *DfsProfile          `json:"dfs,omitempty"`
	Fetch        *FetchProfile        `json:"fetch,omitempty"`
	Id           string               `json:"id"`
	Index        string               `json:"index"`
	NodeId       string               `json:"node_id"`
	// Rescore One entry per rescorer, in the order they ran. Omitted when the request has
	// no rescorer.
	Rescore  []RescoreProfile `json:"rescore,omitempty"`
	Searches []SearchProfile  `json:"searches"`
	ShardId  int              `json:"shard_id"`
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

		case "rescore":

			buf := []json.RawMessage{}
			dec.Decode(&buf)
			for _, rawMsg := range buf {

				source := bytes.NewReader(rawMsg)
				localDec := json.NewDecoder(source)
				kind := make(map[string]string, 0)
				localDec.Decode(&kind)
				source.Seek(0, io.SeekStart)

				switch kind["type"] {

				case "query":
					o := NewQueryRescoreProfile()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "query", err)
					}
					s.Rescore = append(s.Rescore, *o)
				case "script":
					o := NewScriptRescoreProfile()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "script", err)
					}
					s.Rescore = append(s.Rescore, *o)
				case "learning_to_rank":
					o := NewLearningToRankRescoreProfile()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "learning_to_rank", err)
					}
					s.Rescore = append(s.Rescore, *o)
				default:
					o := new(any)
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("Rescore | %w", err)
					}
					s.Rescore = append(s.Rescore, *o)
				}
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
