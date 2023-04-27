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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/acknowledgementoptions"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// AcknowledgeState type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/watcher/_types/Action.ts#L112-L115
type AcknowledgeState struct {
	State     acknowledgementoptions.AcknowledgementOptions `json:"state"`
	Timestamp DateTime                                      `json:"timestamp"`
}

func (s *AcknowledgeState) UnmarshalJSON(data []byte) error {

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

		case "state":
			if err := dec.Decode(&s.State); err != nil {
				return err
			}

		case "timestamp":
			if err := dec.Decode(&s.Timestamp); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewAcknowledgeState returns a AcknowledgeState.
func NewAcknowledgeState() *AcknowledgeState {
	r := &AcknowledgeState{}

	return r
}
