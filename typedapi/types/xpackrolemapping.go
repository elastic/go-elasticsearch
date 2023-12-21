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
// https://github.com/elastic/elasticsearch-specification/tree/5c8fed5fe577b0d5e9fde34fb13795c5a66fe9fe

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// XpackRoleMapping type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5c8fed5fe577b0d5e9fde34fb13795c5a66fe9fe/specification/xpack/usage/types.ts#L270-L273
type XpackRoleMapping struct {
	Enabled int `json:"enabled"`
	Size    int `json:"size"`
}

func (s *XpackRoleMapping) UnmarshalJSON(data []byte) error {

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

		case "enabled":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Enabled = value
			case float64:
				f := int(v)
				s.Enabled = f
			}

		case "size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Size = value
			case float64:
				f := int(v)
				s.Size = f
			}

		}
	}
	return nil
}

// NewXpackRoleMapping returns a XpackRoleMapping.
func NewXpackRoleMapping() *XpackRoleMapping {
	r := &XpackRoleMapping{}

	return r
}
