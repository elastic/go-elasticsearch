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

// SourceOnlyRepositorySettingsForReadOnlyUrl type.
//
// https://github.com/elastic/elasticsearch-specification/blob/37285cbd3fd155f913b50d880b40ec45f9df64b3/specification/snapshot/_types/SnapshotRepository.ts#L434-L438
type SourceOnlyRepositorySettingsForReadOnlyUrl struct {
	DelegateType string `json:"delegate_type,omitempty"`
	// HttpMaxRetries The maximum number of retries for HTTP and HTTPS URLs.
	HttpMaxRetries *int `json:"http_max_retries,omitempty"`
	// HttpSocketTimeout The maximum wait time for data transfers over a connection.
	HttpSocketTimeout Duration `json:"http_socket_timeout,omitempty"`
	// MaxNumberOfSnapshots The maximum number of snapshots the repository can contain. The default is
	// `Integer.MAX_VALUE`, which is 2^31-1 or `2147483647`.
	MaxNumberOfSnapshots *int `json:"max_number_of_snapshots,omitempty"`
	// Url The URL location of the root of the shared filesystem repository. The
	// following protocols are supported:
	//
	//   - `file`
	//   - `ftp`
	//   - `http`
	//   - `https`
	//   - `jar`
	//
	// URLs using the HTTP, HTTPS, or FTP protocols must be explicitly allowed with
	// the `repositories.url.allowed_urls` cluster setting. This setting supports
	// wildcards in the place of a host, path, query, or fragment in the URL.
	//
	// URLs using the file protocol must point to the location of a shared
	// filesystem accessible to all master and data nodes in the cluster. This
	// location must be registered in the `path.repo` setting. You don't need to
	// register URLs using the FTP, HTTP, HTTPS, or JAR protocols in the `path.repo`
	// setting.
	Url string `json:"url"`
}

func (s *SourceOnlyRepositorySettingsForReadOnlyUrl) UnmarshalJSON(data []byte) error {

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

// MarshalJSON override marshalling to include literal value
func (s SourceOnlyRepositorySettingsForReadOnlyUrl) MarshalJSON() ([]byte, error) {
	type innerSourceOnlyRepositorySettingsForReadOnlyUrl SourceOnlyRepositorySettingsForReadOnlyUrl
	tmp := innerSourceOnlyRepositorySettingsForReadOnlyUrl{
		DelegateType:         s.DelegateType,
		HttpMaxRetries:       s.HttpMaxRetries,
		HttpSocketTimeout:    s.HttpSocketTimeout,
		MaxNumberOfSnapshots: s.MaxNumberOfSnapshots,
		Url:                  s.Url,
	}

	tmp.DelegateType = "url"

	return json.Marshal(tmp)
}

// NewSourceOnlyRepositorySettingsForReadOnlyUrl returns a SourceOnlyRepositorySettingsForReadOnlyUrl.
func NewSourceOnlyRepositorySettingsForReadOnlyUrl() *SourceOnlyRepositorySettingsForReadOnlyUrl {
	r := &SourceOnlyRepositorySettingsForReadOnlyUrl{}

	return r
}

type SourceOnlyRepositorySettingsForReadOnlyUrlVariant interface {
	SourceOnlyRepositorySettingsForReadOnlyUrlCaster() *SourceOnlyRepositorySettingsForReadOnlyUrl
}

func (s *SourceOnlyRepositorySettingsForReadOnlyUrl) SourceOnlyRepositorySettingsForReadOnlyUrlCaster() *SourceOnlyRepositorySettingsForReadOnlyUrl {
	return s
}

func (s *SourceOnlyRepositorySettingsForReadOnlyUrl) SourceOnlyRepositorySettingsCaster() *SourceOnlyRepositorySettings {
	if s == nil {
		return nil
	}
	o := SourceOnlyRepositorySettings(s)
	return &o
}
