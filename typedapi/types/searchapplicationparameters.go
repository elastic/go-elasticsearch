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
)

// SearchApplicationParameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/search_application/_types/SearchApplicationParameters.ts#L23-L36
type SearchApplicationParameters struct {
	// AnalyticsCollectionName Analytics collection associated to the Search Application.
	AnalyticsCollectionName *string `json:"analytics_collection_name,omitempty"`
	// Indices Indices that are part of the Search Application.
	Indices []string `json:"indices"`
	// Template Search template to use on search operations.
	Template *SearchApplicationTemplate `json:"template,omitempty"`
}

func (s *SearchApplicationParameters) UnmarshalJSON(data []byte) error {

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

		case "analytics_collection_name":
			if err := dec.Decode(&s.AnalyticsCollectionName); err != nil {
				return fmt.Errorf("%s | %w", "AnalyticsCollectionName", err)
			}

		case "indices":
			if err := dec.Decode(&s.Indices); err != nil {
				return fmt.Errorf("%s | %w", "Indices", err)
			}

		case "template":
			if err := dec.Decode(&s.Template); err != nil {
				return fmt.Errorf("%s | %w", "Template", err)
			}

		}
	}
	return nil
}

// NewSearchApplicationParameters returns a SearchApplicationParameters.
func NewSearchApplicationParameters() *SearchApplicationParameters {
	r := &SearchApplicationParameters{}

	return r
}
