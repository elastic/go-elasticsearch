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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Body type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/esql/list_queries/ListQueriesResponse.ts#L24-L30
type Body struct {
	Id               int64  `json:"id"`
	Node             string `json:"node"`
	Query            string `json:"query"`
	RunningTimeNanos int64  `json:"running_time_nanos"`
	StartTimeMillis  int64  `json:"start_time_millis"`
}

func (s *Body) UnmarshalJSON(data []byte) error {

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

		case "id":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Id", err)
				}
				s.Id = value
			case float64:
				f := int64(v)
				s.Id = f
			}

		case "node":
			if err := dec.Decode(&s.Node); err != nil {
				return fmt.Errorf("%s | %w", "Node", err)
			}

		case "query":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Query = o

		case "running_time_nanos":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "RunningTimeNanos", err)
				}
				s.RunningTimeNanos = value
			case float64:
				f := int64(v)
				s.RunningTimeNanos = f
			}

		case "start_time_millis":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "StartTimeMillis", err)
				}
				s.StartTimeMillis = value
			case float64:
				f := int64(v)
				s.StartTimeMillis = f
			}

		}
	}
	return nil
}

// NewBody returns a Body.
func NewBody() *Body {
	r := &Body{}

	return r
}
