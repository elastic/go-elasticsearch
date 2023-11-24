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
	"strconv"
)

// RuntimeFieldsType type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/xpack/usage/types.ts#L279-L294
type RuntimeFieldsType struct {
	CharsMax        int64    `json:"chars_max"`
	CharsTotal      int64    `json:"chars_total"`
	Count           int64    `json:"count"`
	DocMax          int64    `json:"doc_max"`
	DocTotal        int64    `json:"doc_total"`
	IndexCount      int64    `json:"index_count"`
	Lang            []string `json:"lang"`
	LinesMax        int64    `json:"lines_max"`
	LinesTotal      int64    `json:"lines_total"`
	Name            string   `json:"name"`
	ScriptlessCount int64    `json:"scriptless_count"`
	ShadowedCount   int64    `json:"shadowed_count"`
	SourceMax       int64    `json:"source_max"`
	SourceTotal     int64    `json:"source_total"`
}

func (s *RuntimeFieldsType) UnmarshalJSON(data []byte) error {

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

		case "chars_max":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.CharsMax = value
			case float64:
				f := int64(v)
				s.CharsMax = f
			}

		case "chars_total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.CharsTotal = value
			case float64:
				f := int64(v)
				s.CharsTotal = f
			}

		case "count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Count = value
			case float64:
				f := int64(v)
				s.Count = f
			}

		case "doc_max":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DocMax = value
			case float64:
				f := int64(v)
				s.DocMax = f
			}

		case "doc_total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DocTotal = value
			case float64:
				f := int64(v)
				s.DocTotal = f
			}

		case "index_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.IndexCount = value
			case float64:
				f := int64(v)
				s.IndexCount = f
			}

		case "lang":
			if err := dec.Decode(&s.Lang); err != nil {
				return err
			}

		case "lines_max":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.LinesMax = value
			case float64:
				f := int64(v)
				s.LinesMax = f
			}

		case "lines_total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.LinesTotal = value
			case float64:
				f := int64(v)
				s.LinesTotal = f
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "scriptless_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ScriptlessCount = value
			case float64:
				f := int64(v)
				s.ScriptlessCount = f
			}

		case "shadowed_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ShadowedCount = value
			case float64:
				f := int64(v)
				s.ShadowedCount = f
			}

		case "source_max":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.SourceMax = value
			case float64:
				f := int64(v)
				s.SourceMax = f
			}

		case "source_total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.SourceTotal = value
			case float64:
				f := int64(v)
				s.SourceTotal = f
			}

		}
	}
	return nil
}

// NewRuntimeFieldsType returns a RuntimeFieldsType.
func NewRuntimeFieldsType() *RuntimeFieldsType {
	r := &RuntimeFieldsType{}

	return r
}
