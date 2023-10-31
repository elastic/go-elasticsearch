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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// LikeDocument type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/query_dsl/specialized.ts#L165-L184
type LikeDocument struct {
	// Doc A document not present in the index.
	Doc    json.RawMessage `json:"doc,omitempty"`
	Fields []string        `json:"fields,omitempty"`
	// Id_ ID of a document.
	Id_ *string `json:"_id,omitempty"`
	// Index_ Index of a document.
	Index_           *string                  `json:"_index,omitempty"`
	PerFieldAnalyzer map[string]string        `json:"per_field_analyzer,omitempty"`
	Routing          *string                  `json:"routing,omitempty"`
	Version          *int64                   `json:"version,omitempty"`
	VersionType      *versiontype.VersionType `json:"version_type,omitempty"`
}

func (s *LikeDocument) UnmarshalJSON(data []byte) error {

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

		case "doc":
			if err := dec.Decode(&s.Doc); err != nil {
				return err
			}

		case "fields":
			if err := dec.Decode(&s.Fields); err != nil {
				return err
			}

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return err
			}

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return err
			}

		case "per_field_analyzer":
			if s.PerFieldAnalyzer == nil {
				s.PerFieldAnalyzer = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.PerFieldAnalyzer); err != nil {
				return err
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		case "version_type":
			if err := dec.Decode(&s.VersionType); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewLikeDocument returns a LikeDocument.
func NewLikeDocument() *LikeDocument {
	r := &LikeDocument{
		PerFieldAnalyzer: make(map[string]string, 0),
	}

	return r
}
