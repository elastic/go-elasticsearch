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
// https://github.com/elastic/elasticsearch-specification/tree/50c316c036cf0c3f567011c2bc24e7d2e1b8c781

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// MergeScheduler type.
//
// https://github.com/elastic/elasticsearch-specification/blob/50c316c036cf0c3f567011c2bc24e7d2e1b8c781/specification/indices/_types/IndexSettings.ts#L326-L329
type MergeScheduler struct {
	MaxMergeCount  Stringifiedinteger `json:"max_merge_count,omitempty"`
	MaxThreadCount Stringifiedinteger `json:"max_thread_count,omitempty"`
}

func (s *MergeScheduler) UnmarshalJSON(data []byte) error {

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

		case "max_merge_count":
			if err := dec.Decode(&s.MaxMergeCount); err != nil {
				return err
			}

		case "max_thread_count":
			if err := dec.Decode(&s.MaxThreadCount); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewMergeScheduler returns a MergeScheduler.
func NewMergeScheduler() *MergeScheduler {
	r := &MergeScheduler{}

	return r
}
