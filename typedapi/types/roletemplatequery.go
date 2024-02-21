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
// https://github.com/elastic/elasticsearch-specification/tree/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// RoleTemplateQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/security/_types/Privileges.ts#L141-L151
type RoleTemplateQuery struct {
	// Template When you create a role, you can specify a query that defines the document
	// level security permissions. You can optionally
	// use Mustache templates in the role query to insert the username of the
	// current authenticated user into the role.
	// Like other places in Elasticsearch that support templating or scripting, you
	// can specify inline, stored, or file-based
	// templates and define custom parameters. You access the details for the
	// current authenticated user through the _user parameter.
	Template RoleTemplateScript `json:"template,omitempty"`
}

func (s *RoleTemplateQuery) UnmarshalJSON(data []byte) error {

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

		case "template":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return err
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return err
				}

				switch t {

				case "lang", "options", "source":
					o := NewRoleTemplateInlineScript()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Template = o

				case "id":
					o := NewStoredScriptId()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Template = o

				}
			}

		}
	}
	return nil
}

// NewRoleTemplateQuery returns a RoleTemplateQuery.
func NewRoleTemplateQuery() *RoleTemplateQuery {
	r := &RoleTemplateQuery{}

	return r
}
