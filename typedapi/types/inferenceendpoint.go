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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// InferenceEndpoint type.
//
// https://github.com/elastic/elasticsearch-specification/blob/b1811e10a0722431d79d1c234dd412ff47d8656f/specification/inference/_types/Services.ts#L51-L73
type InferenceEndpoint struct {
	// ChunkingSettings The chunking configuration object.
	// Applies only to the `sparse_embedding` and `text_embedding` task types.
	// Not applicable to the `rerank`, `completion`, or `chat_completion` task
	// types.
	ChunkingSettings *InferenceChunkingSettings `json:"chunking_settings,omitempty"`
	// Service The service type
	Service string `json:"service"`
	// ServiceSettings Settings specific to the service
	ServiceSettings json.RawMessage `json:"service_settings"`
	// TaskSettings Task settings specific to the service and task type
	TaskSettings json.RawMessage `json:"task_settings,omitempty"`
}

func (s *InferenceEndpoint) UnmarshalJSON(data []byte) error {

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

		case "chunking_settings":
			if err := dec.Decode(&s.ChunkingSettings); err != nil {
				return fmt.Errorf("%s | %w", "ChunkingSettings", err)
			}

		case "service":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Service", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Service = o

		case "service_settings":
			if err := dec.Decode(&s.ServiceSettings); err != nil {
				return fmt.Errorf("%s | %w", "ServiceSettings", err)
			}

		case "task_settings":
			if err := dec.Decode(&s.TaskSettings); err != nil {
				return fmt.Errorf("%s | %w", "TaskSettings", err)
			}

		}
	}
	return nil
}

// NewInferenceEndpoint returns a InferenceEndpoint.
func NewInferenceEndpoint() *InferenceEndpoint {
	r := &InferenceEndpoint{}

	return r
}

type InferenceEndpointVariant interface {
	InferenceEndpointCaster() *InferenceEndpoint
}

func (s *InferenceEndpoint) InferenceEndpointCaster() *InferenceEndpoint {
	return s
}
