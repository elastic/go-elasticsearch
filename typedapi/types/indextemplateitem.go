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
// https://github.com/elastic/elasticsearch-specification/tree/5260ec5b7c899ab1a7939f752218cae07ef07dd7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// IndexTemplateItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5260ec5b7c899ab1a7939f752218cae07ef07dd7/specification/indices/get_index_template/IndicesGetIndexTemplateResponse.ts#L29-L32
type IndexTemplateItem struct {
	IndexTemplate IndexTemplate `json:"index_template"`
	Name          string        `json:"name"`
}

func (s *IndexTemplateItem) UnmarshalJSON(data []byte) error {

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

		case "index_template":
			if err := dec.Decode(&s.IndexTemplate); err != nil {
				return err
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewIndexTemplateItem returns a IndexTemplateItem.
func NewIndexTemplateItem() *IndexTemplateItem {
	r := &IndexTemplateItem{}

	return r
}
