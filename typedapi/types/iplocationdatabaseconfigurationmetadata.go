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

// IpLocationDatabaseConfigurationMetadata type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ingest/get_ip_location_database/GetIpLocationDatabaseResponse.ts#L28-L34
type IpLocationDatabaseConfigurationMetadata struct {
	Database           DatabaseConfigurationFull `json:"database"`
	Id                 string                    `json:"id"`
	ModifiedDate       *int64                    `json:"modified_date,omitempty"`
	ModifiedDateMillis *int64                    `json:"modified_date_millis,omitempty"`
	Version            int64                     `json:"version"`
}

func (s *IpLocationDatabaseConfigurationMetadata) UnmarshalJSON(data []byte) error {

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

		case "database":
			if err := dec.Decode(&s.Database); err != nil {
				return fmt.Errorf("%s | %w", "Database", err)
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "modified_date":
			if err := dec.Decode(&s.ModifiedDate); err != nil {
				return fmt.Errorf("%s | %w", "ModifiedDate", err)
			}

		case "modified_date_millis":
			if err := dec.Decode(&s.ModifiedDateMillis); err != nil {
				return fmt.Errorf("%s | %w", "ModifiedDateMillis", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewIpLocationDatabaseConfigurationMetadata returns a IpLocationDatabaseConfigurationMetadata.
func NewIpLocationDatabaseConfigurationMetadata() *IpLocationDatabaseConfigurationMetadata {
	r := &IpLocationDatabaseConfigurationMetadata{}

	return r
}
