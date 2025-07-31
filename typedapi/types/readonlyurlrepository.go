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
)

// ReadOnlyUrlRepository type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/snapshot/_types/SnapshotRepository.ts#L92-L102
type ReadOnlyUrlRepository struct {
	// Settings The repository settings.
	Settings ReadOnlyUrlRepositorySettings `json:"settings"`
	// Type The read-only URL repository type.
	Type string  `json:"type,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

func (s *ReadOnlyUrlRepository) UnmarshalJSON(data []byte) error {

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
func (s ReadOnlyUrlRepository) MarshalJSON() ([]byte, error) {
	type innerReadOnlyUrlRepository ReadOnlyUrlRepository
	tmp := innerReadOnlyUrlRepository{
		Settings: s.Settings,
		Type:     s.Type,
		Uuid:     s.Uuid,
	}

	tmp.Type = "url"

	return json.Marshal(tmp)
}

// NewReadOnlyUrlRepository returns a ReadOnlyUrlRepository.
func NewReadOnlyUrlRepository() *ReadOnlyUrlRepository {
	r := &ReadOnlyUrlRepository{}

	return r
}

type ReadOnlyUrlRepositoryVariant interface {
	ReadOnlyUrlRepositoryCaster() *ReadOnlyUrlRepository
}

func (s *ReadOnlyUrlRepository) ReadOnlyUrlRepositoryCaster() *ReadOnlyUrlRepository {
	return s
}

func (s *ReadOnlyUrlRepository) RepositoryCaster() *Repository {
	o := Repository(s)
	return &o
}
