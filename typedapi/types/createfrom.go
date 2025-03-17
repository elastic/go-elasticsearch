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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// CreateFrom type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ea991724f4dd4f90c496eff547d3cc2e6529f509/specification/indices/create_from/MigrateCreateFromRequest.ts#L46-L60
type CreateFrom struct {
	// MappingsOverride Mappings overrides to be applied to the destination index (optional)
	MappingsOverride *TypeMapping `json:"mappings_override,omitempty"`
	// RemoveIndexBlocks If index blocks should be removed when creating destination index (optional)
	RemoveIndexBlocks *bool `json:"remove_index_blocks,omitempty"`
	// SettingsOverride Settings overrides to be applied to the destination index (optional)
	SettingsOverride *IndexSettings `json:"settings_override,omitempty"`
}

func (s *CreateFrom) UnmarshalJSON(data []byte) error {

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

		case "mappings_override":
			if err := dec.Decode(&s.MappingsOverride); err != nil {
				return fmt.Errorf("%s | %w", "MappingsOverride", err)
			}

		case "remove_index_blocks":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "RemoveIndexBlocks", err)
				}
				s.RemoveIndexBlocks = &value
			case bool:
				s.RemoveIndexBlocks = &v
			}

		case "settings_override":
			if err := dec.Decode(&s.SettingsOverride); err != nil {
				return fmt.Errorf("%s | %w", "SettingsOverride", err)
			}

		}
	}
	return nil
}

// NewCreateFrom returns a CreateFrom.
func NewCreateFrom() *CreateFrom {
	r := &CreateFrom{}

	return r
}

// true

type CreateFromVariant interface {
	CreateFromCaster() *CreateFrom
}

func (s *CreateFrom) CreateFromCaster() *CreateFrom {
	return s
}
