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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/cohereinputtype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/coheretruncatetype"
)

// CohereTaskSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L870-L902
type CohereTaskSettings struct {
	// InputType For a `text_embedding` task, the type of input passed to the model.
	// Valid values are:
	//
	// * `classification`: Use it for embeddings passed through a text classifier.
	// * `clustering`: Use it for the embeddings run through a clustering algorithm.
	// * `ingest`: Use it for storing document embeddings in a vector database.
	// * `search`: Use it for storing embeddings of search queries run against a
	// vector database to find relevant documents.
	//
	// IMPORTANT: The `input_type` field is required when using embedding models
	// `v3` and higher.
	InputType *cohereinputtype.CohereInputType `json:"input_type,omitempty"`
	// ReturnDocuments For a `rerank` task, return doc text within the results.
	ReturnDocuments *bool `json:"return_documents,omitempty"`
	// TopN For a `rerank` task, the number of most relevant documents to return.
	// It defaults to the number of the documents.
	// If this inference endpoint is used in a `text_similarity_reranker` retriever
	// query and `top_n` is set, it must be greater than or equal to
	// `rank_window_size` in the query.
	TopN *int `json:"top_n,omitempty"`
	// Truncate For a `text_embedding` task, the method to handle inputs longer than the
	// maximum token length.
	// Valid values are:
	//
	// * `END`: When the input exceeds the maximum input token length, the end of
	// the input is discarded.
	// * `NONE`: When the input exceeds the maximum input token length, an error is
	// returned.
	// * `START`: When the input exceeds the maximum input token length, the start
	// of the input is discarded.
	Truncate *coheretruncatetype.CohereTruncateType `json:"truncate,omitempty"`
}

func (s *CohereTaskSettings) UnmarshalJSON(data []byte) error {

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
			if err := dec.Decode(&s.InputType); err != nil {
				return fmt.Errorf("%s | %w", "InputType", err)
			}

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

		case "top_n":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TopN", err)
				}
				s.TopN = &value
			case float64:
				f := int(v)
				s.TopN = &f
			}

		case "truncate":
			if err := dec.Decode(&s.Truncate); err != nil {
				return fmt.Errorf("%s | %w", "Truncate", err)
			}

		}
	}
	return nil
}

// NewCohereTaskSettings returns a CohereTaskSettings.
func NewCohereTaskSettings() *CohereTaskSettings {
	r := &CohereTaskSettings{}

	return r
}
