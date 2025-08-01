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

// MappingCharFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/analysis/char_filters.ts#L51-L55
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
				return fmt.Errorf("%s | %w", "Mappings", err)
			}

		case "mappings_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MappingsPath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MappingsPath = &o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
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

type MappingCharFilterVariant interface {
	MappingCharFilterCaster() *MappingCharFilter
}

func (s *MappingCharFilter) MappingCharFilterCaster() *MappingCharFilter {
	return s
}

func (s *MappingCharFilter) CharFilterDefinitionCaster() *CharFilterDefinition {
	o := CharFilterDefinition(s)
	return &o
}
