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
)

// DatabaseConfiguration type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ingest/_types/Database.ts#L22-L37
type DatabaseConfiguration struct {
	AdditionalDatabaseConfigurationProperty map[string]json.RawMessage `json:"-"`
	Ipinfo                                  *Ipinfo                    `json:"ipinfo,omitempty"`
	Maxmind                                 *Maxmind                   `json:"maxmind,omitempty"`
	// Name The provider-assigned name of the IP geolocation database to download.
	Name string `json:"name"`
}

func (s *DatabaseConfiguration) UnmarshalJSON(data []byte) error {

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

		case "ipinfo":
			if err := dec.Decode(&s.Ipinfo); err != nil {
				return fmt.Errorf("%s | %w", "Ipinfo", err)
			}

		case "maxmind":
			if err := dec.Decode(&s.Maxmind); err != nil {
				return fmt.Errorf("%s | %w", "Maxmind", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		default:

			if key, ok := t.(string); ok {
				if s.AdditionalDatabaseConfigurationProperty == nil {
					s.AdditionalDatabaseConfigurationProperty = make(map[string]json.RawMessage, 0)
				}
				raw := new(json.RawMessage)
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "AdditionalDatabaseConfigurationProperty", err)
				}
				s.AdditionalDatabaseConfigurationProperty[key] = *raw
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s DatabaseConfiguration) MarshalJSON() ([]byte, error) {
	type opt DatabaseConfiguration
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.AdditionalDatabaseConfigurationProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalDatabaseConfigurationProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewDatabaseConfiguration returns a DatabaseConfiguration.
func NewDatabaseConfiguration() *DatabaseConfiguration {
	r := &DatabaseConfiguration{
		AdditionalDatabaseConfigurationProperty: make(map[string]json.RawMessage),
	}

	return r
}
