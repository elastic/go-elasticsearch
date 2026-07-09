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

// SourceOnlyRepositorySettingsForAzure type.
//
// https://github.com/elastic/elasticsearch-specification/blob/37285cbd3fd155f913b50d880b40ec45f9df64b3/specification/snapshot/_types/SnapshotRepository.ts#L440-L444
type SourceOnlyRepositorySettingsForAzure struct {
	// BasePath The path to the repository data within the container. It defaults to the root
	// directory.
	//
	// NOTE: Don't set `base_path` when configuring a snapshot repository for
	// Elastic Cloud Enterprise. Elastic Cloud Enterprise automatically generates
	// the `base_path` for each deployment so that multiple deployments can share
	// the same bucket.
	BasePath *string `json:"base_path,omitempty"`
	// Client The name of the Azure repository client to use.
	Client *string `json:"client,omitempty"`
	// Container The Azure container.
	Container    *string `json:"container,omitempty"`
	DelegateType string  `json:"delegate_type,omitempty"`
	// DeleteObjectsMaxSize The maxmimum batch size, between 1 and 256, used for `BlobBatch` requests.
	// Defaults to 256 which is the maximum number supported by the Azure blob batch
	// API.
	DeleteObjectsMaxSize *int `json:"delete_objects_max_size,omitempty"`
	// LocationMode Either `primary_only` or `secondary_only`. Note that if you set it to
	// `secondary_only`, it will force `readonly` to `true`.
	LocationMode *string `json:"location_mode,omitempty"`
	// MaxConcurrentBatchDeletes The maximum number of concurrent batch delete requests that will be submitted
	// for any individual bulk delete with `BlobBatch`. Note that the effective
	// number of concurrent deletes is further limited by the Azure client
	// connection and event loop thread limits. Defaults to 10, minimum is 1,
	// maximum is 100.
	MaxConcurrentBatchDeletes *int `json:"max_concurrent_batch_deletes,omitempty"`
	// Readonly If `true`, the repository is read-only. The cluster can retrieve and restore
	// snapshots from the repository but not write to the repository or create
	// snapshots in it.
	//
	// Only a cluster with write access can create snapshots in the repository. All
	// other clusters connected to the repository should have the `readonly`
	// parameter set to `true`. If `false`, the cluster can write to the repository
	// and create snapshots in it.
	//
	// IMPORTANT: If you register the same snapshot repository with multiple
	// clusters, only one cluster should have write access to the repository. Having
	// multiple clusters write to the repository at the same time risks corrupting
	// the contents of the repository.
	Readonly *bool `json:"readonly,omitempty"`
}

func (s *SourceOnlyRepositorySettingsForAzure) UnmarshalJSON(data []byte) error {

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

		case "container":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Container", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Container = &o

		case "delegate_type":
			if err := dec.Decode(&s.DelegateType); err != nil {
				return fmt.Errorf("%s | %w", "DelegateType", err)
			}

		case "delete_objects_max_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DeleteObjectsMaxSize", err)
				}
				s.DeleteObjectsMaxSize = &value
			case float64:
				f := int(v)
				s.DeleteObjectsMaxSize = &f
			}

		case "location_mode":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "LocationMode", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.LocationMode = &o

		case "max_concurrent_batch_deletes":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxConcurrentBatchDeletes", err)
				}
				s.MaxConcurrentBatchDeletes = &value
			case float64:
				f := int(v)
				s.MaxConcurrentBatchDeletes = &f
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
func (s SourceOnlyRepositorySettingsForAzure) MarshalJSON() ([]byte, error) {
	type innerSourceOnlyRepositorySettingsForAzure SourceOnlyRepositorySettingsForAzure
	tmp := innerSourceOnlyRepositorySettingsForAzure{
		BasePath:                  s.BasePath,
		Client:                    s.Client,
		Container:                 s.Container,
		DelegateType:              s.DelegateType,
		DeleteObjectsMaxSize:      s.DeleteObjectsMaxSize,
		LocationMode:              s.LocationMode,
		MaxConcurrentBatchDeletes: s.MaxConcurrentBatchDeletes,
		Readonly:                  s.Readonly,
	}

	tmp.DelegateType = "azure"

	return json.Marshal(tmp)
}

// NewSourceOnlyRepositorySettingsForAzure returns a SourceOnlyRepositorySettingsForAzure.
func NewSourceOnlyRepositorySettingsForAzure() *SourceOnlyRepositorySettingsForAzure {
	r := &SourceOnlyRepositorySettingsForAzure{}

	return r
}

type SourceOnlyRepositorySettingsForAzureVariant interface {
	SourceOnlyRepositorySettingsForAzureCaster() *SourceOnlyRepositorySettingsForAzure
}

func (s *SourceOnlyRepositorySettingsForAzure) SourceOnlyRepositorySettingsForAzureCaster() *SourceOnlyRepositorySettingsForAzure {
	return s
}

func (s *SourceOnlyRepositorySettingsForAzure) SourceOnlyRepositorySettingsCaster() *SourceOnlyRepositorySettings {
	if s == nil {
		return nil
	}
	o := SourceOnlyRepositorySettings(s)
	return &o
}
