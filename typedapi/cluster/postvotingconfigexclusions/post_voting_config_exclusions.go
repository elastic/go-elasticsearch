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

// Updates the cluster voting config exclusions by node ids or node names.
package postvotingconfigexclusions

import (
	gobytes "bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PostVotingConfigExclusions struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int
}

// NewPostVotingConfigExclusions type alias for index.
type NewPostVotingConfigExclusions func() *PostVotingConfigExclusions

// NewPostVotingConfigExclusionsFunc returns a new instance of PostVotingConfigExclusions with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPostVotingConfigExclusionsFunc(tp elastictransport.Interface) NewPostVotingConfigExclusions {
	return func() *PostVotingConfigExclusions {
		n := New(tp)

		return n
	}
}

// Updates the cluster voting config exclusions by node ids or node names.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/voting-config-exclusions.html
func New(tp elastictransport.Interface) *PostVotingConfigExclusions {
	r := &PostVotingConfigExclusions{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PostVotingConfigExclusions) HttpRequest(ctx context.Context) (*http.Request, error) {
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
func (r PostVotingConfigExclusions) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PostVotingConfigExclusions query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a postvotingconfigexclusions.Response
func (r PostVotingConfigExclusions) Do(ctx context.Context) (bool, error) {
	return r.IsSuccess(ctx)
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r PostVotingConfigExclusions) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the PostVotingConfigExclusions headers map.
func (r *PostVotingConfigExclusions) Header(key, value string) *PostVotingConfigExclusions {
	r.headers.Set(key, value)

	return r
}

// NodeNames A comma-separated list of the names of the nodes to exclude from the
// voting configuration. If specified, you may not also specify node_ids.
// API name: node_names
func (r *PostVotingConfigExclusions) NodeNames(names ...string) *PostVotingConfigExclusions {
	r.values.Set("node_names", strings.Join(names, ","))

	return r
}

// NodeIds A comma-separated list of the persistent ids of the nodes to exclude
// from the voting configuration. If specified, you may not also specify
// node_names.
// API name: node_ids
func (r *PostVotingConfigExclusions) NodeIds(ids ...string) *PostVotingConfigExclusions {
	r.values.Set("node_ids", strings.Join(ids, ","))

	return r
}

// Timeout When adding a voting configuration exclusion, the API waits for the
// specified nodes to be excluded from the voting configuration before
// returning. If the timeout expires before the appropriate condition
// is satisfied, the request fails and returns an error.
// API name: timeout
func (r *PostVotingConfigExclusions) Timeout(duration string) *PostVotingConfigExclusions {
	r.values.Set("timeout", duration)

	return r
}
