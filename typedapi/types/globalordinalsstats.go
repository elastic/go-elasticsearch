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

// GlobalOrdinalsStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/Stats.ts#L131-L135
type GlobalOrdinalsStats struct {
	BuildTime         *string                            `json:"build_time,omitempty"`
	BuildTimeInMillis int64                              `json:"build_time_in_millis"`
	Fields            map[string]GlobalOrdinalFieldStats `json:"fields,omitempty"`
}

func (s *GlobalOrdinalsStats) UnmarshalJSON(data []byte) error {

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

		case "build_time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "BuildTime", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BuildTime = &o

		case "build_time_in_millis":
			if err := dec.Decode(&s.BuildTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "BuildTimeInMillis", err)
			}

		case "fields":
			if s.Fields == nil {
				s.Fields = make(map[string]GlobalOrdinalFieldStats, 0)
			}
			if err := dec.Decode(&s.Fields); err != nil {
				return fmt.Errorf("%s | %w", "Fields", err)
			}

		}
	}
	return nil
}

// NewGlobalOrdinalsStats returns a GlobalOrdinalsStats.
func NewGlobalOrdinalsStats() *GlobalOrdinalsStats {
	r := &GlobalOrdinalsStats{
		Fields: make(map[string]GlobalOrdinalFieldStats),
	}

	return r
}
