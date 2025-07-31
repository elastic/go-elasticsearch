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

// ModelPackageConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/_types/TrainedModel.ts#L257-L272
type ModelPackageConfig struct {
	CreateTime           *int64                     `json:"create_time,omitempty"`
	Description          *string                    `json:"description,omitempty"`
	InferenceConfig      map[string]json.RawMessage `json:"inference_config,omitempty"`
	Metadata             Metadata                   `json:"metadata,omitempty"`
	MinimumVersion       *string                    `json:"minimum_version,omitempty"`
	ModelRepository      *string                    `json:"model_repository,omitempty"`
	ModelType            *string                    `json:"model_type,omitempty"`
	PackagedModelId      string                     `json:"packaged_model_id"`
	PlatformArchitecture *string                    `json:"platform_architecture,omitempty"`
	PrefixStrings        *TrainedModelPrefixStrings `json:"prefix_strings,omitempty"`
	Sha256               *string                    `json:"sha256,omitempty"`
	Size                 ByteSize                   `json:"size,omitempty"`
	Tags                 []string                   `json:"tags,omitempty"`
	VocabularyFile       *string                    `json:"vocabulary_file,omitempty"`
}

func (s *ModelPackageConfig) UnmarshalJSON(data []byte) error {

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

		case "create_time":
			if err := dec.Decode(&s.CreateTime); err != nil {
				return fmt.Errorf("%s | %w", "CreateTime", err)
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "inference_config":
			if s.InferenceConfig == nil {
				s.InferenceConfig = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.InferenceConfig); err != nil {
				return fmt.Errorf("%s | %w", "InferenceConfig", err)
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
			}

		case "minimum_version":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MinimumVersion", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MinimumVersion = &o

		case "model_repository":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ModelRepository", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelRepository = &o

		case "model_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ModelType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelType = &o

		case "packaged_model_id":
			if err := dec.Decode(&s.PackagedModelId); err != nil {
				return fmt.Errorf("%s | %w", "PackagedModelId", err)
			}

		case "platform_architecture":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "PlatformArchitecture", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PlatformArchitecture = &o

		case "prefix_strings":
			if err := dec.Decode(&s.PrefixStrings); err != nil {
				return fmt.Errorf("%s | %w", "PrefixStrings", err)
			}

		case "sha256":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Sha256", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Sha256 = &o

		case "size":
			if err := dec.Decode(&s.Size); err != nil {
				return fmt.Errorf("%s | %w", "Size", err)
			}

		case "tags":
			if err := dec.Decode(&s.Tags); err != nil {
				return fmt.Errorf("%s | %w", "Tags", err)
			}

		case "vocabulary_file":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "VocabularyFile", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.VocabularyFile = &o

		}
	}
	return nil
}

// NewModelPackageConfig returns a ModelPackageConfig.
func NewModelPackageConfig() *ModelPackageConfig {
	r := &ModelPackageConfig{
		InferenceConfig: make(map[string]json.RawMessage),
	}

	return r
}
