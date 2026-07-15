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
// https://github.com/elastic/elasticsearch-specification/tree/37285cbd3fd155f913b50d880b40ec45f9df64b3

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// RequestEmbedding type.
//
// https://github.com/elastic/elasticsearch-specification/blob/37285cbd3fd155f913b50d880b40ec45f9df64b3/specification/inference/_types/CommonTypes.ts#L527-L609
type RequestEmbedding struct {
	// Input Inference input. Either a string, an array of strings, a `content` object, or
	// an array of `content` objects. `content` objects may contain a single item or
	// an array of items. Models that support multiple items per `content` object
	// will return a single embedding for each `content` object, regardless of how
	// many items it contains.
	//
	// string example:
	//
	//	"input": "Some text"
	//
	// string array example:
	//
	//	"input": ["Some text", "Some more text"]
	//
	// `content` object example:
	//
	//	"input": {
	//	    "content": {
	//	      "type": "image",
	//	      "format": "base64",
	//	      "value": "data:image/jpeg;base64,..."
	//	    }
	//	  }
	//
	// `content` object array example:
	//
	//	"input": [
	//	  {
	//	    "content": {
	//	      "type": "text",
	//	      "format": "text",
	//	      "value": "Some text to generate an embedding"
	//	    }
	//	  },
	//	  {
	//	    "content": {
	//	      "type": "image",
	//	      "format": "base64",
	//	      "value": "data:image/jpeg;base64,..."
	//	    }
	//	  }
	//	]
	//
	// Multiple items in one `content` object example:
	//
	//	"input": [
	//	  {
	//	    "content": [
	//	      {
	//	        "type": "image",
	//	        "format": "base64",
	//	        "value": "data:image/jpeg;base64,..."
	//	      },
	//	      {
	//	        "type": "text",
	//	        "value": "Some text to create an embedding"
	//	      }
	//	    ]
	//	  }
	//	]
	Input EmbeddingInput `json:"input"`
	// InputType The input data type for the embedding model. Possible values include:
	//
	//   - `SEARCH`
	//   - `INGEST`
	//   - `CLASSIFICATION`
	//   - `CLUSTERING`
	//
	// Not all models support all values. Unsupported values will trigger a
	// validation exception. Accepted values depend on the configured inference
	// service, refer to the relevant service-specific documentation for more info.
	//
	// > info > The `input_type` parameter specified on the root level of the
	// request body will take precedence over the `input_type` parameter specified
	// in `task_settings`.
	InputType *string `json:"input_type,omitempty"`
	// TaskSettings Task settings for the individual inference request. These settings are
	// specific to the <task_type> you specified and override the task settings
	// specified when initializing the service.
	TaskSettings json.RawMessage `json:"task_settings,omitempty"`
}

func (s *RequestEmbedding) UnmarshalJSON(data []byte) error {

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

		case "task_settings":
			if err := dec.Decode(&s.TaskSettings); err != nil {
				return fmt.Errorf("%s | %w", "TaskSettings", err)
			}

		}
	}
	return nil
}

// NewRequestEmbedding returns a RequestEmbedding.
func NewRequestEmbedding() *RequestEmbedding {
	r := &RequestEmbedding{}

	return r
}

type RequestEmbeddingVariant interface {
	RequestEmbeddingCaster() *RequestEmbedding
}

func (s *RequestEmbedding) RequestEmbeddingCaster() *RequestEmbedding {
	return s
}
