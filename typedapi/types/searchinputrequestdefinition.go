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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/searchtype"
)

// SearchInputRequestDefinition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/watcher/_types/Input.ts#L118-L125
type SearchInputRequestDefinition struct {
	Body               *SearchInputRequestBody    `json:"body,omitempty"`
	Indices            []string                   `json:"indices,omitempty"`
	IndicesOptions     *IndicesOptions            `json:"indices_options,omitempty"`
	RestTotalHitsAsInt *bool                      `json:"rest_total_hits_as_int,omitempty"`
	SearchType         *searchtype.SearchType     `json:"search_type,omitempty"`
	Template           *SearchTemplateRequestBody `json:"template,omitempty"`
}

func (s *SearchInputRequestDefinition) UnmarshalJSON(data []byte) error {

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

		case "body":
			if err := dec.Decode(&s.Body); err != nil {
				return fmt.Errorf("%s | %w", "Body", err)
			}

		case "indices":
			if err := dec.Decode(&s.Indices); err != nil {
				return fmt.Errorf("%s | %w", "Indices", err)
			}

		case "indices_options":
			if err := dec.Decode(&s.IndicesOptions); err != nil {
				return fmt.Errorf("%s | %w", "IndicesOptions", err)
			}

		case "rest_total_hits_as_int":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "RestTotalHitsAsInt", err)
				}
				s.RestTotalHitsAsInt = &value
			case bool:
				s.RestTotalHitsAsInt = &v
			}

		case "search_type":
			if err := dec.Decode(&s.SearchType); err != nil {
				return fmt.Errorf("%s | %w", "SearchType", err)
			}

		case "template":
			if err := dec.Decode(&s.Template); err != nil {
				return fmt.Errorf("%s | %w", "Template", err)
			}

		}
	}
	return nil
}

// NewSearchInputRequestDefinition returns a SearchInputRequestDefinition.
func NewSearchInputRequestDefinition() *SearchInputRequestDefinition {
	r := &SearchInputRequestDefinition{}

	return r
}
