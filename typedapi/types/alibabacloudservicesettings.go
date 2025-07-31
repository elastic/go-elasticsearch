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

// AlibabaCloudServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L292-L337
type AlibabaCloudServiceSettings struct {
	// ApiKey A valid API key for the AlibabaCloud AI Search API.
	ApiKey string `json:"api_key"`
	// Host The name of the host address used for the inference task.
	// You can find the host address in the API keys section of the documentation.
	Host string `json:"host"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// AlibabaCloud AI Search.
	// By default, the `alibabacloud-ai-search` service sets the number of requests
	// allowed per minute to `1000`.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
	// ServiceId The name of the model service to use for the inference task.
	// The following service IDs are available for the `completion` task:
	//
	// * `ops-qwen-turbo`
	// * `qwen-turbo`
	// * `qwen-plus`
	// * `qwen-max ÷ qwen-max-longcontext`
	//
	// The following service ID is available for the `rerank` task:
	//
	// * `ops-bge-reranker-larger`
	//
	// The following service ID is available for the `sparse_embedding` task:
	//
	// * `ops-text-sparse-embedding-001`
	//
	// The following service IDs are available for the `text_embedding` task:
	//
	// `ops-text-embedding-001`
	// `ops-text-embedding-zh-001`
	// `ops-text-embedding-en-001`
	// `ops-text-embedding-002`
	ServiceId string `json:"service_id"`
	// Workspace The name of the workspace used for the inference task.
	Workspace string `json:"workspace"`
}

func (s *AlibabaCloudServiceSettings) UnmarshalJSON(data []byte) error {

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

		case "host":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Host", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Host = o

		case "rate_limit":
			if err := dec.Decode(&s.RateLimit); err != nil {
				return fmt.Errorf("%s | %w", "RateLimit", err)
			}

		case "service_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ServiceId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ServiceId = o

		case "workspace":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Workspace", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Workspace = o

		}
	}
	return nil
}

// NewAlibabaCloudServiceSettings returns a AlibabaCloudServiceSettings.
func NewAlibabaCloudServiceSettings() *AlibabaCloudServiceSettings {
	r := &AlibabaCloudServiceSettings{}

	return r
}
