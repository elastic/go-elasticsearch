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

// Watcher type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/xpack/usage/types.ts#L466-L470
type Watcher struct {
	Available bool           `json:"available"`
	Count     Counter        `json:"count"`
	Enabled   bool           `json:"enabled"`
	Execution WatcherActions `json:"execution"`
	Watch     WatcherWatch   `json:"watch"`
}

func (s *Watcher) UnmarshalJSON(data []byte) error {

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

		case "available":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Available", err)
				}
				s.Available = value
			case bool:
				s.Available = v
			}

		case "count":
			if err := dec.Decode(&s.Count); err != nil {
				return fmt.Errorf("%s | %w", "Count", err)
			}

		case "enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Enabled", err)
				}
				s.Enabled = value
			case bool:
				s.Enabled = v
			}

		case "execution":
			if err := dec.Decode(&s.Execution); err != nil {
				return fmt.Errorf("%s | %w", "Execution", err)
			}

		case "watch":
			if err := dec.Decode(&s.Watch); err != nil {
				return fmt.Errorf("%s | %w", "Watch", err)
			}

		}
	}
	return nil
}

// NewWatcher returns a Watcher.
func NewWatcher() *Watcher {
	r := &Watcher{}

	return r
}
