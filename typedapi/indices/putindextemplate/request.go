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
// https://github.com/elastic/elasticsearch-specification/tree/f1932ce6b46a53a8342db522b1a7883bcc9e0996

package putindextemplate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putindextemplate
//
// https://github.com/elastic/elasticsearch-specification/blob/f1932ce6b46a53a8342db522b1a7883bcc9e0996/specification/indices/put_index_template/IndicesPutIndexTemplateRequest.ts#L37-L157
type Request struct {

	// AllowAutoCreate This setting overrides the value of the `action.auto_create_index` cluster
	// setting.
	// If set to `true` in a template, then indices can be automatically created
	// using that template even if auto-creation of indices is disabled via
	// `actions.auto_create_index`.
	// If set to `false`, then indices or data streams matching the template must
	// always be explicitly created, and may never be automatically created.
	AllowAutoCreate *bool `json:"allow_auto_create,omitempty"`
	// ComposedOf An ordered list of component template names.
	// Component templates are merged in the order specified, meaning that the last
	// component template specified has the highest precedence.
	ComposedOf []string `json:"composed_of,omitempty"`
	// DataStream If this object is included, the template is used to create data streams and
	// their backing indices.
	// Supports an empty object.
	// Data streams require a matching index template with a `data_stream` object.
	DataStream *types.DataStreamVisibility `json:"data_stream,omitempty"`
	// Deprecated Marks this index template as deprecated. When creating or updating a
	// non-deprecated index template
	// that uses deprecated components, Elasticsearch will emit a deprecation
	// warning.
	Deprecated *bool `json:"deprecated,omitempty"`
	// IgnoreMissingComponentTemplates The configuration option ignore_missing_component_templates can be used when
	// an index template
	// references a component template that might not exist
	IgnoreMissingComponentTemplates []string `json:"ignore_missing_component_templates,omitempty"`
	// IndexPatterns Name of the index template to create.
	IndexPatterns []string `json:"index_patterns,omitempty"`
	// Meta_ Optional user metadata about the index template.
	// It may have any contents.
	// It is not automatically generated or used by Elasticsearch.
	// This user-defined object is stored in the cluster state, so keeping it short
	// is preferable
	// To unset the metadata, replace the template without specifying it.
	Meta_ types.Metadata `json:"_meta,omitempty"`
	// Priority Priority to determine index template precedence when a new data stream or
	// index is created.
	// The index template with the highest priority is chosen.
	// If no priority is specified the template is treated as though it is of
	// priority 0 (lowest priority).
	// This number is not automatically generated by Elasticsearch.
	Priority *int64 `json:"priority,omitempty"`
	// Template Template to be applied.
	// It may optionally include an `aliases`, `mappings`, or `settings`
	// configuration.
	Template *types.IndexTemplateMapping `json:"template,omitempty"`
	// Version Version number used to manage index templates externally.
	// This number is not automatically generated by Elasticsearch.
	// External systems can use these version numbers to simplify template
	// management.
	// To unset a version, replace the template without specifying one.
	Version *int64 `json:"version,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putindextemplate request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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

		case "allow_auto_create":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllowAutoCreate", err)
				}
				s.AllowAutoCreate = &value
			case bool:
				s.AllowAutoCreate = &v
			}

		case "composed_of":
			if err := dec.Decode(&s.ComposedOf); err != nil {
				return fmt.Errorf("%s | %w", "ComposedOf", err)
			}

		case "data_stream":
			if err := dec.Decode(&s.DataStream); err != nil {
				return fmt.Errorf("%s | %w", "DataStream", err)
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

		case "ignore_missing_component_templates":
			if err := dec.Decode(&s.IgnoreMissingComponentTemplates); err != nil {
				return fmt.Errorf("%s | %w", "IgnoreMissingComponentTemplates", err)
			}

		case "index_patterns":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "IndexPatterns", err)
				}

				s.IndexPatterns = append(s.IndexPatterns, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.IndexPatterns); err != nil {
					return fmt.Errorf("%s | %w", "IndexPatterns", err)
				}
			}

		case "_meta":
			if err := dec.Decode(&s.Meta_); err != nil {
				return fmt.Errorf("%s | %w", "Meta_", err)
			}

		case "priority":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Priority", err)
				}
				s.Priority = &value
			case float64:
				f := int64(v)
				s.Priority = &f
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
