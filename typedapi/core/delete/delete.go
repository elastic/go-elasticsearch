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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


// Removes a document from the index.
package delete

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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

const (
	idMask = iota + 1

	indexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Delete struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	id    string
	index string
}

// NewDelete type alias for index.
type NewDelete func(index, id string) *Delete

// NewDeleteFunc returns a new instance of Delete with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDeleteFunc(tp elastictransport.Interface) NewDelete {
	return func(index, id string) *Delete {
		n := New(tp)

		n.Id(id)

		n.Index(index)

		return n
	}
}

// Removes a document from the index.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-delete.html
func New(tp elastictransport.Interface) *Delete {
	r := &Delete{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Delete) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|idMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_doc")
		path.WriteString("/")

		path.WriteString(r.id)

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

// Do runs the http.Request through the provided transport.
func (r Delete) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Delete query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r Delete) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Do(ctx)

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

// Header set a key, value pair in the Delete headers map.
func (r *Delete) Header(key, value string) *Delete {
	r.headers.Set(key, value)

	return r
}

// Id The document ID
// API Name: id
func (r *Delete) Id(v string) *Delete {
	r.paramSet |= idMask
	r.id = v

	return r
}

// Index The name of the index
// API Name: index
func (r *Delete) Index(v string) *Delete {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// IfPrimaryTerm only perform the delete operation if the last operation that has changed the
// document has the specified primary term
// API name: if_primary_term
func (r *Delete) IfPrimaryTerm(value string) *Delete {
	r.values.Set("if_primary_term", value)

	return r
}

// IfSeqNo only perform the delete operation if the last operation that has changed the
// document has the specified sequence number
// API name: if_seq_no
func (r *Delete) IfSeqNo(value string) *Delete {
	r.values.Set("if_seq_no", value)

	return r
}

// Refresh If `true` then refresh the affected shards to make this operation visible to
// search, if `wait_for` then wait for a refresh to make this operation visible
// to search, if `false` (the default) then do nothing with refreshes.
// API name: refresh
func (r *Delete) Refresh(enum refresh.Refresh) *Delete {
	r.values.Set("refresh", enum.String())

	return r
}

// Routing Specific routing value
// API name: routing
func (r *Delete) Routing(value string) *Delete {
	r.values.Set("routing", value)

	return r
}

// Timeout Explicit operation timeout
// API name: timeout
func (r *Delete) Timeout(value string) *Delete {
	r.values.Set("timeout", value)

	return r
}

// Version Explicit version number for concurrency control
// API name: version
func (r *Delete) Version(value string) *Delete {
	r.values.Set("version", value)

	return r
}

// VersionType Specific version type
// API name: version_type
func (r *Delete) VersionType(enum versiontype.VersionType) *Delete {
	r.values.Set("version_type", enum.String())

	return r
}

// WaitForActiveShards Sets the number of shard copies that must be active before proceeding with
// the delete operation. Defaults to 1, meaning the primary shard only. Set to
// `all` for all shard copies, otherwise set to any non-negative value less than
// or equal to the total number of copies for the shard (number of replicas + 1)
// API name: wait_for_active_shards
func (r *Delete) WaitForActiveShards(value string) *Delete {
	r.values.Set("wait_for_active_shards", value)

	return r
}
