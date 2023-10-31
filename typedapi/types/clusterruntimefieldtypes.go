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
	"strconv"
)

// ClusterRuntimeFieldTypes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cluster/stats/types.ts#L169-L226
type ClusterRuntimeFieldTypes struct {
	// CharsMax Maximum number of characters for a single runtime field script.
	CharsMax int `json:"chars_max"`
	// CharsTotal Total number of characters for the scripts that define the current runtime
	// field data type.
	CharsTotal int `json:"chars_total"`
	// Count Number of runtime fields mapped to the field data type in selected nodes.
	Count int `json:"count"`
	// DocMax Maximum number of accesses to doc_values for a single runtime field script
	DocMax int `json:"doc_max"`
	// DocTotal Total number of accesses to doc_values for the scripts that define the
	// current runtime field data type.
	DocTotal int `json:"doc_total"`
	// IndexCount Number of indices containing a mapping of the runtime field data type in
	// selected nodes.
	IndexCount int `json:"index_count"`
	// Lang Script languages used for the runtime fields scripts.
	Lang []string `json:"lang"`
	// LinesMax Maximum number of lines for a single runtime field script.
	LinesMax int `json:"lines_max"`
	// LinesTotal Total number of lines for the scripts that define the current runtime field
	// data type.
	LinesTotal int `json:"lines_total"`
	// Name Field data type used in selected nodes.
	Name string `json:"name"`
	// ScriptlessCount Number of runtime fields that don’t declare a script.
	ScriptlessCount int `json:"scriptless_count"`
	// ShadowedCount Number of runtime fields that shadow an indexed field.
	ShadowedCount int `json:"shadowed_count"`
	// SourceMax Maximum number of accesses to _source for a single runtime field script.
	SourceMax int `json:"source_max"`
	// SourceTotal Total number of accesses to _source for the scripts that define the current
	// runtime field data type.
	SourceTotal int `json:"source_total"`
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
