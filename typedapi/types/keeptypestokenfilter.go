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
// https://github.com/elastic/elasticsearch-specification/tree/07bf82537a186562d8699685e3704ea338b268ef

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/keeptypesmode"
)

// KeepTypesTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/07bf82537a186562d8699685e3704ea338b268ef/specification/_types/analysis/token_filters.ts#L220-L224
type KeepTypesTokenFilter struct {
	Mode    *keeptypesmode.KeepTypesMode `json:"mode,omitempty"`
	Type    string                       `json:"type,omitempty"`
	Types   []string                     `json:"types,omitempty"`
	Version *string                      `json:"version,omitempty"`
}

func (s *KeepTypesTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "mode":
			if err := dec.Decode(&s.Mode); err != nil {
				return fmt.Errorf("%s | %w", "Mode", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "types":
			if err := dec.Decode(&s.Types); err != nil {
				return fmt.Errorf("%s | %w", "Types", err)
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
func (s KeepTypesTokenFilter) MarshalJSON() ([]byte, error) {
	type innerKeepTypesTokenFilter KeepTypesTokenFilter
	tmp := innerKeepTypesTokenFilter{
		Mode:    s.Mode,
		Type:    s.Type,
		Types:   s.Types,
		Version: s.Version,
	}

	tmp.Type = "keep_types"

	return json.Marshal(tmp)
}

// NewKeepTypesTokenFilter returns a KeepTypesTokenFilter.
func NewKeepTypesTokenFilter() *KeepTypesTokenFilter {
	r := &KeepTypesTokenFilter{}

	return r
}
