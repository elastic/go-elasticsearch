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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// The configuration necessary to identify which IP geolocation provider to use
// to download a database, as well as any provider-specific configuration
// necessary for such downloading. At present, the only supported providers are
// `maxmind` and `ipinfo`, and the `maxmind` provider requires that an
// `account_id` (string) is configured. A provider (either `maxmind` or
// `ipinfo`) must be specified. The web and local providers can be returned as
// read only configurations.
//
// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/ingest/_types/Database.ts#L22-L37
type DatabaseConfiguration struct {
	Ipinfo  *Ipinfo  `json:"ipinfo,omitempty"`
	Maxmind *Maxmind `json:"maxmind,omitempty"`
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

		}
	}
	return nil
}

// NewDatabaseConfiguration returns a DatabaseConfiguration.
func NewDatabaseConfiguration() *DatabaseConfiguration {
	r := &DatabaseConfiguration{}

	return r
}

type DatabaseConfigurationVariant interface {
	DatabaseConfigurationCaster() *DatabaseConfiguration
}

func (s *DatabaseConfiguration) DatabaseConfigurationCaster() *DatabaseConfiguration {
	return s
}
