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

// ReloadDetails type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/reload_search_analyzers/types.ts#L27-L31
type ReloadDetails struct {
	Index             string   `json:"index"`
	ReloadedAnalyzers []string `json:"reloaded_analyzers"`
	ReloadedNodeIds   []string `json:"reloaded_node_ids"`
}

func (s *ReloadDetails) UnmarshalJSON(data []byte) error {

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

		case "index":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Index = o

		case "reloaded_analyzers":
			if err := dec.Decode(&s.ReloadedAnalyzers); err != nil {
				return fmt.Errorf("%s | %w", "ReloadedAnalyzers", err)
			}

		case "reloaded_node_ids":
			if err := dec.Decode(&s.ReloadedNodeIds); err != nil {
				return fmt.Errorf("%s | %w", "ReloadedNodeIds", err)
			}

		}
	}
	return nil
}

// NewReloadDetails returns a ReloadDetails.
func NewReloadDetails() *ReloadDetails {
	r := &ReloadDetails{}

	return r
}
