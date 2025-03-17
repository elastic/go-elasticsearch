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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

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
// https://github.com/elastic/elasticsearch-specification/blob/ea991724f4dd4f90c496eff547d3cc2e6529f509/specification/snapshot/_types/SnapshotRepository.ts#L64-L78
type S3Repository struct {
	// Settings The repository settings.
	//
	// NOTE: In addition to the specified settings, you can also use all non-secure
	// client settings in the repository settings.
	// In this case, the client settings found in the repository settings will be
	// merged with those of the named client used by the repository.
	// Conflicts between client and repository settings are resolved by the
	// repository settings taking precedence over client settings.
	Settings S3RepositorySettings `json:"settings"`
	// Type The S3 repository type.
	Type string  `json:"type,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
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

// true

type S3RepositoryVariant interface {
	S3RepositoryCaster() *S3Repository
}

func (s *S3Repository) S3RepositoryCaster() *S3Repository {
	return s
}
