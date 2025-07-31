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

// AzureAiStudioTaskSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L685-L713
type AzureAiStudioTaskSettings struct {
	// DoSample For a `completion` task, instruct the inference process to perform sampling.
	// It has no effect unless `temperature` or `top_p` is specified.
	DoSample *float32 `json:"do_sample,omitempty"`
	// MaxNewTokens For a `completion` task, provide a hint for the maximum number of output
	// tokens to be generated.
	MaxNewTokens *int `json:"max_new_tokens,omitempty"`
	// Temperature For a `completion` task, control the apparent creativity of generated
	// completions with a sampling temperature.
	// It must be a number in the range of 0.0 to 2.0.
	// It should not be used if `top_p` is specified.
	Temperature *float32 `json:"temperature,omitempty"`
	// TopP For a `completion` task, make the model consider the results of the tokens
	// with nucleus sampling probability.
	// It is an alternative value to `temperature` and must be a number in the range
	// of 0.0 to 2.0.
	// It should not be used if `temperature` is specified.
	TopP *float32 `json:"top_p,omitempty"`
	// User For a `text_embedding` task, specify the user issuing the request.
	// This information can be used for abuse detection.
	User *string `json:"user,omitempty"`
}

func (s *AzureAiStudioTaskSettings) UnmarshalJSON(data []byte) error {

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

		case "do_sample":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "DoSample", err)
				}
				f := float32(value)
				s.DoSample = &f
			case float64:
				f := float32(v)
				s.DoSample = &f
			}

		case "max_new_tokens":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxNewTokens", err)
				}
				s.MaxNewTokens = &value
			case float64:
				f := int(v)
				s.MaxNewTokens = &f
			}

		case "temperature":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Temperature", err)
				}
				f := float32(value)
				s.Temperature = &f
			case float64:
				f := float32(v)
				s.Temperature = &f
			}

		case "top_p":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "TopP", err)
				}
				f := float32(value)
				s.TopP = &f
			case float64:
				f := float32(v)
				s.TopP = &f
			}

		case "user":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "User", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.User = &o

		}
	}
	return nil
}

// NewAzureAiStudioTaskSettings returns a AzureAiStudioTaskSettings.
func NewAzureAiStudioTaskSettings() *AzureAiStudioTaskSettings {
	r := &AzureAiStudioTaskSettings{}

	return r
}
