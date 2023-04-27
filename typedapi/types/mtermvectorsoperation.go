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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"

	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// MTermVectorsOperation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_global/mtermvectors/types.ts#L35-L49
type MTermVectorsOperation struct {
	Doc             json.RawMessage          `json:"doc,omitempty"`
	FieldStatistics *bool                    `json:"field_statistics,omitempty"`
	Fields          []string                 `json:"fields,omitempty"`
	Filter          *TermVectorsFilter       `json:"filter,omitempty"`
	Id_             string                   `json:"_id"`
	Index_          *string                  `json:"_index,omitempty"`
	Offsets         *bool                    `json:"offsets,omitempty"`
	Payloads        *bool                    `json:"payloads,omitempty"`
	Positions       *bool                    `json:"positions,omitempty"`
	Routing         *string                  `json:"routing,omitempty"`
	TermStatistics  *bool                    `json:"term_statistics,omitempty"`
	Version         *int64                   `json:"version,omitempty"`
	VersionType     *versiontype.VersionType `json:"version_type,omitempty"`
}

func (s *MTermVectorsOperation) UnmarshalJSON(data []byte) error {

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

		case "field_statistics":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.FieldStatistics = &value
			case bool:
				s.FieldStatistics = &v
			}

		case "fields":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Fields = append(s.Fields, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Fields); err != nil {
					return err
				}
			}

		case "filter":
			if err := dec.Decode(&s.Filter); err != nil {
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

		case "offsets":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Offsets = &value
			case bool:
				s.Offsets = &v
			}

		case "payloads":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Payloads = &value
			case bool:
				s.Payloads = &v
			}

		case "positions":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Positions = &value
			case bool:
				s.Positions = &v
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return err
			}

		case "term_statistics":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.TermStatistics = &value
			case bool:
				s.TermStatistics = &v
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

// NewMTermVectorsOperation returns a MTermVectorsOperation.
func NewMTermVectorsOperation() *MTermVectorsOperation {
	r := &MTermVectorsOperation{}

	return r
}
