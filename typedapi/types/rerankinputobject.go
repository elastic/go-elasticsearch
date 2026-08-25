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
// https://github.com/elastic/elasticsearch-specification/tree/abf9c2c6bb21328339daa197aae15af2ecbc46f0

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rerankinputformat"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rerankinputtype"
)

// An object describing a single input for the `rerank` task, which additionally
// allows specifying non-text inputs, such as images.
//
// https://github.com/elastic/elasticsearch-specification/blob/abf9c2c6bb21328339daa197aae15af2ecbc46f0/specification/inference/rerank/RerankRequest.ts#L168-L185
type RerankInputObject struct {
	// Format The format of the input. For the `text` type this must be `text`. For the
	// `image` type this must be `base64`. If not specified, this defaults to `text`
	// for the `text` type and `base64` for the `image` type.
	Format *rerankinputformat.RerankInputFormat `json:"format,omitempty"`
	// Type The type of input. Not all services and models support all input types.
	Type rerankinputtype.RerankInputType `json:"type"`
	// Value The value of the input. For images, this must be a base64-encoded data URI,
	// that is, "data:content/type;base64,...".
	Value string `json:"value"`
}

func (s *RerankInputObject) UnmarshalJSON(data []byte) error {

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

		case "format":
			if err := dec.Decode(&s.Format); err != nil {
				return fmt.Errorf("%s | %w", "Format", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Value", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Value = o

		}
	}
	return nil
}

// NewRerankInputObject returns a RerankInputObject.
func NewRerankInputObject() *RerankInputObject {
	r := &RerankInputObject{}

	return r
}

type RerankInputObjectVariant interface {
	RerankInputObjectCaster() *RerankInputObject
}

func (s *RerankInputObject) RerankInputObjectCaster() *RerankInputObject {
	return s
}

func (s *RerankInputObject) RerankObjectInputCaster() *RerankObjectInput {
	if s == nil {
		return nil
	}
	o := RerankObjectInput{*s}
	return &o
}

func (s *RerankInputObject) RerankQueryCaster() *RerankQuery {
	if s == nil {
		return nil
	}
	o := RerankQuery(s)
	return &o
}
