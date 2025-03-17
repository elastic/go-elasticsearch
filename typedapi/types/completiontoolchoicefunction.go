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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// CompletionToolChoiceFunction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ea991724f4dd4f90c496eff547d3cc2e6529f509/specification/inference/chat_completion_unified/UnifiedRequest.ts#L167-L176
type CompletionToolChoiceFunction struct {
	// Name The name of the function to call.
	Name string `json:"name"`
}

func (s *CompletionToolChoiceFunction) UnmarshalJSON(data []byte) error {

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

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = o

		}
	}
	return nil
}

// NewCompletionToolChoiceFunction returns a CompletionToolChoiceFunction.
func NewCompletionToolChoiceFunction() *CompletionToolChoiceFunction {
	r := &CompletionToolChoiceFunction{}

	return r
}

// true

type CompletionToolChoiceFunctionVariant interface {
	CompletionToolChoiceFunctionCaster() *CompletionToolChoiceFunction
}

func (s *CompletionToolChoiceFunction) CompletionToolChoiceFunctionCaster() *CompletionToolChoiceFunction {
	return s
}
