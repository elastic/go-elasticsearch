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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

// Clears cluster voting config exclusions.
package deletevotingconfigexclusions

import (
	gobytes "bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteVotingConfigExclusions struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int
}

// NewDeleteVotingConfigExclusions type alias for index.
type NewDeleteVotingConfigExclusions func() *DeleteVotingConfigExclusions

// NewDeleteVotingConfigExclusionsFunc returns a new instance of DeleteVotingConfigExclusions with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDeleteVotingConfigExclusionsFunc(tp elastictransport.Interface) NewDeleteVotingConfigExclusions {
	return func() *DeleteVotingConfigExclusions {
		n := New(tp)

		return n
	}
}

// Clears cluster voting config exclusions.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/voting-config-exclusions.html
func New(tp elastictransport.Interface) *DeleteVotingConfigExclusions {
	r := &DeleteVotingConfigExclusions{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *DeleteVotingConfigExclusions) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_cluster")
		path.WriteString("/")
		path.WriteString("voting_config_exclusions")

		method = http.MethodDelete
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
func (r DeleteVotingConfigExclusions) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the DeleteVotingConfigExclusions query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r DeleteVotingConfigExclusions) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the DeleteVotingConfigExclusions headers map.
func (r *DeleteVotingConfigExclusions) Header(key, value string) *DeleteVotingConfigExclusions {
	r.headers.Set(key, value)

	return r
}

// WaitForRemoval Specifies whether to wait for all excluded nodes to be removed from the
// cluster before clearing the voting configuration exclusions list.
// Defaults to true, meaning that all excluded nodes must be removed from
// the cluster before this API takes any action. If set to false then the
// voting configuration exclusions list is cleared even if some excluded
// nodes are still in the cluster.
// API name: wait_for_removal
func (r *DeleteVotingConfigExclusions) WaitForRemoval(b bool) *DeleteVotingConfigExclusions {
	r.values.Set("wait_for_removal", strconv.FormatBool(b))

	return r
}
