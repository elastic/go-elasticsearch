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
// https://github.com/elastic/elasticsearch-specification/tree/3a94b6715915b1e9311724a2614c643368eece90

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// DatabaseConfigurationFull type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3a94b6715915b1e9311724a2614c643368eece90/specification/ingest/_types/Database.ts#L39-L53
type DatabaseConfigurationFull struct {
	Ipinfo  *Ipinfo  `json:"ipinfo,omitempty"`
	Local   *Local   `json:"local,omitempty"`
	Maxmind *Maxmind `json:"maxmind,omitempty"`
	// Name The provider-assigned name of the IP geolocation database to download.
	Name string `json:"name"`
	Web  *Web   `json:"web,omitempty"`
}

func (s *DatabaseConfigurationFull) UnmarshalJSON(data []byte) error {

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

		case "local":
			if err := dec.Decode(&s.Local); err != nil {
				return fmt.Errorf("%s | %w", "Local", err)
			}

		case "maxmind":
			if err := dec.Decode(&s.Maxmind); err != nil {
				return fmt.Errorf("%s | %w", "Maxmind", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "web":
			if err := dec.Decode(&s.Web); err != nil {
				return fmt.Errorf("%s | %w", "Web", err)
			}

		}
	}
	return nil
}

// NewDatabaseConfigurationFull returns a DatabaseConfigurationFull.
func NewDatabaseConfigurationFull() *DatabaseConfigurationFull {
	r := &DatabaseConfigurationFull{}

	return r
}
