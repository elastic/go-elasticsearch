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
// https://github.com/elastic/elasticsearch-specification/tree/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67

package mount

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Request holds the request body struct for the package mount
//
// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/searchable_snapshots/mount/SearchableSnapshotsMountRequest.ts#L26-L49
type Request struct {
	IgnoreIndexSettings []string                   `json:"ignore_index_settings,omitempty"`
	Index               string                     `json:"index"`
	IndexSettings       map[string]json.RawMessage `json:"index_settings,omitempty"`
	RenamedIndex        *string                    `json:"renamed_index,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		IndexSettings: make(map[string]json.RawMessage, 0),
	}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Mount request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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

		case "ignore_index_settings":
			if err := dec.Decode(&s.IgnoreIndexSettings); err != nil {
				return err
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "index_settings":
			if s.IndexSettings == nil {
				s.IndexSettings = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.IndexSettings); err != nil {
				return err
			}

		case "renamed_index":
			if err := dec.Decode(&s.RenamedIndex); err != nil {
				return err
			}

		}
	}
	return nil
}
