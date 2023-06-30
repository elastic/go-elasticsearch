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

// ClusterRuntimeFieldTypes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/cluster/stats/types.ts#L116-L131
type ClusterRuntimeFieldTypes struct {
	CharsMax        int      `json:"chars_max"`
	CharsTotal      int      `json:"chars_total"`
	Count           int      `json:"count"`
	DocMax          int      `json:"doc_max"`
	DocTotal        int      `json:"doc_total"`
	IndexCount      int      `json:"index_count"`
	Lang            []string `json:"lang"`
	LinesMax        int      `json:"lines_max"`
	LinesTotal      int      `json:"lines_total"`
	Name            string   `json:"name"`
	ScriptlessCount int      `json:"scriptless_count"`
	ShadowedCount   int      `json:"shadowed_count"`
	SourceMax       int      `json:"source_max"`
	SourceTotal     int      `json:"source_total"`
}

func (s *ClusterRuntimeFieldTypes) UnmarshalJSON(data []byte) error {

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
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.CharsMax = value
			case float64:
				f := int(v)
				s.CharsMax = f
			}

		case "chars_total":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.CharsTotal = value
			case float64:
				f := int(v)
				s.CharsTotal = f
			}

		case "count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Count = value
			case float64:
				f := int(v)
				s.Count = f
			}

		case "doc_max":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.DocMax = value
			case float64:
				f := int(v)
				s.DocMax = f
			}

		case "doc_total":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.DocTotal = value
			case float64:
				f := int(v)
				s.DocTotal = f
			}

		case "index_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.IndexCount = value
			case float64:
				f := int(v)
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
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.LinesMax = value
			case float64:
				f := int(v)
				s.LinesMax = f
			}

		case "lines_total":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.LinesTotal = value
			case float64:
				f := int(v)
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
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ScriptlessCount = value
			case float64:
				f := int(v)
				s.ScriptlessCount = f
			}

		case "shadowed_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ShadowedCount = value
			case float64:
				f := int(v)
				s.ShadowedCount = f
			}

		case "source_max":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.SourceMax = value
			case float64:
				f := int(v)
				s.SourceMax = f
			}

		case "source_total":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.SourceTotal = value
			case float64:
				f := int(v)
				s.SourceTotal = f
			}

		}
	}
	return nil
}

// NewClusterRuntimeFieldTypes returns a ClusterRuntimeFieldTypes.
func NewClusterRuntimeFieldTypes() *ClusterRuntimeFieldTypes {
	r := &ClusterRuntimeFieldTypes{}

	return r
}
