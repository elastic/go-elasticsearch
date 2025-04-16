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
// https://github.com/elastic/elasticsearch-specification/tree/f1932ce6b46a53a8342db522b1a7883bcc9e0996

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// InProgress type.
//
// https://github.com/elastic/elasticsearch-specification/blob/f1932ce6b46a53a8342db522b1a7883bcc9e0996/specification/slm/_types/SnapshotLifecycle.ts#L141-L146
type InProgress struct {
	Name            string `json:"name"`
	StartTimeMillis int64  `json:"start_time_millis"`
	State           string `json:"state"`
	Uuid            string `json:"uuid"`
}

func (s *InProgress) UnmarshalJSON(data []byte) error {

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

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "start_time_millis":
			if err := dec.Decode(&s.StartTimeMillis); err != nil {
				return fmt.Errorf("%s | %w", "StartTimeMillis", err)
			}

		case "state":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "State", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.State = o

		case "uuid":
			if err := dec.Decode(&s.Uuid); err != nil {
				return fmt.Errorf("%s | %w", "Uuid", err)
			}

		}
	}
	return nil
}

// NewInProgress returns a InProgress.
func NewInProgress() *InProgress {
	r := &InProgress{}

	return r
}

// false
