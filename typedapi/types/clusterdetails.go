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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/clustersearchstatus"
)

// ClusterDetails type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/Stats.ts#L45-L52
type ClusterDetails struct {
	Failures []ShardFailure                          `json:"failures,omitempty"`
	Indices  string                                  `json:"indices"`
	Shards_  *ShardStatistics                        `json:"_shards,omitempty"`
	Status   clustersearchstatus.ClusterSearchStatus `json:"status"`
	TimedOut bool                                    `json:"timed_out"`
	Took     *int64                                  `json:"took,omitempty"`
}

func (s *ClusterDetails) UnmarshalJSON(data []byte) error {

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

		case "failures":
			if err := dec.Decode(&s.Failures); err != nil {
				return err
			}

		case "indices":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Indices = o

		case "_shards":
			if err := dec.Decode(&s.Shards_); err != nil {
				return err
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return err
			}

		case "timed_out":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.TimedOut = value
			case bool:
				s.TimedOut = v
			}

		case "took":
			if err := dec.Decode(&s.Took); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewClusterDetails returns a ClusterDetails.
func NewClusterDetails() *ClusterDetails {
	r := &ClusterDetails{}

	return r
}
