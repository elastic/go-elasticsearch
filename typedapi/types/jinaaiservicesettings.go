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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jinaaielementtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jinaaisimilaritytype"
)

// JinaAIServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e196f9953fa743572ee46884835f1934bce9a16b/specification/inference/_types/CommonTypes.ts#L1593-L1633
type JinaAIServiceSettings struct {
	// ApiKey A valid API key of your JinaAI account.
	//
	// IMPORTANT: You need to provide the API key only once, during the inference
	// model creation.
	// The get inference endpoint API does not retrieve your API key.
	ApiKey string `json:"api_key"`
	// Dimensions For a `text_embedding` task, the number of dimensions the resulting output
	// embeddings should have.
	// By default, the model's standard output dimension is used.
	// Refer to the Jina documentation for more information.
	Dimensions *int `json:"dimensions,omitempty"`
	// ElementType For a `text_embedding` task, the data type returned by the model.
	// Use `bit` for binary embeddings, which are encoded as bytes with signed int8
	// precision.
	// Use `binary` for binary embeddings, which are encoded as bytes with signed
	// int8 precision (this is a synonym of `bit`).
	// Use `float` for the default float embeddings.
	ElementType *jinaaielementtype.JinaAIElementType `json:"element_type,omitempty"`
	// ModelId The name of the model to use for the inference task.
	ModelId string `json:"model_id"`
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

		case "element_type":
			if err := dec.Decode(&s.ElementType); err != nil {
				return fmt.Errorf("%s | %w", "ElementType", err)
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

type JinaAIServiceSettingsVariant interface {
	JinaAIServiceSettingsCaster() *JinaAIServiceSettings
}

func (s *JinaAIServiceSettings) JinaAIServiceSettingsCaster() *JinaAIServiceSettings {
	return s
}
