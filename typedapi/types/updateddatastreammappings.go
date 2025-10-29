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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// UpdatedDataStreamMappings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/d520d9e8cf14cad487de5e0654007686c395b494/specification/indices/put_data_stream_mappings/IndicesPutDataStreamMappingsResponse.ts#L30-L51
type UpdatedDataStreamMappings struct {
	// AppliedToDataStream If the mappings were successfully applied to the data stream (or would have
	// been, if running in `dry_run`
	// mode), it is `true`. If an error occurred, it is `false`.
	AppliedToDataStream bool `json:"applied_to_data_stream"`
	// EffectiveMappings The mappings that are effective on this data stream, taking into account the
	// mappings from the matching index
	// template and the mappings specific to this data stream.
	EffectiveMappings *TypeMapping `json:"effective_mappings,omitempty"`
	// Error A message explaining why the mappings could not be applied to the data
	// stream.
	Error *string `json:"error,omitempty"`
	// Mappings The mappings that are specfic to this data stream that will override any
	// mappings from the matching index template.
	Mappings *TypeMapping `json:"mappings,omitempty"`
	// Name The data stream name.
	Name string `json:"name"`
}

func (s *UpdatedDataStreamMappings) UnmarshalJSON(data []byte) error {

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

		case "effective_mappings":
			if err := dec.Decode(&s.EffectiveMappings); err != nil {
				return fmt.Errorf("%s | %w", "EffectiveMappings", err)
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

		case "mappings":
			if err := dec.Decode(&s.Mappings); err != nil {
				return fmt.Errorf("%s | %w", "Mappings", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		}
	}
	return nil
}

// NewUpdatedDataStreamMappings returns a UpdatedDataStreamMappings.
func NewUpdatedDataStreamMappings() *UpdatedDataStreamMappings {
	r := &UpdatedDataStreamMappings{}

	return r
}
