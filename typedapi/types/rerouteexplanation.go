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

// RerouteExplanation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/f1932ce6b46a53a8342db522b1a7883bcc9e0996/specification/cluster/reroute/types.ts#L92-L96
type RerouteExplanation struct {
	Command    string            `json:"command"`
	Decisions  []RerouteDecision `json:"decisions"`
	Parameters RerouteParameters `json:"parameters"`
}

func (s *RerouteExplanation) UnmarshalJSON(data []byte) error {

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

		case "command":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Command", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Command = o

		case "decisions":
			if err := dec.Decode(&s.Decisions); err != nil {
				return fmt.Errorf("%s | %w", "Decisions", err)
			}

		case "parameters":
			if err := dec.Decode(&s.Parameters); err != nil {
				return fmt.Errorf("%s | %w", "Parameters", err)
			}

		}
	}
	return nil
}

// NewRerouteExplanation returns a RerouteExplanation.
func NewRerouteExplanation() *RerouteExplanation {
	r := &RerouteExplanation{}

	return r
}

// false
