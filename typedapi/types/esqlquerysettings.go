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
// https://github.com/elastic/elasticsearch-specification/tree/7560c979602e6941815872bdaec801200bc7ec4e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Per-query settings supplied through the request body. This is the
// request-body equivalent of the in-query `SET` command. Only settings that are
// exposed as request-body parameters can be set here; other `SET`-only settings
// (such as `unmapped_fields`) must be supplied in the query itself.
//
// https://github.com/elastic/elasticsearch-specification/blob/7560c979602e6941815872bdaec801200bc7ec4e/specification/esql/_types/types.ts#L177-L214
type EsqlQuerySettings struct {
	// Approximation Enables query approximation if possible for the query. `false` (the default)
	// disables query approximation and `true` enables it with default settings. A
	// map value enables query approximation with custom settings.
	Approximation EsqlApproximation `json:"approximation,omitempty"`
	// ColumnMetadata When enabled, column metadata is added to the query response as additional
	// `_meta` properties. Currently, only `_meta.bucket` is added for columns
	// corresponding to the `BUCKET` function, containing the bucket interval and
	// unit for queries where it can be determined.
	ColumnMetadata Stringifiedboolean `json:"column_metadata,omitempty"`
	// ProjectRouting Limits the scope of a cross-project search (CPS) to specific projects before
	// query execution, based on a Lucene query expression evaluated against project
	// tags. Excluded projects are not queried, which can reduce cost and latency.
	ProjectRouting *string `json:"project_routing,omitempty"`
	// TimeZone The default timezone to be used in the query. It defaults to UTC and
	// overrides the `time_zone` request parameter.
	TimeZone *string `json:"time_zone,omitempty"`
}

func (s *EsqlQuerySettings) UnmarshalJSON(data []byte) error {

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

		case "approximation":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Approximation", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		approximation_field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Approximation", err)
				}

				switch t {

				case "confidence_level", "rows":
					o := NewEsqlApproximationSettings()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Approximation", err)
					}
					s.Approximation = o
					break approximation_field

				}
			}
			if s.Approximation == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.Approximation); err != nil {
					return fmt.Errorf("%s | %w", "Approximation", err)
				}
			}

		case "column_metadata":
			if err := dec.Decode(&s.ColumnMetadata); err != nil {
				return fmt.Errorf("%s | %w", "ColumnMetadata", err)
			}

		case "project_routing":
			if err := dec.Decode(&s.ProjectRouting); err != nil {
				return fmt.Errorf("%s | %w", "ProjectRouting", err)
			}

		case "time_zone":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "TimeZone", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TimeZone = &o

		}
	}
	return nil
}

// NewEsqlQuerySettings returns a EsqlQuerySettings.
func NewEsqlQuerySettings() *EsqlQuerySettings {
	r := &EsqlQuerySettings{}

	return r
}

type EsqlQuerySettingsVariant interface {
	EsqlQuerySettingsCaster() *EsqlQuerySettings
}

func (s *EsqlQuerySettings) EsqlQuerySettingsCaster() *EsqlQuerySettings {
	return s
}
