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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/amazonsagemakerapi"
)

// AmazonSageMakerServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L445-L499
type AmazonSageMakerServiceSettings struct {
	// AccessKey A valid AWS access key that has permissions to use Amazon SageMaker and
	// access to models for invoking requests.
	AccessKey string `json:"access_key"`
	// Api The API format to use when calling SageMaker.
	// Elasticsearch will convert the POST _inference request to this data format
	// when invoking the SageMaker endpoint.
	Api amazonsagemakerapi.AmazonSageMakerApi `json:"api"`
	// BatchSize The maximum number of inputs in each batch. This value is used by inference
	// ingestion pipelines
	// when processing semantic values. It correlates to the number of times the
	// SageMaker endpoint is
	// invoked (one per batch of input).
	BatchSize *int `json:"batch_size,omitempty"`
	// Dimensions The number of dimensions returned by the text embedding models. If this value
	// is not provided, then
	// it is guessed by making invoking the endpoint for the `text_embedding` task.
	Dimensions *int `json:"dimensions,omitempty"`
	// EndpointName The name of the SageMaker endpoint.
	EndpointName string `json:"endpoint_name"`
	// InferenceComponentName The inference component to directly invoke when calling a multi-component
	// endpoint.
	InferenceComponentName *string `json:"inference_component_name,omitempty"`
	// Region The region that your endpoint or Amazon Resource Name (ARN) is deployed in.
	// The list of available regions per model can be found in the Amazon SageMaker
	// documentation.
	Region string `json:"region"`
	// SecretKey A valid AWS secret key that is paired with the `access_key`.
	// For information about creating and managing access and secret keys, refer to
	// the AWS documentation.
	SecretKey string `json:"secret_key"`
	// TargetContainerHostname The container to directly invoke when calling a multi-container endpoint.
	TargetContainerHostname *string `json:"target_container_hostname,omitempty"`
	// TargetModel The model ID when calling a multi-model endpoint.
	TargetModel *string `json:"target_model,omitempty"`
}

func (s *AmazonSageMakerServiceSettings) UnmarshalJSON(data []byte) error {

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

		case "access_key":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "AccessKey", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.AccessKey = o

		case "api":
			if err := dec.Decode(&s.Api); err != nil {
				return fmt.Errorf("%s | %w", "Api", err)
			}

		case "batch_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "BatchSize", err)
				}
				s.BatchSize = &value
			case float64:
				f := int(v)
				s.BatchSize = &f
			}

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

		case "endpoint_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "EndpointName", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.EndpointName = o

		case "inference_component_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "InferenceComponentName", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.InferenceComponentName = &o

		case "region":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Region", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Region = o

		case "secret_key":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "SecretKey", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SecretKey = o

		case "target_container_hostname":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "TargetContainerHostname", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TargetContainerHostname = &o

		case "target_model":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "TargetModel", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TargetModel = &o

		}
	}
	return nil
}

// NewAmazonSageMakerServiceSettings returns a AmazonSageMakerServiceSettings.
func NewAmazonSageMakerServiceSettings() *AmazonSageMakerServiceSettings {
	r := &AmazonSageMakerServiceSettings{}

	return r
}
