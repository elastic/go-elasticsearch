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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// VoyageAIServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L1616-L1647
type VoyageAIServiceSettings struct {
	// Dimensions The number of dimensions for resulting output embeddings.
	// This setting maps to `output_dimension` in the VoyageAI documentation.
	// Only for the `text_embedding` task type.
	Dimensions *int `json:"dimensions,omitempty"`
	// EmbeddingType The data type for the embeddings to be returned.
	// This setting maps to `output_dtype` in the VoyageAI documentation.
	// Permitted values: float, int8, bit.
	// `int8` is a synonym of `byte` in the VoyageAI documentation.
	// `bit` is a synonym of `binary` in the VoyageAI documentation.
	// Only for the `text_embedding` task type.
	EmbeddingType *float32 `json:"embedding_type,omitempty"`
	// ModelId The name of the model to use for the inference task.
	// Refer to the VoyageAI documentation for the list of available text embedding
	// and rerank models.
	ModelId string `json:"model_id"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// VoyageAI.
	// The `voyageai` service sets a default number of requests allowed per minute
	// depending on the task type.
	// For both `text_embedding` and `rerank`, it is set to `2000`.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
}

func (s *VoyageAIServiceSettings) UnmarshalJSON(data []byte) error {

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

		case "dimensions":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Dimensions", err)
				}
				s.Dimensions = &value
			case float64:
				f := int(v)
				s.Dimensions = &f
			}

		case "embedding_type":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "EmbeddingType", err)
				}
				f := float32(value)
				s.EmbeddingType = &f
			case float64:
				f := float32(v)
				s.EmbeddingType = &f
			}

		case "model_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ModelId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelId = o

		case "rate_limit":
			if err := dec.Decode(&s.RateLimit); err != nil {
				return fmt.Errorf("%s | %w", "RateLimit", err)
			}

		}
	}
	return nil
}

// NewVoyageAIServiceSettings returns a VoyageAIServiceSettings.
func NewVoyageAIServiceSettings() *VoyageAIServiceSettings {
	r := &VoyageAIServiceSettings{}

	return r
}
