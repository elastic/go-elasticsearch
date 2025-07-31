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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// UpdatedDataStreamSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/put_data_stream_settings/IndicesPutDataStreamSettingsResponse.ts#L30-L55
type UpdatedDataStreamSettings struct {
	// AppliedToDataStream If the settings were successfully applied to the data stream (or would have
	// been, if running in `dry_run`
	// mode), it is `true`. If an error occurred, it is `false`.
	AppliedToDataStream bool `json:"applied_to_data_stream"`
	// EffectiveSettings The settings that are effective on this data stream, taking into account the
	// settings from the matching index
	// template and the settings specific to this data stream.
	EffectiveSettings IndexSettings `json:"effective_settings"`
	// Error A message explaining why the settings could not be applied to the data
	// stream.
	Error *string `json:"error,omitempty"`
	// IndexSettingsResults Information about whether and where each setting was applied.
	IndexSettingsResults IndexSettingResults `json:"index_settings_results"`
	// Name The data stream name.
	Name string `json:"name"`
	// Settings The settings that are specfic to this data stream that will override any
	// settings from the matching index template.
	Settings IndexSettings `json:"settings"`
}

func (s *UpdatedDataStreamSettings) UnmarshalJSON(data []byte) error {

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

		case "applied_to_data_stream":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AppliedToDataStream", err)
				}
				s.AppliedToDataStream = value
			case bool:
				s.AppliedToDataStream = v
			}

		case "effective_settings":
			if err := dec.Decode(&s.EffectiveSettings); err != nil {
				return fmt.Errorf("%s | %w", "EffectiveSettings", err)
			}

		case "error":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Error", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Error = &o

		case "index_settings_results":
			if err := dec.Decode(&s.IndexSettingsResults); err != nil {
				return fmt.Errorf("%s | %w", "IndexSettingsResults", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "settings":
			if err := dec.Decode(&s.Settings); err != nil {
				return fmt.Errorf("%s | %w", "Settings", err)
			}

		}
	}
	return nil
}

// NewUpdatedDataStreamSettings returns a UpdatedDataStreamSettings.
func NewUpdatedDataStreamSettings() *UpdatedDataStreamSettings {
	r := &UpdatedDataStreamSettings{}

	return r
}
