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

// FilteringConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/connector/_types/Connector.ts#L209-L213
type FilteringConfig struct {
	Active FilteringRules `json:"active"`
	Domain *string        `json:"domain,omitempty"`
	Draft  FilteringRules `json:"draft"`
}

func (s *FilteringConfig) UnmarshalJSON(data []byte) error {

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

		case "active":
			if err := dec.Decode(&s.Active); err != nil {
				return fmt.Errorf("%s | %w", "Active", err)
			}

		case "domain":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Domain", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Domain = &o

		case "draft":
			if err := dec.Decode(&s.Draft); err != nil {
				return fmt.Errorf("%s | %w", "Draft", err)
			}

		}
	}
	return nil
}

// NewFilteringConfig returns a FilteringConfig.
func NewFilteringConfig() *FilteringConfig {
	r := &FilteringConfig{}

	return r
}

type FilteringConfigVariant interface {
	FilteringConfigCaster() *FilteringConfig
}

func (s *FilteringConfig) FilteringConfigCaster() *FilteringConfig {
	return s
}
