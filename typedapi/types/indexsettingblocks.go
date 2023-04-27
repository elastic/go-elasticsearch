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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// IndexSettingBlocks type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/indices/_types/IndexSettings.ts#L245-L251
type IndexSettingBlocks struct {
	Metadata            Stringifiedboolean `json:"metadata,omitempty"`
	Read                *bool              `json:"read,omitempty"`
	ReadOnly            *bool              `json:"read_only,omitempty"`
	ReadOnlyAllowDelete *bool              `json:"read_only_allow_delete,omitempty"`
	Write               string             `json:"write,omitempty"`
}

func (s *IndexSettingBlocks) UnmarshalJSON(data []byte) error {

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

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return err
			}

		case "read":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Read = &value
			case bool:
				s.Read = &v
			}

		case "read_only":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.ReadOnly = &value
			case bool:
				s.ReadOnly = &v
			}

		case "read_only_allow_delete":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.ReadOnlyAllowDelete = &value
			case bool:
				s.ReadOnlyAllowDelete = &v
			}

		case "write":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Write = o

		}
	}
	return nil
}

// NewIndexSettingBlocks returns a IndexSettingBlocks.
func NewIndexSettingBlocks() *IndexSettingBlocks {
	r := &IndexSettingBlocks{}

	return r
}
