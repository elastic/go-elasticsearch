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

package inference

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Request holds the request body struct for the package inference
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/inference/InferenceRequest.ts#L26-L104
type Request struct {

	// Input The text on which you want to perform the inference task.
	// It can be a single string or an array.
	//
	// > info
	// > Inference endpoints for the `completion` task type currently only support a
	// single string as input.
	Input []string `json:"input"`
	// InputType Specifies the input data type for the text embedding model. The `input_type`
	// parameter only applies to Inference Endpoints with the `text_embedding` task
	// type. Possible values include:
	// * `SEARCH`
	// * `INGEST`
	// * `CLASSIFICATION`
	// * `CLUSTERING`
	// Not all services support all values. Unsupported values will trigger a
	// validation exception.
	// Accepted values depend on the configured inference service, refer to the
	// relevant service-specific documentation for more info.
	//
	// > info
	// > The `input_type` parameter specified on the root level of the request body
	// will take precedence over the `input_type` parameter specified in
	// `task_settings`.
	InputType *string `json:"input_type,omitempty"`
	// Query The query input, which is required only for the `rerank` task.
	// It is not required for other tasks.
	Query *string `json:"query,omitempty"`
	// TaskSettings Task settings for the individual inference request.
	// These settings are specific to the task type you specified and override the
	// task settings specified when initializing the service.
	TaskSettings json.RawMessage `json:"task_settings,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Inference request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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

		case "input":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Input", err)
				}

				s.Input = append(s.Input, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Input); err != nil {
					return fmt.Errorf("%s | %w", "Input", err)
				}
			}

		case "input_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "InputType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.InputType = &o

		case "query":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Query = &o

		case "task_settings":
			if err := dec.Decode(&s.TaskSettings); err != nil {
				return fmt.Errorf("%s | %w", "TaskSettings", err)
			}

		}
	}
	return nil
}
