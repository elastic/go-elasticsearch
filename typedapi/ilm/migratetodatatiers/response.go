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

package migratetodatatiers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Response holds the response body struct for the package migratetodatatiers
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ilm/migrate_to_data_tiers/Response.ts#L22-L51
type Response struct {
	DryRun bool `json:"dry_run"`
	// MigratedComponentTemplates The component templates that were updated to not contain custom routing
	// settings for the provided data attribute.
	MigratedComponentTemplates []string `json:"migrated_component_templates"`
	// MigratedComposableTemplates The composable index templates that were updated to not contain custom
	// routing settings for the provided data attribute.
	MigratedComposableTemplates []string `json:"migrated_composable_templates"`
	// MigratedIlmPolicies The ILM policies that were updated.
	MigratedIlmPolicies []string `json:"migrated_ilm_policies"`
	// MigratedIndices The indices that were migrated to tier preference routing.
	MigratedIndices []string `json:"migrated_indices"`
	// MigratedLegacyTemplates The legacy index templates that were updated to not contain custom routing
	// settings for the provided data attribute.
	MigratedLegacyTemplates []string `json:"migrated_legacy_templates"`
	// RemovedLegacyTemplate The name of the legacy index template that was deleted.
	// This information is missing if no legacy index templates were deleted.
	RemovedLegacyTemplate string `json:"removed_legacy_template"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}

func (s *Response) UnmarshalJSON(data []byte) error {
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

		case "dry_run":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DryRun", err)
				}
				s.DryRun = value
			case bool:
				s.DryRun = v
			}

		case "migrated_component_templates":
			if err := dec.Decode(&s.MigratedComponentTemplates); err != nil {
				return fmt.Errorf("%s | %w", "MigratedComponentTemplates", err)
			}

		case "migrated_composable_templates":
			if err := dec.Decode(&s.MigratedComposableTemplates); err != nil {
				return fmt.Errorf("%s | %w", "MigratedComposableTemplates", err)
			}

		case "migrated_ilm_policies":
			if err := dec.Decode(&s.MigratedIlmPolicies); err != nil {
				return fmt.Errorf("%s | %w", "MigratedIlmPolicies", err)
			}

		case "migrated_indices":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "MigratedIndices", err)
				}

				s.MigratedIndices = append(s.MigratedIndices, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.MigratedIndices); err != nil {
					return fmt.Errorf("%s | %w", "MigratedIndices", err)
				}
			}

		case "migrated_legacy_templates":
			if err := dec.Decode(&s.MigratedLegacyTemplates); err != nil {
				return fmt.Errorf("%s | %w", "MigratedLegacyTemplates", err)
			}

		case "removed_legacy_template":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "RemovedLegacyTemplate", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RemovedLegacyTemplate = o

		}
	}
	return nil
}
