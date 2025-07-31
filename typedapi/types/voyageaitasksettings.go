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

// VoyageAITaskSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L1649-L1673
type VoyageAITaskSettings struct {
	// InputType Type of the input text.
	// Permitted values: `ingest` (maps to `document` in the VoyageAI
	// documentation), `search` (maps to `query` in the VoyageAI documentation).
	// Only for the `text_embedding` task type.
	InputType *string `json:"input_type,omitempty"`
	// ReturnDocuments Whether to return the source documents in the response.
	// Only for the `rerank` task type.
	ReturnDocuments *bool `json:"return_documents,omitempty"`
	// TopK The number of most relevant documents to return.
	// If not specified, the reranking results of all documents will be returned.
	// Only for the `rerank` task type.
	TopK *int `json:"top_k,omitempty"`
	// Truncation Whether to truncate the input texts to fit within the context length.
	Truncation *bool `json:"truncation,omitempty"`
}

func (s *VoyageAITaskSettings) UnmarshalJSON(data []byte) error {

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

		case "input_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "InputType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.InputType = &o

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

		case "truncation":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Truncation", err)
				}
				s.Truncation = &value
			case bool:
				s.Truncation = &v
			}

		}
	}
	return nil
}

// NewVoyageAITaskSettings returns a VoyageAITaskSettings.
func NewVoyageAITaskSettings() *VoyageAITaskSettings {
	r := &VoyageAITaskSettings{}

	return r
}
