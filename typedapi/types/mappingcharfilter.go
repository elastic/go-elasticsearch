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
// https://github.com/elastic/elasticsearch-specification/tree/b89646a75dd9e8001caf92d22bd8b3704c59dfdf

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// MappingCharFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/b89646a75dd9e8001caf92d22bd8b3704c59dfdf/specification/_types/analysis/char_filters.ts#L47-L51
type MappingCharFilter struct {
	Mappings     []string `json:"mappings,omitempty"`
	MappingsPath *string  `json:"mappings_path,omitempty"`
	Type         string   `json:"type,omitempty"`
	Version      *string  `json:"version,omitempty"`
}

func (s *MappingCharFilter) UnmarshalJSON(data []byte) error {

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

		case "mappings":
			if err := dec.Decode(&s.Mappings); err != nil {
				return err
			}

		case "mappings_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MappingsPath = &o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s MappingCharFilter) MarshalJSON() ([]byte, error) {
	type innerMappingCharFilter MappingCharFilter
	tmp := innerMappingCharFilter{
		Mappings:     s.Mappings,
		MappingsPath: s.MappingsPath,
		Type:         s.Type,
		Version:      s.Version,
	}

	tmp.Type = "mapping"

	return json.Marshal(tmp)
}

// NewMappingCharFilter returns a MappingCharFilter.
func NewMappingCharFilter() *MappingCharFilter {
	r := &MappingCharFilter{}

	return r
}
