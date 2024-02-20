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
// https://github.com/elastic/elasticsearch-specification/tree/6e0fb6b929f337b62bf0676bdf503e061121fad2

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// IndicesBlockStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/6e0fb6b929f337b62bf0676bdf503e061121fad2/specification/indices/add_block/IndicesAddBlockResponse.ts#L30-L33
type IndicesBlockStatus struct {
	Blocked bool   `json:"blocked"`
	Name    string `json:"name"`
}

func (s *IndicesBlockStatus) UnmarshalJSON(data []byte) error {

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

		case "blocked":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Blocked = value
			case bool:
				s.Blocked = v
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewIndicesBlockStatus returns a IndicesBlockStatus.
func NewIndicesBlockStatus() *IndicesBlockStatus {
	r := &IndicesBlockStatus{}

	return r
}
