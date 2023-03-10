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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

// Sets a cluster wide upgrade_mode setting that prepares machine learning
// indices for an upgrade.
package setupgrademode

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SetUpgradeMode struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int
}

// NewSetUpgradeMode type alias for index.
type NewSetUpgradeMode func() *SetUpgradeMode

// NewSetUpgradeModeFunc returns a new instance of SetUpgradeMode with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSetUpgradeModeFunc(tp elastictransport.Interface) NewSetUpgradeMode {
	return func() *SetUpgradeMode {
		n := New(tp)

		return n
	}
}

// Sets a cluster wide upgrade_mode setting that prepares machine learning
// indices for an upgrade.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-set-upgrade-mode.html
func New(tp elastictransport.Interface) *SetUpgradeMode {
	r := &SetUpgradeMode{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *SetUpgradeMode) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("set_upgrade_mode")

		method = http.MethodPost
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r SetUpgradeMode) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the SetUpgradeMode query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a setupgrademode.Response
func (r SetUpgradeMode) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil

	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	return nil, errorResponse
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r SetUpgradeMode) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(ioutil.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	return false, nil
}

// Header set a key, value pair in the SetUpgradeMode headers map.
func (r *SetUpgradeMode) Header(key, value string) *SetUpgradeMode {
	r.headers.Set(key, value)

	return r
}

// Enabled When `true`, it enables `upgrade_mode` which temporarily halts all job
// and datafeed tasks and prohibits new job and datafeed tasks from
// starting.
// API name: enabled
func (r *SetUpgradeMode) Enabled(b bool) *SetUpgradeMode {
	r.values.Set("enabled", strconv.FormatBool(b))

	return r
}

// Timeout The time to wait for the request to be completed.
// API name: timeout
func (r *SetUpgradeMode) Timeout(v string) *SetUpgradeMode {
	r.values.Set("timeout", v)

	return r
}
