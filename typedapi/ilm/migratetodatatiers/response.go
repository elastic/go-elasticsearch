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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package migratetodatatiers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// Response holds the response body struct for the package migratetodatatiers
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ilm/migrate_to_data_tiers/Response.ts#L22-L32

type Response struct {
	DryRun                      bool     `json:"dry_run"`
	MigratedComponentTemplates  []string `json:"migrated_component_templates"`
	MigratedComposableTemplates []string `json:"migrated_composable_templates"`
	MigratedIlmPolicies         []string `json:"migrated_ilm_policies"`
	MigratedIndices             []string `json:"migrated_indices"`
	MigratedLegacyTemplates     []string `json:"migrated_legacy_templates"`
	RemovedLegacyTemplate       string   `json:"removed_legacy_template"`
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
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.DryRun = value
			case bool:
				s.DryRun = v
			}

		case "migrated_component_templates":
			if err := dec.Decode(&s.MigratedComponentTemplates); err != nil {
				return err
			}

		case "migrated_composable_templates":
			if err := dec.Decode(&s.MigratedComposableTemplates); err != nil {
				return err
			}

		case "migrated_ilm_policies":
			if err := dec.Decode(&s.MigratedIlmPolicies); err != nil {
				return err
			}

		case "migrated_indices":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.MigratedIndices = append(s.MigratedIndices, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.MigratedIndices); err != nil {
					return err
				}
			}

		case "migrated_legacy_templates":
			if err := dec.Decode(&s.MigratedLegacyTemplates); err != nil {
				return err
			}

		case "removed_legacy_template":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
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
