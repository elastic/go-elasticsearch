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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// RemoteSource type.
//
// https://github.com/elastic/elasticsearch-specification/blob/6785a6caa1fa3ca5ab3308963d79dce923a3469f/specification/_global/reindex/types.ts#L115-L151
type RemoteSource struct {
	// ApiKey The API key to use for authentication with the remote host (as an alternative
	// to basic auth when the remote cluster is in Elastic Cloud).
	// (It is not permitted to set this and also to set an `Authorization` header
	// via `headers`.)
	ApiKey *string `json:"api_key,omitempty"`
	// ConnectTimeout The remote connection timeout.
	ConnectTimeout Duration `json:"connect_timeout,omitempty"`
	// Headers An object containing the headers of the request.
	Headers map[string]string `json:"headers,omitempty"`
	// Host The URL for the remote instance of Elasticsearch that you want to index from.
	// This information is required when you're indexing from remote.
	Host string `json:"host"`
	// Password The password to use for authentication with the remote host (required when
	// using basic auth).
	Password *string `json:"password,omitempty"`
	// SocketTimeout The remote socket read timeout.
	SocketTimeout Duration `json:"socket_timeout,omitempty"`
	// Username The username to use for authentication with the remote host (required when
	// using basic auth).
	Username *string `json:"username,omitempty"`
}

func (s *RemoteSource) UnmarshalJSON(data []byte) error {

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

		case "api_key":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ApiKey", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ApiKey = &o

		case "connect_timeout":
			if err := dec.Decode(&s.ConnectTimeout); err != nil {
				return fmt.Errorf("%s | %w", "ConnectTimeout", err)
			}

		case "headers":
			if s.Headers == nil {
				s.Headers = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Headers); err != nil {
				return fmt.Errorf("%s | %w", "Headers", err)
			}

		case "host":
			if err := dec.Decode(&s.Host); err != nil {
				return fmt.Errorf("%s | %w", "Host", err)
			}

		case "password":
			if err := dec.Decode(&s.Password); err != nil {
				return fmt.Errorf("%s | %w", "Password", err)
			}

		case "socket_timeout":
			if err := dec.Decode(&s.SocketTimeout); err != nil {
				return fmt.Errorf("%s | %w", "SocketTimeout", err)
			}

		case "username":
			if err := dec.Decode(&s.Username); err != nil {
				return fmt.Errorf("%s | %w", "Username", err)
			}

		}
	}
	return nil
}

// NewRemoteSource returns a RemoteSource.
func NewRemoteSource() *RemoteSource {
	r := &RemoteSource{
		Headers: make(map[string]string),
	}

	return r
}

type RemoteSourceVariant interface {
	RemoteSourceCaster() *RemoteSource
}

func (s *RemoteSource) RemoteSourceCaster() *RemoteSource {
	return s
}
