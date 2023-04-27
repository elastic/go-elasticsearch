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

	"encoding/json"
)

// Latest type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/transform/_types/Transform.ts#L47-L52
type Latest struct {
	// Sort Specifies the date field that is used to identify the latest documents.
	Sort string `json:"sort"`
	// UniqueKey Specifies an array of one or more fields that are used to group the data.
	UniqueKey []string `json:"unique_key"`
}

func (s *Latest) UnmarshalJSON(data []byte) error {

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

		case "sort":
			if err := dec.Decode(&s.Sort); err != nil {
				return err
			}

		case "unique_key":
			if err := dec.Decode(&s.UniqueKey); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewLatest returns a Latest.
func NewLatest() *Latest {
	r := &Latest{}

	return r
}
