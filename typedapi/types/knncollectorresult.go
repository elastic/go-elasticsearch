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

// KnnCollectorResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/search/_types/profile.ts#L222-L228
type KnnCollectorResult struct {
	Children    []KnnCollectorResult `json:"children,omitempty"`
	Name        string               `json:"name"`
	Reason      string               `json:"reason"`
	Time        Duration             `json:"time,omitempty"`
	TimeInNanos int64                `json:"time_in_nanos"`
}

func (s *KnnCollectorResult) UnmarshalJSON(data []byte) error {

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

		case "children":
			if err := dec.Decode(&s.Children); err != nil {
				return fmt.Errorf("%s | %w", "Children", err)
			}

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = o

		case "reason":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Reason", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Reason = o

		case "time":
			if err := dec.Decode(&s.Time); err != nil {
				return fmt.Errorf("%s | %w", "Time", err)
			}

		case "time_in_nanos":
			if err := dec.Decode(&s.TimeInNanos); err != nil {
				return fmt.Errorf("%s | %w", "TimeInNanos", err)
			}

		}
	}
	return nil
}

// NewKnnCollectorResult returns a KnnCollectorResult.
func NewKnnCollectorResult() *KnnCollectorResult {
	r := &KnnCollectorResult{}

	return r
}
