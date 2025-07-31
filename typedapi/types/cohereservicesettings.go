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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/cohereembeddingtype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/coheresimilaritytype"
)

// CohereServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L788-L831
type CohereServiceSettings struct {
	// ApiKey A valid API key for your Cohere account.
	// You can find or create your Cohere API keys on the Cohere API key settings
	// page.
	//
	// IMPORTANT: You need to provide the API key only once, during the inference
	// model creation.
	// The get inference endpoint API does not retrieve your API key.
	// After creating the inference model, you cannot change the associated API key.
	// If you want to use a different API key, delete the inference model and
	// recreate it with the same name and the updated API key.
	ApiKey string `json:"api_key"`
	// EmbeddingType For a `text_embedding` task, the types of embeddings you want to get back.
	// Use `binary` for binary embeddings, which are encoded as bytes with signed
	// int8 precision.
	// Use `bit` for binary embeddings, which are encoded as bytes with signed int8
	// precision (this is a synonym of `binary`).
	// Use `byte` for signed int8 embeddings (this is a synonym of `int8`).
	// Use `float` for the default float embeddings.
	// Use `int8` for signed int8 embeddings.
	EmbeddingType *cohereembeddingtype.CohereEmbeddingType `json:"embedding_type,omitempty"`
	// ModelId For a `completion`, `rerank`, or `text_embedding` task, the name of the model
	// to use for the inference task.
	//
	// * For the available `completion` models, refer to the [Cohere command
	// docs](https://docs.cohere.com/docs/models#command).
	// * For the available `rerank` models, refer to the [Cohere rerank
	// docs](https://docs.cohere.com/reference/rerank-1).
	// * For the available `text_embedding` models, refer to [Cohere embed
	// docs](https://docs.cohere.com/reference/embed).
	//
	// The default value for a text embedding task is `embed-english-v2.0`.
	ModelId *string `json:"model_id,omitempty"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// Cohere.
	// By default, the `cohere` service sets the number of requests allowed per
	// minute to 10000.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
	// Similarity The similarity measure.
	// If the `embedding_type` is `float`, the default value is `dot_product`.
	// If the `embedding_type` is `int8` or `byte`, the default value is `cosine`.
	Similarity *coheresimilaritytype.CohereSimilarityType `json:"similarity,omitempty"`
}

func (s *CohereServiceSettings) UnmarshalJSON(data []byte) error {

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

		case "api_key":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ApiKey", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ApiKey = o

		case "embedding_type":
			if err := dec.Decode(&s.EmbeddingType); err != nil {
				return fmt.Errorf("%s | %w", "EmbeddingType", err)
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
			s.ModelId = &o

		case "rate_limit":
			if err := dec.Decode(&s.RateLimit); err != nil {
				return fmt.Errorf("%s | %w", "RateLimit", err)
			}

		case "similarity":
			if err := dec.Decode(&s.Similarity); err != nil {
				return fmt.Errorf("%s | %w", "Similarity", err)
			}

		}
	}
	return nil
}

// NewCohereServiceSettings returns a CohereServiceSettings.
func NewCohereServiceSettings() *CohereServiceSettings {
	r := &CohereServiceSettings{}

	return r
}
