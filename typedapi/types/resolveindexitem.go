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

// ResolveIndexItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/indices/resolve_index/ResolveIndexResponse.ts#L30-L35
type ResolveIndexItem struct {
	Aliases    []string `json:"aliases,omitempty"`
	Attributes []string `json:"attributes"`
	DataStream *string  `json:"data_stream,omitempty"`
	Name       string   `json:"name"`
}

func (s *ResolveIndexItem) UnmarshalJSON(data []byte) error {

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
			if err := dec.Decode(&s.Aliases); err != nil {
				return err
			}

		case "attributes":
			if err := dec.Decode(&s.Attributes); err != nil {
				return err
			}

		case "data_stream":
			if err := dec.Decode(&s.DataStream); err != nil {
				return err
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewResolveIndexItem returns a ResolveIndexItem.
func NewResolveIndexItem() *ResolveIndexItem {
	r := &ResolveIndexItem{}

	return r
}
