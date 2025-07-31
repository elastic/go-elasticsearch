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
	"strconv"
)

// ReadOnlyUrlRepositorySettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/snapshot/_types/SnapshotRepository.ts#L377-L412
type ReadOnlyUrlRepositorySettings struct {
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
	// HttpMaxRetries The maximum number of retries for HTTP and HTTPS URLs.
	HttpMaxRetries *int `json:"http_max_retries,omitempty"`
	// HttpSocketTimeout The maximum wait time for data transfers over a connection.
	HttpSocketTimeout Duration `json:"http_socket_timeout,omitempty"`
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
	// Url The URL location of the root of the shared filesystem repository.
	// The following protocols are supported:
	//
	// * `file`
	// * `ftp`
	// * `http`
	// * `https`
	// * `jar`
	//
	// URLs using the HTTP, HTTPS, or FTP protocols must be explicitly allowed with
	// the `repositories.url.allowed_urls` cluster setting.
	// This setting supports wildcards in the place of a host, path, query, or
	// fragment in the URL.
	//
	// URLs using the file protocol must point to the location of a shared
	// filesystem accessible to all master and data nodes in the cluster.
	// This location must be registered in the `path.repo` setting.
	// You don't need to register URLs using the FTP, HTTP, HTTPS, or JAR protocols
	// in the `path.repo` setting.
	Url string `json:"url"`
}

func (s *ReadOnlyUrlRepositorySettings) UnmarshalJSON(data []byte) error {

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

		case "http_max_retries":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "HttpMaxRetries", err)
				}
				s.HttpMaxRetries = &value
			case float64:
				f := int(v)
				s.HttpMaxRetries = &f
			}

		case "http_socket_timeout":
			if err := dec.Decode(&s.HttpSocketTimeout); err != nil {
				return fmt.Errorf("%s | %w", "HttpSocketTimeout", err)
			}

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

		case "url":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Url", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Url = o

		}
	}
	return nil
}

// NewReadOnlyUrlRepositorySettings returns a ReadOnlyUrlRepositorySettings.
func NewReadOnlyUrlRepositorySettings() *ReadOnlyUrlRepositorySettings {
	r := &ReadOnlyUrlRepositorySettings{}

	return r
}

type ReadOnlyUrlRepositorySettingsVariant interface {
	ReadOnlyUrlRepositorySettingsCaster() *ReadOnlyUrlRepositorySettings
}

func (s *ReadOnlyUrlRepositorySettings) ReadOnlyUrlRepositorySettingsCaster() *ReadOnlyUrlRepositorySettings {
	return s
}
