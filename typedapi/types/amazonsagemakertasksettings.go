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

// AmazonSageMakerTaskSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L535-L564
type AmazonSageMakerTaskSettings struct {
	// CustomAttributes The AWS custom attributes passed verbatim through to the model running in the
	// SageMaker Endpoint.
	// Values will be returned in the `X-elastic-sagemaker-custom-attributes`
	// header.
	CustomAttributes *string `json:"custom_attributes,omitempty"`
	// EnableExplanations The optional JMESPath expression used to override the EnableExplanations
	// provided during endpoint creation.
	EnableExplanations *string `json:"enable_explanations,omitempty"`
	// InferenceId The capture data ID when enabled in the endpoint.
	InferenceId *string `json:"inference_id,omitempty"`
	// SessionId The stateful session identifier for a new or existing session.
	// New sessions will be returned in the `X-elastic-sagemaker-new-session-id`
	// header.
	// Closed sessions will be returned in the
	// `X-elastic-sagemaker-closed-session-id` header.
	SessionId *string `json:"session_id,omitempty"`
	// TargetVariant Specifies the variant when running with multi-variant Endpoints.
	TargetVariant *string `json:"target_variant,omitempty"`
}

func (s *AmazonSageMakerTaskSettings) UnmarshalJSON(data []byte) error {

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

		case "custom_attributes":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "CustomAttributes", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CustomAttributes = &o

		case "enable_explanations":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "EnableExplanations", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.EnableExplanations = &o

		case "inference_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "InferenceId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.InferenceId = &o

		case "session_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "SessionId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SessionId = &o

		case "target_variant":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "TargetVariant", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TargetVariant = &o

		}
	}
	return nil
}

// NewAmazonSageMakerTaskSettings returns a AmazonSageMakerTaskSettings.
func NewAmazonSageMakerTaskSettings() *AmazonSageMakerTaskSettings {
	r := &AmazonSageMakerTaskSettings{}

	return r
}
