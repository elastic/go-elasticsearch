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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// IndexTemplateDataStreamConfiguration type.
//
// https://github.com/elastic/elasticsearch-specification/blob/52c473efb1fb5320a5bac12572d0b285882862fb/specification/indices/_types/IndexTemplate.ts#L83-L94
type IndexTemplateDataStreamConfiguration struct {
	// AllowCustomRouting If true, the data stream supports custom routing.
	AllowCustomRouting *bool `json:"allow_custom_routing,omitempty"`
	// Hidden If true, the data stream is hidden.
	Hidden *bool `json:"hidden,omitempty"`
}

func (s *IndexTemplateDataStreamConfiguration) UnmarshalJSON(data []byte) error {

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

		case "allow_custom_routing":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllowCustomRouting", err)
				}
				s.AllowCustomRouting = &value
			case bool:
				s.AllowCustomRouting = &v
			}

		case "hidden":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Hidden", err)
				}
				s.Hidden = &value
			case bool:
				s.Hidden = &v
			}

		}
	}
	return nil
}

// NewIndexTemplateDataStreamConfiguration returns a IndexTemplateDataStreamConfiguration.
func NewIndexTemplateDataStreamConfiguration() *IndexTemplateDataStreamConfiguration {
	r := &IndexTemplateDataStreamConfiguration{}

	return r
}

// true

type IndexTemplateDataStreamConfigurationVariant interface {
	IndexTemplateDataStreamConfigurationCaster() *IndexTemplateDataStreamConfiguration
}

func (s *IndexTemplateDataStreamConfiguration) IndexTemplateDataStreamConfigurationCaster() *IndexTemplateDataStreamConfiguration {
	return s
}
