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

// SourceOnlyRepositorySettingsForGcs type.
//
// https://github.com/elastic/elasticsearch-specification/blob/37285cbd3fd155f913b50d880b40ec45f9df64b3/specification/snapshot/_types/SnapshotRepository.ts#L446-L450
type SourceOnlyRepositorySettingsForGcs struct {
	// ApplicationName The name used by the client when it uses the Google Cloud Storage service.
	ApplicationName *string `json:"application_name,omitempty"`
	// BasePath The path to the repository data within the bucket. It defaults to the root of
	// the bucket.
	//
	// NOTE: Don't set `base_path` when configuring a snapshot repository for
	// Elastic Cloud Enterprise. Elastic Cloud Enterprise automatically generates
	// the `base_path` for each deployment so that multiple deployments can share
	// the same bucket.
	BasePath *string `json:"base_path,omitempty"`
	// Bucket The name of the bucket to be used for snapshots.
	Bucket string `json:"bucket"`
	// Client The name of the client to use to connect to Google Cloud Storage.
	Client       *string `json:"client,omitempty"`
	DelegateType string  `json:"delegate_type,omitempty"`
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

func (s *SourceOnlyRepositorySettingsForGcs) UnmarshalJSON(data []byte) error {

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

		case "application_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ApplicationName", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ApplicationName = &o

		case "base_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "BasePath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BasePath = &o

		case "bucket":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Bucket", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Bucket = o

		case "client":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Client", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Client = &o

		case "delegate_type":
			if err := dec.Decode(&s.DelegateType); err != nil {
				return fmt.Errorf("%s | %w", "DelegateType", err)
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
func (s SourceOnlyRepositorySettingsForGcs) MarshalJSON() ([]byte, error) {
	type innerSourceOnlyRepositorySettingsForGcs SourceOnlyRepositorySettingsForGcs
	tmp := innerSourceOnlyRepositorySettingsForGcs{
		ApplicationName: s.ApplicationName,
		BasePath:        s.BasePath,
		Bucket:          s.Bucket,
		Client:          s.Client,
		DelegateType:    s.DelegateType,
		Readonly:        s.Readonly,
	}

	tmp.DelegateType = "gcs"

	return json.Marshal(tmp)
}

// NewSourceOnlyRepositorySettingsForGcs returns a SourceOnlyRepositorySettingsForGcs.
func NewSourceOnlyRepositorySettingsForGcs() *SourceOnlyRepositorySettingsForGcs {
	r := &SourceOnlyRepositorySettingsForGcs{}

	return r
}

type SourceOnlyRepositorySettingsForGcsVariant interface {
	SourceOnlyRepositorySettingsForGcsCaster() *SourceOnlyRepositorySettingsForGcs
}

func (s *SourceOnlyRepositorySettingsForGcs) SourceOnlyRepositorySettingsForGcsCaster() *SourceOnlyRepositorySettingsForGcs {
	return s
}

func (s *SourceOnlyRepositorySettingsForGcs) SourceOnlyRepositorySettingsCaster() *SourceOnlyRepositorySettings {
	if s == nil {
		return nil
	}
	o := SourceOnlyRepositorySettings(s)
	return &o
}
