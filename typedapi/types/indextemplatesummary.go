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
)

// IndexTemplateSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/_types/IndexTemplate.ts#L97-L124
type IndexTemplateSummary struct {
	// Aliases Aliases to add.
	// If the index template includes a `data_stream` object, these are data stream
	// aliases.
	// Otherwise, these are index aliases.
	// Data stream aliases ignore the `index_routing`, `routing`, and
	// `search_routing` options.
	Aliases           map[string]Alias                 `json:"aliases,omitempty"`
	DataStreamOptions *DataStreamOptionsTemplate       `json:"data_stream_options,omitempty"`
	Lifecycle         *DataStreamLifecycleWithRollover `json:"lifecycle,omitempty"`
	// Mappings Mapping for fields in the index.
	// If specified, this mapping can include field names, field data types, and
	// mapping parameters.
	Mappings *TypeMapping `json:"mappings,omitempty"`
	// Settings Configuration options for the index.
	Settings *IndexSettings `json:"settings,omitempty"`
}

func (s *IndexTemplateSummary) UnmarshalJSON(data []byte) error {

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
				s.Aliases = make(map[string]Alias, 0)
			}
			if err := dec.Decode(&s.Aliases); err != nil {
				return fmt.Errorf("%s | %w", "Aliases", err)
			}

		case "data_stream_options":
			if err := dec.Decode(&s.DataStreamOptions); err != nil {
				return fmt.Errorf("%s | %w", "DataStreamOptions", err)
			}

		case "lifecycle":
			if err := dec.Decode(&s.Lifecycle); err != nil {
				return fmt.Errorf("%s | %w", "Lifecycle", err)
			}

		case "mappings":
			if err := dec.Decode(&s.Mappings); err != nil {
				return fmt.Errorf("%s | %w", "Mappings", err)
			}

		case "settings":
			if err := dec.Decode(&s.Settings); err != nil {
				return fmt.Errorf("%s | %w", "Settings", err)
			}

		}
	}
	return nil
}

// NewIndexTemplateSummary returns a IndexTemplateSummary.
func NewIndexTemplateSummary() *IndexTemplateSummary {
	r := &IndexTemplateSummary{
		Aliases: make(map[string]Alias),
	}

	return r
}

type IndexTemplateSummaryVariant interface {
	IndexTemplateSummaryCaster() *IndexTemplateSummary
}

func (s *IndexTemplateSummary) IndexTemplateSummaryCaster() *IndexTemplateSummary {
	return s
}
