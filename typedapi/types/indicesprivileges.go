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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexprivilege"
)

// IndicesPrivileges type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/security/_types/Privileges.ts#L82-L105
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
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AllowRestrictedIndices = &value
			case bool:
				s.AllowRestrictedIndices = &v
			}

		case "field_security":
			if err := dec.Decode(&s.FieldSecurity); err != nil {
				return err
			}

		case "names":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Names = append(s.Names, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Names); err != nil {
					return err
				}
			}

		case "privileges":
			if err := dec.Decode(&s.Privileges); err != nil {
				return err
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return err
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
