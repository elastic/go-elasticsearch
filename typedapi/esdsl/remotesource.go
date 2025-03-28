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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _remoteSource struct {
	v *types.RemoteSource
}

func NewRemoteSource() *_remoteSource {

	return &_remoteSource{v: types.NewRemoteSource()}

}

// The remote connection timeout.
func (s *_remoteSource) ConnectTimeout(duration types.DurationVariant) *_remoteSource {

	s.v.ConnectTimeout = *duration.DurationCaster()

	return s
}

// An object containing the headers of the request.
func (s *_remoteSource) Headers(headers map[string]string) *_remoteSource {

	s.v.Headers = headers
	return s
}

func (s *_remoteSource) AddHeader(key string, value string) *_remoteSource {

	var tmp map[string]string
	if s.v.Headers == nil {
		s.v.Headers = make(map[string]string)
	} else {
		tmp = s.v.Headers
	}

	tmp[key] = value

	s.v.Headers = tmp
	return s
}

// The URL for the remote instance of Elasticsearch that you want to index from.
// This information is required when you're indexing from remote.
func (s *_remoteSource) Host(host string) *_remoteSource {

	s.v.Host = host

	return s
}

// The password to use for authentication with the remote host.
func (s *_remoteSource) Password(password string) *_remoteSource {

	s.v.Password = &password

	return s
}

// The remote socket read timeout.
func (s *_remoteSource) SocketTimeout(duration types.DurationVariant) *_remoteSource {

	s.v.SocketTimeout = *duration.DurationCaster()

	return s
}

// The username to use for authentication with the remote host.
func (s *_remoteSource) Username(username string) *_remoteSource {

	s.v.Username = &username

	return s
}

func (s *_remoteSource) RemoteSourceCaster() *types.RemoteSource {
	return s.v
}
