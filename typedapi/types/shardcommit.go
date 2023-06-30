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

// ShardCommit type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/indices/stats/types.ts#L103-L108
type ShardCommit struct {
	Generation int               `json:"generation"`
	Id         string            `json:"id"`
	NumDocs    int64             `json:"num_docs"`
	UserData   map[string]string `json:"user_data"`
}

func (s *ShardCommit) UnmarshalJSON(data []byte) error {

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

		case "generation":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Generation = value
			case float64:
				f := int(v)
				s.Generation = f
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return err
			}

		case "num_docs":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.NumDocs = value
			case float64:
				f := int64(v)
				s.NumDocs = f
			}

		case "user_data":
			if s.UserData == nil {
				s.UserData = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.UserData); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewShardCommit returns a ShardCommit.
func NewShardCommit() *ShardCommit {
	r := &ShardCommit{
		UserData: make(map[string]string, 0),
	}

	return r
}
