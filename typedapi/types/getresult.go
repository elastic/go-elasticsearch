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
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// GetResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_global/get/types.ts#L25-L35
type GetResult struct {
	Fields       map[string]json.RawMessage `json:"fields,omitempty"`
	Found        bool                       `json:"found"`
	Id_          string                     `json:"_id"`
	Index_       string                     `json:"_index"`
	PrimaryTerm_ *int64                     `json:"_primary_term,omitempty"`
	Routing_     *string                    `json:"_routing,omitempty"`
	SeqNo_       *int64                     `json:"_seq_no,omitempty"`
	Source_      json.RawMessage            `json:"_source,omitempty"`
	Version_     *int64                     `json:"_version,omitempty"`
}

func (s *GetResult) UnmarshalJSON(data []byte) error {

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

		case "fields":
			if s.Fields == nil {
				s.Fields = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Fields); err != nil {
				return err
			}

		case "found":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Found = value
			case bool:
				s.Found = v
			}

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return err
			}

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return err
			}

		case "_primary_term":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.PrimaryTerm_ = &value
			case float64:
				f := int64(v)
				s.PrimaryTerm_ = &f
			}

		case "_routing":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Routing_ = &o

		case "_seq_no":
			if err := dec.Decode(&s.SeqNo_); err != nil {
				return err
			}

		case "_source":
			if err := dec.Decode(&s.Source_); err != nil {
				return err
			}

		case "_version":
			if err := dec.Decode(&s.Version_); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewGetResult returns a GetResult.
func NewGetResult() *GetResult {
	r := &GetResult{
		Fields: make(map[string]json.RawMessage, 0),
	}

	return r
}
