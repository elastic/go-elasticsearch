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

// SLMPolicy type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/slm/_types/SnapshotLifecycle.ts#L86-L92
type SLMPolicy struct {
	Config     *Configuration `json:"config,omitempty"`
	Name       string         `json:"name"`
	Repository string         `json:"repository"`
	Retention  *Retention     `json:"retention,omitempty"`
	Schedule   string         `json:"schedule"`
}

func (s *SLMPolicy) UnmarshalJSON(data []byte) error {

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

		case "config":
			if err := dec.Decode(&s.Config); err != nil {
				return fmt.Errorf("%s | %w", "Config", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "repository":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Repository", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Repository = o

		case "retention":
			if err := dec.Decode(&s.Retention); err != nil {
				return fmt.Errorf("%s | %w", "Retention", err)
			}

		case "schedule":
			if err := dec.Decode(&s.Schedule); err != nil {
				return fmt.Errorf("%s | %w", "Schedule", err)
			}

		}
	}
	return nil
}

// NewSLMPolicy returns a SLMPolicy.
func NewSLMPolicy() *SLMPolicy {
	r := &SLMPolicy{}

	return r
}
