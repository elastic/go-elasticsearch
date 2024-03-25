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
// https://github.com/elastic/elasticsearch-specification/tree/5fb8f1ce9c4605abcaa44aa0f17dbfc60497a757

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// MappingLimitSettingsNestedObjects type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fb8f1ce9c4605abcaa44aa0f17dbfc60497a757/specification/indices/_types/IndexSettings.ts#L453-L460
type MappingLimitSettingsNestedObjects struct {
	// Limit The maximum number of nested JSON objects that a single document can contain
	// across all nested types. This limit helps
	// to prevent out of memory errors when a document contains too many nested
	// objects.
	Limit *int `json:"limit,omitempty"`
}

func (s *MappingLimitSettingsNestedObjects) UnmarshalJSON(data []byte) error {

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

			var tmp interface{}
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

// NewMappingLimitSettingsNestedObjects returns a MappingLimitSettingsNestedObjects.
func NewMappingLimitSettingsNestedObjects() *MappingLimitSettingsNestedObjects {
	r := &MappingLimitSettingsNestedObjects{}

	return r
}
