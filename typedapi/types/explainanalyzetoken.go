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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ExplainAnalyzeToken type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/indices/analyze/types.ts#L52-L64
type ExplainAnalyzeToken struct {
	Bytes               string                     `json:"bytes"`
	EndOffset           int64                      `json:"end_offset"`
	ExplainAnalyzeToken map[string]json.RawMessage `json:"-"`
	Keyword             *bool                      `json:"keyword,omitempty"`
	Position            int64                      `json:"position"`
	PositionLength      int64                      `json:"positionLength"`
	StartOffset         int64                      `json:"start_offset"`
	TermFrequency       int64                      `json:"termFrequency"`
	Token               string                     `json:"token"`
	Type                string                     `json:"type"`
}

func (s *ExplainAnalyzeToken) UnmarshalJSON(data []byte) error {

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

		case "bytes":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Bytes = o

		case "end_offset":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.EndOffset = value
			case float64:
				f := int64(v)
				s.EndOffset = f
			}

		case "keyword":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Keyword = &value
			case bool:
				s.Keyword = &v
			}

		case "position":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Position = value
			case float64:
				f := int64(v)
				s.Position = f
			}

		case "positionLength":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.PositionLength = value
			case float64:
				f := int64(v)
				s.PositionLength = f
			}

		case "start_offset":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.StartOffset = value
			case float64:
				f := int64(v)
				s.StartOffset = f
			}

		case "termFrequency":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TermFrequency = value
			case float64:
				f := int64(v)
				s.TermFrequency = f
			}

		case "token":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Token = o

		case "type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Type = o

		default:

			if key, ok := t.(string); ok {
				if s.ExplainAnalyzeToken == nil {
					s.ExplainAnalyzeToken = make(map[string]json.RawMessage, 0)
				}
				raw := new(json.RawMessage)
				if err := dec.Decode(&raw); err != nil {
					return err
				}
				s.ExplainAnalyzeToken[key] = *raw
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s ExplainAnalyzeToken) MarshalJSON() ([]byte, error) {
	type opt ExplainAnalyzeToken
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
	for key, value := range s.ExplainAnalyzeToken {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "ExplainAnalyzeToken")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewExplainAnalyzeToken returns a ExplainAnalyzeToken.
func NewExplainAnalyzeToken() *ExplainAnalyzeToken {
	r := &ExplainAnalyzeToken{
		ExplainAnalyzeToken: make(map[string]json.RawMessage, 0),
	}

	return r
}
