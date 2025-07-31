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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// RateLimitSetting type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/inference/_types/Services.ts#L323-L349
type RateLimitSetting struct {
	// RequestsPerMinute The number of requests allowed per minute.
	// By default, the number of requests allowed per minute is set by each service
	// as follows:
	//
	// * `alibabacloud-ai-search` service: `1000`
	// * `anthropic` service: `50`
	// * `azureaistudio` service: `240`
	// * `azureopenai` service and task type `text_embedding`: `1440`
	// * `azureopenai` service and task type `completion`: `120`
	// * `cohere` service: `10000`
	// * `elastic` service and task type `chat_completion`: `240`
	// * `googleaistudio` service: `360`
	// * `googlevertexai` service: `30000`
	// * `hugging_face` service: `3000`
	// * `jinaai` service: `2000`
	// * `mistral` service: `240`
	// * `openai` service and task type `text_embedding`: `3000`
	// * `openai` service and task type `completion`: `500`
	// * `voyageai` service: `2000`
	// * `watsonxai` service: `120`
	RequestsPerMinute *int `json:"requests_per_minute,omitempty"`
}

func (s *RateLimitSetting) UnmarshalJSON(data []byte) error {

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

		case "requests_per_minute":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "RequestsPerMinute", err)
				}
				s.RequestsPerMinute = &value
			case float64:
				f := int(v)
				s.RequestsPerMinute = &f
			}

		}
	}
	return nil
}

// NewRateLimitSetting returns a RateLimitSetting.
func NewRateLimitSetting() *RateLimitSetting {
	r := &RateLimitSetting{}

	return r
}

type RateLimitSettingVariant interface {
	RateLimitSettingCaster() *RateLimitSetting
}

func (s *RateLimitSetting) RateLimitSettingCaster() *RateLimitSetting {
	return s
}
