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
// https://github.com/elastic/elasticsearch-specification/tree/9a0362eb2579c6604966a8fb307caee92de04270

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// MappingLimitSettingsNestedFields type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9a0362eb2579c6604966a8fb307caee92de04270/specification/indices/_types/IndexSettings.ts#L443-L451
type MappingLimitSettingsNestedFields struct {
	// Limit The maximum number of distinct nested mappings in an index. The nested type
	// should only be used in special cases, when
	// arrays of objects need to be queried independently of each other. To
	// safeguard against poorly designed mappings, this
	// setting limits the number of unique nested types per index.
	Limit *int `json:"limit,omitempty"`
}

func (s *MappingLimitSettingsNestedFields) UnmarshalJSON(data []byte) error {

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

		case "limit":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Limit", err)
				}
				s.Limit = &value
			case float64:
				f := int(v)
				s.Limit = &f
			}

		}
	}
	return nil
}

// NewMappingLimitSettingsNestedFields returns a MappingLimitSettingsNestedFields.
func NewMappingLimitSettingsNestedFields() *MappingLimitSettingsNestedFields {
	r := &MappingLimitSettingsNestedFields{}

	return r
}
