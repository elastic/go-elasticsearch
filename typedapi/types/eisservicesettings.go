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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// EisServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/cd5cc9962e79198ac2daf9110c00808293977f13/specification/inference/_types/CommonTypes.ts#L684-L694
type EisServiceSettings struct {
	// ModelId The name of the model to use for the inference task.
	ModelId string `json:"model_id"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned.
	// By default, the `elastic` service sets the number of requests allowed per
	// minute to `240` in case of `chat_completion`.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
}

func (s *EisServiceSettings) UnmarshalJSON(data []byte) error {

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

// NewEisServiceSettings returns a EisServiceSettings.
func NewEisServiceSettings() *EisServiceSettings {
	r := &EisServiceSettings{}

	return r
}

// true

type EisServiceSettingsVariant interface {
	EisServiceSettingsCaster() *EisServiceSettings
}

func (s *EisServiceSettings) EisServiceSettingsCaster() *EisServiceSettings {
	return s
}
