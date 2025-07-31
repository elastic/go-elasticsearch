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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/jinaaisimilaritytype"
)

// JinaAIServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L1439-L1468
type JinaAIServiceSettings struct {
	// ApiKey A valid API key of your JinaAI account.
	//
	// IMPORTANT: You need to provide the API key only once, during the inference
	// model creation.
	// The get inference endpoint API does not retrieve your API key.
	// After creating the inference model, you cannot change the associated API key.
	// If you want to use a different API key, delete the inference model and
	// recreate it with the same name and the updated API key.
	ApiKey string `json:"api_key"`
	// ModelId The name of the model to use for the inference task.
	// For a `rerank` task, it is required.
	// For a `text_embedding` task, it is optional.
	ModelId *string `json:"model_id,omitempty"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// JinaAI.
	// By default, the `jinaai` service sets the number of requests allowed per
	// minute to 2000 for all task types.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
	// Similarity For a `text_embedding` task, the similarity measure. One of cosine,
	// dot_product, l2_norm.
	// The default values varies with the embedding type.
	// For example, a float embedding type uses a `dot_product` similarity measure
	// by default.
	Similarity *jinaaisimilaritytype.JinaAISimilarityType `json:"similarity,omitempty"`
}

func (s *JinaAIServiceSettings) UnmarshalJSON(data []byte) error {

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

// NewJinaAIServiceSettings returns a JinaAIServiceSettings.
func NewJinaAIServiceSettings() *JinaAIServiceSettings {
	r := &JinaAIServiceSettings{}

	return r
}
