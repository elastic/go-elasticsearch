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
)

// ComponentTemplateSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/_types/ComponentTemplate.ts#L43-L55
type ComponentTemplateSummary struct {
	Aliases   map[string]AliasDefinition       `json:"aliases,omitempty"`
	Lifecycle *DataStreamLifecycleWithRollover `json:"lifecycle,omitempty"`
	Mappings  *TypeMapping                     `json:"mappings,omitempty"`
	Meta_     Metadata                         `json:"_meta,omitempty"`
	Settings  map[string]IndexSettings         `json:"settings,omitempty"`
	Version   *int64                           `json:"version,omitempty"`
}

func (s *ComponentTemplateSummary) UnmarshalJSON(data []byte) error {

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

		case "aliases":
			if s.Aliases == nil {
				s.Aliases = make(map[string]AliasDefinition, 0)
			}
			if err := dec.Decode(&s.Aliases); err != nil {
				return fmt.Errorf("%s | %w", "Aliases", err)
			}

		case "lifecycle":
			if err := dec.Decode(&s.Lifecycle); err != nil {
				return fmt.Errorf("%s | %w", "Lifecycle", err)
			}

		case "mappings":
			if err := dec.Decode(&s.Mappings); err != nil {
				return fmt.Errorf("%s | %w", "Mappings", err)
			}

		case "_meta":
			if err := dec.Decode(&s.Meta_); err != nil {
				return fmt.Errorf("%s | %w", "Meta_", err)
			}

		case "settings":
			if s.Settings == nil {
				s.Settings = make(map[string]IndexSettings, 0)
			}
			if err := dec.Decode(&s.Settings); err != nil {
				return fmt.Errorf("%s | %w", "Settings", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewComponentTemplateSummary returns a ComponentTemplateSummary.
func NewComponentTemplateSummary() *ComponentTemplateSummary {
	r := &ComponentTemplateSummary{
		Aliases:  make(map[string]AliasDefinition),
		Settings: make(map[string]IndexSettings),
	}

	return r
}
