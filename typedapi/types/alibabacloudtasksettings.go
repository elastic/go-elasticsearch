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

// AlibabaCloudTaskSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L339-L353
type AlibabaCloudTaskSettings struct {
	// InputType For a `sparse_embedding` or `text_embedding` task, specify the type of input
	// passed to the model.
	// Valid values are:
	//
	// * `ingest` for storing document embeddings in a vector database.
	// * `search` for storing embeddings of search queries run against a vector
	// database to find relevant documents.
	InputType *string `json:"input_type,omitempty"`
	// ReturnToken For a `sparse_embedding` task, it affects whether the token name will be
	// returned in the response.
	// It defaults to `false`, which means only the token ID will be returned in the
	// response.
	ReturnToken *bool `json:"return_token,omitempty"`
}

func (s *AlibabaCloudTaskSettings) UnmarshalJSON(data []byte) error {

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

		case "return_token":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ReturnToken", err)
				}
				s.ReturnToken = &value
			case bool:
				s.ReturnToken = &v
			}

		}
	}
	return nil
}

// NewAlibabaCloudTaskSettings returns a AlibabaCloudTaskSettings.
func NewAlibabaCloudTaskSettings() *AlibabaCloudTaskSettings {
	r := &AlibabaCloudTaskSettings{}

	return r
}
