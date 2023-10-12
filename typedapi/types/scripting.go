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
// https://github.com/elastic/elasticsearch-specification/tree/3b09f9d8e90178243f8a340a7bc324aab152c602

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// Scripting type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3b09f9d8e90178243f8a340a7bc324aab152c602/specification/nodes/_types/Stats.ts#L389-L395
type Scripting struct {
	CacheEvictions            *int64           `json:"cache_evictions,omitempty"`
	CompilationLimitTriggered *int64           `json:"compilation_limit_triggered,omitempty"`
	Compilations              *int64           `json:"compilations,omitempty"`
	CompilationsHistory       map[string]int64 `json:"compilations_history,omitempty"`
	Contexts                  []NodesContext   `json:"contexts,omitempty"`
}

func (s *Scripting) UnmarshalJSON(data []byte) error {

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

		case "cache_evictions":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.CacheEvictions = &value
			case float64:
				f := int64(v)
				s.CacheEvictions = &f
			}

		case "compilation_limit_triggered":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.CompilationLimitTriggered = &value
			case float64:
				f := int64(v)
				s.CompilationLimitTriggered = &f
			}

		case "compilations":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Compilations = &value
			case float64:
				f := int64(v)
				s.Compilations = &f
			}

		case "compilations_history":
			if s.CompilationsHistory == nil {
				s.CompilationsHistory = make(map[string]int64, 0)
			}
			if err := dec.Decode(&s.CompilationsHistory); err != nil {
				return err
			}

		case "contexts":
			if err := dec.Decode(&s.Contexts); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewScripting returns a Scripting.
func NewScripting() *Scripting {
	r := &Scripting{
		CompilationsHistory: make(map[string]int64, 0),
	}

	return r
}
