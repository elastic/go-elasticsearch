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
// https://github.com/elastic/elasticsearch-specification/tree/19027dbdd366978ccae41842a040a636730e7c10

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// MappingLimitSettingsTotalFields type.
//
// https://github.com/elastic/elasticsearch-specification/blob/19027dbdd366978ccae41842a040a636730e7c10/specification/indices/_types/IndexSettings.ts#L426-L443
type MappingLimitSettingsTotalFields struct {
	// IgnoreDynamicBeyondLimit This setting determines what happens when a dynamically mapped field would
	// exceed the total fields limit. When set
	// to false (the default), the index request of the document that tries to add a
	// dynamic field to the mapping will fail
	// with the message Limit of total fields [X] has been exceeded. When set to
	// true, the index request will not fail.
	// Instead, fields that would exceed the limit are not added to the mapping,
	// similar to dynamic: false.
	// The fields that were not added to the mapping will be added to the _ignored
	// field.
	IgnoreDynamicBeyondLimit *bool `json:"ignore_dynamic_beyond_limit,omitempty"`
	// Limit The maximum number of fields in an index. Field and object mappings, as well
	// as field aliases count towards this limit.
	// The limit is in place to prevent mappings and searches from becoming too
	// large. Higher values can lead to performance
	// degradations and memory issues, especially in clusters with a high load or
	// few resources.
	Limit *int64 `json:"limit,omitempty"`
}

func (s *MappingLimitSettingsTotalFields) UnmarshalJSON(data []byte) error {

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

		case "ignore_dynamic_beyond_limit":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreDynamicBeyondLimit", err)
				}
				s.IgnoreDynamicBeyondLimit = &value
			case bool:
				s.IgnoreDynamicBeyondLimit = &v
			}

		case "limit":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Limit", err)
				}
				s.Limit = &value
			case float64:
				f := int64(v)
				s.Limit = &f
			}

		}
	}
	return nil
}

// NewMappingLimitSettingsTotalFields returns a MappingLimitSettingsTotalFields.
func NewMappingLimitSettingsTotalFields() *MappingLimitSettingsTotalFields {
	r := &MappingLimitSettingsTotalFields{}

	return r
}
