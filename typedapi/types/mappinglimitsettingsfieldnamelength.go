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

// MappingLimitSettingsFieldNameLength type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fb8f1ce9c4605abcaa44aa0f17dbfc60497a757/specification/indices/_types/IndexSettings.ts#L462-L469
type MappingLimitSettingsFieldNameLength struct {
	// Limit Setting for the maximum length of a field name. This setting isn’t really
	// something that addresses mappings explosion but
	// might still be useful if you want to limit the field length. It usually
	// shouldn’t be necessary to set this setting. The
	// default is okay unless a user starts to add a huge number of fields with
	// really long names. Default is `Long.MAX_VALUE` (no limit).
	Limit *int64 `json:"limit,omitempty"`
}

func (s *MappingLimitSettingsFieldNameLength) UnmarshalJSON(data []byte) error {

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

// NewMappingLimitSettingsFieldNameLength returns a MappingLimitSettingsFieldNameLength.
func NewMappingLimitSettingsFieldNameLength() *MappingLimitSettingsFieldNameLength {
	r := &MappingLimitSettingsFieldNameLength{}

	return r
}
