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

package mocktransport

import "net/http"

// RoundTripFunc is a function adapter for http.RoundTripper.
type RoundTripFunc func(*http.Request) (*http.Response, error)

// Transport is a mock implementation of http.RoundTripper.
type Transport struct {
	RoundTripFn        RoundTripFunc
	DefaultRoundTripFn RoundTripFunc
}

// RoundTrip executes the configured round-trip function.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.RoundTripFn != nil {
		return t.RoundTripFn(req)
	}

	if t.DefaultRoundTripFn != nil {
		return t.DefaultRoundTripFn(req)
	}

	return nil, nil
}

// Perform executes the configured round-trip function.
func (t *Transport) Perform(req *http.Request) (*http.Response, error) {
	return t.RoundTrip(req)
}
