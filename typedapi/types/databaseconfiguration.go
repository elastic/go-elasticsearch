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
// https://github.com/elastic/elasticsearch-specification/tree/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64

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
// https://github.com/elastic/elasticsearch-specification/blob/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64/specification/ingest/_types/Database.ts#L22-L29
type DatabaseConfiguration struct {
	// Maxmind The configuration necessary to identify which IP geolocation provider to use
	// to download the database, as well as any provider-specific configuration
	// necessary for such downloading.
	// At present, the only supported provider is maxmind, and the maxmind provider
	// requires that an account_id (string) is configured.
	Maxmind Maxmind `json:"maxmind"`
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
