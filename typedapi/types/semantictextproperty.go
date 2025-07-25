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
// https://github.com/elastic/elasticsearch-specification/tree/a0b0db20330063a6d11f7997ff443fd2a1a827d1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// SemanticTextProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0b0db20330063a6d11f7997ff443fd2a1a827d1/specification/_types/mapping/core.ts#L229-L245
type SemanticTextProperty struct {
	// InferenceId Inference endpoint that will be used to generate embeddings for the field.
	// This parameter cannot be updated. Use the Create inference API to create the
	// endpoint.
	// If `search_inference_id` is specified, the inference endpoint will only be
	// used at index time.
	InferenceId *string           `json:"inference_id,omitempty"`
	Meta        map[string]string `json:"meta,omitempty"`
	// SearchInferenceId Inference endpoint that will be used to generate embeddings at query time.
	// You can update this parameter by using the Update mapping API. Use the Create
	// inference API to create the endpoint.
	// If not specified, the inference endpoint defined by inference_id will be used
	// at both index and query time.
	SearchInferenceId *string `json:"search_inference_id,omitempty"`
	Type              string  `json:"type,omitempty"`
}

func (s *SemanticTextProperty) UnmarshalJSON(data []byte) error {

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

		case "inference_id":
			if err := dec.Decode(&s.InferenceId); err != nil {
				return fmt.Errorf("%s | %w", "InferenceId", err)
			}

		case "meta":
			if s.Meta == nil {
				s.Meta = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Meta); err != nil {
				return fmt.Errorf("%s | %w", "Meta", err)
			}

		case "search_inference_id":
			if err := dec.Decode(&s.SearchInferenceId); err != nil {
				return fmt.Errorf("%s | %w", "SearchInferenceId", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s SemanticTextProperty) MarshalJSON() ([]byte, error) {
	type innerSemanticTextProperty SemanticTextProperty
	tmp := innerSemanticTextProperty{
		InferenceId:       s.InferenceId,
		Meta:              s.Meta,
		SearchInferenceId: s.SearchInferenceId,
		Type:              s.Type,
	}

	tmp.Type = "semantic_text"

	return json.Marshal(tmp)
}

// NewSemanticTextProperty returns a SemanticTextProperty.
func NewSemanticTextProperty() *SemanticTextProperty {
	r := &SemanticTextProperty{
		Meta: make(map[string]string),
	}

	return r
}

type SemanticTextPropertyVariant interface {
	SemanticTextPropertyCaster() *SemanticTextProperty
}

func (s *SemanticTextProperty) SemanticTextPropertyCaster() *SemanticTextProperty {
	return s
}

func (s *SemanticTextProperty) PropertyCaster() *Property {
	o := Property(s)
	return &o
}
