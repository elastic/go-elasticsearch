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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// CreateOperation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_global/bulk/types.ts#L130-L130
type CreateOperation struct {
	// DynamicTemplates A map from the full name of fields to the name of dynamic templates.
	// Defaults to an empty map.
	// If a name matches a dynamic template, then that template will be applied
	// regardless of other match predicates defined in the template.
	// If a field is already defined in the mapping, then this parameter won’t be
	// used.
	DynamicTemplates map[string]string `json:"dynamic_templates,omitempty"`
	// Id_ The document ID.
	Id_           *string `json:"_id,omitempty"`
	IfPrimaryTerm *int64  `json:"if_primary_term,omitempty"`
	IfSeqNo       *int64  `json:"if_seq_no,omitempty"`
	// Index_ Name of the index or index alias to perform the action on.
	Index_ *string `json:"_index,omitempty"`
	// Pipeline ID of the pipeline to use to preprocess incoming documents.
	// If the index has a default ingest pipeline specified, then setting the value
	// to `_none` disables the default ingest pipeline for this request.
	// If a final pipeline is configured it will always run, regardless of the value
	// of this parameter.
	Pipeline *string `json:"pipeline,omitempty"`
	// RequireAlias If `true`, the request’s actions must target an index alias.
	RequireAlias *bool `json:"require_alias,omitempty"`
	// Routing Custom value used to route operations to a specific shard.
	Routing     *string                  `json:"routing,omitempty"`
	Version     *int64                   `json:"version,omitempty"`
	VersionType *versiontype.VersionType `json:"version_type,omitempty"`
}

func (s *CreateOperation) UnmarshalJSON(data []byte) error {

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

		case "dynamic_templates":
			if s.DynamicTemplates == nil {
				s.DynamicTemplates = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.DynamicTemplates); err != nil {
				return err
			}

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return err
			}

		case "if_primary_term":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.IfPrimaryTerm = &value
			case float64:
				f := int64(v)
				s.IfPrimaryTerm = &f
			}

		case "if_seq_no":
			if err := dec.Decode(&s.IfSeqNo); err != nil {
				return err
			}

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return err
			}

		case "pipeline":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Pipeline = &o

		case "require_alias":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.RequireAlias = &value
			case bool:
				s.RequireAlias = &v
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		case "version_type":
			if err := dec.Decode(&s.VersionType); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewCreateOperation returns a CreateOperation.
func NewCreateOperation() *CreateOperation {
	r := &CreateOperation{
		DynamicTemplates: make(map[string]string, 0),
	}

	return r
}
