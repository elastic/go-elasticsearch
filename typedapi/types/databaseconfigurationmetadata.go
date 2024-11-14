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
// https://github.com/elastic/elasticsearch-specification/tree/4fcf747dfafc951e1dcf3077327e3dcee9107db3

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// DatabaseConfigurationMetadata type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4fcf747dfafc951e1dcf3077327e3dcee9107db3/specification/ingest/get_geoip_database/GetGeoipDatabaseResponse.ts#L29-L34
type DatabaseConfigurationMetadata struct {
	Database           DatabaseConfiguration `json:"database"`
	Id                 string                `json:"id"`
	ModifiedDateMillis int64                 `json:"modified_date_millis"`
	Version            int64                 `json:"version"`
}

func (s *DatabaseConfigurationMetadata) UnmarshalJSON(data []byte) error {

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

		case "modified_date_millis":
			if err := dec.Decode(&s.ModifiedDateMillis); err != nil {
				return fmt.Errorf("%s | %w", "ModifiedDateMillis", err)
			}

		case "version":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Version", err)
				}
				s.Version = value
			case float64:
				f := int64(v)
				s.Version = f
			}

		}
	}
	return nil
}

// NewDatabaseConfigurationMetadata returns a DatabaseConfigurationMetadata.
func NewDatabaseConfigurationMetadata() *DatabaseConfigurationMetadata {
	r := &DatabaseConfigurationMetadata{}

	return r
}
