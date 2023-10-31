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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// SearchApplicationListItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/search_application/list/SearchApplicationsListResponse.ts#L31-L48
type SearchApplicationListItem struct {
	// AnalyticsCollectionName Analytics collection associated to the Search Application
	AnalyticsCollectionName *string `json:"analytics_collection_name,omitempty"`
	// Indices Indices that are part of the Search Application
	Indices []string `json:"indices"`
	// Name Search Application name
	Name string `json:"name"`
	// UpdatedAtMillis Last time the Search Application was updated
	UpdatedAtMillis int64 `json:"updated_at_millis"`
}

func (s *SearchApplicationListItem) UnmarshalJSON(data []byte) error {

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

		case "updated_at_millis":
			if err := dec.Decode(&s.UpdatedAtMillis); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewSearchApplicationListItem returns a SearchApplicationListItem.
func NewSearchApplicationListItem() *SearchApplicationListItem {
	r := &SearchApplicationListItem{}

	return r
}
