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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// DataStreamMappings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/2514615770f18dbb4e3887cc1a279995dbfd0724/specification/indices/get_data_stream_mappings/IndicesGetDataStreamMappingsResponse.ts#L29-L39
type DataStreamMappings struct {
	// EffectiveMappings The settings specific to this data stream merged with the settings from its
	// template. These `effective_settings`
	// are the settings that will be used when a new index is created for this data
	// stream.
	EffectiveMappings TypeMapping `json:"effective_mappings"`
	// Mappings The settings specific to this data stream
	Mappings TypeMapping `json:"mappings"`
	// Name The name of the data stream.
	Name string `json:"name"`
}

func (s *DataStreamMappings) UnmarshalJSON(data []byte) error {

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

		case "effective_mappings":
			if err := dec.Decode(&s.EffectiveMappings); err != nil {
				return fmt.Errorf("%s | %w", "EffectiveMappings", err)
			}

		case "mappings":
			if err := dec.Decode(&s.Mappings); err != nil {
				return fmt.Errorf("%s | %w", "Mappings", err)
			}

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = o

		}
	}
	return nil
}

// NewDataStreamMappings returns a DataStreamMappings.
func NewDataStreamMappings() *DataStreamMappings {
	r := &DataStreamMappings{}

	return r
}
