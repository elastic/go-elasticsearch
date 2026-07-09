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
// https://github.com/elastic/elasticsearch-specification/tree/37285cbd3fd155f913b50d880b40ec45f9df64b3

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// SourceOnlyRepository type.
//
// https://github.com/elastic/elasticsearch-specification/blob/37285cbd3fd155f913b50d880b40ec45f9df64b3/specification/snapshot/_types/SnapshotRepository.ts#L105-L115
type SourceOnlyRepository struct {
	// Settings The repository settings.
	Settings SourceOnlyRepositorySettings `json:"settings"`
	// Type The source-only repository type.
	Type string  `json:"type,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

func (s *SourceOnlyRepository) UnmarshalJSON(data []byte) error {

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

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			kind := make(map[string]string, 0)
			localDec := json.NewDecoder(source)
			localDec.Decode(&kind)
			source.Seek(0, io.SeekStart)

			switch kind["delegate_type"] {

			case "fs":
				o := NewSourceOnlyRepositorySettingsForSharedFileSystem()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "fs", err)
				}
				s.Settings = *o
			case "url":
				o := NewSourceOnlyRepositorySettingsForReadOnlyUrl()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "url", err)
				}
				s.Settings = *o
			case "azure":
				o := NewSourceOnlyRepositorySettingsForAzure()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "azure", err)
				}
				s.Settings = *o
			case "gcs":
				o := NewSourceOnlyRepositorySettingsForGcs()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "gcs", err)
				}
				s.Settings = *o
			case "s3":
				o := NewSourceOnlyRepositorySettingsForS3()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "s3", err)
				}
				s.Settings = *o
			default:
				if err := localDec.Decode(&s.Settings); err != nil {
					return fmt.Errorf("Settings | %w", err)
				}
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
func (s SourceOnlyRepository) MarshalJSON() ([]byte, error) {
	type innerSourceOnlyRepository SourceOnlyRepository
	tmp := innerSourceOnlyRepository{
		Settings: s.Settings,
		Type:     s.Type,
		Uuid:     s.Uuid,
	}

	tmp.Type = "source"

	return json.Marshal(tmp)
}

// NewSourceOnlyRepository returns a SourceOnlyRepository.
func NewSourceOnlyRepository() *SourceOnlyRepository {
	r := &SourceOnlyRepository{}

	return r
}

type SourceOnlyRepositoryVariant interface {
	SourceOnlyRepositoryCaster() *SourceOnlyRepository
}

func (s *SourceOnlyRepository) SourceOnlyRepositoryCaster() *SourceOnlyRepository {
	return s
}

func (s *SourceOnlyRepository) RepositoryCaster() *Repository {
	if s == nil {
		return nil
	}
	o := Repository(s)
	return &o
}
