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
	"io"
	"strconv"
)

// TransformProgress type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/transform/get_transform_stats/types.ts#L45-L51
type TransformProgress struct {
	DocsIndexed     int64   `json:"docs_indexed"`
	DocsProcessed   int64   `json:"docs_processed"`
	DocsRemaining   int64   `json:"docs_remaining"`
	PercentComplete Float64 `json:"percent_complete"`
	TotalDocs       int64   `json:"total_docs"`
}

func (s *TransformProgress) UnmarshalJSON(data []byte) error {

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

		case "docs_indexed":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DocsIndexed = value
			case float64:
				f := int64(v)
				s.DocsIndexed = f
			}

		case "docs_processed":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DocsProcessed = value
			case float64:
				f := int64(v)
				s.DocsProcessed = f
			}

		case "docs_remaining":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DocsRemaining = value
			case float64:
				f := int64(v)
				s.DocsRemaining = f
			}

		case "percent_complete":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.PercentComplete = f
			case float64:
				f := Float64(v)
				s.PercentComplete = f
			}

		case "total_docs":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalDocs = value
			case float64:
				f := int64(v)
				s.TotalDocs = f
			}

		}
	}
	return nil
}

// NewTransformProgress returns a TransformProgress.
func NewTransformProgress() *TransformProgress {
	r := &TransformProgress{}

	return r
}
