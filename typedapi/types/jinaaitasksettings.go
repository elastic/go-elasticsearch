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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jinaaitextembeddingtask"
)

// JinaAITaskSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/6785a6caa1fa3ca5ab3308963d79dce923a3469f/specification/inference/_types/CommonTypes.ts#L1699-L1729
type JinaAITaskSettings struct {
	// LateChunking For a `text_embedding` task, controls when text is split into chunks.
	// When set to `true`, a request from Elasticsearch contains only chunks related
	// to a single document. Instead of batching chunks across documents,
	// Elasticsearch sends them in separate requests. This ensures that chunk
	// embeddings retain context from the entire document, improving semantic
	// quality.
	//
	// If a document exceeds the model's context limits, late chunking is
	// automatically disabled for that document only and standard chunking is used
	// instead.
	//
	// If not specified, defaults to `false`.
	LateChunking *bool `json:"late_chunking,omitempty"`
	// ReturnDocuments For a `rerank` task, return the doc text within the results.
	ReturnDocuments *bool `json:"return_documents,omitempty"`
	// Task For a `text_embedding` task, the task passed to the model.
	// Valid values are:
	//
	// * `classification`: Use it for embeddings passed through a text classifier.
	// * `clustering`: Use it for the embeddings run through a clustering algorithm.
	// * `ingest`: Use it for storing document embeddings in a vector database.
	// * `search`: Use it for storing embeddings of search queries run against a
	// vector database to find relevant documents.
	Task *jinaaitextembeddingtask.JinaAITextEmbeddingTask `json:"task,omitempty"`
	// TopN For a `rerank` task, the number of most relevant documents to return.
	// It defaults to the number of the documents.
	// If this inference endpoint is used in a `text_similarity_reranker` retriever
	// query and `top_n` is set, it must be greater than or equal to
	// `rank_window_size` in the query.
	TopN *int `json:"top_n,omitempty"`
}

func (s *JinaAITaskSettings) UnmarshalJSON(data []byte) error {

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

		case "late_chunking":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "LateChunking", err)
				}
				s.LateChunking = &value
			case bool:
				s.LateChunking = &v
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

		case "task":
			if err := dec.Decode(&s.Task); err != nil {
				return fmt.Errorf("%s | %w", "Task", err)
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

		}
	}
	return nil
}

// NewJinaAITaskSettings returns a JinaAITaskSettings.
func NewJinaAITaskSettings() *JinaAITaskSettings {
	r := &JinaAITaskSettings{}

	return r
}

type JinaAITaskSettingsVariant interface {
	JinaAITaskSettingsCaster() *JinaAITaskSettings
}

func (s *JinaAITaskSettings) JinaAITaskSettingsCaster() *JinaAITaskSettings {
	return s
}
