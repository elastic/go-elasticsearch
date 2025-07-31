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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// IndexOperation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/bulk/types.ts#L142-L142
type IndexOperation struct {
	// DynamicTemplates A map from the full name of fields to the name of dynamic templates.
	// It defaults to an empty map.
	// If a name matches a dynamic template, that template will be applied
	// regardless of other match predicates defined in the template.
	// If a field is already defined in the mapping, then this parameter won't be
	// used.
	DynamicTemplates map[string]string `json:"dynamic_templates,omitempty"`
	// Id_ The document ID.
	Id_           *string `json:"_id,omitempty"`
	IfPrimaryTerm *int64  `json:"if_primary_term,omitempty"`
	IfSeqNo       *int64  `json:"if_seq_no,omitempty"`
	// Index_ The name of the index or index alias to perform the action on.
	Index_ *string `json:"_index,omitempty"`
	// Pipeline The ID of the pipeline to use to preprocess incoming documents.
	// If the index has a default ingest pipeline specified, setting the value to
	// `_none` turns off the default ingest pipeline for this request.
	// If a final pipeline is configured, it will always run regardless of the value
	// of this parameter.
	Pipeline *string `json:"pipeline,omitempty"`
	// RequireAlias If `true`, the request's actions must target an index alias.
	RequireAlias *bool `json:"require_alias,omitempty"`
	// Routing A custom value used to route operations to a specific shard.
	Routing     *string                  `json:"routing,omitempty"`
	Version     *int64                   `json:"version,omitempty"`
	VersionType *versiontype.VersionType `json:"version_type,omitempty"`
}

func (s *IndexOperation) UnmarshalJSON(data []byte) error {

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
				return fmt.Errorf("%s | %w", "DynamicTemplates", err)
			}

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return fmt.Errorf("%s | %w", "Id_", err)
			}

		case "if_primary_term":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "IfPrimaryTerm", err)
				}
				s.IfPrimaryTerm = &value
			case float64:
				f := int64(v)
				s.IfPrimaryTerm = &f
			}

		case "if_seq_no":
			if err := dec.Decode(&s.IfSeqNo); err != nil {
				return fmt.Errorf("%s | %w", "IfSeqNo", err)
			}

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return fmt.Errorf("%s | %w", "Index_", err)
			}

		case "pipeline":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Pipeline", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Pipeline = &o

		case "require_alias":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "RequireAlias", err)
				}
				s.RequireAlias = &value
			case bool:
				s.RequireAlias = &v
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return fmt.Errorf("%s | %w", "Routing", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		case "version_type":
			if err := dec.Decode(&s.VersionType); err != nil {
				return fmt.Errorf("%s | %w", "VersionType", err)
			}

		}
	}
	return nil
}

// NewIndexOperation returns a IndexOperation.
func NewIndexOperation() *IndexOperation {
	r := &IndexOperation{
		DynamicTemplates: make(map[string]string),
	}

	return r
}
