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
// https://github.com/elastic/elasticsearch-specification/tree/9fcf6a64c550d2e8090c8134867f200b56fd7fc7

package rerank

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Request holds the request body struct for the package rerank
//
// https://github.com/elastic/elasticsearch-specification/blob/9fcf6a64c550d2e8090c8134867f200b56fd7fc7/specification/inference/rerank/RerankRequest.ts#L26-L135
type Request struct {
	// Input The documents to rank. The input can be specified as a single string or an
	// array of strings, or as an object or an array of objects. The object form
	// additionally allows specifying non-text inputs, such as images.
	//
	// > info > Only the `elastic` service currently supports non-text inputs for
	// the `rerank` task. For all other services, the input must be a string or an
	// array of strings.
	//
	// string example:
	//
	// 	"input": "some document text"
	//
	// string array example:
	//
	// 	"input": ["some document text", "some more document text"]
	//
	// object example:
	//
	// 	"input": {
	// 	  "type": "image",
	// 	  "format": "base64",
	// 	  "value": "data:image/jpeg;base64,..."
	// 	}
	//
	// object array example:
	//
	// 	"input": [
	// 	  {
	// 	    "type": "text",
	// 	    "format": "text",
	// 	    "value": "some document text"
	// 	  },
	// 	  {
	// 	    "type": "image",
	// 	    "format": "base64",
	// 	    "value": "data:image/jpeg;base64,..."
	// 	  }
	// 	]
	Input types.RerankInput `json:"input"`
	// Query Query input. The query can be specified as a single string, or as an object.
	// The object form additionally allows specifying non-text inputs, such as
	// images.
	//
	// > info > Only the `elastic` service currently supports non-text queries for
	// the `rerank` task. For all other services, the query must be a string.
	//
	// string example:
	//
	// 	"query": "some query text"
	//
	// object example:
	//
	// 	"query": {
	// 	  "type": "image",
	// 	  "format": "base64",
	// 	  "value": "data:image/jpeg;base64,..."
	// 	}
	Query types.RerankQuery `json:"query"`
	// ReturnDocuments Include the document text in the response.
	ReturnDocuments *bool `json:"return_documents,omitempty"`
	// TaskSettings Task settings for the individual inference request. These settings are
	// specific to the task type you specified and override the task settings
	// specified when initializing the service.
	TaskSettings json.RawMessage `json:"task_settings,omitempty"`
	// TopN Limit the response to the top N documents.
	TopN *int `json:"top_n,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Rerank request: %w", err)
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
			if err := dec.Decode(&s.Input); err != nil {
				return fmt.Errorf("%s | %w", "Input", err)
			}

		case "query":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		query_field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Query", err)
				}

				switch t {

				case "format", "type", "value":
					o := types.NewRerankInputObject()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Query", err)
					}
					s.Query = o
					break query_field

				}
			}
			if s.Query == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.Query); err != nil {
					return fmt.Errorf("%s | %w", "Query", err)
				}
			}

		case "return_documents":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ReturnDocuments", err)
				}
				s.ReturnDocuments = &value
			case bool:
				s.ReturnDocuments = &v
			}

		case "task_settings":
			if err := dec.Decode(&s.TaskSettings); err != nil {
				return fmt.Errorf("%s | %w", "TaskSettings", err)
			}

		case "top_n":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TopN", err)
				}
				s.TopN = &value
			case float64:
				f := int(v)
				s.TopN = &f
			}

		}
	}
	return nil
}
