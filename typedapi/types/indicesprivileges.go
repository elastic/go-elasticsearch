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
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexprivilege"
)

// IndicesPrivileges type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/_types/Privileges.ts#L216-L242
type IndicesPrivileges struct {
	// AllowRestrictedIndices Set to `true` if using wildcard or regular expressions for patterns that
	// cover restricted indices. Implicitly, restricted indices have limited
	// privileges that can cause pattern tests to fail. If restricted indices are
	// explicitly included in the `names` list, Elasticsearch checks privileges
	// against these indices regardless of the value set for
	// `allow_restricted_indices`.
	AllowRestrictedIndices *bool `json:"allow_restricted_indices,omitempty"`
	// FieldSecurity The document fields that the owners of the role have read access to.
	FieldSecurity *FieldSecurity `json:"field_security,omitempty"`
	// Names A list of indices (or index name patterns) to which the permissions in this
	// entry apply.
	Names []string `json:"names"`
	// Privileges The index level privileges that owners of the role have on the specified
	// indices.
	Privileges []indexprivilege.IndexPrivilege `json:"privileges"`
	// Query A search query that defines the documents the owners of the role have access
	// to. A document within the specified indices must match this query for it to
	// be accessible by the owners of the role.
	Query IndicesPrivilegesQuery `json:"query,omitempty"`
}

func (s *IndicesPrivileges) UnmarshalJSON(data []byte) error {

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

		case "allow_restricted_indices":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllowRestrictedIndices", err)
				}
				s.AllowRestrictedIndices = &value
			case bool:
				s.AllowRestrictedIndices = &v
			}

		case "field_security":
			if err := dec.Decode(&s.FieldSecurity); err != nil {
				return fmt.Errorf("%s | %w", "FieldSecurity", err)
			}

		case "names":
			if err := dec.Decode(&s.Names); err != nil {
				return fmt.Errorf("%s | %w", "Names", err)
			}

		case "privileges":
			if err := dec.Decode(&s.Privileges); err != nil {
				return fmt.Errorf("%s | %w", "Privileges", err)
			}

		case "query":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		query_field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Query", err)
				}

				switch t {

				case "AdditionalQueryProperty", "bool", "boosting", "combined_fields", "common", "constant_score", "dis_max", "distance_feature", "exists", "function_score", "fuzzy", "geo_bounding_box", "geo_distance", "geo_grid", "geo_polygon", "geo_shape", "has_child", "has_parent", "ids", "intervals", "knn", "match", "match_all", "match_bool_prefix", "match_none", "match_phrase", "match_phrase_prefix", "more_like_this", "multi_match", "nested", "parent_id", "percolate", "pinned", "prefix", "query_string", "range", "rank_feature", "regexp", "rule", "script", "script_score", "semantic", "shape", "simple_query_string", "span_containing", "span_field_masking", "span_first", "span_multi", "span_near", "span_not", "span_or", "span_term", "span_within", "sparse_vector", "term", "terms", "terms_set", "text_expansion", "type", "weighted_tokens", "wildcard", "wrapper":
					o := NewQuery()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Query", err)
					}
					s.Query = o
					break query_field

				case "template":
					o := NewRoleTemplateQuery()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Query", err)
					}
					s.Query = o
					break query_field

				}
			}
			if s.Query == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.Query); err != nil {
					return fmt.Errorf("%s | %w", "Query", err)
				}
			}

		}
	}
	return nil
}

// NewIndicesPrivileges returns a IndicesPrivileges.
func NewIndicesPrivileges() *IndicesPrivileges {
	r := &IndicesPrivileges{}

	return r
}
