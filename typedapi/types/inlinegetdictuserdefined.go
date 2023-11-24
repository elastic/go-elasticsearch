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
	"fmt"
	"io"
	"strconv"
)

// InlineGetDictUserDefined type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/common.ts#L321-L330
type InlineGetDictUserDefined struct {
	Fields                   map[string]json.RawMessage `json:"fields,omitempty"`
	Found                    bool                       `json:"found"`
	InlineGetDictUserDefined map[string]json.RawMessage `json:"-"`
	PrimaryTerm_             *int64                     `json:"_primary_term,omitempty"`
	Routing_                 *string                    `json:"_routing,omitempty"`
	SeqNo_                   *int64                     `json:"_seq_no,omitempty"`
	Source_                  map[string]json.RawMessage `json:"_source"`
}

func (s *InlineGetDictUserDefined) UnmarshalJSON(data []byte) error {

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
			if err := dec.Decode(&s.Routing_); err != nil {
				return err
			}

		case "_seq_no":
			if err := dec.Decode(&s.SeqNo_); err != nil {
				return err
			}

		case "_source":
			if s.Source_ == nil {
				s.Source_ = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Source_); err != nil {
				return err
			}

		default:

			if key, ok := t.(string); ok {
				if s.InlineGetDictUserDefined == nil {
					s.InlineGetDictUserDefined = make(map[string]json.RawMessage, 0)
				}
				raw := new(json.RawMessage)
				if err := dec.Decode(&raw); err != nil {
					return err
				}
				s.InlineGetDictUserDefined[key] = *raw
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s InlineGetDictUserDefined) MarshalJSON() ([]byte, error) {
	type opt InlineGetDictUserDefined
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.InlineGetDictUserDefined {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "InlineGetDictUserDefined")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewInlineGetDictUserDefined returns a InlineGetDictUserDefined.
func NewInlineGetDictUserDefined() *InlineGetDictUserDefined {
	r := &InlineGetDictUserDefined{
		Fields:                   make(map[string]json.RawMessage, 0),
		InlineGetDictUserDefined: make(map[string]json.RawMessage, 0),
		Source_:                  make(map[string]json.RawMessage, 0),
	}

	return r
}
