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
	"strconv"
)

// SourceOnlyRepositorySettingsForSharedFileSystem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/37285cbd3fd155f913b50d880b40ec45f9df64b3/specification/snapshot/_types/SnapshotRepository.ts#L428-L432
type SourceOnlyRepositorySettingsForSharedFileSystem struct {
	DelegateType string `json:"delegate_type,omitempty"`
	// Location The location of the shared filesystem used to store and retrieve snapshots.
	// This location must be registered in the `path.repo` setting on all master and
	// data nodes in the cluster. Unlike `path.repo`, this setting supports only a
	// single file path.
	Location string `json:"location"`
	// MaxNumberOfSnapshots The maximum number of snapshots the repository can contain. The default is
	// `Integer.MAX_VALUE`, which is 2^31-1 or `2147483647`.
	MaxNumberOfSnapshots *int `json:"max_number_of_snapshots,omitempty"`
	// Readonly If `true`, the repository is read-only. The cluster can retrieve and restore
	// snapshots from the repository but not write to the repository or create
	// snapshots in it.
	//
	// Only a cluster with write access can create snapshots in the repository. All
	// other clusters connected to the repository should have the `readonly`
	// parameter set to `true`.
	//
	// If `false`, the cluster can write to the repository and create snapshots in
	// it.
	//
	// IMPORTANT: If you register the same snapshot repository with multiple
	// clusters, only one cluster should have write access to the repository. Having
	// multiple clusters write to the repository at the same time risks corrupting
	// the contents of the repository.
	Readonly *bool `json:"readonly,omitempty"`
}

func (s *SourceOnlyRepositorySettingsForSharedFileSystem) UnmarshalJSON(data []byte) error {

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

		case "delegate_type":
			if err := dec.Decode(&s.DelegateType); err != nil {
				return fmt.Errorf("%s | %w", "DelegateType", err)
			}

		case "location":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Location", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Location = o

		case "max_number_of_snapshots":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxNumberOfSnapshots", err)
				}
				s.MaxNumberOfSnapshots = &value
			case float64:
				f := int(v)
				s.MaxNumberOfSnapshots = &f
			}

		case "readonly":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Readonly", err)
				}
				s.Readonly = &value
			case bool:
				s.Readonly = &v
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s SourceOnlyRepositorySettingsForSharedFileSystem) MarshalJSON() ([]byte, error) {
	type innerSourceOnlyRepositorySettingsForSharedFileSystem SourceOnlyRepositorySettingsForSharedFileSystem
	tmp := innerSourceOnlyRepositorySettingsForSharedFileSystem{
		DelegateType:         s.DelegateType,
		Location:             s.Location,
		MaxNumberOfSnapshots: s.MaxNumberOfSnapshots,
		Readonly:             s.Readonly,
	}

	tmp.DelegateType = "fs"

	return json.Marshal(tmp)
}

// NewSourceOnlyRepositorySettingsForSharedFileSystem returns a SourceOnlyRepositorySettingsForSharedFileSystem.
func NewSourceOnlyRepositorySettingsForSharedFileSystem() *SourceOnlyRepositorySettingsForSharedFileSystem {
	r := &SourceOnlyRepositorySettingsForSharedFileSystem{}

	return r
}

type SourceOnlyRepositorySettingsForSharedFileSystemVariant interface {
	SourceOnlyRepositorySettingsForSharedFileSystemCaster() *SourceOnlyRepositorySettingsForSharedFileSystem
}

func (s *SourceOnlyRepositorySettingsForSharedFileSystem) SourceOnlyRepositorySettingsForSharedFileSystemCaster() *SourceOnlyRepositorySettingsForSharedFileSystem {
	return s
}

func (s *SourceOnlyRepositorySettingsForSharedFileSystem) SourceOnlyRepositorySettingsCaster() *SourceOnlyRepositorySettings {
	if s == nil {
		return nil
	}
	o := SourceOnlyRepositorySettings(s)
	return &o
}
