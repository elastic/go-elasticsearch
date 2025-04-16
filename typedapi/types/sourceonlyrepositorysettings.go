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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// SourceOnlyRepositorySettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/52c473efb1fb5320a5bac12572d0b285882862fb/specification/snapshot/_types/SnapshotRepository.ts#L414-L441
type SourceOnlyRepositorySettings struct {
	// ChunkSize Big files can be broken down into multiple smaller blobs in the blob store
	// during snapshotting.
	// It is not recommended to change this value from its default unless there is
	// an explicit reason for limiting the size of blobs in the repository.
	// Setting a value lower than the default can result in an increased number of
	// API calls to the blob store during snapshot create and restore operations
	// compared to using the default value and thus make both operations slower and
	// more costly.
	// Specify the chunk size as a byte unit, for example: `10MB`, `5KB`, 500B.
	// The default varies by repository type.
	ChunkSize ByteSize `json:"chunk_size,omitempty"`
	// Compress When set to `true`, metadata files are stored in compressed format.
	// This setting doesn't affect index files that are already compressed by
	// default.
	Compress *bool `json:"compress,omitempty"`
	// DelegateType The delegated repository type. For valid values, refer to the `type`
	// parameter.
	// Source repositories can use `settings` properties for its delegated
	// repository type.
	DelegateType *string `json:"delegate_type,omitempty"`
	// MaxNumberOfSnapshots The maximum number of snapshots the repository can contain.
	// The default is `Integer.MAX_VALUE`, which is 2^31-1 or `2147483647`.
	MaxNumberOfSnapshots *int `json:"max_number_of_snapshots,omitempty"`
	// MaxRestoreBytesPerSec The maximum snapshot restore rate per node.
	// It defaults to unlimited.
	// Note that restores are also throttled through recovery settings.
	MaxRestoreBytesPerSec ByteSize `json:"max_restore_bytes_per_sec,omitempty"`
	// MaxSnapshotBytesPerSec The maximum snapshot creation rate per node.
	// It defaults to 40mb per second.
	// Note that if the recovery settings for managed services are set, then it
	// defaults to unlimited, and the rate is additionally throttled through
	// recovery settings.
	MaxSnapshotBytesPerSec ByteSize `json:"max_snapshot_bytes_per_sec,omitempty"`
	// ReadOnly If `true`, the repository is read-only.
	// The cluster can retrieve and restore snapshots from the repository but not
	// write to the repository or create snapshots in it.
	//
	// Only a cluster with write access can create snapshots in the repository.
	// All other clusters connected to the repository should have the `readonly`
	// parameter set to `true`.
	//
	// If `false`, the cluster can write to the repository and create snapshots in
	// it.
	//
	// IMPORTANT: If you register the same snapshot repository with multiple
	// clusters, only one cluster should have write access to the repository.
	// Having multiple clusters write to the repository at the same time risks
	// corrupting the contents of the repository.
	ReadOnly *bool `json:"read_only,omitempty"`
}

func (s *SourceOnlyRepositorySettings) UnmarshalJSON(data []byte) error {

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

		case "chunk_size":
			if err := dec.Decode(&s.ChunkSize); err != nil {
				return fmt.Errorf("%s | %w", "ChunkSize", err)
			}

		case "compress":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Compress", err)
				}
				s.Compress = &value
			case bool:
				s.Compress = &v
			}

		case "delegate_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "DelegateType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DelegateType = &o

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

		case "max_restore_bytes_per_sec":
			if err := dec.Decode(&s.MaxRestoreBytesPerSec); err != nil {
				return fmt.Errorf("%s | %w", "MaxRestoreBytesPerSec", err)
			}

		case "max_snapshot_bytes_per_sec":
			if err := dec.Decode(&s.MaxSnapshotBytesPerSec); err != nil {
				return fmt.Errorf("%s | %w", "MaxSnapshotBytesPerSec", err)
			}

		case "read_only", "readonly":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ReadOnly", err)
				}
				s.ReadOnly = &value
			case bool:
				s.ReadOnly = &v
			}

		}
	}
	return nil
}

// NewSourceOnlyRepositorySettings returns a SourceOnlyRepositorySettings.
func NewSourceOnlyRepositorySettings() *SourceOnlyRepositorySettings {
	r := &SourceOnlyRepositorySettings{}

	return r
}

// true

type SourceOnlyRepositorySettingsVariant interface {
	SourceOnlyRepositorySettingsCaster() *SourceOnlyRepositorySettings
}

func (s *SourceOnlyRepositorySettings) SourceOnlyRepositorySettingsCaster() *SourceOnlyRepositorySettings {
	return s
}
