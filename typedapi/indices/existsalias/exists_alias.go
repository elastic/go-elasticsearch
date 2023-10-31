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

// Returns information about whether a particular alias exists.
package existsalias

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
)

const (
	nameMask = iota + 1

	indexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ExistsAlias struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	name  string
	index string
}

// NewExistsAlias type alias for index.
type NewExistsAlias func(name string) *ExistsAlias

// NewExistsAliasFunc returns a new instance of ExistsAlias with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewExistsAliasFunc(tp elastictransport.Interface) NewExistsAlias {
	return func(name string) *ExistsAlias {
		n := New(tp)

		n._name(name)

		return n
	}
}

// Returns information about whether a particular alias exists.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-aliases.html
func New(tp elastictransport.Interface) *ExistsAlias {
	r := &ExistsAlias{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ExistsAlias) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_alias")
		path.WriteString("/")

		path.WriteString(r.name)

		method = http.MethodHead
	case r.paramSet == indexMask|nameMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_alias")
		path.WriteString("/")

		path.WriteString(r.name)

		method = http.MethodHead
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
func (r ExistsAlias) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the ExistsAlias query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a existsalias.Response
func (r ExistsAlias) Do(ctx context.Context) (bool, error) {
	return r.IsSuccess(ctx)
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r ExistsAlias) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the ExistsAlias headers map.
func (r *ExistsAlias) Header(key, value string) *ExistsAlias {
	r.headers.Set(key, value)

	return r
}

// Name Comma-separated list of aliases to check. Supports wildcards (`*`).
// API Name: name
func (r *ExistsAlias) _name(name string) *ExistsAlias {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// Index Comma-separated list of data streams or indices used to limit the request.
// Supports wildcards (`*`).
// To target all data streams and indices, omit this parameter or use `*` or
// `_all`.
// API Name: index
func (r *ExistsAlias) Index(index string) *ExistsAlias {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices If `false`, the request returns an error if any wildcard expression, index
// alias, or `_all` value targets only missing or closed indices.
// This behavior applies even if the request targets other open indices.
// API name: allow_no_indices
func (r *ExistsAlias) AllowNoIndices(allownoindices bool) *ExistsAlias {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match.
// If the request can target data streams, this argument determines whether
// wildcard expressions match hidden data streams.
// Supports comma-separated values, such as `open,hidden`.
// Valid values are: `all`, `open`, `closed`, `hidden`, `none`.
// API name: expand_wildcards
func (r *ExistsAlias) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *ExistsAlias {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// IgnoreUnavailable If `false`, requests that include a missing data stream or index in the
// target indices or data streams return an error.
// API name: ignore_unavailable
func (r *ExistsAlias) IgnoreUnavailable(ignoreunavailable bool) *ExistsAlias {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// Local If `true`, the request retrieves information from the local node only.
// API name: local
func (r *ExistsAlias) Local(local bool) *ExistsAlias {
	r.values.Set("local", strconv.FormatBool(local))

	return r
}
