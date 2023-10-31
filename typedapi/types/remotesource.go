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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// RemoteSource type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_global/reindex/types.ts#L99-L125
type RemoteSource struct {
	// ConnectTimeout The remote connection timeout.
	// Defaults to 30 seconds.
	ConnectTimeout Duration `json:"connect_timeout,omitempty"`
	// Headers An object containing the headers of the request.
	Headers map[string]string `json:"headers,omitempty"`
	// Host The URL for the remote instance of Elasticsearch that you want to index from.
	Host string `json:"host"`
	// Password The password to use for authentication with the remote host.
	Password *string `json:"password,omitempty"`
	// SocketTimeout The remote socket read timeout. Defaults to 30 seconds.
	SocketTimeout Duration `json:"socket_timeout,omitempty"`
	// Username The username to use for authentication with the remote host.
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

		case "connect_timeout":
			if err := dec.Decode(&s.ConnectTimeout); err != nil {
				return err
			}

		case "headers":
			if s.Headers == nil {
				s.Headers = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Headers); err != nil {
				return err
			}

		case "host":
			if err := dec.Decode(&s.Host); err != nil {
				return err
			}

		case "password":
			if err := dec.Decode(&s.Password); err != nil {
				return err
			}

		case "socket_timeout":
			if err := dec.Decode(&s.SocketTimeout); err != nil {
				return err
			}

		case "username":
			if err := dec.Decode(&s.Username); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewRemoteSource returns a RemoteSource.
func NewRemoteSource() *RemoteSource {
	r := &RemoteSource{
		Headers: make(map[string]string, 0),
	}

	return r
}
