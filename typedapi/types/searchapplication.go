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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// SearchApplication type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/search_application/_types/SearchApplication.ts#L24-L45
type SearchApplication struct {
	// AnalyticsCollectionName Analytics collection associated to the Search Application.
	AnalyticsCollectionName *string `json:"analytics_collection_name,omitempty"`
	// Indices Indices that are part of the Search Application.
	Indices []string `json:"indices"`
	// Name Search Application name.
	Name string `json:"name"`
	// Template Search template to use on search operations.
	Template *SearchApplicationTemplate `json:"template,omitempty"`
	// UpdatedAtMillis Last time the Search Application was updated.
	UpdatedAtMillis int64 `json:"updated_at_millis"`
}

func (s *SearchApplication) UnmarshalJSON(data []byte) error {

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
				return err
			}

		case "indices":
			if err := dec.Decode(&s.Indices); err != nil {
				return err
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "template":
			if err := dec.Decode(&s.Template); err != nil {
				return err
			}

		case "updated_at_millis":
			if err := dec.Decode(&s.UpdatedAtMillis); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewSearchApplication returns a SearchApplication.
func NewSearchApplication() *SearchApplication {
	r := &SearchApplication{}

	return r
}
