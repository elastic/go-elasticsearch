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
// https://github.com/elastic/elasticsearch-specification/tree/9665eef0d78c41f20c4c83e69b7c8155efd58f24

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

// IndexResultRequestSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9665eef0d78c41f20c4c83e69b7c8155efd58f24/specification/watcher/_types/Actions.ts#L289-L294
type IndexResultRequestSummary struct {
	DocId   *string          `json:"doc_id,omitempty"`
	Index   string           `json:"index"`
	Refresh *refresh.Refresh `json:"refresh,omitempty"`
	Source  json.RawMessage  `json:"source,omitempty"`
}

func (s *IndexResultRequestSummary) UnmarshalJSON(data []byte) error {

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

		case "doc_id":
			if err := dec.Decode(&s.DocId); err != nil {
				return fmt.Errorf("%s | %w", "DocId", err)
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "refresh":
			if err := dec.Decode(&s.Refresh); err != nil {
				return fmt.Errorf("%s | %w", "Refresh", err)
			}

		case "source":
			if err := dec.Decode(&s.Source); err != nil {
				return fmt.Errorf("%s | %w", "Source", err)
			}

		}
	}
	return nil
}

// NewIndexResultRequestSummary returns a IndexResultRequestSummary.
func NewIndexResultRequestSummary() *IndexResultRequestSummary {
	r := &IndexResultRequestSummary{}

	return r
}
