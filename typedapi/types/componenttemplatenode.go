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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ComponentTemplateNode type.
//
// https://github.com/elastic/elasticsearch-specification/blob/6785a6caa1fa3ca5ab3308963d79dce923a3469f/specification/cluster/_types/ComponentTemplate.ts#L34-L67
type ComponentTemplateNode struct {
	// CreatedDate Date and time when the component template was created. Only returned if the
	// `human` query parameter is `true`.
	CreatedDate DateTime `json:"created_date,omitempty"`
	// CreatedDateMillis Date and time when the component template was created, in milliseconds since
	// the epoch.
	CreatedDateMillis *int64   `json:"created_date_millis,omitempty"`
	Deprecated        *bool    `json:"deprecated,omitempty"`
	Meta_             Metadata `json:"_meta,omitempty"`
	// ModifiedDate Date and time when the component template was last modified. Only returned if
	// the `human` query parameter is `true`.
	ModifiedDate DateTime `json:"modified_date,omitempty"`
	// ModifiedDateMillis Date and time when the component template was last modified, in milliseconds
	// since the epoch.
	ModifiedDateMillis *int64                   `json:"modified_date_millis,omitempty"`
	Template           ComponentTemplateSummary `json:"template"`
	Version            *int64                   `json:"version,omitempty"`
}

func (s *ComponentTemplateNode) UnmarshalJSON(data []byte) error {

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

		case "created_date":
			if err := dec.Decode(&s.CreatedDate); err != nil {
				return fmt.Errorf("%s | %w", "CreatedDate", err)
			}

		case "created_date_millis":
			if err := dec.Decode(&s.CreatedDateMillis); err != nil {
				return fmt.Errorf("%s | %w", "CreatedDateMillis", err)
			}

		case "deprecated":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Deprecated", err)
				}
				s.Deprecated = &value
			case bool:
				s.Deprecated = &v
			}

		case "_meta":
			if err := dec.Decode(&s.Meta_); err != nil {
				return fmt.Errorf("%s | %w", "Meta_", err)
			}

		case "modified_date":
			if err := dec.Decode(&s.ModifiedDate); err != nil {
				return fmt.Errorf("%s | %w", "ModifiedDate", err)
			}

		case "modified_date_millis":
			if err := dec.Decode(&s.ModifiedDateMillis); err != nil {
				return fmt.Errorf("%s | %w", "ModifiedDateMillis", err)
			}

		case "template":
			if err := dec.Decode(&s.Template); err != nil {
				return fmt.Errorf("%s | %w", "Template", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewComponentTemplateNode returns a ComponentTemplateNode.
func NewComponentTemplateNode() *ComponentTemplateNode {
	r := &ComponentTemplateNode{}

	return r
}

type ComponentTemplateNodeVariant interface {
	ComponentTemplateNodeCaster() *ComponentTemplateNode
}

func (s *ComponentTemplateNode) ComponentTemplateNodeCaster() *ComponentTemplateNode {
	return s
}
