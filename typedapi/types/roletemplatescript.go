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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"
)

// RoleTemplateScript type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/_types/Privileges.ts#L363-L381
type RoleTemplateScript struct {
	// Id The `id` for a stored script.
	Id *string `json:"id,omitempty"`
	// Lang Specifies the language the script is written in.
	Lang    *scriptlanguage.ScriptLanguage `json:"lang,omitempty"`
	Options map[string]string              `json:"options,omitempty"`
	// Params Specifies any named parameters that are passed into the script as variables.
	// Use parameters instead of hard-coded values to decrease compile time.
	Params map[string]json.RawMessage `json:"params,omitempty"`
	Source RoleTemplateInlineQuery    `json:"source,omitempty"`
}

func (s *RoleTemplateScript) UnmarshalJSON(data []byte) error {

	if !bytes.HasPrefix(data, []byte(`{`)) {
		err := json.NewDecoder(bytes.NewReader(data)).Decode(&s.Source)
		return err
	}

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

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "lang":
			if err := dec.Decode(&s.Lang); err != nil {
				return fmt.Errorf("%s | %w", "Lang", err)
			}

		case "options":
			if s.Options == nil {
				s.Options = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Options); err != nil {
				return fmt.Errorf("%s | %w", "Options", err)
			}

		case "params":
			if s.Params == nil {
				s.Params = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Params); err != nil {
				return fmt.Errorf("%s | %w", "Params", err)
			}

		case "source":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Source", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		source_field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Source", err)
				}

				switch t {

				case "AdditionalQueryProperty", "bool", "boosting", "combined_fields", "common", "constant_score", "dis_max", "distance_feature", "exists", "function_score", "fuzzy", "geo_bounding_box", "geo_distance", "geo_grid", "geo_polygon", "geo_shape", "has_child", "has_parent", "ids", "intervals", "knn", "match", "match_all", "match_bool_prefix", "match_none", "match_phrase", "match_phrase_prefix", "more_like_this", "multi_match", "nested", "parent_id", "percolate", "pinned", "prefix", "query_string", "range", "rank_feature", "regexp", "rule", "script", "script_score", "semantic", "shape", "simple_query_string", "span_containing", "span_field_masking", "span_first", "span_multi", "span_near", "span_not", "span_or", "span_term", "span_within", "sparse_vector", "term", "terms", "terms_set", "text_expansion", "type", "weighted_tokens", "wildcard", "wrapper":
					o := NewQuery()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Source", err)
					}
					s.Source = o
					break source_field

				}
			}
			if s.Source == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.Source); err != nil {
					return fmt.Errorf("%s | %w", "Source", err)
				}
			}

		}
	}
	return nil
}

// NewRoleTemplateScript returns a RoleTemplateScript.
func NewRoleTemplateScript() *RoleTemplateScript {
	r := &RoleTemplateScript{
		Options: make(map[string]string),
		Params:  make(map[string]json.RawMessage),
	}

	return r
}
