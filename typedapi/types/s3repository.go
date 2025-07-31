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

// S3Repository type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/snapshot/_types/SnapshotRepository.ts#L50-L53
type S3Repository struct {
	Settings S3RepositorySettings `json:"settings"`
	Type     string               `json:"type,omitempty"`
	Uuid     *string              `json:"uuid,omitempty"`
}

func (s *S3Repository) UnmarshalJSON(data []byte) error {

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

		case "settings":
			if err := dec.Decode(&s.Settings); err != nil {
				return fmt.Errorf("%s | %w", "Settings", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "uuid":
			if err := dec.Decode(&s.Uuid); err != nil {
				return fmt.Errorf("%s | %w", "Uuid", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s S3Repository) MarshalJSON() ([]byte, error) {
	type innerS3Repository S3Repository
	tmp := innerS3Repository{
		Settings: s.Settings,
		Type:     s.Type,
		Uuid:     s.Uuid,
	}

	tmp.Type = "s3"

	return json.Marshal(tmp)
}

// NewS3Repository returns a S3Repository.
func NewS3Repository() *S3Repository {
	r := &S3Repository{}

	return r
}
