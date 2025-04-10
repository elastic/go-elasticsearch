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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package types

import (
	"encoding/json"
	"fmt"
)

// QueryVectorBuilder type.
//
// https://github.com/elastic/elasticsearch-specification/blob/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9/specification/_types/Knn.ts#L89-L92
type QueryVectorBuilder struct {
	AdditionalQueryVectorBuilderProperty map[string]json.RawMessage `json:"-"`
	TextEmbedding                        *TextEmbedding             `json:"text_embedding,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s QueryVectorBuilder) MarshalJSON() ([]byte, error) {
	type opt QueryVectorBuilder
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.AdditionalQueryVectorBuilderProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalQueryVectorBuilderProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewQueryVectorBuilder returns a QueryVectorBuilder.
func NewQueryVectorBuilder() *QueryVectorBuilder {
	r := &QueryVectorBuilder{
		AdditionalQueryVectorBuilderProperty: make(map[string]json.RawMessage),
	}

	return r
}

// true

type QueryVectorBuilderVariant interface {
	QueryVectorBuilderCaster() *QueryVectorBuilder
}

func (s *QueryVectorBuilder) QueryVectorBuilderCaster() *QueryVectorBuilder {
	return s
}
