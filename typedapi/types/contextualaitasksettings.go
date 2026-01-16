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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ContextualAITaskSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27/specification/inference/_types/CommonTypes.ts#L1251-L1269
type ContextualAITaskSettings struct {
	// Instruction Instructions for the reranking model. Refer to
	// <https://docs.contextual.ai/api-reference/rerank/rerank#body-instruction>
	// Only for the `rerank` task type.
	Instruction *string `json:"instruction,omitempty"`
	// ReturnDocuments Whether to return the source documents in the response.
	// Only for the `rerank` task type.
	ReturnDocuments *bool `json:"return_documents,omitempty"`
	// TopK The number of most relevant documents to return.
	// If not specified, the reranking results of all documents will be returned.
	// Only for the `rerank` task type.
	TopK *int `json:"top_k,omitempty"`
}

func (s *ContextualAITaskSettings) UnmarshalJSON(data []byte) error {

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

		case "instruction":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Instruction", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Instruction = &o

		case "return_documents":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ReturnDocuments", err)
				}
				s.ReturnDocuments = &value
			case bool:
				s.ReturnDocuments = &v
			}

		case "top_k":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TopK", err)
				}
				s.TopK = &value
			case float64:
				f := int(v)
				s.TopK = &f
			}

		}
	}
	return nil
}

// NewContextualAITaskSettings returns a ContextualAITaskSettings.
func NewContextualAITaskSettings() *ContextualAITaskSettings {
	r := &ContextualAITaskSettings{}

	return r
}

type ContextualAITaskSettingsVariant interface {
	ContextualAITaskSettingsCaster() *ContextualAITaskSettings
}

func (s *ContextualAITaskSettings) ContextualAITaskSettingsCaster() *ContextualAITaskSettings {
	return s
}
